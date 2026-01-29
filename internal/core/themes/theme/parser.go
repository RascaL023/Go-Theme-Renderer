package theme

import "encoding/json"

func Parse(raw json.RawMessage) (r ResolvedTheme, err error){
	err = json.Unmarshal(raw, &r);
	return r, err;
}
