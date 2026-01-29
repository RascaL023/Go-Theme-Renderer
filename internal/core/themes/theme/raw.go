package theme

import "encoding/json"

type RawTheme struct {
	Theme  json.RawMessage `json:"theme"`
	Domain map[string]json.RawMessage `json:"domain"`
	Tools  map[string]json.RawMessage `json:"tools"`
}


