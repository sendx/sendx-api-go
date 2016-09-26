# \CustomfieldApi

All URIs are relative to *http://app.sendx.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomfieldCustomfieldIdGet**](CustomfieldApi.md#CustomfieldCustomfieldIdGet) | **Get** /customfield/{customfieldId} | Find customfield by ID
[**CustomfieldGet**](CustomfieldApi.md#CustomfieldGet) | **Get** /customfield | Get information about all customfield
[**CustomfieldPost**](CustomfieldApi.md#CustomfieldPost) | **Post** /customfield | Add a new Customfield


# **CustomfieldCustomfieldIdGet**
> Customfield CustomfieldCustomfieldIdGet($apiKey, $customfieldId)

Find customfield by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **customfieldId** | **int64**| ID of customfield that needs to be fetched | 

### Return type

[**Customfield**](Customfield.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomfieldGet**
> InlineResponse2006 CustomfieldGet($apiKey)

Get information about all customfield




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomfieldPost**
> InlineResponse2007 CustomfieldPost($apiKey, $body)

Add a new Customfield




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**CustomfieldAddUpdate**](CustomfieldAddUpdate.md)| Customfield object that needs to be added | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

