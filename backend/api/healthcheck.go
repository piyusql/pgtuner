package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type ServerStatus struct {
	appname string `json:"appname"`
	host    string `json:"hostname"`
	url     string `json:"url"`
	status  string `json:"status"`
	uptime  string `json:"uptime"`
}

var startTime time.Time

func HealthCheckInit() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func doHealthCheck(req *http.Request) ServerStatus {
	host, _ := os.Hostname()
	return ServerStatus{
		appname: "PGTuner API",
		host:    host,
		url:     req.URL.Path,
		status:  "OK",
		uptime:  uptime().String(),
	}
}
