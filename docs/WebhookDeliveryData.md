# WebhookDeliveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**SignerId** | Pointer to **string** | Present for signer.* events | [optional] 
**SignerLabel** | Pointer to **string** | Present for signer.* events | [optional] 
**RecipientEmail** | Pointer to **string** | Present for signer.* events | [optional] 

## Methods

### NewWebhookDeliveryData

`func NewWebhookDeliveryData() *WebhookDeliveryData`

NewWebhookDeliveryData instantiates a new WebhookDeliveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryDataWithDefaults

`func NewWebhookDeliveryDataWithDefaults() *WebhookDeliveryData`

NewWebhookDeliveryDataWithDefaults instantiates a new WebhookDeliveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *WebhookDeliveryData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookDeliveryData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookDeliveryData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WebhookDeliveryData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTemplateId

`func (o *WebhookDeliveryData) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *WebhookDeliveryData) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *WebhookDeliveryData) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *WebhookDeliveryData) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetSignerId

`func (o *WebhookDeliveryData) GetSignerId() string`

GetSignerId returns the SignerId field if non-nil, zero value otherwise.

### GetSignerIdOk

`func (o *WebhookDeliveryData) GetSignerIdOk() (*string, bool)`

GetSignerIdOk returns a tuple with the SignerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerId

`func (o *WebhookDeliveryData) SetSignerId(v string)`

SetSignerId sets SignerId field to given value.

### HasSignerId

`func (o *WebhookDeliveryData) HasSignerId() bool`

HasSignerId returns a boolean if a field has been set.

### GetSignerLabel

`func (o *WebhookDeliveryData) GetSignerLabel() string`

GetSignerLabel returns the SignerLabel field if non-nil, zero value otherwise.

### GetSignerLabelOk

`func (o *WebhookDeliveryData) GetSignerLabelOk() (*string, bool)`

GetSignerLabelOk returns a tuple with the SignerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerLabel

`func (o *WebhookDeliveryData) SetSignerLabel(v string)`

SetSignerLabel sets SignerLabel field to given value.

### HasSignerLabel

`func (o *WebhookDeliveryData) HasSignerLabel() bool`

HasSignerLabel returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *WebhookDeliveryData) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *WebhookDeliveryData) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *WebhookDeliveryData) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *WebhookDeliveryData) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


