package flow

import "context"

type Context struct {
	context.Context
	//Event *Event
	Node      Node
	EventUUID string
}

var NilContext = Context{}
