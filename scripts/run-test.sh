#!/bin/bash
set -euxo pipefail

OPERATOR_ROOT=${OPERATOR_ROOT:-/root/go/src/github.com/openshift/sriov-network-operator}
NAMESPACE=${NAMESPACE:-openshift-sriov-network-operator}
DIR=$PWD

export TEST_NAMESPACE=${NAMESPACE}
export KUBECONFIG=${KUBECONFIG:-/root/dev-scripts/ocp/auth/kubeconfig}

cd $OPERATOR_ROOT
EXCLUSIONS="operator.yaml" $OPERATOR_ROOT/hack/deploy-setup.sh ${NAMESPACE}

source $OPERATOR_ROOT/hack/env.sh
echo ${SRIOV_CNI_IMAGE}
echo ${SRIOV_DEVICE_PLUGIN_IMAGE}
echo ${NETWORK_RESOURCES_INJECTOR_IMAGE}
echo ${SRIOV_NETWORK_CONFIG_DAEMON_IMAGE}
echo ${SRIOV_NETWORK_OPERATOR_IMAGE}
echo ${SRIOV_NETWORK_WEBHOOK_IMAGE}
envsubst < deploy/operator.yaml  > $OPERATOR_ROOT/deploy/operator-init.yaml

cd $DIR
# GO111MODULE=on go test ./tests/operator/...  -root=$OPERATOR_ROOT -kubeconfig=$KUBECONFIG -globalMan $OPERATOR_ROOT/deploy/crds/sriovnetwork.openshift.io_sriovnetworks_crd.yaml -namespacedMan $OPERATOR_ROOT/deploy/operator-init.yaml -v -singleNamespace true
ginkgo -v --progress ./tests/$1 -- -root=$OPERATOR_ROOT -kubeconfig=$KUBECONFIG -globalMan $OPERATOR_ROOT/deploy/crds/sriovnetwork.openshift.io_sriovnetworks_crd.yaml -namespacedMan $OPERATOR_ROOT/deploy/operator-init.yaml -singleNamespace true

cd $OPERATOR_ROOT
make undeploy
