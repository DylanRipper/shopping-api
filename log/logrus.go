package log

import (
	"shopping-api/pkg/util"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	maxAgeStr := util.GetEnv("LOG_MAX_AGE", "7")
	maxAge, _ := strconv.Atoi(maxAgeStr)
	rotationStr := util.GetEnv("LOG_ROTATION_TIME", "1")
	rotation, _ := strconv.Atoi(rotationStr)
	writer, err := rotatelogs.New(
		"log/logs/log_shopping_api.log"+"%Y%m%d",
		rotatelogs.WithLinkName(util.GetEnv("LOG_PATH", "log/logs/log_shopping_api.log")),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(rotation*24)*time.Hour),
	)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	customFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	}
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
		},
		customFormatter,
	))

}
