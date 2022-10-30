package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/s-vvardenfell/psyllium/config"
	core "github.com/s-vvardenfell/psyllium/internal"
	"github.com/s-vvardenfell/psyllium/internal/linux_agent"
	"github.com/s-vvardenfell/psyllium/internal/windows_agent"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cnfg    config.Config
)

var rootCmd = &cobra.Command{
	Use:   "agent",
	Short: "The agent performs initial processing and filtering, as well as collecting security events.",
	Run: func(cmd *cobra.Command, args []string) {
		// d, err := cmd.Flags().GetBool("debug")
		// cobra.CheckErr(err)
		// ss := somepackage.New(&cnfg, d)
		// ss.Work()

		fmt.Println(cnfg)

		fmt.Println("Agent works!")

		os := runtime.GOOS

		c := core.NewCore()
		var a core.Agent

		if os == "linux" {
			a = linux_agent.NewLinuxAgent()
		} else if os == "windows" {
			a = windows_agent.NewWindowsAgent()
		}

		// тоже сделать для ключей-фронта: запуск фронта с ядром

		if err := c.WithAgent(a).Start(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/config.yml)")
	rootCmd.Flags().BoolP("debug", "d", false, "Runs in debug-mode")
}

func initConfig() {
	ConfigureViper()
	ConfigureLogrus()
}

func ConfigureViper() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(wd, "config"))
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		cobra.CheckErr(err)
	}

	if err := viper.Unmarshal(&cnfg); err != nil {
		cobra.CheckErr(err)
	}
}

func ConfigureLogrus() {
	if cnfg.Logrus.ToFile {
		if err := os.Mkdir(filepath.Dir(cnfg.Logrus.LogDir), 0777); err != nil && !errors.Is(err, os.ErrExist) {
			cobra.CheckErr(err)
		}

		file, err := os.OpenFile(cnfg.Logrus.LogDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			cobra.CheckErr(err)
		}
	}

	if cnfg.Logrus.ToJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetLevel(logrus.Level(cnfg.Logrus.LogLvl))
}
