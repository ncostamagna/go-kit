package cast

import (
	"encoding/json"
)

func ByteToMap[v any](data string) (map[string]v, error) {
	body := make(map[string]v)
	if err := json.Unmarshal([]byte(data), &body); err != nil {
		return nil, err
	}
	return body, nil
}

func ResponseToMap(bytes []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func ResponseToStruct[D any](data interface{}) (D, error) {
	jsonStr, err := json.Marshal(data.(map[string]interface{}))
	var d D
	if err != nil {
		return d, err
	}

	if err := json.Unmarshal(jsonStr, &d); err != nil {
		return d, err
	}

	return d, nil
}
