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
	"github.com/Fred78290/vmrest-go-client/client/model"
)

type Client interface {
	Patch(string, interface{}, interface{}) error
	Post(string, interface{}, interface{}) error
	Put(string, interface{}, interface{}) error
	Get(string, interface{}) error
	Delete(string, interface{}) error
}

type HostNetworksManagementApiService interface {
	/*
	   CreateNetwork Creates a virtual network
	     - @param parameters Host network to be created

	   @return Network
	*/
	CreateNetwork(parameters *model.CreateVmnetParameter) (*model.Network, error)

	/*
	   DeletePortforward Deletes port forwarding
	     - @param vmnet NAT type of virtual network
	     - @param protocol Protocol type: tcp, udp
	     - @param port Host port number
	*/
	DeletePortforward(vmnet string, protocol string, port int) error

	/*
	   GetAllNetworks Returns all virtual networks

	   @return Networks
	*/
	GetAllNetworks() (*model.Networks, error)

	/*
	   GetMACToIPs Returns all MAC-to-IP settings for DHCP service
	     - @param vmnet Virtual network that has DHCP enabled

	   @return MactoIps
	*/
	GetMACToIPs(vmnet string) (*model.MactoIps, error)

	/*
	   GetPortforwards Returns all port forwardings
	     - @param vmnet NAT type of virtual network

	   @return Portforwards
	*/
	GetPortforwards(vmnet string) (*model.Portforwards, error)

	/*
	   UpdateMacToIP Updates the MAC-to-IP binding
	     - @param vmnet Virtual network that enabled DHCP
	     - @param mac Mac address that want to be mapped with a given IP
	     - @param parameters IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted

	   @return ErrorModel
	*/
	UpdateMacToIP(vmnet string, mac string, parameters *model.MacToIpParameter) (*model.ErrorModel, error)

	/*
	   UpdatePortforward Updates port forwarding
	     - @param vmnet NAT type of virtual network
	     - @param protocol Protocol type: tcp, udp
	     - @param port Host port number
	     - @param parameters Guest to forward to

	   @return ErrorModel
	*/
	UpdatePortforward(vmnet string, protocol string, port int, parameters *model.PortforwardParameter) (*model.ErrorModel, error)
}

type ManagementApiService interface {
	/*
	   ConfigVMParams update the vm config params
	     - @param id ID of VM
	     - @param parameters Parameters set to the VM

	   @return ErrorModel
	*/
	ConfigVMParams(id string, parameters *model.ConfigVmParamsParameter) (*model.ErrorModel, error)

	/*
	   CreateVM Creates a copy of the VM
	     - @param params Parameters of VM to create

	   @return VmInformation
	*/
	CreateVM(parameters *model.VmCloneParameter) (*model.VmInformation, error)

	/*
	   DeleteVM Deletes a VM
	     - @param id ID of VM
	*/
	DeleteVM(id string) error

	/*
	   GetAllVMs Returns a list of VM IDs and paths for all VMs

	   @return []Vmid
	*/
	GetAllVMs() ([]model.Vmid, error)

	/*
	   VMManagementApiService Returns the VM setting information of a VM
	     - @param id ID of VM

	   @return VmInformation
	*/
	GetVM(id string) (*model.VmInformation, error)

	/*
	   GetVMParams Get the VM config params
	     - @param id ID of VM
	     - @param name Name of the param

	   @return ConfigVmParamsParameter
	*/
	GetVMParams(id string, name string) (*model.ConfigVmParamsParameter, error)

	/*
	   GetVMRestrictions Returns the restrictions information of the VM
	     - @param id ID of VM

	   @return VmRestrictionsInformation
	*/
	GetVMRestrictions(id string) (*model.VmRestrictionsInformation, error)

	/*
	   RegisterVM Register VM to VM Library
	     - @param parameters Parameters of the VM to register

	   @return VmRegistrationInformation
	*/
	RegisterVM(parameters *model.VmRegisterParameter) (*model.VmRegistrationInformation, error)

	/*
	   UpdateVM Updates the VM settings
	     - @param id ID of VM
	     - @param parameters VM definition

	   @return VmInformation
	*/
	UpdateVM(id string, parameters *model.VmParameter) (*model.VmInformation, error)
}

type PowerManagementApiService interface {
	/*
	   ChangePowerState Changes the VM power state
	     - @param id ID of VM
	     - @param operation VM power operation: on, off, shutdown, suspend, pause, unpause

	   @return VmPowerState
	*/
	ChangePowerState(id string, operation model.VmPowerOperation) (*model.VmPowerState, error)

	/*
	   GetPowerState Returns the power state of the VM
	     - @param id ID of VM

	   @return VmPowerState
	*/
	GetPowerState(id string) (*model.VmPowerState, error)
}

type NetworkAdaptersManagementApiService interface {
	/*
	   CreateNICDevice Creates a network adapter in the VM
	     - @param id ID of VM
	     - @param parameters Parameters of network adapter to create

	   @return NicDevice
	*/
	CreateNICDevice(id string, parameters *model.NicDeviceParameter) (*model.NicDevice, error)

	/*
	   DeleteNICDevice Deletes a VM network adapter
	     - @param id ID of VM
	     - @param index Index of VM network adapter
	*/
	DeleteNICDevice(id string, index int) error

	/*
	   GetAllNICDevices Returns all network adapters in the VM
	     - @param id ID of VM

	   @return NicDevices
	*/
	GetAllNICDevices(id string) (*model.NicDevices, error)

	/*
	   GetIPAddress Returns the IP address of a VM
	     - @param id ID of VM

	   @return InlineResponse200
	*/
	GetIPAddress(id string) (*model.InlineResponse200, error)

	/*
	   VMNetworkAdaptersManagementApiService Returns the IP stack configuration of all NICs of a VM
	     - @param id ID of VM

	   @return NicIpStackAll
	*/
	GetNicInfo(id string) (*model.NicIpStackAll, error)

	/*
	   UpdateNICDevice Updates a network adapter in the VM
	     - @param id ID of VM
	     - @param index Index of VM network adapter
	     - @param parameters Parameters of network adapter to update to

	   @return NicDevice
	*/
	UpdateNICDevice(id string, index int, parameters *model.NicDeviceParameter) (*model.NicDevice, error)
}

type SharedFoldersManagementApiService interface {
	/*
	   CreateSharedFolder Mounts a new shared folder in the VM
	     - @param id ID of VM
	     - @param parameters Parameters of the shared folder to mount

	   @return SharedFolders
	*/
	CreateSharedFolder(id string, parameters *model.SharedFolder) (model.SharedFolders, error)

	/*
	   DeleteSharedFolder Deletes a shared folder
	     - @param id ID of VM
	     - @param folderId ID of shared folder
	*/
	DeleteSharedFolder(id string, folderId string) error

	/*
	   GetAllSharedFolders Returns all shared folders mounted in the VM
	     - @param id ID of VM

	   @return SharedFolders
	*/
	GetAllSharedFolders(id string) (model.SharedFolders, error)

	/*
	   UpdateSharedFolder Updates a shared folder mounted in the VM
	     - @param id ID of VM
	     - @param folderId ID of VM shared folder
	     - @param parameters Parameters of the shared folder to update to

	   @return SharedFolders
	*/
	UpdateSharedFolder(id string, folderId string, parameters *model.SharedFolderParameter) (model.SharedFolders, error)
}
