# \VMNetworkAdaptersManagementApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNICDevice**](VMNetworkAdaptersManagementApi.md#CreateNICDevice) | **Post** /vms/{id}/nic | Creates a network adapter in the VM
[**DeleteNICDevice**](VMNetworkAdaptersManagementApi.md#DeleteNICDevice) | **Delete** /vms/{id}/nic/{index} | Deletes a VM network adapter
[**GetAllNICDevices**](VMNetworkAdaptersManagementApi.md#GetAllNICDevices) | **Get** /vms/{id}/nic | Returns all network adapters in the VM
[**GetIPAddress**](VMNetworkAdaptersManagementApi.md#GetIPAddress) | **Get** /vms/{id}/ip | Returns the IP address of a VM
[**GetNicInfo**](VMNetworkAdaptersManagementApi.md#GetNicInfo) | **Get** /vms/{id}/nicips | Returns the IP stack configuration of all NICs of a VM
[**UpdateNICDevice**](VMNetworkAdaptersManagementApi.md#UpdateNICDevice) | **Put** /vms/{id}/nic/{index} | Updates a network adapter in the VM


# **CreateNICDevice**
> NicDevice CreateNICDevice(ctx, id, parameters)
Creates a network adapter in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **parameters** | [**NicDeviceParameter**](NicDeviceParameter.md)| Parameters of network adapter to create | 

### Return type

[**NicDevice**](NICDevice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNICDevice**
> DeleteNICDevice(ctx, id, index)
Deletes a VM network adapter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **index** | **string**| Index of VM network adapter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllNICDevices**
> NicDevices GetAllNICDevices(ctx, id)
Returns all network adapters in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 

### Return type

[**NicDevices**](NICDevices.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAddress**
> InlineResponse200 GetIPAddress(ctx, id)
Returns the IP address of a VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNicInfo**
> NicIpStackAll GetNicInfo(ctx, id)
Returns the IP stack configuration of all NICs of a VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 

### Return type

[**NicIpStackAll**](NicIpStackAll.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNICDevice**
> NicDevice UpdateNICDevice(ctx, id, index, parameters)
Updates a network adapter in the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **index** | **string**| Index of VM network adapter | 
  **parameters** | [**NicDeviceParameter**](NicDeviceParameter.md)| Parameters of network adapter to update to | 

### Return type

[**NicDevice**](NICDevice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

