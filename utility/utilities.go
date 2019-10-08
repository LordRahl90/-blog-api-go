package utility

import (
	"encoding/json"
)

//APIResponse - This manages the APIResponse. This is in utility because its meant to be used across the application
func APIResponse(code int, success bool, message string, data interface{}) ([]byte, error) {
	r := make(map[string]interface{})
	r["code"] = code
	r["success"] = success
	r["message"] = message
	r["data"] = data

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}
