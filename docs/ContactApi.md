# \ContactApi

All URIs are relative to *http://app.sendx.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactIdentifyPost**](ContactApi.md#ContactIdentifyPost) | **Post** /contact/identify | Identify a contact as user
[**ContactTrackPost**](ContactApi.md#ContactTrackPost) | **Post** /contact/track | Add tracking info using tags to a contact


# **ContactIdentifyPost**
> ContactResponse ContactIdentifyPost($apiKey, $teamId, $contactDetails)

Identify a contact as user




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **string**|  | 
 **contactDetails** | [**ContactRequest**](ContactRequest.md)| Contact details | 

### Return type

[**ContactResponse**](ContactResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactTrackPost**
> TrackResponse ContactTrackPost($apiKey, $teamId, $email, $trackDetails)

Add tracking info using tags to a contact




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string**|  | 
 **teamId** | **string**|  | 
 **email** | **string**|  | 
 **trackDetails** | [**TrackRequest**](TrackRequest.md)| Track Details | 

### Return type

[**TrackResponse**](TrackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

