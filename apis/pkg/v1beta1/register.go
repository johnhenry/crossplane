/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "pkg.crossplane.io"
	Version = "v1beta1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds all registered types to the scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// Configuation type metadata.
var (
	ConfigurationKind             = reflect.TypeOf(Configuration{}).Name()
	ConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigurationKind}.String()
	ConfigurationKindAPIVersion   = ConfigurationKind + "." + SchemeGroupVersion.String()
	ConfigurationGroupVersionKind = SchemeGroupVersion.WithKind(ConfigurationKind)
)

// ConfigurationRevision type metadata.
var (
	ConfigurationRevisionKind             = reflect.TypeOf(ConfigurationRevision{}).Name()
	ConfigurationRevisionGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigurationRevisionKind}.String()
	ConfigurationRevisionKindAPIVersion   = ConfigurationRevisionKind + "." + SchemeGroupVersion.String()
	ConfigurationRevisionGroupVersionKind = SchemeGroupVersion.WithKind(ConfigurationRevisionKind)
)

// Provider type metadata.
var (
	ProviderKind             = reflect.TypeOf(Provider{}).Name()
	ProviderGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderKind}.String()
	ProviderKindAPIVersion   = ProviderKind + "." + SchemeGroupVersion.String()
	ProviderGroupVersionKind = SchemeGroupVersion.WithKind(ProviderKind)
)

// ProviderRevision type metadata.
var (
	ProviderRevisionKind             = reflect.TypeOf(ProviderRevision{}).Name()
	ProviderRevisionGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderRevisionKind}.String()
	ProviderRevisionKindAPIVersion   = ProviderRevisionKind + "." + SchemeGroupVersion.String()
	ProviderRevisionGroupVersionKind = SchemeGroupVersion.WithKind(ProviderRevisionKind)
)

// Lock type metadata.
var (
	LockKind             = reflect.TypeOf(Lock{}).Name()
	LockGroupKind        = schema.GroupKind{Group: Group, Kind: LockKind}.String()
	LockKindAPIVersion   = LockKind + "." + SchemeGroupVersion.String()
	LockGroupVersionKind = SchemeGroupVersion.WithKind(LockKind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
	SchemeBuilder.Register(&ConfigurationRevision{}, &ConfigurationRevisionList{})
	SchemeBuilder.Register(&Provider{}, &ProviderList{})
	SchemeBuilder.Register(&ProviderRevision{}, &ProviderRevisionList{})
	SchemeBuilder.Register(&Lock{}, &LockList{})
}
