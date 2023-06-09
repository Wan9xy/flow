package flow

import (
	"github.com/google/uuid"
)

type Event struct {
	//gorm.Model
	UUID      string
	ProcessID string
	Nodes     []Node
	ExpectEnd bool
	StartJump bool
}

func EventStart(process Process) (event *Event, node *Node, error error) {
	e := &Event{
		ProcessID: process.UUID,
		Nodes:     process.Nodes,
		ExpectEnd: process.ExpectEnd,
		StartJump: process.StartJump,
	}
	e.UUID = uuid.New().String()
	ctx := Context{
		Event: e,
	}
	node = NodeNew(e.Nodes[0].Hook, e.Nodes[0].Action)
	err := node.Run(ctx)
	if err != nil {
		return nil, nil, err
	}
	return e, node, nil
}
