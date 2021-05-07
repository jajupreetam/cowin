package config

import (
	"fmt"
	slog "github.com/go-eden/slf4go"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

type customLogDriver struct {
	profiles         []string
	logfilePath      string
	level            string
	lumberjackLogger lumberjack.Logger
}

func (*customLogDriver) Name() string {
	return "my_driver"
}

func (d *customLogDriver) Print(l *slog.Log) {

	dateTime := formatTime(l.Time)
	var msg string
	if l.Format == nil {
		msg = fmt.Sprint(l.Args...)
	} else {
		msg = fmt.Sprintf(*l.Format, l.Args...)
	}
	sid := formatSidCid(l.Fields["sid"])
	cid := formatSidCid(l.Fields["cid"])
	logLine := fmt.Sprintf("%s %s [%s] SID=%s CID=%s [%x] %s\n", dateTime, l.Level, l.Logger, sid, cid, l.Gid, msg)

	_, l2stdout := findStringInSlice(d.profiles, "l2stdout")
	_, l2file := findStringInSlice(d.profiles, "l2file")

	if l2stdout {
		fmt.Fprintf(os.Stdout, logLine)
	}
	if l2file {
		fmt.Fprintf(&d.lumberjackLogger, logLine)
	}
}

func formatSidCid(id interface{}) string {
	if id == nil {
		return ""
	}
	return fmt.Sprintf("%v", id)
}

func formatTime(microsUnix int64) string {
	seconds := microsUnix / 1e6
	nanoseconds := microsUnix % int64(1e6) * int64(1e3)
	dateTime := time.Unix(seconds, nanoseconds).Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf("%s,%s", dateTime[0:len("2006-01-02 15:04:05")], dateTime[len("2006-01-02 15:04:05."):])
}

func (d *customLogDriver) GetLevel(logger string) (sl slog.Level) {
	switch d.level {
	case "debug":
		sl = slog.DebugLevel
	case "info":
		sl = slog.InfoLevel
	case "warn":
		sl = slog.WarnLevel
	case "error":
		sl = slog.ErrorLevel
	case "panic":
		sl = slog.PanicLevel
	case "fatal":
		sl = slog.PanicLevel
	default:
		sl = slog.TraceLevel
	}
	return
}

func InitializeLogs(profiles []string, logfilePath string, level string) {
	logPath, logPathExists := os.LookupEnv("LOG_PATH")
	if logPathExists {
		logfilePath = path.Join(logPath, logfilePath)
	}
	slog.Info("Log file path: ", logfilePath)
	d := new(customLogDriver)
	slog.SetDriver(d)
	d.lumberjackLogger = lumberjack.Logger{
		Filename:   logfilePath,
		MaxSize:    1024, // megabytes
		MaxBackups: 180,
		MaxAge:     180,  // days
		Compress:   true, // disabled by default
	}
	d.level = level
	d.profiles = profiles
}

func findStringInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
