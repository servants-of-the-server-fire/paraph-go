# CreateWebhook201ResponseWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Events** | [**[]WebhookEvent**](WebhookEvent.md) |  | 
**Active** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Secret** | Pointer to **string** | HMAC secret for verifying webhook deliveries. Only returned once, when the webhook is created. Store it securely.  | [optional] 

## Methods

### NewCreateWebhook201ResponseWebhook

`func NewCreateWebhook201ResponseWebhook(id string, url string, events []WebhookEvent, active bool, createdAt time.Time, updatedAt time.Time, ) *CreateWebhook201ResponseWebhook`

NewCreateWebhook201ResponseWebhook instantiates a new CreateWebhook201ResponseWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhook201ResponseWebhookWithDefaults

`func NewCreateWebhook201ResponseWebhookWithDefaults() *CreateWebhook201ResponseWebhook`

NewCreateWebhook201ResponseWebhookWithDefaults instantiates a new CreateWebhook201ResponseWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateWebhook201ResponseWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWebhook201ResponseWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWebhook201ResponseWebhook) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *CreateWebhook201ResponseWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhook201ResponseWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhook201ResponseWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *CreateWebhook201ResponseWebhook) GetEvents() []WebhookEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateWebhook201ResponseWebhook) GetEventsOk() (*[]WebhookEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateWebhook201ResponseWebhook) SetEvents(v []WebhookEvent)`

SetEvents sets Events field to given value.


### GetActive

`func (o *CreateWebhook201ResponseWebhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateWebhook201ResponseWebhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateWebhook201ResponseWebhook) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *CreateWebhook201ResponseWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateWebhook201ResponseWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateWebhook201ResponseWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateWebhook201ResponseWebhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateWebhook201ResponseWebhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateWebhook201ResponseWebhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSecret

`func (o *CreateWebhook201ResponseWebhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateWebhook201ResponseWebhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateWebhook201ResponseWebhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateWebhook201ResponseWebhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


