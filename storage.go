package flow

// Storage 存储接口
// TODO: 未来可以抽取为支持多种储存方式，除gorm外，还可以支持redis、mongo等
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
