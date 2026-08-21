package exe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gerdreiss/mgit/config"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

func PrintConfig(cmd *cobra.Command, args []string) {
	displayJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		fmt.Printf("❌ 'json' flag of indeterminate value - using default value: %v\n", err)
		displayJson = false
	}

	appConfig := config.GetAppConfig()

	if len(args) == 0 {
		if displayJson {
			prettyJSON, err := json.MarshalIndent(appConfig, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling config: %v\n", err)
				return
			}
			fmt.Println(string(prettyJSON))
		} else {
			yamlData, err := yaml.Marshal(appConfig)
			if err != nil {
				fmt.Printf("Error marshaling config to YAML: %v\n", err)
				return
			}
			fmt.Println(string(yamlData))
		}
	} else {
		for _, key := range args {
			value, err := config.Get(key, displayJson)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}

			fmt.Println(value)
		}
	}
}

func WriteConfig(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		err := config.Set(parts[0], parts[1])
		if err != nil {
			fmt.Printf("Error setting value %s for key %s: %v\n", parts[1], parts[0], err)
		}

	}
}
