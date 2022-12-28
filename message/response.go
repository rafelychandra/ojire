package message

import "time"

func RenderResponse(data interface{}, statusCode int, success bool, message string) interface{} {
	outputData := Response{
		AccessTime: time.Now(),
		Output:     data,
		StatusCode: statusCode,
		Success:    success,
		Message:    message,
	}
	return outputData
}
