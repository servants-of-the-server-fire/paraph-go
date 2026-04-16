# CreateRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | ID of the template to fill | 
**Fields** | Pointer to **map[string]string** | Map of field names to values. Text fields accept strings. Checkbox fields accept \&quot;checked\&quot; or \&quot;unchecked\&quot;. | [optional] 
**Signers** | Pointer to [**map[string]SignerInput**](SignerInput.md) | Map of signer label to signer configuration. Each key is a signer label defined in the template&#39;s signature placements. Omit entirely for fill-only requests.  | [optional] 
**Title** | Pointer to **string** | Display title for the request. Defaults to the template name. Shown in emails and dashboards. | [optional] 
**Message** | Pointer to **string** | Custom message included in signing emails sent to signers. | [optional] 
**Metadata** | Pointer to **map[string]string** | Arbitrary key-value pairs for your own use (max 10 keys, key max 128 chars, value max 1024 chars) | [optional] 
**AllowTypedSignature** | Pointer to **bool** | Whether signers can type their name as a signature instead of drawing or uploading one. Defaults to true.  | [optional] [default to true]

## Methods

### NewCreateRequestRequest

`func NewCreateRequestRequest(templateId string, ) *CreateRequestRequest`

NewCreateRequestRequest instantiates a new CreateRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestRequestWithDefaults

`func NewCreateRequestRequestWithDefaults() *CreateRequestRequest`

NewCreateRequestRequestWithDefaults instantiates a new CreateRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *CreateRequestRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateRequestRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateRequestRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetFields

`func (o *CreateRequestRequest) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateRequestRequest) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateRequestRequest) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateRequestRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSigners

`func (o *CreateRequestRequest) GetSigners() map[string]SignerInput`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CreateRequestRequest) GetSignersOk() (*map[string]SignerInput, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CreateRequestRequest) SetSigners(v map[string]SignerInput)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CreateRequestRequest) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetTitle

`func (o *CreateRequestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateRequestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateRequestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateRequestRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *CreateRequestRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateRequestRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateRequestRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateRequestRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateRequestRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateRequestRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateRequestRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateRequestRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAllowTypedSignature

`func (o *CreateRequestRequest) GetAllowTypedSignature() bool`

GetAllowTypedSignature returns the AllowTypedSignature field if non-nil, zero value otherwise.

### GetAllowTypedSignatureOk

`func (o *CreateRequestRequest) GetAllowTypedSignatureOk() (*bool, bool)`

GetAllowTypedSignatureOk returns a tuple with the AllowTypedSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTypedSignature

`func (o *CreateRequestRequest) SetAllowTypedSignature(v bool)`

SetAllowTypedSignature sets AllowTypedSignature field to given value.

### HasAllowTypedSignature

`func (o *CreateRequestRequest) HasAllowTypedSignature() bool`

HasAllowTypedSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


