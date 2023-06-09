package flow

import "context"

type Context struct {
	context.Context
	Event *Event
}

var NilContext = Context{}
