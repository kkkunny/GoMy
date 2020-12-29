package log

import (
	"errors"
	"io"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 日志类型
type logType string

const (
	Info    logType = "info"
	Warning logType = "warning"
	Error   logType = "error"
)

// 新建一个日志管理器
func New(writers ...io.Writer) *Logger {
	return &Logger{lock: sync.Mutex{}, write: io.MultiWriter(writers...)}
}

// 日志管理器
type Logger struct {
	lock  sync.Mutex
	write io.Writer
}

// 获取当前文件名和行号
func (this *Logger) getCurFileAndLine(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = -1
	}
	return file, line
}

// 获取当前时间
func (this *Logger) getCurTime() string {
	return time.Now().Format("2006/1/2 15:04:05")
}

// 写入日志
func (this *Logger) WriteLog(msg string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, err := io.WriteString(this.write, msg)
	return err
}

// 获取日志文本
func (this *Logger) GetLogText(tp logType, msg string, time bool, file bool, skip int) string {
	var text = "[" + string(tp) + "]"
	if time {
		text += "[" + this.getCurTime() + "]"
	}
	if file {
		file, line := this.getCurFileAndLine(skip)
		text += ":{" + file + " " + strconv.Itoa(line) + "}"
	}
	text += ":" + msg + "\n"
	return text
}

// 写入信息日志
func (this *Logger) WriteInfoLog(msg string) error {
	out := this.GetLogText(Info, msg, true, true, 3)
	return this.WriteLog(out)
}

// 写入警报日志
func (this *Logger) WriteWarningLog(msg string) error {
	out := this.GetLogText(Warning, msg, true, true, 3)
	return this.WriteLog(out)
}

// 写入错误日志
func (this *Logger) WriteErrorLog(msg string) error {
	out := this.GetLogText(Error, msg, true, true, 3)
	return this.WriteLog(out)
}

// 写入错误.
func (this *Logger) WriteError(err error) error {
	if err != nil {
		out := this.GetLogText(Error, err.Error(), true, true, 4)
		return this.WriteLog(out)
	} else {
		return errors.New("error is nil")
	}
}
