package installer

// Package categories
const (
    CategoryContainerization = "Containerization"
    CategoryCI              = "CI/CD"
    CategoryMonitoring      = "Monitoring"
    CategoryIaC            = "Infrastructure as Code"
    CategoryDatabase       = "Database"
    CategoryOrchestration  = "Orchestration"
    CategoryVersionControl = "Version Control"
    CategoryCloudTools     = "Cloud Tools"
    CategoryNetworking     = "Networking"
)

// Package represents a software package that can be installed
type Package struct {
    Name        string
    Description string
    Category    string
    Command     string
}

// InstallationStep represents a single step in the installation process
type InstallationStep struct {
    Description string
    Command     string
}