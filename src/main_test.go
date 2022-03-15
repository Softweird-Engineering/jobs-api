package main

import (
	"kinza/config"
	router "kinza/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var conf = config.Config()

func TestHealthcheck(t *testing.T) {
	app, err := InitApp()
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", conf.Server.BasePath+"/healthcheck", nil)
	app.ServeHTTP(w, req)

	response := &router.Healthcheck{
		Status: "healthy",
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response.String(), w.Body.String())
}
