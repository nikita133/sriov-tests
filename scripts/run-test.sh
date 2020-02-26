#!/bin/bash
set -euxo pipefail

NAMESPACE=${NAMESPACE:-openshift-sriov-network-operator}
DIR=$PWD

export TEST_NAMESPACE=${NAMESPACE}
export KUBECONFIG=${KUBECONFIG:-/root/dev-scripts/ocp/auth/kubeconfig}


cd $DIR
# GO111MODULE=on go test ./tests/operator/...  -root=$OPERATOR_ROOT -kubeconfig=$KUBECONFIG -globalMan $OPERATOR_ROOT/deploy/crds/sriovnetwork.openshift.io_sriovnetworks_crd.yaml -namespacedMan $OPERATOR_ROOT/deploy/operator-init.yaml -v -singleNamespace true
ginkgo -v --progress ./tests/$1 -- -root=$DIR -kubeconfig=$KUBECONFIG -globalMan $DIR/scripts/dummy.yaml -namespacedMan $DIR/scripts/dummy.yaml -singleNamespace true
