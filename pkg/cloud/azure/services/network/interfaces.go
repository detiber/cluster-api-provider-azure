/*
Copyright 2018 The Kubernetes Authors.
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

package network

import "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-01-01/network"

// DeleteNetworkInterface deletes the NIC resource.
func (s *Service) DeleteNetworkInterface(resourceGroup string, networkInterfaceName string) (network.InterfacesDeleteFuture, error) {
	return s.InterfacesClient.Delete(s.ctx, resourceGroup, networkInterfaceName)
}

// WaitForNetworkInterfacesDeleteFuture waits for the DeleteNetworkInterface operation to complete.
func (s *Service) WaitForNetworkInterfacesDeleteFuture(future network.InterfacesDeleteFuture) error {
	return future.Future.WaitForCompletionRef(s.ctx, s.InterfacesClient.Client)
}
