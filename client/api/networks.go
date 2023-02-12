/*
Copyright 2023 Fred78290.
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

package api

import (
	"fmt"
	"net/url"

	"github.com/Fred78290/vmrest-go-client/client/model"
)

type hostNetworksManagementApiService struct {
	client Client
}

// NewHostNetworksManagementApiService create service for HostNetworksManagementApiService
func NewHostNetworksManagementApiService(client Client) HostNetworksManagementApiService {
	return &hostNetworksManagementApiService{
		client: client,
	}
}

// CreateNetwork Creates a virtual network
//   - @param parameters Host network to be created
//
// @return Network
func (a *hostNetworksManagementApiService) CreateNetwork(parameters *model.CreateVmnetParameter) (*model.Network, error) {
	var returnValue model.Network

	// body params
	if err := a.client.Post("/api/vmnets", parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// DeletePortforward Deletes port forwarding
//   - @param vmnet NAT type of virtual network
//   - @param protocol Protocol type: tcp, udp
//   - @param port Host port number
func (a *hostNetworksManagementApiService) DeletePortforward(vmnet string, protocol string, port int) error {
	return a.client.Delete(fmt.Sprintf("/api/vmnet/%s/portforward/%s/%d", vmnet, protocol, port), nil)
}

// GetAllNetworks Returns all virtual networks
//
//	@return Networks
func (a *hostNetworksManagementApiService) GetAllNetworks() (*model.Networks, error) {
	var returnValue model.Networks

	if err := a.client.Get("/api/vmnet", &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

//		GetMACToIPs Returns all MAC-to-IP settings for DHCP service
//
//	 - @param vmnet Virtual network that has DHCP enabled
//
// @return MactoIps
func (a *hostNetworksManagementApiService) GetMACToIPs(vmnet string) (*model.MactoIps, error) {
	var returnValue model.MactoIps

	if err := a.client.Get(fmt.Sprintf("/api/vmnet/%s/mactoip", vmnet), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// GetPortforwards Returns all port forwardings
//   - @param vmnet NAT type of virtual network
//
// @return Portforwards
func (a *hostNetworksManagementApiService) GetPortforwards(vmnet string) (*model.Portforwards, error) {
	var returnValue model.Portforwards

	if err := a.client.Get(fmt.Sprintf("/api/vmnet/%s/portforward", vmnet), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// UpdateMacToIP Updates the MAC-to-IP binding
//   - @param vmnet Virtual network that enabled DHCP
//   - @param mac Mac address that want to be mapped with a given IP
//   - @param parameters IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted
//
// @return ErrorModel
func (a *hostNetworksManagementApiService) UpdateMacToIP(vmnet string, mac string, parameters *model.MacToIpParameter) (*model.ErrorModel, error) {
	var returnValue model.ErrorModel

	if err := a.client.Put(fmt.Sprintf("/api/vmnet/%s/mactoip/%s", vmnet, url.PathEscape(mac)), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// UpdatePortforward Updates port forwarding
//   - @param vmnet NAT type of virtual network
//   - @param protocol Protocol type: tcp, udp
//   - @param port Host port number
//   - @param parameters Guest to forward to
//
// @return ErrorModel
func (a *hostNetworksManagementApiService) UpdatePortforward(vmnet string, protocol string, port int, parameters *model.PortforwardParameter) (*model.ErrorModel, error) {
	var returnValue model.ErrorModel

	if err := a.client.Put(fmt.Sprintf("/api/vmnet/%s/portforward/%s/%d", vmnet, protocol, port), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}
