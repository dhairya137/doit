package installer

var defaultPackages = map[string]Package{
    // Containerization Tools
    "docker": {
        Name:        "docker",
        Description: "Container runtime engine with Docker Compose",
        Category:    CategoryContainerization,
    },
    "containerd": {
        Name:        "containerd",
        Description: "Industry-standard container runtime",
        Category:    CategoryContainerization,
        Command:     "apt-get install -y containerd",
    },

    // CI/CD Tools
    "jenkins": {
        Name:        "jenkins",
        Description: "Leading open source automation server",
        Category:    CategoryCI,
    },
    "gitlab-runner": {
        Name:        "gitlab-runner",
        Description: "GitLab CI/CD runner",
        Category:    CategoryCI,
    },

    // Infrastructure as Code
    "terraform": {
        Name:        "terraform",
        Description: "Infrastructure as Code tool by HashiCorp",
        Category:    CategoryIaC,
    },
    "ansible": {
        Name:        "ansible",
        Description: "Automation tool for configuration management",
        Category:    CategoryIaC,
        Command:     "apt-get install -y ansible",
    },
    "packer": {
        Name:        "packer",
        Description: "Machine image creation tool by HashiCorp",
        Category:    CategoryIaC,
    },

    // Monitoring Tools
    "prometheus": {
        Name:        "prometheus",
        Description: "Monitoring and alerting toolkit",
        Category:    CategoryMonitoring,
    },
    "grafana": {
        Name:        "grafana",
        Description: "Analytics and monitoring solution",
        Category:    CategoryMonitoring,
    },

    // Kubernetes Tools
    "kubernetes": {
        Name:        "kubernetes",
        Description: "Container orchestration system",
        Category:    CategoryOrchestration,
    },
    "kubectl": {
        Name:        "kubectl",
        Description: "Kubernetes command-line tool",
        Category:    CategoryOrchestration,
    },
    "minikube": {
        Name:        "minikube",
        Description: "Local Kubernetes cluster",
        Category:    CategoryOrchestration,
    },
    "helm": {
        Name:        "helm",
        Description: "Kubernetes package manager",
        Category:    CategoryOrchestration,
    },

    // Cloud Tools
    "aws-cli": {
        Name:        "aws-cli",
        Description: "AWS Command Line Interface",
        Category:    CategoryCloudTools,
    },
    "azure-cli": {
        Name:        "azure-cli",
        Description: "Azure Command Line Interface",
        Category:    CategoryCloudTools,
    },

    // Database Tools
    "postgresql": {
        Name:        "postgresql",
        Description: "PostgreSQL database server",
        Category:    CategoryDatabase,
        Command:     "apt-get install -y postgresql postgresql-contrib",
    },
    "mysql": {
        Name:        "mysql",
        Description: "MySQL database server",
        Category:    CategoryDatabase,
        Command:     "apt-get install -y mysql-server",
    },
    "mongodb": {
        Name:        "mongodb",
        Description: "MongoDB database server",
        Category:    CategoryDatabase,
    },
}