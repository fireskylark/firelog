package firelog

import (
	"path"
	"strconv"
	"strings"
	"time"
)

// Record defines a log message event.
type Record struct {
	Time    string  `json:"timesatmp,omitempty"` //时间
	Level   uint8   `json:"level,omitempty"`     //日志级别
	Module  string  `json:"module,omitempty"`    //日志钩子(loghandler),默认是"__ROOT__"
	FuncPtr uintptr `json:"funcPtr,omitempty"`   //函数名
	File    string  `json:"file,omitempty"`      //文件名
	Line    int     `json:"line,omitempty"`      //日志行号
	Msg     string  `json:"message,omitempty"`
}

// Format formats LogRecord to a string as required format.
// TIME: time field
// LEVEL: log message level
// MODULE: log handler, if used. Default is __default__
// FUNCNAME: log function name
// PATH: log file path
// FILE: log filename
// LINE: log line
// MESSAGE: log message

func (rcd *Record) Format(str string) string {
	str = strings.Replace(str, "TIME", time.Now().Format("2006-01-02 15:04:05"), -1)
	str = strings.Replace(str, "LEVEL", LevelMap[rcd.Level], -1)
	str = strings.Replace(str, "MODULE", rcd.Module, -1)
	str = strings.Replace(str, "FUNCNAME", FuncName(rcd.FuncPtr), -1)
	str = strings.Replace(str, "PATH", path.Dir(rcd.File), -1)
	str = strings.Replace(str, "FILE", path.Base(rcd.File), -1)
	str = strings.Replace(str, "LINE", strconv.Itoa(rcd.Line), -1)
	str = strings.Replace(str, "MESSAGE", rcd.Msg, -1)
	return str
}
