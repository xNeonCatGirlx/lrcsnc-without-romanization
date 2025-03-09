package config

import (
	"errors"
	"os"
	"strings"

	"lrcsnc/internal/log"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"

	"github.com/pelletier/go-toml/v2"
)

func Read(path string) error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	configFile, err := os.ReadFile(os.ExpandEnv(path))
	if err != nil {
		return ErrFileUnreachable
	}

	var config structs.Config

	if err := toml.Unmarshal(configFile, &config); err != nil {
		var decodeErr *toml.DecodeError
		if errors.As(err, &decodeErr) {
			lines := strings.Join(strings.Split(decodeErr.String(), "\n"), "\n\t")
			log.Error("config/Read", "Error parsing the config file: \n\t"+lines)
			return ErrFileInvalid
		}
	}

	errs := Validate(&config)
	fatal := false
	for _, v := range errs {
		if v.Fatal {
			log.Error("config: "+v.Path, v.Message)
			fatal = true
		} else {
			log.Warn("config: "+v.Path, v.Message)
		}
	}

	if !fatal {
		global.Config.C = config
		global.Config.Path = path
	} else {
		log.Error("config/Read", "Fatal errors in the config were detected")
		return errors.New("fatal validation errors")
	}

	return nil
}

func ReadUserWide() error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	userConfigDir += "/lrcsnc"

	if _, err := os.Stat(userConfigDir + "/config.toml"); err != nil {
		return errors.New("user config file doesn't exist")
	}

	return Read(userConfigDir + "/config.toml")
}

func ReadSystemWide() error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	sysWideConfigPath := "/etc/lrcsnc/config.toml"
	_, err := os.Stat(sysWideConfigPath)
	if err != nil {
		log.Error("config/ReadSystemWide", "The system-wide config doesn't exist")
	}

	return Read(sysWideConfigPath)
}

func Update() {
	if err := Read(global.Config.Path); err != nil {
		switch {
		case errors.Is(err, ErrFileUnreachable):
			log.Error("config/Update", "The config file is now unreachable. The configuration will remain the same")
		default:
			log.Error("config/Update", "Unknown error: "+err.Error())
		}
	}
}
