// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetwork":                 schema_pkg_apis_sriovnetwork_v1_SriovNetwork(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicy":       schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicy(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec":   schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicySpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus": schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicyStatus(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeState":        schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeState(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec":    schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateSpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus":  schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateStatus(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec":             schema_pkg_apis_sriovnetwork_v1_SriovNetworkSpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus":           schema_pkg_apis_sriovnetwork_v1_SriovNetworkStatus(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfig":          schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfig(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigSpec":      schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfigSpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigStatus":    schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfigStatus(ref),
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetwork(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetwork is the Schema for the sriovnetworks API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicy is the Schema for the sriovnetworknodepolicies API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicySpec defines the desired state of SriovNetworkNodePolicy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Description: "SRIOV Network device plugin endpoint resource name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector selects the nodes to be configured",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"priority": {
						SchemaProps: spec.SchemaProps{
							Description: "Priority of the policy, higher priority policies can override lower ones.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"mtu": {
						SchemaProps: spec.SchemaProps{
							Description: "MTU of VF",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"numVfs": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of VFs for each PF",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"nicSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NicSelector selects the NICs to be configured",
							Ref:         ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNicSelector"),
						},
					},
					"deviceType": {
						SchemaProps: spec.SchemaProps{
							Description: "The driver type for configured VFs. Allowed value \"netdevice\", \"vfio-pci\". Defaults to netdevice.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"isRdma": {
						SchemaProps: spec.SchemaProps{
							Description: "RDMA mode. Defaults to false.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"resourceName", "nodeSelector", "numVfs", "nicSelector"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNicSelector"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicyStatus defines the observed state of SriovNetworkNodePolicy",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeState is the Schema for the sriovnetworknodestates API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeStateSpec defines the desired state of SriovNetworkNodeState",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"dpConfigVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"interfaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.Interface"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.Interface"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeStateStatus defines the observed state of SriovNetworkNodeState",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"interfaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.InterfaceExt"),
									},
								},
							},
						},
					},
					"syncStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastSyncError": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.InterfaceExt"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkSpec defines the desired state of SriovNetwork",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"networkNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the NetworkAttachmentDefinition custom resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Description: "SRIOV Network device plugin endpoint resource name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"capabilities": {
						SchemaProps: spec.SchemaProps{
							Description: "Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{\"mac\": true}'",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ipam": {
						SchemaProps: spec.SchemaProps{
							Description: "IPAM configuration to be used for this network.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"vlan": {
						SchemaProps: spec.SchemaProps{
							Description: "VLAN ID to assign for the VF. Defaults to 0.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"vlanQoS": {
						SchemaProps: spec.SchemaProps{
							Description: "VLAN QoS ID to assign for the VF. Defaults to 0.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"spoofChk": {
						SchemaProps: spec.SchemaProps{
							Description: "VF spoof check, (on|off)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"trust": {
						SchemaProps: spec.SchemaProps{
							Description: "VF trust mode (on|off)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"linkState": {
						SchemaProps: spec.SchemaProps{
							Description: "VF link state (enable|disable|auto)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"resourceName"},
			},
		},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkStatus defines the observed state of SriovNetwork",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovOperatorConfig is the Schema for the sriovoperatorconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigSpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovOperatorConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovOperatorConfigSpec defines the desired state of SriovOperatorConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configDaemonNodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector selects the nodes to be configured",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"enableInjector": {
						SchemaProps: spec.SchemaProps{
							Description: "Flag to control whether the network resource injector webhook shall be deployed",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"enableOperatorWebhook": {
						SchemaProps: spec.SchemaProps{
							Description: "Flag to control whether the operator admission controller webhook shall be deployed",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovOperatorConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovOperatorConfigStatus defines the observed state of SriovOperatorConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"injector": {
						SchemaProps: spec.SchemaProps{
							Description: "Show the runtime status of the network resource injector webhook",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"operatorWebhook": {
						SchemaProps: spec.SchemaProps{
							Description: "Show the runtime status of the operator admission controller webhook",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}
