# \VMSharedFoldersManagementApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedFolder**](VMSharedFoldersManagementApi.md#CreateSharedFolder) | **Post** /vms/{id}/sharedfolders | Mounts a new shared folder in the VM
[**DeleteSharedFolder**](VMSharedFoldersManagementApi.md#DeleteSharedFolder) | **Delete** /vms/{id}/sharedfolders/{folder id} | Deletes a shared folder
[**GetAllSharedFolders**](VMSharedFoldersManagementApi.md#GetAllSharedFolders) | **Get** /vms/{id}/sharedfolders | Returns all shared folders mounted in the VM
[**UpdataSharedFolder**](VMSharedFoldersManagementApi.md#UpdataSharedFolder) | **Put** /vms/{id}/sharedfolders/{folder id} | Updates a shared folder mounted in the VM


# **CreateSharedFolder**
> SharedFolders CreateSharedFolder(ctx, id, parameters)
Mounts a new shared folder in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **parameters** | [**SharedFolder**](SharedFolder.md)| Parameters of the shared folder to mount | 

### Return type

[**SharedFolders**](SharedFolders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSharedFolder**
> DeleteSharedFolder(ctx, id, folderId)
Deletes a shared folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **folderId** | **string**| ID of shared folder | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSharedFolders**
> SharedFolders GetAllSharedFolders(ctx, id)
Returns all shared folders mounted in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 

### Return type

[**SharedFolders**](SharedFolders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdataSharedFolder**
> SharedFolders UpdataSharedFolder(ctx, id, folderId, parameters)
Updates a shared folder mounted in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **folderId** | **string**| ID of VM shared folder | 
  **parameters** | [**SharedFolderParameter**](SharedFolderParameter.md)| Parameters of the shared folder to update to | 

### Return type

[**SharedFolders**](SharedFolders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

