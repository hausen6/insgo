package command

import "github.com/BurntSushi/toml"

// InstallConfig is a all settings definition
type InstallConfig struct {
	Targets []InstallTargetConfig `toml:"targets"`
}

// InstallTargetConfig is a
type InstallTargetConfig struct {
	Name     string   `toml:"name"`
	Commands []string `toml:"commands"`
	Depends  []string `toml:"depends"`
}

// loadToml is a load setting file (written by toml) function.
func loadToml(filename string) (*InstallConfig, error) {
	var config InstallConfig
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return &config, err
	}
	return &config, err
}

// parseDepends is a function.
func parseDependTree(config *InstallConfig) {
	// str, err := osext.Executable()
	// fmt.Println(str, err)
}
