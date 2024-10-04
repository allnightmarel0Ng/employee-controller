package logger

import "log"

func Info(format string, args ...any) {
	log.Printf("[INFO] "+format+"\n", args...)
}

func Debug(format string, args ...any) {
	log.Printf("[DEBUG] "+format+"\n", args...)
}

func Warning(format string, args ...any) {
	log.Printf("[WARNING] "+format+"\n", args...)
}

func Trace(format string, args ...any) {
	log.Printf("[TRACE] "+format+"\n", args...)
}

func Error(format string, args ...any) {
	log.Fatalf("[ERROR] "+format+"\n", args...)
}
