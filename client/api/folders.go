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

type sharedFoldersManagementApiService struct {
	client Client
}

func NewSharedFoldersManagementApiService(client Client) SharedFoldersManagementApiService {
	return &sharedFoldersManagementApiService{
		client: client,
	}
}

/*
CreateSharedFolder Mounts a new shared folder in the VM
  - @param id ID of VM
  - @param parameters Parameters of the shared folder to mount

@return SharedFolders
*/
func (a *sharedFoldersManagementApiService) CreateSharedFolder(id string, parameters *model.SharedFolder) (model.SharedFolders, error) {
	var returnValue model.SharedFolders

	if err := a.client.Post(fmt.Sprintf("/api/vms/%s/sharedfolders", id), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, err
	}
}

/*
DeleteSharedFolder Deletes a shared folder
  - @param id ID of VM
  - @param folderId ID of shared folder
*/
func (a *sharedFoldersManagementApiService) DeleteSharedFolder(id string, folderId string) error {
	return a.client.Delete(fmt.Sprintf("/api/vms/%s/sharedfolders/%s", id, folderId), nil)
}

/*
GetAllSharedFolders Returns all shared folders mounted in the VM
  - @param id ID of VM

@return SharedFolders
*/
func (a *sharedFoldersManagementApiService) GetAllSharedFolders(id string) (model.SharedFolders, error) {
	var returnValue model.SharedFolders

	if err := a.client.Get(fmt.Sprintf("/api/vms/%s/sharedfolders", id), &returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, err
	}
}

/*
UpdateSharedFolder Updates a shared folder mounted in the VM
  - @param id ID of VM
  - @param folderId ID of VM shared folder
  - @param parameters Parameters of the shared folder to update to

@return SharedFolders
*/
func (a *sharedFoldersManagementApiService) UpdateSharedFolder(id string, folderId string, parameters *model.SharedFolderParameter) (model.SharedFolders, error) {
	var returnValue model.SharedFolders

	if err := a.client.Put(fmt.Sprintf("/api/vms/%s/sharedfolders/%s", id, folderId), parameters, &returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, err
	}
}
