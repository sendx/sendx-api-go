# \ContactApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactContactIdDelete**](ContactApi.md#ContactContactIdDelete) | **Delete** /contact/{contactId} | Deletes a contact
[**ContactContactIdGet**](ContactApi.md#ContactContactIdGet) | **Get** /contact/{contactId} | Find contact by ID
[**ContactContactIdPut**](ContactApi.md#ContactContactIdPut) | **Put** /contact/{contactId} | Update a contact by ID
[**ContactGet**](ContactApi.md#ContactGet) | **Get** /contact | Get information about all contacts
[**ContactPost**](ContactApi.md#ContactPost) | **Post** /contact | Add a new contact


# **ContactContactIdDelete**
> ContactContactIdDelete($apiKey, $contactId)

Deletes a contact




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **contactId** | **int64**| Contact ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactContactIdGet**
> Contact ContactContactIdGet($apiKey, $contactId)

Find contact by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **contactId** | **int64**| ID of contact that needs to be fetched | 

### Return type

[**Contact**](Contact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactContactIdPut**
> InlineResponse2002 ContactContactIdPut($apiKey, $contactId, $body)

Update a contact by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **contactId** | **int64**| ID of contact that needs to be updated | 
 **body** | [**ContactAddUpdate**](ContactAddUpdate.md)| Contact object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactGet**
> InlineResponse2003 ContactGet($apiKey, $limit, $offset)

Get information about all contacts




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **limit** | **int32**| Maximum number of contacts to be returned. Note that limit maximum value is 100 and default value is 10. | [optional] 
 **offset** | **int32**| Offset from where we contacts should be retrieved. Default value is 0. | [optional] 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactPost**
> InlineResponse2004 ContactPost($apiKey, $body)

Add a new contact




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**ContactAddUpdate**](ContactAddUpdate.md)| Contact object that needs to be added | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

