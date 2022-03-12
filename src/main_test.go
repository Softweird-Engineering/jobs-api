package main

import (
	"kinza/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var conf = config.Config()

func TestHealthcheck(t *testing.T) {
	router := InitApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conf.Server.BasePath+"/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"healthy\"}", w.Body.String())
}
