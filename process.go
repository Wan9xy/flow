package flow

import (
	"github.com/google/uuid"
)

type Process struct {
	//gorm.Model
	Name      string
	UUID      string
	Enable    bool
	Nodes     []Node
	ExpectEnd bool
	StartJump bool
}

func NewProcess(name string, nodes []Node, expectEnd bool, startJump bool) *Process {
	p := &Process{
		Name:      name,
		Nodes:     nodes,
		ExpectEnd: expectEnd,
		StartJump: startJump,
	}
	p.UUID = uuid.New().String()
	return p
}
