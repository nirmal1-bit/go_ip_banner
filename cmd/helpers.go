package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type envelope map[string]any

func writeJson(w http.ResponseWriter, data envelope, headers http.Header, status int) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value // this does like hard replace (for this exact condition)

	}

	// set are for the headers with one value standard authorization
	w.Header().Set("content-type", "application/json")

	// add for the headers value which can be multiple like set cookies accept
	// so this writes headers meaning finilizes the headers even if we don't call this if we write
	// then go auto does w.WriteHeader(http.StatusOk)
	// can be called only once
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func GetBanDuration(bancount int) time.Duration {
	switch {
	case bancount <= 3:
		return time.Minute
	case bancount <= 6:
		return 5 * time.Minute
	case bancount <= 10:
		return time.Hour
	default:
		return time.Hour * 24
	}

}
