# DocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TemplateId** | **string** |  | 
**Title** | Pointer to **string** | Display title for the request | [optional] 
**Message** | Pointer to **string** | Custom message included in signing emails | [optional] 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**HasSigning** | **bool** | Whether this request includes signers | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Inputs** | Pointer to **map[string]string** | Field values used to fill the template | [optional] 
**Metadata** | Pointer to **map[string]string** | Arbitrary key-value pairs for your own use (max 10 keys, key max 128 chars, value max 1024 chars) | [optional] 
**Signers** | Pointer to [**[]Signer**](Signer.md) | Signers attached to this request, if any | [optional] 

## Methods

### NewDocumentRequest

`func NewDocumentRequest(id string, templateId string, status RequestStatus, hasSigning bool, createdAt time.Time, updatedAt time.Time, ) *DocumentRequest`

NewDocumentRequest instantiates a new DocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentRequestWithDefaults

`func NewDocumentRequestWithDefaults() *DocumentRequest`

NewDocumentRequestWithDefaults instantiates a new DocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *DocumentRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DocumentRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DocumentRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetTitle

`func (o *DocumentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *DocumentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DocumentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DocumentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DocumentRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentRequest) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentRequest) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentRequest) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetHasSigning

`func (o *DocumentRequest) GetHasSigning() bool`

GetHasSigning returns the HasSigning field if non-nil, zero value otherwise.

### GetHasSigningOk

`func (o *DocumentRequest) GetHasSigningOk() (*bool, bool)`

GetHasSigningOk returns a tuple with the HasSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSigning

`func (o *DocumentRequest) SetHasSigning(v bool)`

SetHasSigning sets HasSigning field to given value.


### GetCreatedAt

`func (o *DocumentRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DocumentRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInputs

`func (o *DocumentRequest) GetInputs() map[string]string`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *DocumentRequest) GetInputsOk() (*map[string]string, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *DocumentRequest) SetInputs(v map[string]string)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *DocumentRequest) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMetadata

`func (o *DocumentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSigners

`func (o *DocumentRequest) GetSigners() []Signer`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *DocumentRequest) GetSignersOk() (*[]Signer, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *DocumentRequest) SetSigners(v []Signer)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *DocumentRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


