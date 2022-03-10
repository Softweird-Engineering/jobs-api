package utils

import (
	"net/http"
	"os"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Logger_JSON(filename string, environment string) gin.HandlerFunc {
	// Handling release/debug environments
	w_file := environment == "release"

	log.SetFormatter(&log.TextFormatter{
		DisableColors: w_file,
		FullTimestamp: true,
	})

	if w_file {
		log.SetLevel(log.InfoLevel)
		log.SetReportCaller(true)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.WithFields(log.Fields{
					"filename": filename,
			}).Info("Outputing logs to file.")
			log.SetOutput(file)
		} else {
			log.Error("Failed to log to file, using default stderr")
		}

		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	return func(c *gin.Context) {

		start := time.Now().UnixNano()
		c.Next() // Request is processed here
		stop := time.Now().UnixNano()

		fields := log.WithFields(log.Fields{
			"IP":             c.ClientIP(),
			"Method":         c.Request.Method,
			"Host":           c.Request.Host,
			"Path":           c.Request.URL.EscapedPath(),
			"Query":          c.Request.URL.RawQuery,
			"Protocol":       c.Request.Proto,
			"ContentType":    c.ContentType(),
			"ContentLength":  c.Request.ContentLength,
			"ResponseStatus": c.Writer.Status(),
			"ResponseSize":   c.Writer.Size(),
			"Headers":        c.Request.Header,
			"RequestProcessingTime": stop - start,
		})

		if c.Writer.Status() == http.StatusOK {
			fields.Info("Request")
		} else {
			fields.Warn("Request")
		}
	}
}
