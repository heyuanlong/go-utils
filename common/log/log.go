package log

import (
	"os"
	syslog "log"
)

const(
	LOG_LEVEL_ERROR     = 0  //错误
	LOG_LEVEL_WARN	    = 1  //警告
	LOG_LEVEL_INFO      = 2  //普通
	LOG_LEVEL_DEBUG     = 3  //调试
)
var setLevel = LOG_LEVEL_DEBUG

type EntryWrite struct{
}
func (EntryWrite) Write(p []byte)(n int ,err error)  {
	return 0,nil
}

var Error 			*syslog.Logger
var Warn 			*syslog.Logger
var Info 			*syslog.Logger
var Debug 			*syslog.Logger

func init() {
	Error = 	syslog.New(os.Stdout,"[Debug]",syslog.LstdFlags | syslog.Lshortfile )
	Warn = 		syslog.New(os.Stdout,"[Warn]",syslog.LstdFlags | syslog.Lshortfile )
	Info = 		syslog.New(os.Stdout,"[Info]",syslog.LstdFlags | syslog.Lshortfile )
	Debug = 	syslog.New(os.Stdout,"[Debug]",syslog.LstdFlags | syslog.Lshortfile )
}

func SetLogLevel(level int)  {
	setLevel = level
	entry :=  EntryWrite{}
	if LOG_LEVEL_DEBUG > setLevel{
		Debug = syslog.New(entry,"",syslog.LstdFlags | syslog.Lshortfile )
	}
	if LOG_LEVEL_INFO > setLevel{
		Info = syslog.New(entry,"",syslog.LstdFlags | syslog.Lshortfile )
	}
	if LOG_LEVEL_WARN > setLevel{
		Warn = syslog.New(entry,"",syslog.LstdFlags | syslog.Lshortfile )
	}
	if LOG_LEVEL_ERROR > setLevel{
		Error = syslog.New(entry,"",syslog.LstdFlags | syslog.Lshortfile )
	}

}



func SetLogfile(logfile string ) error  {
	if logfile != ""{
		logFile,err  := os.OpenFile(logfile,os.O_CREATE | os.O_APPEND,0644)
		if err != nil {
			syslog.Fatalln("open file error")
		}
		if LOG_LEVEL_DEBUG <= setLevel{
			Debug = syslog.New(logFile,"[Debug]",syslog.LstdFlags | syslog.Lshortfile )
		}
		if LOG_LEVEL_INFO <= setLevel{
			Info = syslog.New(logFile,"[Info]",syslog.LstdFlags | syslog.Lshortfile )
		}
		if LOG_LEVEL_WARN <= setLevel{
			Warn = syslog.New(logFile,"[Warn]",syslog.LstdFlags | syslog.Lshortfile )
		}
		if LOG_LEVEL_ERROR <= setLevel{
			Error = syslog.New(logFile,"[Error]",syslog.LstdFlags | syslog.Lshortfile )
		}
	}
	return nil
}
