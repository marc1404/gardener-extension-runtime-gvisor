package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// GVisorConfiguration defines the configuration for the GVisor runtime extension.
type GVisorConfiguration struct {
	metav1.TypeMeta

	// ConfigFlags is a map of additional flags that are passed to the runsc binary.
	ConfigFlags *map[string]interface{}
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
