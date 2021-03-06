package api

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
)

// Config contains a set of Kubernetes resources to be applied.
// TODO: Unify with Kubernetes Config
//       https://github.com/GoogleCloudPlatform/kubernetes/pull/1007
type Config struct {
	kapi.TypeMeta   `json:",inline"`
	kapi.ObjectMeta `json:"metadata,omitempty"`

	// Required: Items is an array of Kubernetes resources of Service,
	// Pod and/or ReplicationController kind.
	Items []runtime.RawExtension `json:"items"`
}
