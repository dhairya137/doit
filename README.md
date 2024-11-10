# DevOps Installation Tools (doit)

A command-line utility to simplify the installation of DevOps tools on Ubuntu servers.

## Installation

### Prerequisites

Make sure you have Go installed (version 1.16 or higher):

```bash
go version
```

If you don't have Go installed, you can install it:

```bash
# On Ubuntu
sudo apt-get update
sudo apt-get install golang-go
```

### Install doit

Simply run:

```bash
go install github.com/yourusername/doit@latest
```

Make sure your PATH includes the Go bin directory:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

To make this permanent, add it to your ~/.bashrc:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

## Usage

List available packages:

```bash
doit list
```

Install a package (requires sudo):

```bash
sudo doit install docker
sudo doit install kubernetes
sudo doit install terraform
```

## Available Tools

- Containerization
  - Docker
  - Containerd
- CI/CD
  - Jenkins
  - GitLab Runner
- Infrastructure as Code
  - Terraform
  - Ansible
  - Packer
- Monitoring
  - Prometheus
  - Grafana
- Orchestration
  - Kubernetes
  - Helm
  - Minikube
- Databases
  - MongoDB
  - PostgreSQL
  - MySQL
- Cloud Tools
  - AWS CLI
  - Azure CLI
