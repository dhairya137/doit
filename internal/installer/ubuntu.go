package installer

import (
	"fmt"
	"os/exec"
	"strings"
)

type UbuntuInstaller struct {
    packages map[string]Package
}

func NewUbuntuInstaller() *UbuntuInstaller {
    return &UbuntuInstaller{
        packages: defaultPackages,
    }
}

// Base installer methods
func (u *UbuntuInstaller) executeSteps(steps []InstallationStep) error {
    for _, step := range steps {
        fmt.Printf("📦 %s...\n", step.Description)
        
        cmdParts := strings.Fields("sudo " + step.Command)
        if strings.Contains(step.Command, "sh -c") {
            cmdParts = []string{"sudo", "sh", "-c", strings.TrimPrefix(step.Command, "sh -c ")}
        }
        
        cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("failed during step '%s': %v\nOutput: %s", 
                step.Description, err, string(output))
        }
    }
    return nil
}

func (u *UbuntuInstaller) Install(packageName string) error {
    pkg, exists := u.packages[packageName]
    if !exists {
        return fmt.Errorf("package %s not found", packageName)
    }

    // Check for specialized installation functions
    switch packageName {
    case "docker":
        return u.dockerInstall()
    case "kubernetes":
        return u.kubernetesInstall()
    case "terraform":
        return u.terraformInstall()
    case "jenkins":
        return u.jenkinsInstall()
    case "prometheus":
        return u.prometheusInstall()
    case "grafana":
        return u.grafanaInstall()
    case "helm":
        return u.helmInstall()
    case "gitlab-runner":
        return u.gitlabRunnerInstall()
    case "aws-cli":
        return u.awsCliInstall()
    case "mongodb":
        return u.mongoDBInstall()
    case "minikube":
        return u.minikubeInstall()
    }

    // Default installation for simple packages
    if pkg.Command != "" {
        fmt.Printf("📦 Installing %s...\n", packageName)
        return u.executeSteps([]InstallationStep{{
            Description: fmt.Sprintf("Installing %s", packageName),
            Command:     pkg.Command,
        }})
    }

    return fmt.Errorf("no installation method defined for package %s", packageName)
}

func (u *UbuntuInstaller) ListAvailablePackages() []Package {
    packages := make([]Package, 0, len(u.packages))
    for _, pkg := range u.packages {
        packages = append(packages, pkg)
    }
    return packages
}

// Complex package installations
func (u *UbuntuInstaller) dockerInstall() error {
    steps := []InstallationStep{
        {"Updating package list", "apt-get update"},
        {"Installing prerequisites", "apt-get install -y ca-certificates curl gnupg"},
        {"Creating keyring directory", "install -m 0755 -d /etc/apt/keyrings"},
        {"Downloading Docker's GPG key", "curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc"},
        {"Setting permissions for Docker's GPG key", "chmod a+r /etc/apt/keyrings/docker.asc"},
        {"Adding Docker repository", `sh -c 'echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null'`},
        {"Updating package list with Docker repository", "apt-get update"},
        {"Installing Docker", "apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin"},
        {"Adding current user to docker group", "usermod -aG docker $USER"},
    }
    fmt.Println("🔧 Docker installed successfully! You may need to log out and back in for group changes to take effect.")
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) kubernetesInstall() error {
    steps := []InstallationStep{
        {"Installing prerequisites", "apt-get install -y apt-transport-https ca-certificates curl"},
        {"Downloading Kubernetes GPG key", "curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-archive-keyring.gpg"},
        {"Adding Kubernetes repository", `sh -c 'echo "deb [signed-by=/etc/apt/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing kubelet, kubeadm and kubectl", "apt-get install -y kubelet kubeadm kubectl"},
        {"Marking packages to hold version", "apt-mark hold kubelet kubeadm kubectl"},
    }
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) terraformInstall() error {
    steps := []InstallationStep{
        {"Installing prerequisites", "apt-get install -y gnupg software-properties-common"},
        {"Adding HashiCorp GPG key", "wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg"},
        {"Adding HashiCorp repository", `sh -c 'echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing Terraform", "apt-get install -y terraform"},
    }
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) jenkinsInstall() error {
    steps := []InstallationStep{
        {"Installing Java", "apt-get install -y openjdk-11-jdk"},
        {"Adding Jenkins key", "curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key | sudo tee /usr/share/keyrings/jenkins-keyring.asc > /dev/null"},
        {"Adding Jenkins repository", `sh -c 'echo "deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] https://pkg.jenkins.io/debian-stable binary/" | tee /etc/apt/sources.list.d/jenkins.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing Jenkins", "apt-get install -y jenkins"},
        {"Starting Jenkins service", "systemctl start jenkins"},
        {"Enabling Jenkins service", "systemctl enable jenkins"},
    }
    fmt.Println("🔧 After installation, get your initial admin password with: sudo cat /var/lib/jenkins/secrets/initialAdminPassword")
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) prometheusInstall() error {
    steps := []InstallationStep{
        {"Creating Prometheus user", "useradd --no-create-home --shell /bin/false prometheus"},
        {"Creating directories", "mkdir -p /etc/prometheus /var/lib/prometheus"},
        {"Downloading Prometheus", "curl -L https://github.com/prometheus/prometheus/releases/download/v2.47.0/prometheus-2.47.0.linux-amd64.tar.gz -o prometheus.tar.gz"},
        {"Extracting Prometheus", "tar xvf prometheus.tar.gz"},
        {"Moving binary", "cp prometheus-*/prometheus /usr/local/bin/"},
        {"Moving configuration", "cp prometheus-*/promtool /usr/local/bin/"},
        {"Setting permissions", "chown prometheus:prometheus /usr/local/bin/prometheus /usr/local/bin/promtool"},
        {"Configuring service", `sh -c 'echo "[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecStart=/usr/local/bin/prometheus

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/prometheus.service'`},
        {"Reloading systemd", "systemctl daemon-reload"},
        {"Starting Prometheus", "systemctl start prometheus"},
        {"Enabling Prometheus", "systemctl enable prometheus"},
        {"Cleanup", "rm -rf prometheus.tar.gz prometheus-*"},
    }
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) grafanaInstall() error {
    steps := []InstallationStep{
        {"Installing prerequisites", "apt-get install -y software-properties-common wget"},
        {"Adding Grafana GPG key", "wget -q -O - https://packages.grafana.com/gpg.key | apt-key add -"},
        {"Adding Grafana repository", `sh -c 'echo "deb https://packages.grafana.com/oss/deb stable main" | tee /etc/apt/sources.list.d/grafana.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing Grafana", "apt-get install -y grafana"},
        {"Starting Grafana service", "systemctl start grafana-server"},
        {"Enabling Grafana service", "systemctl enable grafana-server"},
    }
    fmt.Println("🔧 Grafana will be available at http://localhost:3000 (default credentials: admin/admin)")
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) helmInstall() error {
    steps := []InstallationStep{
        {"Adding Helm GPG key", "curl https://baltocdn.com/helm/signing.asc | gpg --dearmor -o /usr/share/keyrings/helm.gpg"},
        {"Adding Helm repository", `sh -c 'echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing Helm", "apt-get install -y helm"},
    }
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) gitlabRunnerInstall() error {
    steps := []InstallationStep{
        {"Adding GitLab Runner repository", `curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | bash`},
        {"Installing GitLab Runner", "apt-get install -y gitlab-runner"},
    }
    fmt.Println("🔧 Remember to register your runner with: gitlab-runner register")
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) awsCliInstall() error {
    steps := []InstallationStep{
        {"Installing prerequisites", "apt-get install -y unzip"},
        {"Downloading AWS CLI", "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o 'awscliv2.zip'"},
        {"Extracting AWS CLI", "unzip awscliv2.zip"},
        {"Installing AWS CLI", "./aws/install"},
        {"Cleanup", "rm -rf aws awscliv2.zip"},
    }
    fmt.Println("🔧 Configure AWS CLI with: aws configure")
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) mongoDBInstall() error {
    steps := []InstallationStep{
        {"Installing prerequisites", "apt-get install -y gnupg curl"},
        {"Adding MongoDB GPG key", "curl -fsSL https://www.mongodb.org/static/pgp/server-6.0.asc | sudo gpg --dearmor -o /usr/share/keyrings/mongodb.gpg"},
        {"Adding MongoDB repository", `sh -c 'echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb.gpg ] https://repo.mongodb.org/apt/ubuntu $(lsb_release -cs)/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list'`},
        {"Updating package list", "apt-get update"},
        {"Installing MongoDB", "apt-get install -y mongodb-org"},
        {"Starting MongoDB service", "systemctl start mongod"},
        {"Enabling MongoDB service", "systemctl enable mongod"},
    }
    return u.executeSteps(steps)
}

func (u *UbuntuInstaller) minikubeInstall() error {
    steps := []InstallationStep{
        {"Downloading Minikube", "curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64"},
        {"Installing Minikube", "install minikube-linux-amd64 /usr/local/bin/minikube"},
        {"Cleanup", "rm minikube-linux-amd64"},
    }
    fmt.Println("🔧 Start Minikube with: minikube start")
    return u.executeSteps(steps)
}