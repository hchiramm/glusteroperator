/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const GlusterClusterResourcePlural = "glusterclusters"

type GlusterCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              GlusterClusterSpec   `json:"spec"`
	Status            GlusterClusterStatus `json:"status,omitempty"`
}

type GlusterClusterSpec struct {
	Node    string `json:"node"`
	Devices bool   `json:"device"`
}

type GlusterClusterStatus struct {
	State   GlusterClusterState `json:"state,omitempty"`
	Message string              `json:"message,omitempty"`
}

type GlusterClusterState string

const (
	GlusterClusterStateCreated GlusterClusterState = "Created"
	GlusterClusterProcessed    GlusterClusterState = "Processed"
)

type GlusterClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GlusterCluster `json:"glustercluster"`
}
