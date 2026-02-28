package tool

import "encoding/json"

type Tool interface {
	Call(args json.RawMessage) error
}
