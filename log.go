package hlog

import (
	"fmt"
	"log"
	"os"
)

var (
	errorLogger   = log.New(os.Stdout, red("ERROR: "), log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	warningLogger = log.New(os.Stdout, yellow("WARNING: "), log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	infoLogger    = log.New(os.Stdout, green("INFO: "), log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
)

func Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	_ = errorLogger.Output(2, red(msg))
}

func Warning(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	_ = warningLogger.Output(2, yellow(msg))
}

func Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	_ = infoLogger.Output(2, green(msg))
}
