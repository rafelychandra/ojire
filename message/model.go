package message

import "time"

type Response struct {
	AccessTime time.Time   `json:"accessTime,omitempty"`
	Output     interface{} `json:"output,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Success    bool        `json:"success,omitempty"`
	Message    string      `json:"message,omitempty"`
}
