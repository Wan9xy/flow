package flow

import (
	"fmt"
)

type Hook struct {
	Before Action
	After  Action
	Run    Action
}

type Node struct {
	Name   string
	UUID   string
	Hook   Hook
	Action Action
	Next   []string
}

type node interface {
	Run(ctx Context) error
}

func NodeNew(hook Hook, action Action) *Node {
	n := &Node{
		Hook:   hook,
		Action: action,
	}
	return n
}

func (n *Node) Run(ctx Context) error {
	if n.Hook.Before != nil {
		if err := n.Hook.Before.Do(ctx); err != nil {
			return err
		}
	}
	if n.Hook.Run != nil {
		go func() {
			if err := n.Hook.Run.Do(ctx); err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	if err := n.Action.Do(ctx); err != nil {
		return err
	}
	if n.Hook.After != nil {
		if err := n.Hook.After.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}
