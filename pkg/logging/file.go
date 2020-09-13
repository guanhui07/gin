package logging

import (
	"fmt"
	"time"

	"gin-orm/pkg/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.RuntimeRootPath, setting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSaveName,
		time.Now().Format(setting.TimeFormat),
		setting.LogFileExt,
	)
}
