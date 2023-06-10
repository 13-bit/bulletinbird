package util

import (
	"encoding/json"
	"log"
	"os/exec"
	"path/filepath"
)

func GoEnv() map[string]string {
	cmd := exec.Command("go", "env", "-json")

	envJson, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	var envMap map[string]string

	json.Unmarshal([]byte(envJson), &envMap)

	return envMap
}

func GoModDir() string {
	goEnv := GoEnv()

	goModDir := filepath.Dir(goEnv["GOMOD"])

	return goModDir
}
