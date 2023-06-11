package flow

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Hook 节点钩子
// TODO：暂时未用到
type Hook struct {
	Before Action
	After  Action
	Run    Action
}

// Node 流程节点
// 节点是流程的基本组成单元
// 节点包含一个动作,一个自定义的上下文和若干条边
// 其中Action是节点运行的异步动作，可以是任意一个实现了Action接口的函数，请注意，动作是异步执行的，所以执行动作是否发生意外，并不会影响流程的执行
// Context是节点的上下文，可以在到达当前节点后，将上下文中的数据传递给流程相关函数
// 节点之间通过边连接，边包含一个条件和一个目标节点
type Node struct {
	Name string
	//Hook   *Hook
	Action  Action
	Edges   Array[Edge]
	Context Array[byte]
}

type NodeArray []Node

func (na *NodeArray) Scan(value interface{}) error {
	var dbNodes []DbNode
	err := json.Unmarshal(value.([]byte), &dbNodes)
	var nodes []Node
	for i, _ := range dbNodes {
		nodes = append(nodes, dbNodes[i].toNode())
	}
	*na = nodes
	return err
}

func (na NodeArray) Value() (driver.Value, error) {
	var dbNodes []DbNode
	for i, _ := range na {
		dbNodes = append(dbNodes, na[i].toDbNode())
	}
	return json.Marshal(dbNodes)
}

func (n *Node) toDbNode() DbNode {
	var av *ActionWrapper
	if n.Action != nil {
		a := n.Action.Value()
		av = &a
	}
	return DbNode{
		Name:    n.Name,
		Action:  av,
		Edges:   &n.Edges,
		Context: &n.Context,
	}
}

type DbNode struct {
	Name    string
	Action  *ActionWrapper
	Edges   *Array[Edge]
	Context *Array[byte]
}

func (n *DbNode) toNode() Node {
	var act Action
	var con Array[byte]
	if n.Action != nil {
		if a, ok := _allowedActions[n.Action.ActionName]; ok {
			act = a
			act = act.Scan(n.Action.Meta)
		} else {
			panic(fmt.Sprintf("action %s not found", n.Action.ActionName))
		}
	}
	if n.Context != nil {
		con = *n.Context
	}
	return Node{
		Name:    n.Name,
		Action:  act,
		Edges:   *n.Edges,
		Context: con,
	}
}

type Edge struct {
	Cond string
	To   string
}

func (n *Node) Run(ctx Context) error {
	//if n.Hook.Before != nil {
	//	if err := n.Hook.Before.Do(ctx); err != nil {
	//		return err
	//	}
	//}
	//if n.Hook.Run != nil {
	//	go func() {
	//		if err := n.Hook.Run.Do(ctx); err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//	}()
	//}
	if n.Action == nil {
		return nil
	}
	if err := n.Action.Do(ctx); err != nil {
		return err
	}
	//if n.Hook.After != nil {
	//	if err := n.Hook.After.Do(ctx); err != nil {
	//		return err
	//	}
	//}
	return nil
}
