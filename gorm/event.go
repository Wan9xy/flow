package gorm

import (
	"flow"
	"github.com/google/uuid"
	g "gorm.io/gorm"
)

type Event struct {
	g.Model
	UUID      string
	ProcessID string
	Nodes     Array[flow.Node]
	ExpectEnd bool
	StartJump bool
}

type Hook struct {
	Before Action
	After  Action
	Run    Action
}

func (e *Event) TableName() string {
	return "flow_events"
}

func (g *gorm) NewEvent(process flow.Process) (flow.Event, flow.Node, error) {
	g.Model(&Event{}).Create(&Event{
		UUID:      uuid.New().String(),
		ProcessID: process.UUID,
		Nodes:     process.Nodes,
		ExpectEnd: process.ExpectEnd,
		StartJump: process.StartJump,
	})

	g.Model(&Node{}).Create(&Node{
		UUID:      uuid.New().String(),
		EventID:   process.UUID,
		Hook: 	process.Nodes[0].Hook,
	}
}

func (g *gorm) GetEvents(cond flow.StorageCondition) ([]flow.Event, error) {

}

func (g *gorm) GetEventWithUUID(uuid string) (flow.Event, error) {

}

func (g *gorm) TransferEvent(event flow.Event, node flow.Node) error {

}
