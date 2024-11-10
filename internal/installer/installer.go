package installer

// Installer interface defines methods for package installation
type Installer interface {
    Install(packageName string) error
    ListAvailablePackages() []Package
}

// NewInstaller creates a new installer instance
func NewInstaller() Installer {
    return NewUbuntuInstaller()
}