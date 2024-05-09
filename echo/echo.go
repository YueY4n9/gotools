package echo

import (
	"encoding/json"
	"fmt"
)

func Json(a any) {
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
}
