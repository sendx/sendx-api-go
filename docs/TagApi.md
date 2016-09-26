# \TagApi

All URIs are relative to *http://app.sendx.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagGet**](TagApi.md#TagGet) | **Get** /tag | Get information about all tags
[**TagPost**](TagApi.md#TagPost) | **Post** /tag | Add a new tag
[**TagTagIdContactDelete**](TagApi.md#TagTagIdContactDelete) | **Delete** /tag/{tagId}/contact | Remove a contact from a tag
[**TagTagIdContactPost**](TagApi.md#TagTagIdContactPost) | **Post** /tag/{tagId}/contact | Add a contact to a tag
[**TagTagIdContactsGet**](TagApi.md#TagTagIdContactsGet) | **Get** /tag/{tagId}/contacts | Find contacts belonging to a tag
[**TagTagIdDelete**](TagApi.md#TagTagIdDelete) | **Delete** /tag/{tagId} | Deletes a tag
[**TagTagIdGet**](TagApi.md#TagTagIdGet) | **Get** /tag/{tagId} | Find tag by ID
[**TagTagIdPut**](TagApi.md#TagTagIdPut) | **Put** /tag/{tagId} | Update a tag by ID


# **TagGet**
> InlineResponse20010 TagGet($apiKey)

Get information about all tags




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagPost**
> InlineResponse20011 TagPost($apiKey, $body)

Add a new tag




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**TagAddUpdate**](TagAddUpdate.md)| Tag object that needs to be added | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdContactDelete**
> TagTagIdContactDelete($apiKey, $tagId, $body)

Remove a contact from a tag




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| ID of tag for which contact needs to be remove | 
 **body** | [**TagContact**](TagContact.md)| Contact email and team id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdContactPost**
> InlineResponse20012 TagTagIdContactPost($apiKey, $tagId, $body)

Add a contact to a tag




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| ID of tag for which the contact needs to be added | 
 **body** | [**TagContact**](TagContact.md)| Contact email and team id | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdContactsGet**
> []DeepTeamEmailContact TagTagIdContactsGet($apiKey, $tagId, $limit, $offset, $contactType, $search)

Find contacts belonging to a tag




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| ID of tag that needs to be fetched | 
 **limit** | **int32**| Maximum number of contacts to be returned. Note that limit maximum value is 100 and default value is 10. | [optional] 
 **offset** | **int32**| Offset from where we contacts should be retrieved. Default value is 0. | [optional] 
 **contactType** | **string**| Can be any of the following - all, confirmed, unconfirmed, unsubscribed, bounced, markedspam. Default contact_type is all | [optional] 
 **search** | **string**| search term which shall be run against contact&#39;s first name, last name and email. | [optional] 

### Return type

[**[]DeepTeamEmailContact**](DeepTeamEmailContact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdDelete**
> TagTagIdDelete($apiKey, $tagId)

Deletes a tag




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| Tag ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdGet**
> Tag TagTagIdGet($apiKey, $tagId)

Find tag by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| ID of tag that needs to be fetched | 

### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagTagIdPut**
> InlineResponse2002 TagTagIdPut($apiKey, $tagId, $body)

Update a tag by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **tagId** | **int64**| ID of tag that needs to be updated | 
 **body** | [**TagAddUpdate**](TagAddUpdate.md)| Tag object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

