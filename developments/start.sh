#!/bin/bash

# TODO: variables
INSFRAS=insfras
SERVICES=services
insfras=(
  "bitnami/postgresql"
  "jaeger-all-in-one/jaeger-all-in-one"
  "hashicorp/consul"
  "grafana/grafana"
  "bitnami/kafka"
  "bitnami/redis"
  "nats/nats"
  "minio/minio"
)
srvs=(
  "gateway"
  "order"
  "inventory"
  "third_party"
)

#TODO: kind set up
install_kind() {
  go install sigs.k8s.io/kind@v0.17.0
  kind create cluster --quiet
}

#TODO: install tools
install_tools() {
  go install github.com/spf13/cobra-cli@latest
}
#TODO: helm repositories
install_helm_repos() {
  helm repo add bitnami https://charts.bitnami.com/bitnami --no-update
  helm repo add jaeger-all-in-one https://raw.githubusercontent.com/hansehe/jaeger-all-in-one/master/helm/charts --no-update
  helm repo add hashicorp https://helm.releases.hashicorp.com --no-update
  helm repo add grafana https://grafana.github.io/helm-charts --no-update
  helm repo add nats https://nats-io.github.io/k8s/helm/charts --no-update
  helm repo add minio https://charts.min.io/ --no-update
}

#TODO: install insfras
install_insfras() {
  for insfra in ${insfras[@]}; do
    name=$(echo ${insfra} | sed -r 's:.*/::')
    helm upgrade --install ${name} ${insfra} --namespace ${INSFRAS} --create-namespace
  done
}

#TODO: install services
install_services() {
  for srv in ${srvs[@]}; do
    helm upgrade --install ${srv} ../deployments/helms/${srv} --namespace ${SERVICES} --create-namespace
  done
}

main() {
  install_kind
  install_tools
  install_helm_repos
  install_insfras
  install_services
}

main
