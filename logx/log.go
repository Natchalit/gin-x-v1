package logx

import "github.com/sirupsen/logrus"

func Debugf(format string, v ...any) {
	logrus.Debugf(format, v...)
}

func Infof(format string, v ...any) {
	logrus.Infof(format, v...)
}

func Errorf(format string, v ...any) {
	logrus.Errorf(format, v...)
}

func Warnf(format string, v ...any) {
	logrus.Warnf(format, v...)
}
