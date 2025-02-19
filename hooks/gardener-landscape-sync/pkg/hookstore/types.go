// Copyright 2021 The Kubeswitch authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hookstore

import (
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	"github.com/danielfoehrkn/kubeswitch/types"
)

const (
	KubeconfigStoreFilesystem = "filesystem"
	KubeconfigStoreVault      = "vault"
)

type KubeconfigStore interface {
	GetKind() types.StoreKind
	CreateLandscapeDirectory(dir string) error
	WriteKubeconfigFile(directory, kubeconfigName string, kubeconfigSecret corev1.Secret) error
	CleanExistingKubeconfigs(log *logrus.Entry, dir string) error
}

type FileStore struct {
}

type VaultStore struct {
	Client *vaultapi.Client
}
