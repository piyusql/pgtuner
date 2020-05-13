package api

import (
	"net/http"
	"os"
	"time"
)

// ServerStatus :: Data which represents the health status of the system
type ServerStatus struct {
	AppName  string `json:"appname"`
	HostName string `json:"hostname"`
	URL      string `json:"url"`
	Status   string `json:"status"`
	UpTime   string `json:"uptime"`
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
		AppName:  "PGTuner API",
		HostName: host,
		URL:      req.URL.Path,
		Status:   "OK",
		UpTime:   uptime().String(),
	}
}
