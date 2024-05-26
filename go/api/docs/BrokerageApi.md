# \BrokerageApi

All URIs are relative to *https://moderntrader.mershab.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserBrokerAccount**](BrokerageApi.md#CreateUserBrokerAccount) | **Post** /api/users/{userId}/brokeraccounts | Create a new brokerage account for a user



## CreateUserBrokerAccount

> CreateUserBrokerAccount(ctx, userId).BrokerAccount(brokerAccount).Execute()

Create a new brokerage account for a user

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
    brokerAccount := *openapiclient.NewBrokerAccount() // BrokerAccount |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrokerageApi.CreateUserBrokerAccount(context.Background(), userId).BrokerAccount(brokerAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrokerageApi.CreateUserBrokerAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserBrokerAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brokerAccount** | [**BrokerAccount**](BrokerAccount.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

