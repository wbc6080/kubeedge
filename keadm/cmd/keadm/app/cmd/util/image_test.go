/*
Copyright 2022 The KubeEdge Authors.

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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kubeedge/kubeedge/keadm/cmd/keadm/app/cmd/common"
)

func TestEdgeSet(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		args common.JoinOptions
		want Set
	}{
		{
			name: "repo nil, ver not nil",
			args: common.JoinOptions{
				ImageRepository: "",
				KubeEdgeVersion: "v1.9.1",
			},
			want: Set{
				EdgeCore: "kubeedge/installation-package:v1.9.1",
			},
		},
		{
			name: "repo nil, ver nil",
			args: common.JoinOptions{
				ImageRepository: "",
				KubeEdgeVersion: "",
			},
			want: Set{
				EdgeCore: "kubeedge/installation-package",
			},
		},
		{
			name: "repo not nil, ver not nil",
			args: common.JoinOptions{
				ImageRepository: "kubeedge-test",
				KubeEdgeVersion: "v1.9.1",
			},
			want: Set{
				EdgeCore: "kubeedge-test/installation-package:v1.9.1",
			},
		},
		{
			name: "repo not nil, ver nil",
			args: common.JoinOptions{
				ImageRepository: "kubeedge-test",
				KubeEdgeVersion: "",
			},
			want: Set{
				EdgeCore: "kubeedge-test/installation-package",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EdgeSet(&tt.args)
			assert.Equal(tt.want, got)
		})
	}
}

func TestCloudSet(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		imageRepository string
		version         string
	}
	tests := []struct {
		name string
		args args
		want Set
	}{
		{
			name: "repo nil, ver not nil",
			args: args{
				imageRepository: "",
				version:         "v1.12.0",
			},
			want: Set{
				CloudAdmission:         "kubeedge/admission:v1.12.0",
				CloudCloudcore:         "kubeedge/cloudcore:v1.12.0",
				CloudIptablesManager:   "kubeedge/iptables-manager:v1.12.0",
				CloudControllerManager: "kubeedge/controller-manager:v1.12.0",
			},
		},
		{
			name: "repo nil, ver nil",
			args: args{
				imageRepository: "",
				version:         "",
			},
			want: Set{
				CloudAdmission:         "kubeedge/admission",
				CloudCloudcore:         "kubeedge/cloudcore",
				CloudIptablesManager:   "kubeedge/iptables-manager",
				CloudControllerManager: "kubeedge/controller-manager",
			},
		},
		{
			name: "repo not nil, ver not nil",
			args: args{
				imageRepository: "kubeedge-test",
				version:         "v1.12.0",
			},
			want: Set{
				CloudAdmission:         "kubeedge-test/admission:v1.12.0",
				CloudCloudcore:         "kubeedge-test/cloudcore:v1.12.0",
				CloudIptablesManager:   "kubeedge-test/iptables-manager:v1.12.0",
				CloudControllerManager: "kubeedge-test/controller-manager:v1.12.0",
			},
		},
		{
			name: "repo not nil, ver nil",
			args: args{
				imageRepository: "kubeedge-test",
				version:         "",
			},
			want: Set{
				CloudAdmission:         "kubeedge-test/admission",
				CloudCloudcore:         "kubeedge-test/cloudcore",
				CloudIptablesManager:   "kubeedge-test/iptables-manager",
				CloudControllerManager: "kubeedge-test/controller-manager",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CloudSet(tt.args.imageRepository, tt.args.version)
			assert.Equal(tt.want, got)
		})
	}
}

func TestSet_Get(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		s    Set
		args string
		want string
	}{
		{
			name: "get cloudcore image",
			s:    Set{"cloudcore": "kubeedge-test/cloudcore:1.12.0"},
			args: "cloudcore",
			want: "kubeedge-test/cloudcore:1.12.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Get(tt.args)
			assert.Equal(tt.want, got)
		})
	}
}

func TestSet_List(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		s    Set
		want []string
	}{
		{
			name: "test list",
			s:    Set{"cloudcore": "kubeedge-test/cloudcore:v1.12.0", "admission": "kubeedge-test/admission:v1.12.0", "controller-manager": "kubeedge-test/controller-manager:v1.12.0", "iptables-manager": "kubeedge-test/iptables-manager:v1.12.0"},
			want: []string{"kubeedge-test/cloudcore:v1.12.0", "kubeedge-test/admission:v1.12.0", "kubeedge-test/controller-manager:v1.12.0", "kubeedge-test/iptables-manager:v1.12.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.List()
			// we don't care about array sequence, so convert slice to map and compare it
			gotMap := make(map[string]string)
			for _, v := range got {
				gotMap[v] = ""
			}
			wantMap := make(map[string]string)
			for _, v := range tt.want {
				wantMap[v] = ""
			}
			assert.Equal(wantMap, gotMap)
		})
	}
}

func TestSet_Merge(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		s    Set
		args Set
		want Set
	}{
		{
			name: "no overlapping keys",
			s:    Set{"kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/admission": "v1.12.0"},
			args: Set{"kubeedge-test/iptables-manager": "v1.12.0", "kubeedge-test/controller-manager": "v1.12.0"},
			want: Set{"kubeedge-test/admission": "v1.12.0", "kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/controller-manager": "v1.12.0", "kubeedge-test/iptables-manager": "v1.12.0"},
		},
		{
			name: "all no overlapping keys",
			s:    Set{"kubeedge-test/cloudcore": "v1.9.1", "kubeedge-test/admission": "v1.9.1"},
			args: Set{"kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/admission": "v1.12.0"},
			want: Set{"kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/admission": "v1.12.0"},
		},
		{
			name: "partially overlapping keys",
			s:    Set{"kubeedge-test/cloudcore": "v1.9.1", "kubeedge-test/admission": "v1.9.1", "kubeedge-test/iptables-manager": "v1.12.0"},
			args: Set{"kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/admission": "v1.12.0"},
			want: Set{"kubeedge-test/cloudcore": "v1.12.0", "kubeedge-test/admission": "v1.12.0", "kubeedge-test/iptables-manager": "v1.12.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Merge(tt.args)
			assert.Equal(tt.want, got)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		s    Set
		want Set
	}{
		{
			name: "get edgecore image",
			s: Set{
				EdgeCore: "kubeedge/installation-package:v1.12.2",
				"mqtt":   "kubeedge/eclipse-mosquitto:1.6.15",
			},
			want: Set{
				EdgeCore: "kubeedge/installation-package:v1.12.2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Remove("mqtt")
			assert.Equal(tt.want, got)
		})
	}
}
