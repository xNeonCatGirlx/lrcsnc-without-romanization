package log

import (
	"fmt"
	"lrcsnc/internal/config"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aerogo/log"
)

var logger *log.Log

func Init() {
	if !config.CurrentConfig.Config.Global.Log.Enabled {
		return
	}

	logger = log.New()

	// Check for the file to be writeable, otherwise panic
	if err := os.MkdirAll(path.Dir(os.ExpandEnv(config.CurrentConfig.Config.Global.Log.Destination)), 0777); err != nil {
		panic(fmt.Errorf("[log/Init] FATAL: Log directory could not be created. Either disable logs before launch, or change the directory and try again.\n\nMore:\n%s", err.Error()))
	}
	f, err := os.Create(os.ExpandEnv(config.CurrentConfig.Config.Global.Log.Destination))
	if err != nil {
		panic(fmt.Errorf("[log/Init] FATAL: Log file could not be created. Either disable logs before launch, or change the file path and try again.\n\nMore:\n%s", err.Error()))
	}
	f.Close()

	logger.AddWriter(log.File(os.ExpandEnv(config.CurrentConfig.Config.Global.Log.Destination)))
}

// A log function specifically for detailed debug information.
func Debug(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		config.CurrentConfig.Mutex.Lock()
		defer config.CurrentConfig.Mutex.Unlock()
		if config.CurrentConfig.Config.Global.Log.Level.ToInt() < 4 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " " + modulePath + " " + strings.ToUpper(string(config.LogLevelDebug)) + ": " + message)
	}()
}

// A general-use log function for stuff like start/stop, config updates, etc.
func Info(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		config.CurrentConfig.Mutex.Lock()
		defer config.CurrentConfig.Mutex.Unlock()
		if config.CurrentConfig.Config.Global.Log.Level.ToInt() < 3 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " " + modulePath + " " + strings.ToUpper(string(config.LogLevelInfo)) + ": " + message)
	}()
}

// A log function specifically for warnings.
func Warn(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		config.CurrentConfig.Mutex.Lock()
		defer config.CurrentConfig.Mutex.Unlock()
		if config.CurrentConfig.Config.Global.Log.Level.ToInt() < 2 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " " + modulePath + " " + strings.ToUpper(string(config.LogLevelWarn)) + ": " + message)
	}()
}

// A log function specifically for non-fatal errors.
func Error(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		config.CurrentConfig.Mutex.Lock()
		defer config.CurrentConfig.Mutex.Unlock()
		if config.CurrentConfig.Config.Global.Log.Level.ToInt() < 1 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Error(t + " " + modulePath + " " + strings.ToUpper(string(config.LogLevelError)) + ": " + message)
	}()
}

// A log function specifically for fatal errors. 
// Immediately exits the app.
// Doesn't write to file: instead uses stdout to ensure the message gets to user.
func Fatal(modulePath string, message string) {
	t := time.Now().Format(time.DateTime)
	fmt.Println(t + " " + modulePath + " " + strings.ToUpper(string(config.LogLevelFatal)) + ": " + message)
	os.Exit(1)
}
