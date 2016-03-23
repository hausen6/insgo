package command

import "testing"

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
	} else {
		if len(config.Targets) != len(actualConfig.Targets) {
			t.Error("'test.toml parsed failed.'")
		}
		for i := 0; i < len(config.Targets); i++ {
			// name check
			if config.Targets[i].Name != actualConfig.Targets[i].Name {
				t.Error("'test.toml parsed failed.'")
			}
			// commands check
			if len(config.Targets[i].Commands) != len(config.Targets[i].Commands) {
				t.Error("'test.toml parsed failed.'")
			} else {
				for j := 0; j < len(config.Targets[i].Commands); j++ {
					if config.Targets[i].Commands[j] != actualConfig.Targets[i].Commands[j] {
						t.Error("'test.toml parsed failed.'")
					}
				}
			}
			// depends check
			if len(config.Targets[i].Depends) != len(config.Targets[i].Depends) {
				t.Error("'test.toml parsed failed.'")
			} else {
				for j := 0; j < len(config.Targets[i].Depends); j++ {
					if config.Targets[i].Depends[j] != actualConfig.Targets[i].Depends[j] {
						t.Error("'test.toml parsed failed.'")
					}
				}
			}
		}
	}
	t.Log("success load 'test.toml'")

	// 失敗するはずのテスト
	config, err = LoadToml("not_exist.toml")
	if err == nil {
		t.Error("not_exits.toml is not exists, but error was invalid.")
	} else {
		t.Log("succeed.")
	}
}
