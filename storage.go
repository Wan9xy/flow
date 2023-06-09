package flow

var storage Storage

type StorageCondition map[string]interface{}

type Storage interface {
	NewProcess(process Process) (Process, error)
	GetProcesses(cond StorageCondition) ([]Process, error)
	GetProcessWithUUID(uuid string) (Process, error)
	NewEvent(process Process) (Event, Node, error)
	GetEventWithUUID(uuid string) (Event, error)
	GetEvents(cond StorageCondition) ([]Event, error)
	GetNodesWithEventUUID(uuid string) ([]Node, error)
	GetNodeWithUUID(uuid string) (Node, error)
	TransferEvent(event Event, node Node) error
}
