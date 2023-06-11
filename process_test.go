package flow

import (
	"testing"
)

func TestNewProcess(t *testing.T) {
	TestDb(t)
	request1 := ActionHttpRequest{
		Url:     "test1",
		Method:  "GET",
		Params:  nil,
		Headers: nil,
	}
	edges := &[]Edge{
		{
			Cond: "yes",
			To:   "test1",
		},
		{
			Cond: "no",
			To:   "test2",
		},
	}
	//nodes := []DbNode{
	//	{
	//		Name:   "test",
	//		UUID:   uuid.New().String(),
	//		Action: request1.Value(),
	//		Edges:  (*Array[Edge])(edges),
	//	},
	//	{
	//		Name:   "test1",
	//		UUID:   uuid.New().String(),
	//		Action: request1.Value(),
	//		Edges:  (*Array[Edge])(edges),
	//	},
	//	{
	//		Name:   "test2",
	//		UUID:   uuid.New().String(),
	//		Action: request1.Value(),
	//		Edges:  (*Array[Edge])(edges),
	//	},
	//}
	nodes := []Node{
		{
			Name:    "start",
			Action:  nil,
			Edges:   *edges,
			Context: Array[byte]("test"),
		},
		{
			Name:   "test1",
			Action: nil,
			Edges:  *edges,
		},
		{
			Name:    "test2",
			Action:  nil,
			Edges:   *edges,
			Context: Array[byte]("hello world"),
		},
		{
			Name:   "end",
			Action: &request1,
			Edges:  *edges,
		},
	}
	p := NewProcess("test", nodes, true, false)
	//db.Create(p)
	t.Log(p)
}

func TestGetProcess(t *testing.T) {
	TestDb(t)
	p := GetProcess("0fa8e329-aa59-46fe-8dda-321fdb2435e0")
	t.Log(p)
}
