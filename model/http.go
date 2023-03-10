package model

type Response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

type StatsResponse struct {
	TotalPings           string `json:"total_pings"`
	TotalLocalPings      string `json:"total_local_pings"`
	CompletedLocalPings  string `json:"completed_local_pings"`
	ErrorLocalPings      string `json:"error_local_pings"`
	TotalRemotePings     string `json:"total_remote_pings"`
	CompletedRemotePings string `json:"completed_remote_pings"`
	ErrorRemotePings     string `json:"error_remote_pings"`
}
