package config

type Config struct {
	Sources  EventSources `mapstructure:"sources"`
	Logrus   Logrus       `mapstructure:"logrus"`
	LogFiles []string     `mapstructure:"logfiles"`
}

type EventSources struct {
	EnablePacketsChecking bool `mapstructure:"packets"`
	EnableAppsChecking    bool `mapstructure:"apps"`
	EnableProcChecking    bool `mapstructure:"proc"`
	EnableLogsChecking    bool `mapstructure:"logs"`
}

type Logrus struct {
	LogLvl int    `mapstructure:"log_level"`
	ToFile bool   `mapstructure:"to_file"`
	ToJson bool   `mapstructure:"to_json"`
	LogDir string `mapstructure:"log_dir"`
}
