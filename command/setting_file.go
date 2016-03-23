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

// LoadToml is a load setting file (written by toml) function.
func LoadToml(filename string) (*InstallConfig, error) {
	var config InstallConfig
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return &config, err
	}
	return &config, err
}
