package logger

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-logr/logr"
)



var _ Logger = (*logger)(nil)


type Logger = logr.Logger

type logger struct {
	enabled bool
	level   int
	kitlog log.Logger
}


const (
	DEBUG = iota +1 
	INFO
	WARN
	ERROR
)


func New(out io.Writer) Logger {
	l := log.NewJSONLogger(out)
	return &logger{
		kitlog: l,
		enabled: true,
		level: INFO,
	}
}

func (l *logger) Enabled() bool {
	return l.enabled
}


func (l *logger) V(level int) Logger {
	return &logger{

	}
}


func (l *logger) WithValues(keysAndValues ...interface{}) Logger {
	return &logger{
		
	}
}

func (l *logger) WithName(name string) Logger {
	return nil
}

func (l *logger) Error(err error, msg string, keysAndValues ...interface{}) {

}

func (l *logger) Info( msg string, keysAndValues ...interface{}) {
	
}
