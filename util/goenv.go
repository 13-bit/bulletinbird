package util

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
)

func GoEnv() map[string]string {
	cmd := exec.Command("go", "env", "-json")

	envJson, err := cmd.Output()

	CheckError(err)

	var envMap map[string]string

	json.Unmarshal([]byte(envJson), &envMap)

	return envMap
}

func GoModDir() string {
	goEnv := GoEnv()

	goModDir := filepath.Dir(goEnv["GOMOD"])

	return goModDir
}
