# \BrokerAccountApi

All URIs are relative to *https://moderntrader.mershab.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserBrokerAccounts**](BrokerAccountApi.md#GetUserBrokerAccounts) | **Get** /api/users/{userId}/brokeraccounts | Get all broker accounts for a user



## GetUserBrokerAccounts

> []BrokerAccount GetUserBrokerAccounts(ctx, userId).Execute()

Get all broker accounts for a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerAccountApi.GetUserBrokerAccounts(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerAccountApi.GetUserBrokerAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBrokerAccounts`: []BrokerAccount
    fmt.Fprintf(os.Stdout, "Response from `BrokerAccountApi.GetUserBrokerAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBrokerAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BrokerAccount**](BrokerAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

