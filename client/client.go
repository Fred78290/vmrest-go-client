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

package client

import (
	"github.com/Fred78290/vmrest-go-client/client/api"
	"github.com/Fred78290/vmrest-go-client/client/model"
)

// APIClient manages communication with the VMware Workstation REST API API v1.3.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	Client                        api.Client
	_hostNetworksManagementApi    api.HostNetworksManagementApiService
	_managementApi                api.ManagementApiService
	_networkAdaptersManagementApi api.NetworkAdaptersManagementApiService
	_powerManagementApi           api.PowerManagementApiService
	_sharedFoldersManagementApi   api.SharedFoldersManagementApiService
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) (*APIClient, error) {

	if c, err := NewHttpClient(cfg.Endpoint, cfg.UserAgent, cfg.UserName, cfg.Password, cfg.Timeout, cfg.UnsecureTLS); err != nil {
		return nil, err
	} else {
		return &APIClient{
			Client: c,
		}, nil

	}
}

func (c *APIClient) hostNetworksManagementApi() api.HostNetworksManagementApiService {
	if c._hostNetworksManagementApi == nil {
		c._hostNetworksManagementApi = api.NewHostNetworksManagementApiService(c.Client)
	}

	return c._hostNetworksManagementApi
}

func (c *APIClient) managementApi() api.ManagementApiService {
	if c._managementApi == nil {
		c._managementApi = api.NewManagementApiService(c.Client)
	}

	return c._managementApi
}

func (c *APIClient) networkAdaptersManagementApi() api.NetworkAdaptersManagementApiService {
	if c._networkAdaptersManagementApi == nil {
		c._networkAdaptersManagementApi = api.NewNetworkAdaptersManagementApiService(c.Client)
	}

	return c._networkAdaptersManagementApi
}

func (c *APIClient) powerManagementApi() api.PowerManagementApiService {
	if c._powerManagementApi == nil {
		c._powerManagementApi = api.NewPowerManagementApiService(c.Client)
	}

	return c._powerManagementApi
}

func (c *APIClient) sharedFoldersManagementApi() api.SharedFoldersManagementApiService {
	if c._sharedFoldersManagementApi == nil {
		c._sharedFoldersManagementApi = api.NewSharedFoldersManagementApiService(c.Client)
	}

	return c._sharedFoldersManagementApi
}

/*
CreateNetwork Creates a virtual network
  - @param parameters Host network to be created

@return Network
*/
func (c *APIClient) CreateNetwork(parameters *model.CreateVmnetParameter) (*model.Network, error) {
	return c.hostNetworksManagementApi().CreateNetwork(parameters)
}

/*
DeletePortforward Deletes port forwarding
  - @param vmnet NAT type of virtual network
  - @param protocol Protocol type: tcp, udp
  - @param port Host port number
*/
func (c *APIClient) DeletePortforward(vmnet string, protocol string, port int) error {
	return c.hostNetworksManagementApi().DeletePortforward(vmnet, protocol, port)
}

/*
GetAllNetworks Returns all virtual networks

@return Networks
*/
func (c *APIClient) GetAllNetworks() (*model.Networks, error) {
	return c.hostNetworksManagementApi().GetAllNetworks()
}

/*
GetMACToIPs Returns all MAC-to-IP settings for DHCP service
  - @param vmnet Virtual network that has DHCP enabled

@return MactoIps
*/
func (c *APIClient) GetMACToIPs(vmnet string) (*model.MactoIps, error) {
	return c.hostNetworksManagementApi().GetMACToIPs(vmnet)
}

/*
GetPortforwards Returns all port forwardings
  - @param vmnet NAT type of virtual network

@return Portforwards
*/
func (c *APIClient) GetPortforwards(vmnet string) (*model.Portforwards, error) {
	return c.hostNetworksManagementApi().GetPortforwards(vmnet)
}

/*
UpdateMacToIP Updates the MAC-to-IP binding
  - @param vmnet Virtual network that enabled DHCP
  - @param mac Mac address that want to be mapped with a given IP
  - @param parameters IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted

@return ErrorModel
*/
func (c *APIClient) UpdateMacToIP(vmnet string, mac string, parameters *model.MacToIpParameter) (*model.ErrorModel, error) {
	return c.hostNetworksManagementApi().UpdateMacToIP(vmnet, mac, parameters)
}

/*
UpdatePortforward Updates port forwarding
  - @param vmnet NAT type of virtual network
  - @param protocol Protocol type: tcp, udp
  - @param port Host port number
  - @param parameters Guest to forward to

@return ErrorModel
*/
func (c *APIClient) UpdatePortforward(vmnet string, protocol string, port int, parameters *model.PortforwardParameter) (*model.ErrorModel, error) {
	return c.hostNetworksManagementApi().UpdatePortforward(vmnet, protocol, port, parameters)
}

/*
ConfigVMParams update the vm config params
  - @param id ID of VM
  - @param parameters Parameters set to the VM

@return ErrorModel
*/
func (c *APIClient) ConfigVMParams(id string, parameters *model.ConfigVmParamsParameter) (*model.ErrorModel, error) {
	return c.managementApi().ConfigVMParams(id, parameters)
}

/*
CreateVM Creates a copy of the VM
  - @param params Parameters of VM to create

@return VmInformation
*/
func (c *APIClient) CreateVM(parameters *model.VmCloneParameter) (*model.VmInformation, error) {
	return c.managementApi().CreateVM(parameters)
}

/*
DeleteVM Deletes a VM
  - @param id ID of VM
*/
func (c *APIClient) DeleteVM(id string) error {
	return c.managementApi().DeleteVM(id)
}

/*
GetAllVMs Returns a list of VM IDs and paths for all VMs

@return []Vmid
*/
func (c *APIClient) GetAllVMs() ([]model.Vmid, error) {
	return c.managementApi().GetAllVMs()
}

/*
GetVM Returns the VM setting information of a VM
  - @param id ID of VM

@return VmInformation
*/
func (c *APIClient) GetVM(id string) (*model.VmInformation, error) {
	return c.managementApi().GetVM(id)
}

/*
GetVMParams Get the VM config params
  - @param id ID of VM
  - @param name Name of the param

@return ConfigVmParamsParameter
*/
func (c *APIClient) GetVMParams(id string, name string) (*model.ConfigVmParamsParameter, error) {
	return c.managementApi().GetVMParams(id, name)
}

/*
GetVMRestrictions Returns the restrictions information of the VM
  - @param id ID of VM

@return VmRestrictionsInformation
*/
func (c *APIClient) GetVMRestrictions(id string) (*model.VmRestrictionsInformation, error) {
	return c.managementApi().GetVMRestrictions(id)
}

/*
RegisterVM Register VM to VM Library
  - @param parameters Parameters of the VM to register

@return VmRegistrationInformation
*/
func (c *APIClient) RegisterVM(parameters *model.VmRegisterParameter) (*model.VmRegistrationInformation, error) {
	return c.managementApi().RegisterVM(parameters)
}

/*
UpdateVM Updates the VM settings
  - @param id ID of VM
  - @param parameters VM definition

@return VmInformation
*/
func (c *APIClient) UpdateVM(id string, parameters *model.VmParameter) (*model.VmInformation, error) {
	return c.managementApi().UpdateVM(id, parameters)
}

/*
CreateNICDevice Creates a network adapter in the VM
  - @param id ID of VM
  - @param parameters Parameters of network adapter to create

@return NicDevice
*/
func (c *APIClient) CreateNICDevice(id string, parameters *model.NicDeviceParameter) (*model.NicDevice, error) {
	return c.networkAdaptersManagementApi().CreateNICDevice(id, parameters)
}

/*
DeleteNICDevice Deletes a VM network adapter
  - @param id ID of VM
  - @param index Index of VM network adapter
*/
func (c *APIClient) DeleteNICDevice(id string, index int) error {
	return c.networkAdaptersManagementApi().DeleteNICDevice(id, index)
}

/*
GetAllNICDevices Returns all network adapters in the VM
  - @param id ID of VM

@return NicDevices
*/
func (c *APIClient) GetAllNICDevices(id string) (*model.NicDevices, error) {
	return c.networkAdaptersManagementApi().GetAllNICDevices(id)
}

/*
GetIPAddress Returns the IP address of a VM
  - @param id ID of VM

@return InlineResponse200
*/
func (c *APIClient) GetIPAddress(id string) (*model.InlineResponse200, error) {
	return c.networkAdaptersManagementApi().GetIPAddress(id)
}

/*
GetNicInfo Returns the IP stack configuration of all NICs of a VM
  - @param id ID of VM

@return NicIpStackAll
*/
func (c *APIClient) GetNicInfo(id string) (*model.NicIpStackAll, error) {
	return c.networkAdaptersManagementApi().GetNicInfo(id)
}

/*
UpdateNICDevice Updates a network adapter in the VM
  - @param id ID of VM
  - @param index Index of VM network adapter
  - @param parameters Parameters of network adapter to update to

@return NicDevice
*/
func (c *APIClient) UpdateNICDevice(id string, index int, parameters *model.NicDeviceParameter) (*model.NicDevice, error) {
	return c.networkAdaptersManagementApi().UpdateNICDevice(id, index, parameters)
}

/*
ChangePowerState Changes the VM power state
  - @param id ID of VM
  - @param operation VM power operation: on, off, shutdown, suspend, pause, unpause

@return VmPowerState
*/
func (c *APIClient) ChangePowerState(id string, operation model.VmPowerOperation) (*model.VmPowerState, error) {
	return c.powerManagementApi().ChangePowerState(id, operation)
}

/*
GetPowerState Returns the power state of the VM
  - @param id ID of VM

@return VmPowerState
*/
func (c *APIClient) GetPowerState(id string) (*model.VmPowerState, error) {
	return c.powerManagementApi().GetPowerState(id)
}

/*
CreateSharedFolder Mounts a new shared folder in the VM
  - @param id ID of VM
  - @param parameters Parameters of the shared folder to mount

@return SharedFolders
*/
func (c *APIClient) CreateSharedFolder(id string, parameters *model.SharedFolder) (model.SharedFolders, error) {
	return c.sharedFoldersManagementApi().CreateSharedFolder(id, parameters)
}

/*
DeleteSharedFolder Deletes a shared folder
  - @param id ID of VM
  - @param folderId ID of shared folder
*/
func (c *APIClient) DeleteSharedFolder(id string, folderId string) error {
	return c.sharedFoldersManagementApi().DeleteSharedFolder(id, folderId)
}

/*
GetAllSharedFolders Returns all shared folders mounted in the VM
  - @param id ID of VM

@return SharedFolders
*/
func (c *APIClient) GetAllSharedFolders(id string) (model.SharedFolders, error) {
	return c.sharedFoldersManagementApi().GetAllSharedFolders(id)
}

/*
UpdataSharedFolder Updates a shared folder mounted in the VM
  - @param id ID of VM
  - @param folderId ID of VM shared folder
  - @param parameters Parameters of the shared folder to update to

@return SharedFolders
*/
func (c *APIClient) UpdateSharedFolder(id string, folderId string, parameters *model.SharedFolderParameter) (model.SharedFolders, error) {
	return c.sharedFoldersManagementApi().UpdateSharedFolder(id, folderId, parameters)
}
