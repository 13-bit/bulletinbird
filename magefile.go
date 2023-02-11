//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	if err := sh.Run("go", "build", "-o", "build/", "./cmd/bulletinbird"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", "build/", "./tools/download-taxonomy"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", "build/", "./tools/reset-botd"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", "build/", "./tools/gen-botd"); err != nil {
		return err
	}

	return nil
}

func BuildServer() error {
	return sh.Run("go", "build", "-o", "build/", "./cmd/bulletinbird")
}

func BuildTools() error {
	if err := sh.Run("go", "build", "-o", "build/", "./tools/download-taxonomy"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", "build/", "./tools/reset-botd"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", "build/", "./tools/gen-botd"); err != nil {
		return err
	}

	return nil
}

func Install() error {
	if err := sh.Run("go", "install", "./cmd/bulletinbird"); err != nil {
		return err
	}

	if err := sh.Run("go", "install", "./tools/download-taxonomy"); err != nil {
		return err
	}

	if err := sh.Run("go", "install", "./tools/reset-botd"); err != nil {
		return err
	}

	if err := sh.Run("go", "install", "./tools/gen-botd"); err != nil {
		return err
	}

	return nil
}
