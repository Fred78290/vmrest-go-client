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

const (
	kApiVM = "/api/vms/%s"
)

type managementApiService struct {
	client Client
}

// NewManagementApiService create service for ManagementApiService
func NewManagementApiService(client Client) ManagementApiService {
	return &managementApiService{
		client: client,
	}
}

// ConfigVMParams update the vm config params
//   - @param id ID of VM
//   - @param parameters Parameters set to the VM
//
// @return ErrorModel
func (a *managementApiService) ConfigVMParams(id string, parameters *model.ConfigVmParamsParameter) (*model.ErrorModel, error) {
	var returnValue model.ErrorModel

	if err := a.client.Put(fmt.Sprintf("/api/vms/%s/params", id), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// CreateVM Creates a copy of the VM
//   - @param params Parameters of VM to create
//
// @return VmInformation
func (a *managementApiService) CreateVM(parameters *model.VmCloneParameter) (*model.VmInformation, error) {
	var returnValue model.VmInformation

	if err := a.client.Post("/api/vms", parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// DeleteVM Deletes a VM
//   - @param id ID of VM
func (a *managementApiService) DeleteVM(id string) error {
	return a.client.Delete(fmt.Sprintf(kApiVM, id), nil)
}

// GetAllVMs Returns a list of VM IDs and paths for all VMs
// @return []Vmid
func (a *managementApiService) GetAllVMs() ([]model.Vmid, error) {
	var returnValue []model.Vmid

	if err := a.client.Get("/api/vms", &returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, err
	}
}

// GetVM Returns the VM setting information of a VM
//   - @param id ID of VM
//
// @return VmInformation
func (a *managementApiService) GetVM(id string) (*model.VmInformation, error) {
	var returnValue model.VmInformation

	if err := a.client.Get(fmt.Sprintf(kApiVM, id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// GetVMParams Get the VM config params
//   - @param id ID of VM
//   - @param name Name of the param
//
// @return ConfigVmParamsParameter
func (a *managementApiService) GetVMParams(id string, name string) (*model.ConfigVmParamsParameter, error) {
	var returnValue model.ConfigVmParamsParameter

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/params/%s", id, url.PathEscape(name)), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// GetVMRestrictions Returns the restrictions information of the VM
//   - @param id ID of VM
//
// @return VmRestrictionsInformation
func (a *managementApiService) GetVMRestrictions(id string) (*model.VmRestrictionsInformation, error) {
	var returnValue model.VmRestrictionsInformation

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/restrictions", id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// RegisterVM Register VM to VM Library
//   - @param parameters Parameters of the VM to register
//
// @return VmRegistrationInformation
func (a *managementApiService) RegisterVM(parameters *model.VmRegisterParameter) (*model.VmRegistrationInformation, error) {
	var returnValue model.VmRegistrationInformation

	if err := a.client.Post("/api/vms/registration", parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

// UpdateVM Updates the VM settings
//   - @param id ID of VM
//   - @param parameters VM definition
//
// @return VmInformation
func (a *managementApiService) UpdateVM(id string, parameters *model.VmParameter) (*model.VmInformation, error) {
	var returnValue model.VmInformation

	// body params
	if err := a.client.Put(fmt.Sprintf(kApiVM, id), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}
