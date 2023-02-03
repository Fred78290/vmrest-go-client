# \VMPowerManagementApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePowerState**](VMPowerManagementApi.md#ChangePowerState) | **Put** /vms/{id}/power | Changes the VM power state
[**GetPowerState**](VMPowerManagementApi.md#GetPowerState) | **Get** /vms/{id}/power | Returns the power state of the VM


# **ChangePowerState**
> VmPowerState ChangePowerState(ctx, id, operation)
Changes the VM power state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 
  **operation** | [**VmPowerOperation**](VmPowerOperation.md)| VM power operation: on, off, shutdown, suspend, pause, unpause | 

### Return type

[**VmPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPowerState**
> VmPowerState GetPowerState(ctx, id)
Returns the power state of the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of VM | 

### Return type

[**VmPowerState**](VMPowerState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
 - **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

