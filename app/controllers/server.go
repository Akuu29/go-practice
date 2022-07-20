package controllers

import (
	"net/http"

	"github.com/Akuu29/go-practice/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
