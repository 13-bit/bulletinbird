package util

import "os"

func CopyFile(src string, dst string) {
	bytes, err := os.ReadFile(src)
	CheckError(err)

	err = os.WriteFile(dst, bytes, os.FileMode(0755))
	CheckError(err)
}
