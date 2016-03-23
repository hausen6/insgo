package command

import (
	"fmt"
	"testing"
)

// TestLoadToml test
func TestLoadToml(t *testing.T) {
	// 成功するはずのテスト
	actualConfig := &InstallConfig{
		[]InstallTargetConfig{
			{
				Name:     "hello",
				Depends:  []string{"echo"},
				Commands: []string{"echo hello", "echo world"},
			},
			{
				Name:     "world",
				Depends:  []string{"echo"},
				Commands: []string{"echo HELLO", "echo WORLD"},
			},
		},
	}
	config, err := LoadToml("test.toml")
	if err != nil {
		t.Error(err)
	} else if *config != *actualConfig {
		fmt.Println(config)
		fmt.Println(actualConfig)
		t.Error("failed to load toml file.")
	} else {
		t.Log("success load 'test.toml'")
	}

	// 失敗するはずのテスト
	config, err = LoadToml("not_exist.toml")
	if err == nil {
		t.Error("not_exits.toml is not exists, but error was invalid.")
	} else {
		t.Log("succeed.")
	}
}
