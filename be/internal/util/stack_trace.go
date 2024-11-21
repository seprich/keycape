package util

import "runtime"

func CaptureStackTrace() string {
	buffer := make([]byte, 2048)
	n := runtime.Stack(buffer, false)
	return string(buffer[:n])
}
