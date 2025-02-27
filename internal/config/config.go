package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// TODO: move to TOML for configuration

func ReadConfig(path string) error {
	CurrentConfig.Mutex.Lock()
	defer CurrentConfig.Mutex.Unlock()

	configFile, err := os.ReadFile(os.ExpandEnv(path))
	if err != nil {
		return err
	}

	var config Config

	if err := toml.Unmarshal(configFile, &config); err != nil {
		return err
	}

	errs, fatal := ValidateConfig(&config)

	for _, v := range errs {
		log.Println(v)
	}

	if !fatal {
		CurrentConfig.Config = config
		CurrentConfig.path = path
	} else {
		return fmt.Errorf("FATAL ERRORS IN THE CONFIG WERE DETECTED! Rolling back... ")
	}

	return nil
}

func ReadConfigFromUserPath() error {
	CurrentConfig.Mutex.Lock()
	defer CurrentConfig.Mutex.Unlock()

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	userConfigDir += "/lrcsnc"

	if _, err := os.ReadDir(userConfigDir); err != nil {
		os.Mkdir(userConfigDir, 0777)
		os.Chmod(userConfigDir, 0777)
	}

	if _, err := os.Lstat(userConfigDir + "/config.toml"); err != nil {
		return err
	}

	configFile, err := os.ReadFile(userConfigDir + "/config.toml")
	if err != nil {
		return err
	}

	var config Config

	if err := toml.Unmarshal(configFile, &config); err != nil {
		return err
	}

	errs, fatal := ValidateConfig(&config)

	for _, v := range errs {
		// TODO: logger
		log.Println(v)
	}

	if !fatal {
		CurrentConfig.Config = config
		CurrentConfig.path = userConfigDir + "/config.toml"
	} else {
		return fmt.Errorf("FATAL ERRORS IN THE CONFIG WERE DETECTED! Rolling back... ")
	}

	return nil
}

func ReadDefaultConfig() error {
	CurrentConfig.Mutex.Lock()
	defer CurrentConfig.Mutex.Unlock()

	var defaultConfigPath string = "/usr/share/lrcsnc/config.toml"
	_, err := os.Stat("/usr/share/lrcsnc/config.toml")
	if err != nil {
		defaultConfigPath = "/usr/local/share/lrcsnc/config.toml"
		_, err = os.Stat("/usr/local/share/lrcsnc/config.toml")
		if err != nil {
			return fmt.Errorf("[config/ReadDefaultConfig] ERROR: Couldn't find default config. Was it deleted? Actually, was it there to begin with?")
		}
	}

	configFile, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		return err
	}

	var config Config

	if err := toml.Unmarshal(configFile, &config); err != nil {
		return err
	}

	errs, fatal := ValidateConfig(&config)

	for _, v := range errs {
		// TODO: logger
		log.Println(v)
	}

	if !fatal {
		CurrentConfig.Config = config
		CurrentConfig.path = defaultConfigPath
	} else {
		return fmt.Errorf("[config/ReadDefaultConfig] ERROR: The default config has errors in it. Was it modified?")
	}

	return nil
}

func UpdateConfig() {
	if err := ReadConfig(CurrentConfig.path); err != nil {
		// TODO: logger
		log.Println(err)
	}
}
