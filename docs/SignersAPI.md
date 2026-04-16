# \SignersAPI

All URIs are relative to *https://paraph.dev/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadSignerSignature**](SignersAPI.md#DownloadSignerSignature) | **Get** /requests/{id}/signers/{sid}/signature | Download signer signature
[**ResendSigningLink**](SignersAPI.md#ResendSigningLink) | **Post** /requests/{id}/signers/{sid}/resend | Resend signing link



## DownloadSignerSignature

> *os.File DownloadSignerSignature(ctx, id, sid).Execute()

Download signer signature



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
	sid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Signer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignersAPI.DownloadSignerSignature(context.Background(), id, sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignersAPI.DownloadSignerSignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSignerSignature`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SignersAPI.DownloadSignerSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Request ID | 
**sid** | **string** | Signer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSignerSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendSigningLink

> ResendSigningLink200Response ResendSigningLink(ctx, id, sid).Execute()

Resend signing link



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
	sid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Signer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignersAPI.ResendSigningLink(context.Background(), id, sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignersAPI.ResendSigningLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendSigningLink`: ResendSigningLink200Response
	fmt.Fprintf(os.Stdout, "Response from `SignersAPI.ResendSigningLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Request ID | 
**sid** | **string** | Signer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendSigningLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResendSigningLink200Response**](ResendSigningLink200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

