package log

import (
	"fmt"
	"time"
)

type Logger struct {
	name  string // 日志名称
	level int    // 日志等级
}

func (l *Logger) Debug(a ...interface{}) {
	if l.level < DEBUG {
		return
	}
	fmt.Print(l.GenLogInfo(DEBUG))
	fmt.Print(a...)
}
func (l *Logger) Debugln(a ...interface{}) {
	if l.level < DEBUG {
		return
	}
	fmt.Print(l.GenLogInfo(DEBUG))
	fmt.Println(a...)
}
func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.level < DEBUG {
		return
	}
	fmt.Print(l.GenLogInfo(DEBUG))
	fmt.Printf(format, a...)
}

func (l *Logger) Info(a ...interface{}) {
	if l.level < INFO {
		return
	}
	fmt.Print(l.GenLogInfo(INFO))
	fmt.Print(a...)
}
func (l *Logger) Infoln(a ...interface{}) {
	if l.level < INFO {
		return
	}
	fmt.Print(l.GenLogInfo(INFO))
	fmt.Println(a...)
}
func (l *Logger) Infof(format string, a ...interface{}) {
	if l.level < INFO {
		return
	}
	fmt.Print(l.GenLogInfo(INFO))
	fmt.Printf(format, a...)
}

func (l *Logger) Warn(a ...interface{}) {
	if l.level < WARNNIG {
		return
	}
	fmt.Print(l.GenLogInfo(WARNNIG))
	fmt.Print(a...)
}
func (l *Logger) Warnln(a ...interface{}) {
	if l.level < WARNNIG {
		return
	}
	fmt.Print(l.GenLogInfo(WARNNIG))
	fmt.Println(a...)
}
func (l *Logger) Warnf(format string, a ...interface{}) {
	if l.level < WARNNIG {
		return
	}
	fmt.Print(l.GenLogInfo(WARNNIG))
	fmt.Printf(format, a...)
}

func (l *Logger) Error(a ...interface{}) {
	if l.level < ERROR {
		return
	}
	fmt.Print(l.GenLogInfo(ERROR))
	fmt.Print(a...)
}
func (l *Logger) Errorln(a ...interface{}) {
	if l.level < ERROR {
		return
	}
	fmt.Print(l.GenLogInfo(ERROR))
	fmt.Println(a...)
}
func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.level < ERROR {
		return
	}
	fmt.Print(l.GenLogInfo(ERROR))
	fmt.Printf(format, a...)
}

func (l *Logger) Critical(a ...interface{}) {
	if l.level < CRITICAL {
		return
	}
	fmt.Print(l.GenLogInfo(CRITICAL))
	fmt.Print(a...)
}
func (l *Logger) Criticaln(a ...interface{}) {
	if l.level < CRITICAL {
		return
	}
	fmt.Print(l.GenLogInfo(CRITICAL))
	fmt.Println(a...)
}
func (l *Logger) Criticalf(format string, a ...interface{}) {
	if l.level < CRITICAL {
		return
	}
	fmt.Print(l.GenLogInfo(CRITICAL))
	fmt.Printf(format, a...)
}

func (l *Logger) GenLogInfo(level int) string {
	var lStr, timeStr string
	switch level {
	case DEBUG:
		lStr = "DEBUG"
	case INFO:
		lStr = "INFO"
	case WARNNIG:
		lStr = "WARNNIG"
	case ERROR:
		lStr = "ERROR"
	case CRITICAL:
		lStr = "CRITICAL"
	}
	timeStr = time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s][%s]", l.name, timeStr, lStr)
}

func NewLogger(name string, level int) *Logger {
	return &Logger{
		name:  name,
		level: level,
	}
}
