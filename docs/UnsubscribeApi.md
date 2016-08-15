# \UnsubscribeApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnsubscribeEncryptedListIdPost**](UnsubscribeApi.md#UnsubscribeEncryptedListIdPost) | **Post** /unsubscribe/{encryptedListId} | Unsubscribe a user from list based on email id


# **UnsubscribeEncryptedListIdPost**
> UnsubscribeEncryptedListIdPost($encryptedListId, $body)

Unsubscribe a user from list based on email id




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptedListId** | **string**| encrypted list Id of the list to which you want to add user | 
 **body** | [**Email**](Email.md)| Email object of the user that needs to be unsubscribed. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

