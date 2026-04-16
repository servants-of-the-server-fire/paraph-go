# WebhookListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhooks** | Pointer to [**[]Webhook**](Webhook.md) |  | [optional] 
**ListInfo** | Pointer to [**ListInfo**](ListInfo.md) |  | [optional] 

## Methods

### NewWebhookListResponse

`func NewWebhookListResponse() *WebhookListResponse`

NewWebhookListResponse instantiates a new WebhookListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookListResponseWithDefaults

`func NewWebhookListResponseWithDefaults() *WebhookListResponse`

NewWebhookListResponseWithDefaults instantiates a new WebhookListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhooks

`func (o *WebhookListResponse) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookListResponse) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookListResponse) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *WebhookListResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetListInfo

`func (o *WebhookListResponse) GetListInfo() ListInfo`

GetListInfo returns the ListInfo field if non-nil, zero value otherwise.

### GetListInfoOk

`func (o *WebhookListResponse) GetListInfoOk() (*ListInfo, bool)`

GetListInfoOk returns a tuple with the ListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListInfo

`func (o *WebhookListResponse) SetListInfo(v ListInfo)`

SetListInfo sets ListInfo field to given value.

### HasListInfo

`func (o *WebhookListResponse) HasListInfo() bool`

HasListInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


