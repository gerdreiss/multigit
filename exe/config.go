package exe

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/gerdreiss/mgit/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

func PrintConfig(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		printAllSettings(cmd)
	} else {
		printSelectedSettings(args)
	}
}

func printAllSettings(cmd *cobra.Command) {
	appConfig := config.GetAppConfig()

	displayJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		fmt.Printf("❌ 'json' flag of indeterminate value - using default value: %v\n", err)
		displayJson = false
	}

	if displayJson {
		prettyJSON, err := json.MarshalIndent(appConfig, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling config: %v\n", err)
			return
		}

		// 3. Print the result
		fmt.Println(string(prettyJSON))
	} else {
		yamlData, err := yaml.Marshal(appConfig)
		if err != nil {
			fmt.Printf("Error marshaling config to YAML: %v\n", err)
			return
		}
		fmt.Println(string(yamlData))
	}
}

func printSelectedSettings(args []string) {
	for _, arg := range args {
		if slices.Contains(viper.AllKeys(), arg) {
			fmt.Printf("%s = %s\n", arg, viper.GetString(arg))
		} else {
			fmt.Fprintf(os.Stderr, "key %s not found. make sure you enter the complete key, e.g. git.remote-name\n", arg)
		}
	}
}

func WriteConfig(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		viper.Set(parts[0], parts[1])
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to determine user home directory: %v\n\n", err)
		return
	}
	if _, err := os.Stat(homedir + "/.mgit"); err != nil {
		if err := os.Mkdir(homedir+"/.mgit", 0755); err != nil {
			fmt.Fprintf(os.Stderr, "❌ failed to create .mgit directory: %v\n", err)
			return
		}
	}
	file, err := os.OpenFile(homedir+"/.mgit/config", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to open or create config file: %v\n", err)
		return
	}
	defer file.Close()
	if err := viper.WriteConfigTo(file); err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to write config file: %v\n", err)
		return
	}
	fmt.Println("✅ successfully written configuration to ~/.mgit/config")
}
