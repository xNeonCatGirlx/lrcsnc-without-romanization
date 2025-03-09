package log

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/types"

	"github.com/aerogo/log"
)

var logger *log.Log

func Init() {
	if !global.Config.C.Global.Log.Enabled {
		return
	}

	logger = log.New()

	// Check for the file to be writeable, otherwise panic
	if err := os.MkdirAll(path.Dir(os.ExpandEnv(global.Config.C.Global.Log.Destination)), 0777); err != nil {
		Fatal("log/Init", fmt.Sprintf("Log directory could not be created. Either disable logs before launch, or change the directory and try again.\n\nMore:\n%s", err.Error()))
	}
	f, err := os.Create(os.ExpandEnv(global.Config.C.Global.Log.Destination))
	if err != nil {
		Fatal("log/Init", fmt.Sprintf("Log file could not be created. Either disable logs before launch, or change the file path and try again.\n\nMore:\n%s", err.Error()))
	}
	err = f.Close()
	if err != nil {
		Fatal("log/Init", fmt.Sprintf("Log file could not be closed after creation. Either disable logs before launch, or change the file path and try again.\n\nMore:\n%s", err.Error()))
	}

	logger.AddWriter(log.File(os.ExpandEnv(global.Config.C.Global.Log.Destination)))
}

// Debug is a log function specifically for detailed debug information.
func Debug(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		global.Config.M.Lock()
		defer global.Config.M.Unlock()
		if !global.Config.C.Global.Log.Enabled || global.Config.C.Global.Log.Level.ToInt() < 4 {

			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " [" + modulePath + "] " + strings.ToUpper(string(types.LogLevelDebug)) + ": " + message)
	}()
}

// Info is a general-use log function for stuff like start/stop, config updates, etc.
func Info(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		global.Config.M.Lock()
		defer global.Config.M.Unlock()
		if !global.Config.C.Global.Log.Enabled || global.Config.C.Global.Log.Level.ToInt() < 3 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " [" + modulePath + "] " + strings.ToUpper(string(types.LogLevelInfo)) + ": " + message)
	}()
}

// Warn is a log function specifically for warnings.
func Warn(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		global.Config.M.Lock()
		defer global.Config.M.Unlock()
		if !global.Config.C.Global.Log.Enabled || global.Config.C.Global.Log.Level.ToInt() < 2 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Info(t + " [" + modulePath + "] " + strings.ToUpper(string(types.LogLevelWarn)) + ": " + message)
	}()
}

// Error is a log function specifically for non-fatal errors.
func Error(modulePath string, message string) {
	// Every log should be asynchronous
	go func() {
		global.Config.M.Lock()
		defer global.Config.M.Unlock()
		if !global.Config.C.Global.Log.Enabled || global.Config.C.Global.Log.Level.ToInt() < 1 {
			return
		}
		t := time.Now().Format(time.DateTime)
		logger.Error(t + " [" + modulePath + "] " + strings.ToUpper(string(types.LogLevelError)) + ": " + message)
	}()
}

// Fatal is a log function specifically for fatal errors.
// Immediately exits the app.
// Doesn't write to file: instead uses stdout to ensure the message gets to user.
func Fatal(modulePath string, message string) {
	t := time.Now().Format(time.DateTime)
	fmt.Println(t + " [" + modulePath + "] " + strings.ToUpper(string(types.LogLevelFatal)) + ": " + message)
	os.Exit(1)
}
