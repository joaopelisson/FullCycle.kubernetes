# <img src="https://avatars.githubusercontent.com/u/10950003?s=200&v=4" alt="Full Cycle" width="25" /> FullCycle.Kubernetes

> 🇺🇸 **English** | 🇧🇷 **Português**

---

## 🇺🇸 English

This repository is part of the **Kubernetes** track from **@FullCycle** and demonstrates **how to containerize, deploy, and manage applications in Kubernetes** using industry best practices and real-world scenarios.

---

### 🎯 Repository Purpose

**Kubernetes fundamentals and advanced container orchestration**, covering practical scenarios such as:

- Docker containerization and image building
- Kubernetes core resources (Pods, Deployments, Services)
- ConfigMaps and Secrets for configuration management
- Persistent storage with PersistentVolumes and PersistentVolumeClaims
- Horizontal Pod Autoscaling (HPA) and resource management
- Health checks (liveness, readiness, startup probes)
- Kubernetes networking and service discovery
- StatefulSets for stateful applications
- Local Kubernetes cluster setup with KinD

---

### ⚙️ Key Concepts & Approaches

- **Container Fundamentals**
  - Docker image creation with Dockerfile
  - Go application containerization
  - Image optimization and best practices
  - Container registry usage

- **Kubernetes Core Resources**
  - Pods and Pod lifecycle
  - Deployments and ReplicaSets
  - Services (ClusterIP, LoadBalancer)
  - StatefulSets for stateful workloads
  - Headless services for StatefulSets

- **Configuration & Secrets Management**
  - ConfigMaps for application configuration
  - Secrets for sensitive data
  - Environment variable injection
  - Volume mounting for configuration files

- **Storage & Data Persistence**
  - PersistentVolumes (PV)
  - PersistentVolumeClaims (PVC)
  - volume mounting in containers
  - Storage classes and resource management

- **Application Health & Reliability**
  - Startup probes for initialization
  - Readiness probes for traffic routing
  - Liveness probes for crash detection
  - Resource requests and limits
  - Graceful shutdown handling

- **Scaling & Performance**
  - Horizontal Pod Autoscaler (HPA)
  - CPU utilization monitoring
  - Metrics Server setup and configuration
  - Auto-scaling policies

- **Local Development**
  - KinD (Kubernetes in Docker) cluster setup
  - Multi-node cluster configuration
  - Local development workflow
  - Testing and debugging

---

### 📋 Project Structure

```
.
├── server.go              # Go application with multiple endpoints
├── Dockerfile             # Container image definition
├── kind.yaml              # KinD cluster configuration
├── deployment.yaml        # Kubernetes Deployment specification
├── service.yaml           # LoadBalancer Service
├── replicaset.yaml        # ReplicaSet for pod replication
├── statefulset.yaml       # StatefulSet for MySQL
├── pod.yaml               # Single pod definition
├── configmap-env.yaml     # Environment configuration
├── configmap-family.yaml  # File-based configuration
├── secret.yaml            # Sensitive data management
├── pv.yaml                # PersistentVolume definition
├── pvc.yaml               # PersistentVolumeClaim definition
├── hpa.yaml               # Horizontal Pod Autoscaler
├── mysql-service-h.yaml   # Headless Service for MySQL
└── metrics-server.yaml    # Metrics collection for HPA
```

---

### 🚀 Getting Started

#### Prerequisites
- Docker installed
- kubectl installed
- KinD installed or docker available

#### Setup Local Kubernetes Cluster with KinD

```bash
# Create a local Kubernetes cluster using KinD
kind create cluster --config kind.yaml

# Verify cluster is running
kubectl cluster-info
kubectl get nodes
```

#### Build and Load Docker Image

```bash
# Build the Go application Docker image
docker build -t joaopelisson/hello-go:v3.6 .

# Load image into KinD cluster (if using local image)
kind load docker-image joaopelisson/hello-go:v3.6
```

#### Deploy Application

```bash
# Create ConfigMaps and Secrets
kubectl apply -f configmap-env.yaml
kubectl apply -f configmap-family.yaml
kubectl apply -f secret.yaml

# Create Storage resources
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

# Deploy the application
kubectl apply -f deployment.yaml

# Expose the service
kubectl apply -f service.yaml

# Setup auto-scaling
kubectl apply -f metrics-server.yaml
kubectl apply -f hpa.yaml
```

#### Verify Deployment

```bash
# Check pod status
kubectl get pods

# Check service
kubectl get service goserver-service

# View application logs
kubectl logs -f deployment/goserver

# Port-forward to access application locally
kubectl port-forward service/goserver-service 8080:80
```

#### Access Application Endpoints

- **Home**: http://localhost:8080/ (shows name and age from ConfigMap)
- **Health Check**: http://localhost:8080/healthz (startup probe)
- **ConfigMap Data**: http://localhost:8080/configmap (mounted file)
- **Secret Data**: http://localhost:8080/secret (environment variables)

---

### 📊 Application Features

- **Docker Integration**: Full containerization workflow
- **Health Probes**: Startup, readiness, and liveness probe implementation
- **Configuration Management**: Multiple ConfigMaps with different sources
- **Secrets**: Base64-encoded sensitive data management
- **Resource Management**: CPU and memory requests/limits
- **Persistent Storage**: Volume mounting for application data
- **Service Discovery**: Internal and external service exposure
- **Auto-scaling**: Automatic pod scaling based on CPU metrics

---

## 🇧🇷 Português (PT-BR)

Este repositório faz parte da trilha **Kubernetes** da **@FullCycle** e demonstra **como containerizar, implantar e gerenciar aplicações em Kubernetes** utilizando as melhores práticas da indústria e cenários do mundo real.

---

### 🎯 Objetivo do Repositório

**Fundamentais de Kubernetes e orquestração avançada de containers**, abordando cenários práticos como:

- Containerização Docker e construção de imagens
- Recursos core do Kubernetes (Pods, Deployments, Services)
- ConfigMaps e Secrets para gerenciamento de configuração
- Armazenamento persistente com PersistentVolumes e PersistentVolumeClaims
- Horizontal Pod Autoscaling (HPA) e gerenciamento de recursos
- Health checks (liveness, readiness, startup probes)
- Networking e service discovery no Kubernetes
- StatefulSets para aplicações com estado
- Setup de cluster local com KinD

---

### ⚙️ Principais Conceitos e Abordagens

- **Fundamentos de Container**
  - Criação de imagem Docker com Dockerfile
  - Containerização de aplicação Go
  - Otimização e boas práticas de imagem
  - Uso de registry de container

- **Recursos Core do Kubernetes**
  - Pods e ciclo de vida
  - Deployments e ReplicaSets
  - Services (ClusterIP, LoadBalancer)
  - StatefulSets para workloads com estado
  - Headless services para StatefulSets

- **Gerenciamento de Configuração & Segredos**
  - ConfigMaps para configuração de aplicação
  - Secrets para dados sensíveis
  - Injeção de variáveis de ambiente
  - Volume mounting para arquivos de configuração

- **Armazenamento & Persistência de Dados**
  - PersistentVolumes (PV)
  - PersistentVolumeClaims (PVC)
  - Volume mounting em containers
  - Storage classes e gerenciamento de recursos

- **Health & Confiabilidade da Aplicação**
  - Startup probes para inicialização
  - Readiness probes para roteamento de tráfego
  - Liveness probes para detecção de falhas
  - Resource requests e limits
  - Shutdown gracioso

- **Scaling & Performance**
  - Horizontal Pod Autoscaler (HPA)
  - Monitoramento de utilização de CPU
  - Setup e configuração de Metrics Server
  - Políticas de auto-scaling

- **Desenvolvimento Local**
  - Setup de cluster KinD (Kubernetes in Docker)
  - Configuração multi-node
  - Workflow de desenvolvimento local
  - Testes e debugging

---

### 📋 Estrutura do Projeto

```
.
├── server.go              # Aplicação Go com múltiplos endpoints
├── Dockerfile             # Definição da imagem container
├── kind.yaml              # Configuração do cluster KinD
├── deployment.yaml        # Especificação de Deployment
├── service.yaml           # Service LoadBalancer
├── replicaset.yaml        # ReplicaSet para replicação de pods
├── statefulset.yaml       # StatefulSet para MySQL
├── pod.yaml               # Definição de pod único
├── configmap-env.yaml     # Configuração de ambiente
├── configmap-family.yaml  # Configuração baseada em arquivo
├── secret.yaml            # Gerenciamento de dados sensíveis
├── pv.yaml                # Definição de PersistentVolume
├── pvc.yaml               # Definição de PersistentVolumeClaim
├── hpa.yaml               # Horizontal Pod Autoscaler
├── mysql-service-h.yaml   # Headless Service para MySQL
└── metrics-server.yaml    # Coleta de métricas para HPA
```

---

### 🚀 Primeiros Passos

#### Pré-requisitos
- Docker instalado
- kubectl instalado
- KinD instalado ou Docker disponível

#### Setup de Cluster Local com KinD

```bash
# Criar um cluster Kubernetes local usando KinD
kind create cluster --config kind.yaml

# Verificar se o cluster está rodando
kubectl cluster-info
kubectl get nodes
```

#### Build e Load da Imagem Docker

```bash
# Build da imagem Docker da aplicação Go
docker build -t joaopelisson/hello-go:v3.6 .

# Load da imagem no cluster KinD (se usar imagem local)
kind load docker-image joaopelisson/hello-go:v3.6
```

#### Implantar Aplicação

```bash
# Criar ConfigMaps e Secrets
kubectl apply -f configmap-env.yaml
kubectl apply -f configmap-family.yaml
kubectl apply -f secret.yaml

# Criar recursos de armazenamento
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

# Implantar a aplicação
kubectl apply -f deployment.yaml

# Expor o serviço
kubectl apply -f service.yaml

# Setup de auto-scaling
kubectl apply -f metrics-server.yaml
kubectl apply -f hpa.yaml
```

#### Verificar Implantação

```bash
# Verificar status dos pods
kubectl get pods

# Verificar serviço
kubectl get service goserver-service

# Ver logs da aplicação
kubectl logs -f deployment/goserver

# Port-forward para acessar a aplicação localmente
kubectl port-forward service/goserver-service 8080:80
```

#### Acessar Endpoints da Aplicação

- **Home**: http://localhost:8080/ (mostra nome e idade do ConfigMap)
- **Health Check**: http://localhost:8080/healthz (startup probe)
- **Dados ConfigMap**: http://localhost:8080/configmap (arquivo montado)
- **Dados Secret**: http://localhost:8080/secret (variáveis de ambiente)

---

### 📊 Características da Aplicação

- **Integração Docker**: Workflow completo de containerização
- **Health Probes**: Implementação de startup, readiness e liveness probes
- **Gerenciamento de Configuração**: Múltiplos ConfigMaps com diferentes fontes
- **Secrets**: Gerenciamento de dados sensíveis codificados em Base64
- **Gerenciamento de Recursos**: CPU e memory requests/limits
- **Armazenamento Persistente**: Volume mounting para dados da aplicação
- **Service Discovery**: Exposição interna e externa de serviços
- **Auto-scaling**: Escalabilidade automática de pods baseada em métricas de CPU

---

### 📚 Resources & Learning

- [Official Kubernetes Documentation](https://kubernetes.io/docs/)
- [KinD Documentation](https://kind.sigs.k8s.io/)
- [Go Web Server Best Practices](https://golang.org/doc/)

