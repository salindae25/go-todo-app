package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MyWeather(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["weather"] = "ee jkhello weather"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err:", err)
	} else {
		w.Write(jsonResp)
	}
}
