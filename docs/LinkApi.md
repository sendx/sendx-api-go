# \LinkApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LinkGet**](LinkApi.md#LinkGet) | **Get** /link | Get information about all links
[**LinkLinkIdDelete**](LinkApi.md#LinkLinkIdDelete) | **Delete** /link/{linkId} | Deletes a link
[**LinkLinkIdGet**](LinkApi.md#LinkLinkIdGet) | **Get** /link/{linkId} | Find link by ID
[**LinkLinkIdPut**](LinkApi.md#LinkLinkIdPut) | **Put** /link/{linkId} | Update a link by ID
[**LinkPost**](LinkApi.md#LinkPost) | **Post** /link | Add a new link


# **LinkGet**
> InlineResponse2005 LinkGet($apiKey)

Get information about all links




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLinkIdDelete**
> LinkLinkIdDelete($apiKey, $linkId)

Deletes a link




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **linkId** | **int64**| Link ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLinkIdGet**
> Link LinkLinkIdGet($apiKey, $linkId)

Find link by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **linkId** | **int64**| ID of link that needs to be fetched | 

### Return type

[**Link**](Link.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLinkIdPut**
> InlineResponse2002 LinkLinkIdPut($apiKey, $linkId, $body)

Update a link by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **linkId** | **int64**| ID of link that needs to be updated | 
 **body** | [**LinkAddUpdate**](LinkAddUpdate.md)| Link object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkPost**
> InlineResponse2006 LinkPost($apiKey, $body)

Add a new link




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**LinkAddUpdate**](LinkAddUpdate.md)| Link object that needs to be added | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

