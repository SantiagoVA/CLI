package config

import (
	"os"

	"github.com/Moldy-Community/moldy/utils/functions"
	vp "github.com/spf13/viper"
)

/* Add the default values, paths, aliases and config name and type */
var (
	defaults = map[string]interface{}{
		"adminProjects": map[string]bool{
			"gitInit":               true,
			"conventionalCommits":   true,
			"conventionalWorkflows": true,
			"semverMode":            true,
			"changelogs":            true,
		},
		"aparienceOptions": map[string]bool{
			"progressBar": true,
			"asciiArt":    true,
			"colorsMode":  true,
		},
		"moldyRunner": map[string]string{
			"test": "echo 'Running a example command'",
		},
	}
	paths = []string{
		"./",
	}
	configName = "MoldyFile.toml"
	configType = "toml"
)

/* Creating the config file */
func CreateConfigFile() {
	for _, p := range paths {
		vp.AddConfigPath(p)
	}
	vp.SetConfigName(configName)
	vp.SetConfigType(configType)
	for k, v := range defaults {
		vp.SetDefault(k, v)
	}

	if err := vp.SafeWriteConfigAs(configName); err != nil {
		if os.IsNotExist(err) {
			err = vp.WriteConfigAs(configName)
			functions.CheckErrors(err, "Code 2", "Error in write the config file :(", "Report the error on github or re try the command with new permmisions")
		}
	}
}
