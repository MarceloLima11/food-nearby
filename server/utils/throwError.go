package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func IfErrThrowFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IfErrThrowFatalf(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func IfErrThrowPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func IfErrThrowPanicf(err error, message string) {
	if err != nil {
		log.Panicf(message, err)
	}
}

func IfErrThrowWriteError(err error, w http.ResponseWriter, errorMessage string, statusCode int) bool {
	if err != nil {
		w.WriteHeader(statusCode)
		response := map[string]string{"error": errorMessage}
		json.NewEncoder(w).Encode(response)

		return true
	}

	return false
}
