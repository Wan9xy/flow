package flow

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

// Event 流程事件
// 流程事件是流程的运行实例，每个流程事件都有一个唯一的UUID，用于标识
// 流程事件是由节点组成的，节点之间通过边连接，其中ExpectEnd表示是否允许流程异常结束，StartJump表示是否允许流程跳回开始节点
type Event struct {
	gorm.Model
	UUID        string
	ProcessID   string
	Nodes       NodeArray
	AlwaysNodes NodeArray
	ExpectEnd   bool
	StartJump   bool
	NowAt       string
}

func (e *Event) TableName() string {
	return "t_event"
}

// EventStart 创建一个新的流程事件
func EventStart(process Process) (event *Event, node *Node, err error) {
	if process.Nodes == nil || len(process.Nodes) == 0 {
		err = errors.New("process has no node")
		return
	}
	if process.Nodes[0].Name != "start" {
		err = errors.New("process first node must be start")
		return
	}
	if process.Nodes[len(process.Nodes)-1].Name != "end" && !process.ExpectEnd {
		err = errors.New("process last node must be end")
		return
	}
	if process.ExpectEnd {
		for i, n := range process.Nodes {
			if n.Name == "end" {
				continue
			}
			process.Nodes[i].Edges = append(n.Edges, Edge{
				To:   "expect_end",
				Cond: "expect_end",
			})
		}
	}
	process.Nodes = append(process.Nodes, Node{
		Name: "expect_end",
	})
	if process.StartJump {
		for i, n := range process.Nodes {
			if n.Name == "start" || n.Name == "end" || n.Name == "expect_end" {
				continue
			}
			process.Nodes[i].Edges = append(n.Edges, Edge{
				To:   "start",
				Cond: "restart",
			})
		}
	}
	event = &Event{
		ProcessID:   process.UUID,
		Nodes:       process.Nodes,
		AlwaysNodes: process.AlwaysNodes,
		ExpectEnd:   process.ExpectEnd,
		StartJump:   process.StartJump,
	}
	event.UUID = uuid.New().String()
	event.NowAt = event.Nodes[0].Name
	db.Create(event)
	node = &event.Nodes[0]
	go func() {
		err = node.Run(Context{})
		if err != nil {
			log.Printf("[warning] node %s action run error: %s, event_id is %s", node.Name, err.Error(), event.UUID)
		}
	}()
	return
}

// EventTransfer 事件流转
func EventTransfer(uuid string, action string) (node *Node, err error) {
	var event Event
	db.Where("uuid = ?", uuid).First(&event)
	if event.NowAt == "end" || event.NowAt == "expect_end" {
		return nil, errors.New("event is end")
	}
	to := ""
	// 优先查找当前节点的边
	for _, n := range event.Nodes {
		if n.Name == event.NowAt {
			for _, e := range n.Edges {
				if e.Cond == action {
					to = e.To
				}
			}
		}
	}
	//// 其次查找always节点的边
	//if to == "" {
	//	for _, n := range event.AlwaysNodes {
	//		if n.Name == event.NowAt {
	//			for _, e := range n.Edges {
	//				if e.Cond == action {
	//					to = e.To
	//				}
	//			}
	//		}
	//	}
	//}
	if to != "" {
		for _, n := range event.Nodes {
			if n.Name == to {
				go func() {
					err = n.Run(Context{
						EventUUID: uuid,
						Node:      n,
					})
					if err != nil {
						log.Printf("[warning] node %s action run error: %s, event_id is %s", n.Name, err.Error(), event.UUID)
					}
				}()
				node = &n
			}
		}
		event.NowAt = to
		db.Save(&event)
		return
	}

	return nil, errors.New("no edge found")
}
