package logx

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)

	logrus.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp:             true,
			ForceColors:               true,
			DisableColors:             false,
			DisableQuote:              true,
			EnvironmentOverrideColors: true,
			TimestampFormat:           " 2006-01-02 15:04:05 ",
		},
	)
}
