package flow

import (
	"fmt"
	"testing"
)

func TestEventStart(t *testing.T) {
	TestDb(t)
	process := GetProcess("3e283e05-5383-475f-a933-d4af007ccd65")
	start, n, err := EventStart(*process)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(start)
	t.Log(n)
}

func TestEventTransfer(t *testing.T) {
	TestDb(t)
	t.Log(EventTransfer("5aedb887-ef1b-4560-9f80-b8229f05e93f", "no2"))
}
