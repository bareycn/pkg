package convert

import "github.com/go-jose/go-jose/v4/json"

func StructToMap(obj interface{}) map[string]interface{} {
	var maps = make(map[string]interface{})
	marshal, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, &maps)
	if err != nil {
		return nil
	}
	return maps
}
