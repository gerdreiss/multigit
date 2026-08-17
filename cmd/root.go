/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"log"
	"os"

	"github.com/gerdreiss/mgit/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mgit",
	Short: "mgit: A GIT helper to work with multiple git repositories.",
	Long:  "mgit is a GIT helper for executing a subset of git commands on multiple repositories.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/mgit/")
	viper.AddConfigPath("/usr/local/etc/mgit/")
	viper.AddConfigPath("$HOME/.mgit/")
	viper.AddConfigPath(".")

	viper.SetDefault("git.remote-name", "origin")
	viper.SetDefault("git.auth-method", "none")

	_ = viper.ReadInConfig()

	var conf config.AppConfig
	// Unmarshal the config file into the AppConfig struct
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	config.SetAppConfig(conf)
}
