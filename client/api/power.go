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

type powerManagementApiService struct {
	client Client
}

func NewPowerManagementApiService(client Client) PowerManagementApiService {
	return &powerManagementApiService{
		client: client,
	}
}

/*
ChangePowerState Changes the VM power state
  - @param id ID of VM
  - @param operation VM power operation: on, off, shutdown, suspend, pause, unpause

@return VmPowerState
*/
func (a *powerManagementApiService) ChangePowerState(id string, operation model.VmPowerOperation) (*model.VmPowerState, error) {
	var returnValue model.VmPowerState

	if err := a.client.Put(fmt.Sprintf("/api/vms/%s/power", id), operation, &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}

/*
GetPowerState Returns the power state of the VM
  - @param id ID of VM

@return VmPowerState
*/
func (a *powerManagementApiService) GetPowerState(id string) (*model.VmPowerState, error) {
	var returnValue model.VmPowerState

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/power", id), &returnValue); err != nil {
		return nil, err
	} else {
		return &returnValue, err
	}
}
