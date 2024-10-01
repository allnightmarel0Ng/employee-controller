package logger

import "log"

type Level uint

const (
	Info Level = iota
	Debug
	Warning
	Trace
	Error
)

func levelToString(level Level) string {
	switch level {
	case Info:
		return "[INFO]"
	case Debug:
		return "[DEBUG]"
	case Warning:
		return "[WARNING]"
	case Trace:
		return "[TRACE]"
	case Error:
		return "[ERROR]"
	}

	return ""
}

func Printf(level Level, format string, args ...any) {
	log.Printf(levelToString(level)+" "+format, args...)
}

func Fatalf(level Level, format string, args ...any) {
	log.Fatalf(levelToString(level)+" "+format, args...)
}

func LogInfo(format string, args ...any) {
	Printf(Info, format, args...)
}

func LogDebug(format string, args ...any) {
	Printf(Debug, format, args...)
}

func LogWarning(format string, args ...any) {
	Printf(Warning, format, args...)
}

func LogTrace(format string, args ...any) {
	Printf(Trace, format, args...)
}

func LogError(format string, args ...any) {
	Fatalf(Error, format, args...)
}
