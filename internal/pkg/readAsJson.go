package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ReadAsJson(result []*map[string]interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for j, val := range result {
		item, e := json.MarshalIndent(val, "", " ")
		if e == nil {
			buffer.WriteString(string(item))
		}
		if j < len(result)-1 {
			buffer.WriteString(",")
		}
		fmt.Println(string(item))
	}

	buffer.WriteString("]")

	return buffer.Bytes(), nil
}
