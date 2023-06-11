package flow

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func init() {
	RegisterAction("http_request", &ActionHttpRequest{})
	fmt.Println("init")
}

var _allowedActions = make(map[string]Action)

// Action 节点动作接口
// Action是节点运行的动作，每个节点都有一个动作，动作可以是任意一个实现了本接口的函数，也可以是一个http请求，例如下方的默认实现
// 动作执行时会传入一个上下文，上下文包含了当前事件的信息，以及当前节点的信息
type Action interface {
	// Do 执行动作
	Do(ctx Context) error
	// ActionName 返回动作名称
	ActionName(ctx Context) string
	// Scan 从数据库中读取动作时调用，用于将数据库中的数据转换为动作
	Scan(meta []byte) Action
	// Value 将动作转换为数据库中的数据
	Value() ActionWrapper
	//Next(ctx context.Context) Action
}

type ActionWrapper struct {
	ActionName string
	Meta       Array[byte]
}

func RegisterAction(name string, action Action) {
	_allowedActions[name] = action
}

type ActionHttpRequest struct {
	Url     string
	Method  string
	Params  map[string]interface{}
	Headers map[string]string
}

func (a *ActionHttpRequest) Do(ctx Context) error {
	r := resty.New().R().SetHeaders(a.Headers).SetHeader("Content-Type", "application/json").
		SetHeader("from", "flow").SetHeader("event-id", ctx.EventUUID).SetHeader("now-node", ctx.Node.Name)
	switch a.Method {
	case "GET":
		for s, i := range a.Params {
			r.SetQueryParam(s, fmt.Sprint(i))
		}
		get, err := r.Get(a.Url)
		if err != nil {
			return err
		}
		if get.StatusCode() != 200 {
			return errors.New(fmt.Sprint("http request error:", get.StatusCode()))
		}
	case "POST":
		post, err := r.SetBody(a.Params).Post(a.Url)
		if err != nil {
			return err
		}
		if post.StatusCode() != 200 {
			return errors.New(fmt.Sprint("http request error:", post.StatusCode()))
		}
	}
	return nil
}

func (a *ActionHttpRequest) ActionName(ctx Context) string {
	return "http_request"
}

func (a *ActionHttpRequest) Scan(meta []byte) Action {
	_ = json.Unmarshal(meta, a)
	return a
}

func (a *ActionHttpRequest) Value() ActionWrapper {
	b, _ := json.Marshal(a)
	return ActionWrapper{
		ActionName: "http_request",
		Meta:       b,
	}
}

func (a *ActionWrapper) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, a)
	}
	return nil
}

func (a *ActionWrapper) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	return json.Marshal(a)
}
