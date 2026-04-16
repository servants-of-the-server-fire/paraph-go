# \RequestsAPI

All URIs are relative to *https://paraph.dev/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRequest**](RequestsAPI.md#CancelRequest) | **Post** /requests/{id}/cancel | Cancel request
[**CreateRequest**](RequestsAPI.md#CreateRequest) | **Post** /requests | Create request
[**DownloadRequest**](RequestsAPI.md#DownloadRequest) | **Get** /requests/{id}/download | Download request PDF
[**GetRequest**](RequestsAPI.md#GetRequest) | **Get** /requests/{id} | Get request detail
[**ListRequests**](RequestsAPI.md#ListRequests) | **Get** /requests | List requests



## CancelRequest

> RequestResponse CancelRequest(ctx, id).Execute()

Cancel request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CancelRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CancelRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRequest`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CancelRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestResponse**](RequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> RequestResponse CreateRequest(ctx).CreateRequestRequest(createRequestRequest).IdempotencyKey(idempotencyKey).Execute()

Create request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	createRequestRequest := *openapiclient.NewCreateRequestRequest("TemplateId_example") // CreateRequestRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Optional. If provided, ensures the request is processed at most once within 24 hours. Retries with the same key and body return the original response without creating a duplicate. Using the same key with a different request body returns 409 Conflict.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CreateRequest(context.Background()).CreateRequestRequest(createRequestRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequest`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequestRequest** | [**CreateRequestRequest**](CreateRequestRequest.md) |  | 
 **idempotencyKey** | **string** | Optional. If provided, ensures the request is processed at most once within 24 hours. Retries with the same key and body return the original response without creating a duplicate. Using the same key with a different request body returns 409 Conflict.  | 

### Return type

[**RequestResponse**](RequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRequest

> *os.File DownloadRequest(ctx, id).Execute()

Download request PDF



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.DownloadRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.DownloadRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadRequest`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.DownloadRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequest

> RequestResponse GetRequest(ctx, id).Execute()

Get request detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequest`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestResponse**](RequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequests

> RequestListResponse ListRequests(ctx).Page(page).PageSize(pageSize).Status(status).From(from).To(to).MetadataKey(metadataKey).MetadataValue(metadataValue).Execute()

List requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	status := openapiclient.RequestStatus("success") // RequestStatus | Filter by request status (optional)
	from := time.Now() // time.Time | Only return requests created on or after this timestamp (RFC 3339) (optional)
	to := time.Now() // time.Time | Only return requests created on or before this timestamp (RFC 3339) (optional)
	metadataKey := "metadataKey_example" // string | Filter by metadata key. Must be used together with metadata_value. (optional)
	metadataValue := "metadataValue_example" // string | Filter by metadata value (exact match). Must be used together with metadata_key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.ListRequests(context.Background()).Page(page).PageSize(pageSize).Status(status).From(from).To(to).MetadataKey(metadataKey).MetadataValue(metadataValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.ListRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRequests`: RequestListResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.ListRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **status** | [**RequestStatus**](RequestStatus.md) | Filter by request status | 
 **from** | **time.Time** | Only return requests created on or after this timestamp (RFC 3339) | 
 **to** | **time.Time** | Only return requests created on or before this timestamp (RFC 3339) | 
 **metadataKey** | **string** | Filter by metadata key. Must be used together with metadata_value. | 
 **metadataValue** | **string** | Filter by metadata value (exact match). Must be used together with metadata_key. | 

### Return type

[**RequestListResponse**](RequestListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

