package socket

import (
	"bytes"
	"encoding/json"
)

func BuildJsonMessage(uri string, v interface{}) []byte {

	if v != nil {
		var data, err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
		return BuildRawMessage([]byte(uri), data)
	}
	return BuildRawMessage([]byte(uri), nil)
}

func BuildStringMessage(uri string, v string) []byte {
	return BuildRawMessage([]byte(uri), []byte(v))
}

func BuildRawMessage(uri []byte, data []byte) []byte {
	var buffer = bytes.NewBuffer(uri)

	if data != nil {
		buffer.WriteString(" ")
		buffer.Write(data)
	}

	return buffer.Bytes()
}

func BuildErrorMessage(uri string, err error) []byte {
	return BuildJsonMessage("/error", map[string]interface{}{
		"uri": uri,
		"err": err.Error(),
	})
}
