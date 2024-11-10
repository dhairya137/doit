// internal/utils/system.go
package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// CheckRoot verifies if the current user has sudo privileges
func CheckRoot() error {
	cmd := exec.Command("sudo", "-n", "true")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("this command requires sudo privileges")
	}
	return nil
}

// GetUbuntuVersion returns the Ubuntu version codename
func GetUbuntuVersion() (string, error) {
	cmd := exec.Command("lsb_release", "-cs")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to determine Ubuntu version: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// IsPackageInstalled checks if a package is already installed
func IsPackageInstalled(packageName string) bool {
	cmd := exec.Command("dpkg", "-l", packageName)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// CheckDependencies verifies if all required dependencies are installed
func CheckDependencies(deps []string) error {
	for _, dep := range deps {
		if !IsPackageInstalled(dep) {
			return fmt.Errorf("required dependency '%s' is not installed", dep)
		}
	}
	return nil
}