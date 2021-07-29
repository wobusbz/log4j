package log4j

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

var DefaultLog4j = New(DEBUG)

type Log4j struct {
	level 	 Level 		// 日志等级
	fileName string 	// 日志文件名称
}

func New(level Level) *Log4j{
	return &Log4j{level: level}
}


// @param format string 格式化字符
// @param args  ...interface{} 待格式化字符串
func (l *Log4j) output(format string, args ...interface{}) {
	var content string
	if len(args) != 0 {
		content = fmt.Sprintf(format, args...)
	}
	var buff = &bytes.Buffer{}
	buff.WriteString(GetYYYYMMDDSS00())
	buff.WriteString("\t")
	buff.WriteString(l.GetFileName())
	buff.WriteString("\t")
	buff.WriteString(fmt.Sprintf("%d", GetGoroutines()))
	buff.WriteString("\t")
	buff.WriteString(l.level.String())
	buff.WriteString("\t")
	var file, line = GetCalle()
	buff.WriteString(path.Base(file))
	buff.WriteString(":")
	buff.WriteString(fmt.Sprintf("%d", line))
	buff.WriteString("\t")
	buff.WriteString(content)
	buff.WriteString("\n")
	os.Stdout.WriteString(buff.String())
}


func (l *Log4j) Debug(format string, args ...interface{}) {
	if l.isLevel(DEBUG){
		l.output(format, args...)
	}
}

func (l *Log4j) Info(format string, args ...interface{}) {
	if l.isLevel(INFO){
		l.output(format, args...)
	}
}

func (l *Log4j) Warin(format string, args ...interface{}) {
	if l.isLevel(WARN){
		l.output(format, args...)
	}
}

func (l *Log4j) Error(format string, args ...interface{}) {
	if l.isLevel(ERROR){
		l.output(format, args...)
	}
}

func (l *Log4j) Fatal(format string, args ...interface{}) {
	if l.isLevel(FATAL){
		l.output(format, args...)
	}
	os.Exit(2)
}

func (l *Log4j) isLevel(level Level) bool {
	return l.level <= level 
}

func (l *Log4j) GetFileName() string{
	if l.fileName == "" {
		return "????"
	}
	return l.fileName
}


func Debug(format string, args ...interface{}){
	DefaultLog4j.Debug(format, args...)
}

func Info(format string, args ...interface{}){
	DefaultLog4j.Info(format, args...)
}

func Warin(format string, args ...interface{}){
	DefaultLog4j.Warin(format, args...)
}

func Error(format string, args ...interface{}){
	DefaultLog4j.Error(format, args...)
}

func Fatal(format string, args ...interface{}){
	DefaultLog4j.Fatal(format, args...)
}