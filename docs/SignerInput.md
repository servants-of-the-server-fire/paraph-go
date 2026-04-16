# SignerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**OverrideSignatureUrl** | Pointer to **string** | URL to a PNG signature image. When provided, the image is downloaded and stored, the signer is immediately marked as signed, and no signing email is sent.  | [optional] 

## Methods

### NewSignerInput

`func NewSignerInput(email string, ) *SignerInput`

NewSignerInput instantiates a new SignerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerInputWithDefaults

`func NewSignerInputWithDefaults() *SignerInput`

NewSignerInputWithDefaults instantiates a new SignerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignerInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignerInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignerInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOverrideSignatureUrl

`func (o *SignerInput) GetOverrideSignatureUrl() string`

GetOverrideSignatureUrl returns the OverrideSignatureUrl field if non-nil, zero value otherwise.

### GetOverrideSignatureUrlOk

`func (o *SignerInput) GetOverrideSignatureUrlOk() (*string, bool)`

GetOverrideSignatureUrlOk returns a tuple with the OverrideSignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSignatureUrl

`func (o *SignerInput) SetOverrideSignatureUrl(v string)`

SetOverrideSignatureUrl sets OverrideSignatureUrl field to given value.

### HasOverrideSignatureUrl

`func (o *SignerInput) HasOverrideSignatureUrl() bool`

HasOverrideSignatureUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


