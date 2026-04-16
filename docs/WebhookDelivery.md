# WebhookDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**WebhookEvent**](WebhookEvent.md) |  | 
**Timestamp** | **time.Time** |  | 
**Data** | [**WebhookDeliveryData**](WebhookDeliveryData.md) |  | 

## Methods

### NewWebhookDelivery

`func NewWebhookDelivery(event WebhookEvent, timestamp time.Time, data WebhookDeliveryData, ) *WebhookDelivery`

NewWebhookDelivery instantiates a new WebhookDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryWithDefaults

`func NewWebhookDeliveryWithDefaults() *WebhookDelivery`

NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookDelivery) GetEvent() WebhookEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookDelivery) GetEventOk() (*WebhookEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookDelivery) SetEvent(v WebhookEvent)`

SetEvent sets Event field to given value.


### GetTimestamp

`func (o *WebhookDelivery) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookDelivery) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookDelivery) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *WebhookDelivery) GetData() WebhookDeliveryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookDelivery) GetDataOk() (*WebhookDeliveryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookDelivery) SetData(v WebhookDeliveryData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


