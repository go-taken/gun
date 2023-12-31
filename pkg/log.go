package pkg

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	switch level {
	case "ERROR":
		level = color.RedString("%s ", " ERROR")
	case "WARN":
		level = color.YellowString("%s ", " WARN")
	case "INFO":
		level = color.GreenString("%s ", " INFO")
	case "DEBUG":
		level = color.HiBlueString("%s ", " DEBUG")
	default:
		level = color.HiMagentaString("%s ", level)
	}
	b.WriteString(fmt.Sprintf("%s - %s - [line:%d] - %s msg ▶ %s\n",
		entry.Time.Format("2006-01-02 15:04:05,678"), entry.Caller.File,
		entry.Caller.Line, level, entry.Message))
	return b.Bytes(), nil
}

func NewLogger() *logrus.Logger {

	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(&MyFormatter{})
	return logger
}
