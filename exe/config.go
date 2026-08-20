package exe

import (
	"encoding/json"
	"fmt"

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
}
