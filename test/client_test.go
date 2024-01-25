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

package vmrest_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/Fred78290/vmrest-go-client/client"
	"github.com/Fred78290/vmrest-go-client/client/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	vmnet       = "vmnet4"
	forwardport = 2022
	protocol    = "tcp"
	vmid        = "GUIFEMR099KA2BUH3CRRAIVC04HDPADP"
	home        = "/home/ubuntu"
	folderid    = "12345"
	macaddress  = "00:0c:29:c3:8e:32"
	ip          = "172.16.134.128"
)

type mockVMRestClient struct {
	mock.Mock
	currentTest *testing.T
}

func newMockVMRestClient(t *testing.T) *mockVMRestClient {
	return &mockVMRestClient{
		currentTest: t,
	}
}

func (c *mockVMRestClient) Post(endpoint string, input interface{}, output interface{}) error {
	log.Infof("POST: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockVMRestClient) Patch(endpoint string, input interface{}, output interface{}) error {
	log.Infof("PATCH: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockVMRestClient) Put(endpoint string, input interface{}, output interface{}) error {
	log.Infof("PUT: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockVMRestClient) Get(endpoint string, output interface{}) error {
	log.Infof("GET: %s", endpoint)
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockVMRestClient) Delete(endpoint string, output interface{}) error {
	log.Infof("DELETE: %s", endpoint)
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func newTestEnv(t *testing.T) (*assert.Assertions, *mockVMRestClient, *client.APIClient) {
	mock := newMockVMRestClient(t)
	return assert.New(t), mock, &client.APIClient{
		Client: mock,
	}

}
func TestAPIClient_CreateNetwork(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)

	expected := &model.Network{
		Name:   vmnet,
		Type:   "nat",
		Dhcp:   "true",
		Subnet: "172.16.134.0",
		Mask:   "255.255.255.0",
	}
	parameters := &model.CreateVmnetParameter{
		Name: vmnet,
		Type: "nat",
	}

	mock.On("Post", "/api/vmnets", parameters).Return(expected, nil).Once()

	result, err := vmrest.CreateNetwork(parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)

	mock.AssertExpectations(t)
}

func TestAPIClient_DeletePortforward(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	vmnet := vmnet
	protocol := "tcp"
	port := 2022

	mock.On("Delete", fmt.Sprintf("/api/vmnet/%s/portforward/%s/%d", vmnet, protocol, port)).Return(nil, nil).Once()

	err := vmrest.DeletePortforward(vmnet, protocol, port)

	assert.NoError(err)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetAllNetworks(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.Networks{
		Num: 5,
		Vmnets: []model.Network{
			{
				Name: "vmnet0",
				Type: "bridged",
				Dhcp: "false",
			},
			{
				Name:   "vmnet1",
				Type:   "hostOnly",
				Dhcp:   "true",
				Subnet: "192.168.38.0",
				Mask:   "255.255.255.0",
			},
			{
				Name: "vmnet2",
				Type: "bridged",
				Dhcp: "false",
			},
			{
				Name: "vmnet3",
				Type: "bridged",
				Dhcp: "false",
			},
			{
				Name:   "vmnet8",
				Type:   "nat",
				Dhcp:   "true",
				Subnet: "172.16.134.0",
				Mask:   "255.255.255.0",
			},
		},
	}
	mock.On("Get", "/api/vmnet").Return(expected, nil).Once()

	result, err := vmrest.GetAllNetworks()

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetMACToIPs(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)

	expected := model.MactoIps{
		Num: 1,
		Mactoips: []model.MactoIp{
			{
				Vmnet: vmnet,
				Mac:   "52:54:00:d2:bc:22",
				Ip:    ip,
			},
		},
	}
	mock.On("Get", "/api/vmnet/vmnet4/mactoip").Return(expected, nil).Once()

	result, err := vmrest.GetMACToIPs(vmnet)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetPortforwards(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.Portforwards{
		Num: 1,
		PortForwardings: []model.Portforward{
			{
				Port:     2022,
				Protocol: "tcp",
				Desc:     "",
				Guest: &model.PortforwardGuest{
					Ip:   ip,
					Port: 22,
				},
			},
		},
	}

	mock.On("Get", "/api/vmnet/vmnet4/portforward").Return(expected, nil).Once()

	result, err := vmrest.GetPortforwards(vmnet)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_UpdateMacToIP(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.ErrorModel{
		Code:    0,
		Message: "OK",
	}

	parameters := &model.MacToIpParameter{
		IP: ip,
	}

	mock.On("Put", fmt.Sprintf("/api/vmnet/vmnet4/mactoip/%s", url.PathEscape("52:54:00:d2:bc:22")), parameters).Return(expected, nil).Once()

	result, err := vmrest.UpdateMacToIP(vmnet, "52:54:00:d2:bc:22", parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_UpdatePortforward(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.ErrorModel{
		Code:    0,
		Message: "OK",
	}
	parameters := &model.PortforwardParameter{
		GuestIp:   ip,
		GuestPort: 22,
		Desc:      "",
	}

	mock.On("Put", "/api/vmnet/vmnet4/portforward/tcp/2022", parameters).Return(expected, nil).Once()

	result, err := vmrest.UpdatePortforward(vmnet, "tcp", 2022, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_ConfigVMParams(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.ErrorModel{
		Code:    0,
		Message: "OK",
	}
	parameters := &model.ConfigVmParamsParameter{
		Name:  "displayName",
		Value: "Sample VM",
	}

	mock.On("Put", fmt.Sprintf("/api/vms/%s/params", vmid), parameters).Return(expected, nil).Once()

	result, err := vmrest.ConfigVMParams(vmid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_CreateVM(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := model.VmInformation{
		Id:     "5QRDNB7JF25QORUUUHD408811D75BDVV",
		Memory: 2048,
		Cpu: &model.Vmcpu{
			Processors: 4,
		},
	}
	parameters := &model.VmCloneParameter{
		Name:     "Sample VM",
		ParentId: "2PC9UFLB400IJ0TOB044TQGP8KLVUJ1K",
	}

	mock.On("Post", "/api/vms", parameters).Return(expected, nil).Once()

	result, err := vmrest.CreateVM(parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(*result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_DeleteVM(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)

	mock.On("Delete", "/api/vms/43a5dfb2-0049-4a1c-87fd-abae9146b65f").Return(nil, nil).Once()

	err := vmrest.DeleteVM("43a5dfb2-0049-4a1c-87fd-abae9146b65f")

	assert.NoError(err)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetAllVMs(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := []model.Vmid{
		{
			Id:   vmid,
			Path: fmt.Sprintf("/home/ubuntu/VM/%s", vmid),
		},
	}

	mock.On("Get", "/api/vms").Return(expected, nil).Once()

	result, err := vmrest.GetAllVMs()

	assert.NoError(err)
	assert.NotNil(result)
	assert.ElementsMatch(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetVM(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmInformation{
		Id:     vmid,
		Memory: 2048,
		Cpu: &model.Vmcpu{
			Processors: 2,
		},
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetVM(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetVMParams(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.ConfigVmParamsParameter{
		Name:  "displayName",
		Value: "Sample VM",
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/params/displayName", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetVMParams(vmid, "displayName")

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetVMRestrictions(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmRestrictionsInformation{
		Id:     vmid,
		Memory: 2048,
		Cpu: &model.Vmcpu{
			Processors: 2,
		},
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/restrictions", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetVMRestrictions(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_RegisterVM(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmRegistrationInformation{
		Id:   "5QRDNB7JF25QORUUUHD408811D75BDVV",
		Path: "/vmware/Sample VM/Sample VM.vmx",
	}
	parameters := &model.VmRegisterParameter{
		Name: "Sample VM",
		Path: "/vmware/Sample VM/Sample VM.vmx",
	}

	mock.On("Post", "/api/vms/registration", parameters).Return(expected, nil).Once()

	result, err := vmrest.RegisterVM(parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_UpdateVM(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmInformation{
		Id:     vmid,
		Memory: 4096,
		Cpu: &model.Vmcpu{
			Processors: 4,
		},
	}
	parameters := &model.VmParameter{
		Processors: 4,
		Memory:     4096,
	}

	mock.On("Put", fmt.Sprintf("/api/vms/%s", vmid), parameters).Return(expected, nil).Once()

	result, err := vmrest.UpdateVM(vmid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_CreateNICDevice(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.NicDevice{
		Index:      1,
		Type:       "vmxnet3",
		Vmnet:      vmnet,
		MacAddress: macaddress,
	}
	parameters := &model.NicDeviceParameter{
		Type:  "vmxnet3",
		Vmnet: vmnet,
	}

	mock.On("Post", fmt.Sprintf("/api/vms/%s/nic", vmid), parameters).Return(expected, nil).Once()

	result, err := vmrest.CreateNICDevice(vmid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_DeleteNICDevice(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)

	mock.On("Delete", fmt.Sprintf("/api/vms/%s/nic/%d", vmid, 1)).Return(nil, nil).Once()

	err := vmrest.DeleteNICDevice(vmid, 1)

	assert.NoError(err)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetAllNICDevices(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.NicDevices{
		Num: 1,
		Nics: []model.NicDevice{
			{
				Index:      1,
				Type:       "vmxnet3",
				Vmnet:      vmnet,
				MacAddress: macaddress,
			},
		},
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/nic", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetAllNICDevices(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetIPAddress(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.InlineResponse200{
		Ip: ip,
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/ip", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetIPAddress(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetNicInfo(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.NicIpStackAll{
		Nics: []model.NicIpStack{
			{
				Mac: macaddress,
				Ip: []string{
					"fe80::b7ae:c029:5bce:35f7/64",
					ip + "/24",
				},
				Dns: &model.DnsConfig{
					Domainname: "localdomain",
					Server: []string{
						"172.16.134.2",
					},
				},
				Dhcp4: &model.DhcpConfig{
					Enabled: true,
					Setting: "",
				},
				Dhcp6: &model.DhcpConfig{
					Enabled: false,
					Setting: "",
				},
			},
			{
				Mac: "e8:48:b8:c8:20:00",
				Ip: []string{
					"fe80::3ee3:8dd0:34a0:aebd/64",
					"169.254.11.25/16",
				},
				Dns: &model.DnsConfig{
					Server: []string{
						"fec0:0:0:ffff::1",
						"fec0:0:0:ffff::2",
						"fec0:0:0:ffff::3",
					},
				},
				Dhcp4: &model.DhcpConfig{
					Enabled: true,
					Setting: "",
				},
				Dhcp6: &model.DhcpConfig{
					Enabled: false,
					Setting: "",
				},
			},
		},
		Dns: &model.DnsConfig{
			Hostname:   "DESKTOP-GR412QE",
			Domainname: "",
			Server: []string{
				"172.16.134.2",
			},
		},
		Routes: []model.RouteEntry{
			{
				Dest:      "0.0.0.0",
				Prefix:    0,
				Nexthop:   "172.16.134.2",
				Interface: 0,
				Type:      0,
				Metric:    0,
			},
			{
				Dest:      "172.16.134.0",
				Prefix:    24,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      ip,
				Prefix:    32,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "172.16.134.255",
				Prefix:    32,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "224.0.0.0",
				Prefix:    4,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "255.255.255.255",
				Prefix:    32,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "255.255.255.255",
				Prefix:    32,
				Interface: 1,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "fe80::",
				Prefix:    64,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "fe80::",
				Prefix:    64,
				Interface: 1,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "fe80::3ee3:8dd0:34a0:aebd",
				Prefix:    128,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "fe80::b7ae:c029:5bce:35f7",
				Prefix:    128,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "ff00::",
				Prefix:    8,
				Interface: 0,
				Type:      0,
				Metric:    256,
			},
			{
				Dest:      "ff00::",
				Prefix:    8,
				Interface: 1,
				Type:      0,
				Metric:    256,
			},
		},
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/nicips", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetNicInfo(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_UpdateNICDevice(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.NicDevice{
		Index:      1,
		Type:       "vmxnet3",
		Vmnet:      vmnet,
		MacAddress: macaddress,
	}
	parameters := &model.NicDeviceParameter{
		Type:  "vmxnet3",
		Vmnet: vmnet,
	}

	mock.On("Put", fmt.Sprintf("/api/vms/%s/nic/%d", vmid, 1), parameters).Return(expected, nil).Once()

	result, err := vmrest.UpdateNICDevice(vmid, 1, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_ChangePowerState(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmPowerState{
		PowerState: string(model.VM_OFF),
	}
	parameters := model.VM_OFF

	mock.On("Put", fmt.Sprintf("/api/vms/%s/power", vmid), parameters).Return(expected, nil).Once()

	result, err := vmrest.ChangePowerState(vmid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetPowerState(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := &model.VmPowerState{
		PowerState: string(model.VM_OFF),
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/power", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetPowerState(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_CreateSharedFolder(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := []model.SharedFolder{
		{
			FolderId: folderid,
			HostPath: home,
			Flags:    0600,
		},
	}
	parameters := &model.SharedFolder{
		HostPath: home,
		Flags:    0600,
	}

	mock.On("Post", fmt.Sprintf("/api/vms/%s/sharedfolders", vmid), parameters).Return(expected, nil).Once()

	result, err := vmrest.CreateSharedFolder(vmid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.ElementsMatch(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_DeleteSharedFolder(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)

	mock.On("Delete", fmt.Sprintf("/api/vms/%s/sharedfolders/%s", vmid, folderid)).Return(nil, nil).Once()

	err := vmrest.DeleteSharedFolder(vmid, folderid)

	assert.NoError(err)
	mock.AssertExpectations(t)
}

func TestAPIClient_GetAllSharedFolders(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := []model.SharedFolder{
		{
			FolderId: folderid,
			HostPath: home,
			Flags:    0600,
		},
	}

	mock.On("Get", fmt.Sprintf("/api/vms/%s/sharedfolders", vmid)).Return(expected, nil).Once()

	result, err := vmrest.GetAllSharedFolders(vmid)

	assert.NoError(err)
	assert.NotNil(result)
	assert.ElementsMatch(result, expected)
	mock.AssertExpectations(t)
}

func TestAPIClient_UpdateSharedFolder(t *testing.T) {
	assert, mock, vmrest := newTestEnv(t)
	expected := []model.SharedFolder{
		{
			FolderId: folderid,
			HostPath: home,
			Flags:    0600,
		},
	}
	parameters := &model.SharedFolderParameter{
		HostPath: home,
		Flags:    0600,
	}

	mock.On("Put", fmt.Sprintf("/api/vms/%s/sharedfolders/%s", vmid, folderid), parameters).Return(expected, nil).Once()

	result, err := vmrest.UpdateSharedFolder(vmid, folderid, parameters)

	assert.NoError(err)
	assert.NotNil(result)
	assert.ElementsMatch(result, expected)
	mock.AssertExpectations(t)
}
