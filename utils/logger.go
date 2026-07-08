package utils

import (
	"fmt"
	"strings"
)

type Level string

const (
	INFO    Level = "INFO"
	WARN    Level = "WARN"
	ERROR   Level = "ERROR"
	SUCCESS Level = "SUCCESS"
)

func LogInfo(operation string, msg string, args ...any) {
	log(INFO, operation, msg, args...)
}

func LogSuccess(operation string, msg string, args ...any) {
	log(SUCCESS, operation, msg, args...)
}

func LogError(operation string, msg string, args ...any) {
	log(ERROR, operation, msg, args...)
}

func LogWarn(operation string, msg string, args ...any) {
	log(WARN, operation, msg, args...)
}

func log(level Level, operation string, msg string, args ...any) {
	finalMsg := formatLogMessage(msg, args...)

	fmt.Printf("%s | %s | %s | %s\n",
		NowUTC().Format("2006-01-02 15:04:05"),
		level,
		operation,
		finalMsg,
	)
}

func formatLogMessage(msg string, args ...any) string {
	result := msg

	for i, arg := range args {
		placeholder := fmt.Sprintf("{%d}", i)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprint(arg))
	}

	return result
}
