package log

import (
	"io"
	syslog "log"
	"os"
)

const (
	Ldate         = syslog.Ldate
	Ltime         = syslog.Ltime
	Lmicroseconds = syslog.Lmicroseconds
	Llongfile     = syslog.Llongfile
	Lshortfile    = syslog.Lshortfile
	LUTC          = syslog.LUTC
	LstdFlags     = syslog.LstdFlags
)
const (
	LOG_LEVEL_ERROR = 9 //错误
	LOG_LEVEL_WARN  = 8 //警告
	LOG_LEVEL_INFO  = 7 //普通
	LOG_LEVEL_DEBUG = 6 //调试
)

type EntryWrite struct {
}

func (EntryWrite) Write(p []byte) (n int, err error) {
	return 0, nil
}

type Llog struct {
	L        *syslog.Logger
	objLevel int
	sysLevel int
	logFile  string
	out      io.Writer
	prefix   string
	flag     int
}

func NewLlog(logFileP string, prefixP string, flagP int, objLevelP, sysLevelP int) *Llog {
	ts := &Llog{
		objLevel: objLevelP,
		sysLevel: sysLevelP,
		logFile:  logFileP,
		prefix:   prefixP,
		flag:     flagP,
	}
	if ts.objLevel > ts.sysLevel {
		ts.out = EntryWrite{}
		ts.L = syslog.New(ts.out, "", 0)
	} else {
		if ts.logFile == "" {
			ts.out = os.Stdout
			ts.L = syslog.New(os.Stdout, "["+ts.prefix+"]", ts.flag)
		} else {
			logf, err := os.OpenFile(ts.logFile, os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				syslog.Println("open file error:", err)
				return nil
			}
			ts.out = logf
			ts.L = syslog.New(logf, "["+ts.prefix+"]", ts.flag)
		}
	}
	return ts
}
func (ts *Llog) SetlogFile(logFileP string) {
	ts.logFile = logFileP

	if ts.objLevel < ts.sysLevel {
		ts.out = EntryWrite{}
		ts.L = syslog.New(ts.out, "", 0)
	} else {
		if ts.logFile == "" {
			ts.out = os.Stdout
			ts.L = syslog.New(os.Stdout, "["+ts.prefix+"]", ts.flag)
		} else {
			logf, err := os.OpenFile(ts.logFile, os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				syslog.Println("open file error")
				return
			}
			ts.out = logf
			ts.L = syslog.New(logf, "["+ts.prefix+"]", ts.flag)
		}
	}
}
func (ts *Llog) SetLevel(objLevelP, sysLevelP int) {
	ts.objLevel = objLevelP
	ts.sysLevel = sysLevelP

	if ts.objLevel < ts.sysLevel {
		ts.out = EntryWrite{}
		ts.L = syslog.New(ts.out, "", 0)
	} else {
		if ts.logFile == "" {
			ts.out = os.Stdout
			ts.L = syslog.New(os.Stdout, "["+ts.prefix+"]", ts.flag)
		} else {
			logf, err := os.OpenFile(ts.logFile, os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				syslog.Println("open file error")
				return
			}
			ts.out = logf
			ts.L = syslog.New(logf, "["+ts.prefix+"]", ts.flag)
		}
	}
}
func (ts *Llog) Printf(format string, v ...interface{}) {
	ts.L.Printf(format, v...)
}
func (ts *Llog) Print(v ...interface{}) {
	ts.L.Print(v...)
}
func (ts *Llog) Println(v ...interface{}) {
	ts.L.Println(v...)
}

var Error *Llog
var Warn *Llog
var Info *Llog
var Debug *Llog

func init() {
	Error = NewLlog("", "Error", syslog.LstdFlags|syslog.Lshortfile, LOG_LEVEL_ERROR, LOG_LEVEL_ERROR)
	Warn = NewLlog("", "Warn", syslog.LstdFlags|syslog.Lshortfile, LOG_LEVEL_WARN, LOG_LEVEL_WARN)
	Info = NewLlog("", "Info", syslog.LstdFlags|syslog.Lshortfile, LOG_LEVEL_INFO, LOG_LEVEL_INFO)
	Debug = NewLlog("", "Debug", syslog.LstdFlags|syslog.Lshortfile, LOG_LEVEL_DEBUG, LOG_LEVEL_DEBUG)
}
