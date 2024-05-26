# \StrategyApi

All URIs are relative to *https://moderntrader.mershab.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStrategy**](StrategyApi.md#CreateStrategy) | **Post** /api/strategies | Create a strategy
[**GetAllStrategies**](StrategyApi.md#GetAllStrategies) | **Get** /api/strategies | Get all strategies



## CreateStrategy

> CreateStrategy(ctx).Strategy(strategy).Execute()

Create a strategy



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
    strategy := *openapiclient.NewStrategy() // Strategy | Create a strategy in the system (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StrategyApi.CreateStrategy(context.Background()).Strategy(strategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategyApi.CreateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **strategy** | [**Strategy**](Strategy.md) | Create a strategy in the system | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllStrategies

> []Strategy GetAllStrategies(ctx).Execute()

Get all strategies



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StrategyApi.GetAllStrategies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategyApi.GetAllStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllStrategies`: []Strategy
    fmt.Fprintf(os.Stdout, "Response from `StrategyApi.GetAllStrategies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllStrategiesRequest struct via the builder pattern


### Return type

[**[]Strategy**](Strategy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

