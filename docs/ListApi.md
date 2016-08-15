# \ListApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGet**](ListApi.md#ListGet) | **Get** /list | Get information about all lists
[**ListListIdContactsGet**](ListApi.md#ListListIdContactsGet) | **Get** /list/{listId}/contacts | Find contacts belonging to a list
[**ListListIdDelete**](ListApi.md#ListListIdDelete) | **Delete** /list/{listId} | Deletes a list
[**ListListIdGet**](ListApi.md#ListListIdGet) | **Get** /list/{listId} | Find list by ID
[**ListListIdPut**](ListApi.md#ListListIdPut) | **Put** /list/{listId} | Update a list by ID
[**ListPost**](ListApi.md#ListPost) | **Post** /list | Add a new list


# **ListGet**
> InlineResponse2007 ListGet($apiKey)

Get information about all lists




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdContactsGet**
> []DeepListEmailContact ListListIdContactsGet($apiKey, $listId, $limit, $offset, $contactType, $search)

Find contacts belonging to a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| ID of list that needs to be fetched | 
 **limit** | **int32**| Maximum number of contacts to be returned. Note that limit maximum value is 100 and default value is 10. | [optional] 
 **offset** | **int32**| Offset from where we contacts should be retrieved. Default value is 0. | [optional] 
 **contactType** | **string**| Can be any of the following - all, confirmed, unconfirmed, unsubscribed, bounced, markedspam. Default contact_type is all | [optional] 
 **search** | **string**| search term which shall be run against contact&#39;s first name, last name and email. | [optional] 

### Return type

[**[]DeepListEmailContact**](DeepListEmailContact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdDelete**
> ListListIdDelete($apiKey, $listId)

Deletes a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| List ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdGet**
> List ListListIdGet($apiKey, $listId)

Find list by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| ID of list that needs to be fetched | 

### Return type

[**List**](List.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdPut**
> InlineResponse2002 ListListIdPut($apiKey, $listId, $body)

Update a list by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| ID of list that needs to be updated | 
 **body** | [**ListAddUpdate**](ListAddUpdate.md)| List object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPost**
> InlineResponse2008 ListPost($apiKey, $body)

Add a new list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**ListAddUpdate**](ListAddUpdate.md)| List object that needs to be added | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

