package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient

// GVisorConfiguration defines the configuration for the gVisor runtime extension.
type GVisorConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ConfigFlags is a map of additional flags that are passed to the runsc binary used by gVisor.
	// +optional
	ConfigFlags *map[string]interface{} `json:"configFlags,omitempty"`
}

// DeepCopyObject is a deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GVisorConfiguration) DeepCopyObject() runtime.Object {
	out := &GVisorConfiguration{}
	out.TypeMeta = in.TypeMeta
	if in.ConfigFlags == nil {
		return out
	}
	configFlags := make(map[string]interface{}, len(*in.ConfigFlags))
	out.ConfigFlags = &configFlags
	for key, value := range *in.ConfigFlags {
		configFlags[key] = value
	}
	return out
}
