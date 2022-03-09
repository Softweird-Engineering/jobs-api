package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type LogItems struct {
	ISOTime        time.Time
	UnixTime       int64
	IP             string
	Method         string
	Host           string
	User           string
	Path           string
	Query          string
	Protocol       string
	ContentType    string
	ContentLength  int64
	ResponseStatus int
	ResponseSize   int
	Headers        http.Header
	TLSData        TLSData

	RequestProcessingTime int64
	LogProcessingTime     int64
}

type TLSData struct {
	TLSVersion     uint16
	TLSCipherUsed  uint16
	TLSMutualProto bool
}

var FormatJSON = func(log LogItems) string {
	logline, _ := json.Marshal(log)
	return fmt.Sprintf("%s\n", logline)
}

func Logger_JSON(filename string, w_stdout bool) gin.HandlerFunc {

	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}

	var out = io.Writer(logfile)
	if w_stdout {
		out = io.MultiWriter(logfile, os.Stdout)
	}

	return func(c *gin.Context) {

		start := time.Now()
		stop := time.Now().UnixNano()

		log := LogItems{
			ISOTime:        start,
			UnixTime:       start.UnixNano(),
			IP:             c.ClientIP(),
			Method:         c.Request.Method,
			Host:           c.Request.Host,
			User:           c.Request.URL.User.Username(),
			Path:           c.Request.URL.EscapedPath(),
			Query:          c.Request.URL.RawQuery,
			Protocol:       c.Request.Proto,
			ContentType:    c.ContentType(),
			ContentLength:  c.Request.ContentLength,
			ResponseStatus: c.Writer.Status(),
			ResponseSize:   c.Writer.Size(),
			Headers:        c.Request.Header,
		}

		if c.Request.TLS != nil {
			log.TLSData = TLSData{
				TLSVersion:     c.Request.TLS.Version,
				TLSCipherUsed:  c.Request.TLS.CipherSuite,
				TLSMutualProto: c.Request.TLS.NegotiatedProtocolIsMutual,
			}
		}

		log.RequestProcessingTime = stop - log.UnixTime

		log.LogProcessingTime = time.Now().UnixNano() - stop

		fmt.Fprint(out, FormatJSON(log))
	}
}
