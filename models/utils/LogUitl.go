package utils

import "log"

func LogInfo(format string, v ...interface{}) {

	format = "[info] " + format
	log.Printf(format, v)
}

func LogError(format string, v ...interface{}) {

	format = "[error] " + format
	log.Printf(format, v)
}
