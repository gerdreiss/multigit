package exe

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

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
	settings := viper.AllSettings()

	displayJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		fmt.Printf("❌ 'json' flag of indeterminate value - using default value: %v\n", err)
		displayJson = false
	}

	if displayJson {
		prettyJSON, err := json.MarshalIndent(settings, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling config: %v\n", err)
			return
		}

		// 3. Print the result
		fmt.Println(string(prettyJSON))
	} else {
		yamlData, err := yaml.Marshal(settings)
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
			fmt.Fprintf(os.Stderr, "key %s not found. make sure you enter the whole key, e.g. git.remote-name\n", arg)
		}
	}
}

func WriteConfig(cmd *cobra.Command, args []string) {
	fmt.Printf("Writing configuration from args: %v\n", args)
	fmt.Fprintf(os.Stderr, "not yet implemented")
}
