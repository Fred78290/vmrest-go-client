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

	"github.com/Fred78290/vmrest-go-client/client/model"
)

type networkAdaptersManagementApiService struct {
	client Client
}

func NewNetworkAdaptersManagementApiService(client Client) NetworkAdaptersManagementApiService {
	return &networkAdaptersManagementApiService{
		client: client,
	}
}

/*
CreateNICDevice Creates a network adapter in the VM
  - @param id ID of VM
  - @param parameters Parameters of network adapter to create

@return NicDevice
*/
func (a *networkAdaptersManagementApiService) CreateNICDevice(id string, parameters *model.NicDeviceParameter) (*model.NicDevice, error) {
	var returnValue model.NicDevice

	if err := a.client.Post(fmt.Sprintf("/api/vms/%s/nic", id), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

/*
DeleteNICDevice Deletes a VM network adapter
  - @param id ID of VM
  - @param index Index of VM network adapter
*/
func (a *networkAdaptersManagementApiService) DeleteNICDevice(id string, index int) error {
	if err := a.client.Delete(fmt.Sprintf("/api/vms/%s/nic/%d", id, index), nil); err != nil {
		return err
	} else {
		return err
	}
}

/*
GetAllNICDevices Returns all network adapters in the VM
  - @param id ID of VM

@return NicDevices
*/
func (a *networkAdaptersManagementApiService) GetAllNICDevices(id string) (*model.NicDevices, error) {
	var returnValue model.NicDevices

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/nic", id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

/*
GetIPAddress Returns the IP address of a VM
  - @param id ID of VM

@return InlineResponse200
*/
func (a *networkAdaptersManagementApiService) GetIPAddress(id string) (*model.InlineResponse200, error) {
	var returnValue model.InlineResponse200

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/ip", id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

/*
GetNicInfo Returns the IP stack configuration of all NICs of a VM
  - @param id ID of VM

@return NicIpStackAll
*/
func (a *networkAdaptersManagementApiService) GetNicInfo(id string) (*model.NicIpStackAll, error) {
	var returnValue model.NicIpStackAll

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/nicips", id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

/*
UpdateNICDevice Updates a network adapter in the VM
  - @param id ID of VM
  - @param index Index of VM network adapter
  - @param parameters Parameters of network adapter to update to

@return NicDevice
*/
func (a *networkAdaptersManagementApiService) UpdateNICDevice(id string, index int, parameters *model.NicDeviceParameter) (*model.NicDevice, error) {
	var returnValue model.NicDevice

	if err := a.client.Put(fmt.Sprintf("/api/vms/%s/nic/%d", id, index), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}
