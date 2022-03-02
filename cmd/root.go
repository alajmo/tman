package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/alajmo/tman/core/dao"
)

const (
	appName      = "tman"
	shortAppDesc = "tman is a cli tool used to manage terminal sessions"
	longAppDesc  = `tman is a cli tool used to manage terminal sessions`
)

var (
	config     dao.Config
	configErr  error
	configFile string
	userConfigDir string
	rootCmd    = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (by default it checks current and all parent directories for tman.yaml|yml)")

	defaultUserConfigDir, _ := os.UserConfigDir()
	defaultUserConfigDir = filepath.Join(defaultUserConfigDir, "tman")

	rootCmd.PersistentFlags().StringVar(&userConfigDir, "user-config-dir", defaultUserConfigDir, "user config directory to automatically")

	rootCmd.AddCommand(
		versionCmd(),
		completionCmd(),
		runCmd(&config, &configErr),
	)
}

func initConfig() {
	config, configErr = dao.ReadConfig(configFile, userConfigDir)
}
