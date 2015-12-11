package atomas
import "encoding/json"

func ToJsonString(any interface{}) string {
	json, _ := json.Marshal(any)
	return string(json)
}