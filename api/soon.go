package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "soon.."
	resp["language"] = "go"
	resp["cloud"] = "vercel"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}
