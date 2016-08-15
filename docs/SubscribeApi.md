# \SubscribeApi

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeEncryptedListIdPost**](SubscribeApi.md#SubscribeEncryptedListIdPost) | **Post** /subscribe/{encryptedListId} | Subscribe a new user a list
[**SubscribeEncryptedListIdPut**](SubscribeApi.md#SubscribeEncryptedListIdPut) | **Put** /subscribe/{encryptedListId} | Subscribe an existing user a list


# **SubscribeEncryptedListIdPost**
> SubscribeEncryptedListIdPost($encryptedListId, $body)

Subscribe a new user a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptedListId** | **string**| encrypted list Id of the list to which you want to add user | 
 **body** | [**ContactAddUpdate**](ContactAddUpdate.md)| Contact object that needs to be added | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeEncryptedListIdPut**
> SubscribeEncryptedListIdPut($encryptedListId, $body)

Subscribe an existing user a list




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptedListId** | **string**| encrypted list Id of the list to which you want to add user | 
 **body** | [**Email**](Email.md)| Contact Emil needs to be added | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

