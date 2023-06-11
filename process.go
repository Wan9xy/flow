package flow

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Process 流程
// 流程是由节点组成的，节点之间通过边连接
type Process struct {
	gorm.Model
	Name      string
	UUID      string
	Enable    bool
	Nodes     NodeArray
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
	db.Create(p)
	return p
}

func (p *Process) TableName() string {
	return "t_process"
}

func GetProcess(uuid string) *Process {
	p := &Process{}
	db.Debug().Where("uuid = ?", uuid).First(p)
	return p
}
