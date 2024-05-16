package sh

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gvcgo/version-manager/pkgs/conf"
)

const vmEnvFish = `# Code generated by vmr. DO NOT EDIT.
function _vmr_cdhook --on-variable PWD --description "version manager cd hook"
	if type -q vmr
        vmr use -E
	end
end

fish_add_path --global %s %s/bin
`

type FishShell struct{}

func NewFishShell() *FishShell {
	return &FishShell{}
}

func (f *FishShell) ConfPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config/fish/config.fish")
}

func (f *FishShell) VMEnvConfPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, fmt.Sprintf(".config/fish/conf.d/%s.fish", vmEnvFileName))
}

func (f *FishShell) WriteVMEnvToShell() {
	installPath := conf.GetVersionManagerWorkDir()
	content := fmt.Sprintf(vmEnvFish, installPath, installPath)
	_ = os.WriteFile(f.VMEnvConfPath(), []byte(content), ModePerm)
}

func (f *FishShell) PackPath(path string) string {
	return fmt.Sprintf("fish_add_path --global %s", path)
}

func (f *FishShell) PackEnv(key, value string) string {
	if value == "" {
		return fmt.Sprintf("set --global %s ", key)
	}
	return fmt.Sprintf("set --global %s %s", key, value)
}
