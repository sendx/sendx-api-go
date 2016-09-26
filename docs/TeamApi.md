# \TeamApi

All URIs are relative to *http://app.sendx.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamGet**](TeamApi.md#TeamGet) | **Get** /team | Get information about all teams
[**TeamPost**](TeamApi.md#TeamPost) | **Post** /team | Add a new team
[**TeamTeamIdCampaignsGet**](TeamApi.md#TeamTeamIdCampaignsGet) | **Get** /team/{teamId}/campaigns | Find campaigns of a team by ID
[**TeamTeamIdContactsGet**](TeamApi.md#TeamTeamIdContactsGet) | **Get** /team/{teamId}/contacts | Find contacts of a team by ID
[**TeamTeamIdDelete**](TeamApi.md#TeamTeamIdDelete) | **Delete** /team/{teamId} | Deletes a team
[**TeamTeamIdGet**](TeamApi.md#TeamTeamIdGet) | **Get** /team/{teamId} | Find team by ID
[**TeamTeamIdListsGet**](TeamApi.md#TeamTeamIdListsGet) | **Get** /team/{teamId}/lists | Find lists of a team by ID
[**TeamTeamIdPut**](TeamApi.md#TeamTeamIdPut) | **Put** /team/{teamId} | Update a team by ID
[**TeamTeamIdTagsGet**](TeamApi.md#TeamTeamIdTagsGet) | **Get** /team/{teamId}/tags | Find tags of a team by ID


# **TeamGet**
> InlineResponse20013 TeamGet($apiKey)

Get information about all teams




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamPost**
> InlineResponse20014 TeamPost($apiKey, $body)

Add a new team




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**TeamAddUpdate**](TeamAddUpdate.md)| Team object that needs to be added | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdCampaignsGet**
> []Campaign TeamTeamIdCampaignsGet($apiKey, $teamId)

Find campaigns of a team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team whose campaigns need to be fetched | 

### Return type

[**[]Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdContactsGet**
> []DeepTeamEmailContact TeamTeamIdContactsGet($apiKey, $teamId, $limit, $offset, $contactType, $search)

Find contacts of a team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team whose campaigns need to be fetched | 
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

# **TeamTeamIdDelete**
> TeamTeamIdDelete($apiKey, $teamId)

Deletes a team




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| Team ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdGet**
> Team TeamTeamIdGet($apiKey, $teamId)

Find team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team that needs to be fetched | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdListsGet**
> []List TeamTeamIdListsGet($apiKey, $teamId)

Find lists of a team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team whose campaigns need to be fetched | 

### Return type

[**[]List**](List.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdPut**
> InlineResponse2002 TeamTeamIdPut($apiKey, $teamId, $body)

Update a team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team that needs to be updated | 
 **body** | [**TeamAddUpdate**](TeamAddUpdate.md)| Team object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamTeamIdTagsGet**
> []Tag TeamTeamIdTagsGet($apiKey, $teamId)

Find tags of a team by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **int64**| ID of team whose campaigns need to be fetched | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

