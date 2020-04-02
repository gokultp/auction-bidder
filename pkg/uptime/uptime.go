package uptime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Uptime contains build details and uptime details Uptime the service
// Version,MinVersion and BuildTime are set on build
type Uptime struct {
	Version    string
	MinVersion string
	BuildTime  string
	StartedAt  time.Time
	Uptime     string
}

// NewUptime initialises with server and version details
func NewUptime(version, minVersion, buidTime string) Uptime {
	return Uptime{
		Version:    version,
		MinVersion: minVersion,
		BuildTime:  buidTime,
		StartedAt:  time.Now(),
	}
}

// Heartbeat returns details of the instance running
func (u Uptime) heartbeat() interface{} {
	uptime := time.Since(u.StartedAt)
	u.Uptime = fmt.Sprintf("%d days %s", uptime/(time.Hour*24), time.Time{}.Add(uptime).Format("15:04:05"))
	return u
}

func (u Uptime) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u.heartbeat())
}
