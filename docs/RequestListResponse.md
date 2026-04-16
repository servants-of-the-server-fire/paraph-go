# RequestListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]DocumentRequestSummary**](DocumentRequestSummary.md) |  | [optional] 
**ListInfo** | Pointer to [**ListInfo**](ListInfo.md) |  | [optional] 

## Methods

### NewRequestListResponse

`func NewRequestListResponse() *RequestListResponse`

NewRequestListResponse instantiates a new RequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestListResponseWithDefaults

`func NewRequestListResponseWithDefaults() *RequestListResponse`

NewRequestListResponseWithDefaults instantiates a new RequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RequestListResponse) GetRequests() []DocumentRequestSummary`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestListResponse) GetRequestsOk() (*[]DocumentRequestSummary, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestListResponse) SetRequests(v []DocumentRequestSummary)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RequestListResponse) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetListInfo

`func (o *RequestListResponse) GetListInfo() ListInfo`

GetListInfo returns the ListInfo field if non-nil, zero value otherwise.

### GetListInfoOk

`func (o *RequestListResponse) GetListInfoOk() (*ListInfo, bool)`

GetListInfoOk returns a tuple with the ListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListInfo

`func (o *RequestListResponse) SetListInfo(v ListInfo)`

SetListInfo sets ListInfo field to given value.

### HasListInfo

`func (o *RequestListResponse) HasListInfo() bool`

HasListInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


