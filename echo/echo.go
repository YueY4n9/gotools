package echo

import (
	"encoding/json"
	"fmt"
)

func Json(a any) {
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
}

func Marshal(a any) []byte {
	bytes, _ := json.Marshal(a)
	return bytes
}
