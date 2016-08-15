# \CampaignApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignCampaignIdDelete**](CampaignApi.md#CampaignCampaignIdDelete) | **Delete** /campaign/{campaignId} | Deletes a campaign
[**CampaignCampaignIdGet**](CampaignApi.md#CampaignCampaignIdGet) | **Get** /campaign/{campaignId} | Find campaign by ID
[**CampaignCampaignIdPut**](CampaignApi.md#CampaignCampaignIdPut) | **Put** /campaign/{campaignId} | Update a campaign by ID
[**CampaignGet**](CampaignApi.md#CampaignGet) | **Get** /campaign | Get information about all campaigns
[**CampaignPost**](CampaignApi.md#CampaignPost) | **Post** /campaign | Add a new campaign


# **CampaignCampaignIdDelete**
> CampaignCampaignIdDelete($apiKey, $campaignId)

Deletes a campaign




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **campaignId** | **int64**| Campaign ID to delete | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignCampaignIdGet**
> Campaign CampaignCampaignIdGet($apiKey, $campaignId)

Find campaign by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **campaignId** | **int64**| ID of campaign that needs to be fetched | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignCampaignIdPut**
> InlineResponse2002 CampaignCampaignIdPut($apiKey, $campaignId, $body)

Update a campaign by ID




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **campaignId** | **int64**| ID of campaign that needs to be updated | 
 **body** | [**CampaignAddUpdate**](CampaignAddUpdate.md)| Campaign object that needs to be added | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignGet**
> InlineResponse200 CampaignGet($apiKey)

Get information about all campaigns




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignPost**
> InlineResponse2001 CampaignPost($apiKey, $body)

Add a new campaign




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **body** | [**CampaignAddUpdate**](CampaignAddUpdate.md)| Campaign object that needs to be added | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

