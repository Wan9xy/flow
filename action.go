package flow

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

var _allowedActions = make(map[string]Action)

type Action interface {
	Do(ctx Context) error
	ActionName(ctx Context) string
	//Next(ctx context.Context) Action
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
	return nil
}

func (a *ActionHttpRequest) ActionName(ctx Context) string {
	return "httpRequest"
}

func (a *ActionHttpRequest) Next(ctx Context) Action {
	return nil
}

func (a *ActionHttpRequest) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion failed")
	}
	return json.Unmarshal(b, &a)
}

func (a *ActionHttpRequest) Value() (driver.Value, error) {
	return json.Marshal(a)
}
