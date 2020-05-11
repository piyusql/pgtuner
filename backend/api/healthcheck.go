package api

import (
	"net/http"
	"os"
	"time"
)

// ServerStatus :: Data which represents the health status of the system
type ServerStatus struct {
	appname string `json:"appname"`
	host    string `json:"hostname"`
	url     string `json:"url"`
	status  string `json:"status"`
	uptime  string `json:"uptime"`
}

var startTime time.Time

// HealthCheckInit :: This method will be called while the application will start
// This will keep track of the time service is UP since then.
func HealthCheckInit() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
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
