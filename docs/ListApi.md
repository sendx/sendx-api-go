# \ListApi

All URIs are relative to *http://app.sendx.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGet**](ListApi.md#ListGet) | **Get** /list | Get information about all lists
[**ListListIdContactDelete**](ListApi.md#ListListIdContactDelete) | **Delete** /list/{listId}/contact | Remove a contact from a list
[**ListListIdContactPost**](ListApi.md#ListListIdContactPost) | **Post** /list/{listId}/contact | Add a contact to a list
[**ListListIdContactsGet**](ListApi.md#ListListIdContactsGet) | **Get** /list/{listId}/contacts | Find contacts belonging to a list
[**ListListIdDelete**](ListApi.md#ListListIdDelete) | **Delete** /list/{listId} | Deletes a list
[**ListListIdGet**](ListApi.md#ListListIdGet) | **Get** /list/{listId} | Find list by ID
[**ListListIdPut**](ListApi.md#ListListIdPut) | **Put** /list/{listId} | Update a list by ID
[**ListPost**](ListApi.md#ListPost) | **Post** /list | Add a new list


# **ListGet**
> InlineResponse2008 ListGet($apiKey)

Get information about all lists




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdContactDelete**
> ListListIdContactDelete($apiKey, $listId, $body)

Remove a contact from a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| ID of list for which contact needs to be remove | 
 **body** | [**ListContact**](ListContact.md)| Contact email and team id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListListIdContactPost**
> InlineResponse20015 ListListIdContactPost($apiKey, $listId, $body)

Add a contact to a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **listId** | **int64**| ID of list for which the contact needs to be added | 
 **body** | [**ListContact**](ListContact.md)| Contact email and team id | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

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
> InlineResponse2009 ListPost($apiKey, $body)

Add a new list

Adding a new list with all the fields. List type can be 0 - Single OptIn 1 - Double OptIn


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**ListAddUpdate**](ListAddUpdate.md)| List object that needs to be added | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

