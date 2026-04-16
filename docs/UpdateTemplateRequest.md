# UpdateTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New template name (omit to leave unchanged) | [optional] 
**Metadata** | Pointer to **map[string]string** | Replace all metadata (omit to leave unchanged, key max 128 chars, value max 1024 chars) | [optional] 
**SignaturePlacements** | Pointer to [**[]SignaturePlacementInput**](SignaturePlacementInput.md) | Replace all signature placements (omit to leave unchanged) | [optional] 

## Methods

### NewUpdateTemplateRequest

`func NewUpdateTemplateRequest() *UpdateTemplateRequest`

NewUpdateTemplateRequest instantiates a new UpdateTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateRequestWithDefaults

`func NewUpdateTemplateRequestWithDefaults() *UpdateTemplateRequest`

NewUpdateTemplateRequestWithDefaults instantiates a new UpdateTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateTemplateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateTemplateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateTemplateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateTemplateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSignaturePlacements

`func (o *UpdateTemplateRequest) GetSignaturePlacements() []SignaturePlacementInput`

GetSignaturePlacements returns the SignaturePlacements field if non-nil, zero value otherwise.

### GetSignaturePlacementsOk

`func (o *UpdateTemplateRequest) GetSignaturePlacementsOk() (*[]SignaturePlacementInput, bool)`

GetSignaturePlacementsOk returns a tuple with the SignaturePlacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturePlacements

`func (o *UpdateTemplateRequest) SetSignaturePlacements(v []SignaturePlacementInput)`

SetSignaturePlacements sets SignaturePlacements field to given value.

### HasSignaturePlacements

`func (o *UpdateTemplateRequest) HasSignaturePlacements() bool`

HasSignaturePlacements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


