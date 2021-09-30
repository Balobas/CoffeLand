package httpApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResult(w http.ResponseWriter, result interface{}) {
	b, err := json.Marshal(result)
	if err != nil {
		if _, err := w.Write([]byte("cant marshal json")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}
