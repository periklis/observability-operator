// +groupName=observabilityui.rhobs
// +kubebuilder:rbac:groups=observabilityui.rhobs,resources=observabilityuiplugins,verbs=list;get;watch
// +kubebuilder:rbac:groups=observabilityui.rhobs,resources=observabilityuiplugins/status;observabilityuiplugins/finalizers,verbs=get;update

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoggingStack is the Schema for the observabilityuiplugin API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type ObservabilityUIPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObservabilityUIPluginSpec   `json:"spec,omitempty"`
	Status ObservabilityUIPluginStatus `json:"status,omitempty"`
}

// LoggingStackList contains a list of LoggingStack
// +kubebuilder:resource
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ObservabilityUIPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityUIPlugin `json:"items"`
}

type UIPluginType string

const (
	TypeLogs UIPluginType = "logs"
)

// Observabilityuipluginpec is the specification for desired Logging Stack
type ObservabilityUIPluginSpec struct {
	// +required
	// +kubebuilder:validation:Required
	Type UIPluginType `json:"type"`
	// +required
	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName"`
}

// Observabilityuiplugintatus defines the observed state of ObservabilityUIPlugin.
// It should always be reconstructable from the state of the cluster and/or outside world.
type ObservabilityUIPluginStatus struct {
	// Conditions provide status information about the LoggingStack
	// +listType=atomic
	Conditions []Condition `json:"conditions"`
}

type ConditionStatus string

// +required
// +kubebuilder:validation:Required
// +kubebuilder:validation:Pattern=`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`
// +kubebuilder:validation:MaxLength=316
type ConditionType string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"

	ReconciledCondition        ConditionType = "Reconciled"
	AvailableCondition         ConditionType = "Available"
	ResourceDiscoveryCondition ConditionType = "ResourceDiscovery"
)

type Condition struct {
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	// The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)
	Type ConditionType `json:"type"`
	// observedGeneration represents the .metadata.generation that the condition was set based upon.
	// For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
	// with respect to the current state of the instance.
	// +optional
	// +kubebuilder:validation:Minimum=0
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// lastTransitionTime is the last time the condition transitioned from one status to another.
	// This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// reason contains a programmatic identifier indicating the reason for the condition's last transition.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	// The value should be a CamelCase string.
	// This field may not be empty.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`
	Reason string `json:"reason"`
	// message is a human readable message indicating details about the transition.
	// This may be an empty string.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=32768
	Message string `json:"message"`
	// status of the condition
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=True;False;Unknown;Degraded
	Status ConditionStatus `json:"status"`
}

func (c Condition) Equal(n Condition) bool {
	if c.Reason == n.Reason && c.Status == n.Status && c.Message == n.Message && c.ObservedGeneration == n.ObservedGeneration {
		return true
	}
	return false
}
