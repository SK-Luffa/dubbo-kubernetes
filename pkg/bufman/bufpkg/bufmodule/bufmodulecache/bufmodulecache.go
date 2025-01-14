// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufmodulecache

import (
	"github.com/apache/dubbo-kubernetes/pkg/bufman/bufpkg/bufmodule"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/connect/registry/v1alpha1/registryv1alpha1connect"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/connectclient"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/storage"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/verbose"
	"go.uber.org/zap"
)

type RepositoryServiceClientFactory func(address string) registryv1alpha1connect.RepositoryServiceClient

func NewRepositoryServiceClientFactory(clientConfig *connectclient.Config) RepositoryServiceClientFactory {
	return func(address string) registryv1alpha1connect.RepositoryServiceClient {
		return connectclient.Make(clientConfig, address, registryv1alpha1connect.NewRepositoryServiceClient)
	}
}

// NewModuleReader creates a new module reader using content addressable storage.
func NewModuleReader(
	logger *zap.Logger,
	verbosePrinter verbose.Printer,
	bucket storage.ReadWriteBucket,
	delegate bufmodule.ModuleReader,
	repositoryClientFactory RepositoryServiceClientFactory,
) bufmodule.ModuleReader {
	return newCASModuleReader(
		bucket,
		delegate,
		repositoryClientFactory,
		logger,
		verbosePrinter,
	)
}
