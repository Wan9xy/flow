package gorm

import (
	"context"
	"encoding/json"
	"flow"
	g "gorm.io/gorm"
)

type Action struct {
	Name string
	Meta Array
}

type Node struct {
	g.Model
	UUID      string
	EventID string
	Hook      string
	Action string
	Next   Array[string]
}

func parseNode(node flow.Node) *Node {
	n := &Node{
		UUID:   node.UUID,
		Hook:   node.Hook,
		Action: node.Action,
		Next: node.Next,
	}
	return n
}

func parseHook(hook flow.Hook) *Hook {
	var beforeMeta,afterMeta,runMeta []byte
	beforeMeta,_ = json.Marshal(hook.Before)
	afterMeta,_ = json.Marshal(hook.After)
	runMeta,_ = json.Marshal(hook.Run)
	var b,a,r Array

	h := &Hook{
		Before: Action{
			Name: hook.Before.ActionName(flow.NilContext),
			Meta: ,
		},
		After:  hook.After,
		Run:    hook.Run,
	}
	return h
}

func (g *gorm) GetNodesWithEventUUID(uuid string) ([]flow.Node, error) {

}

func (g *gorm) GetNodeWithUUID(uuid string) (flow.Node, error) {

}
