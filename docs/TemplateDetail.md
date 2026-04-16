# TemplateDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]string** | Arbitrary key-value pairs (max 10 keys, key max 128 chars, value max 1024 chars) | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Fields** | [**[]Field**](Field.md) | Form fields detected in the uploaded PDF | 
**SignaturePlacements** | [**[]SignaturePlacement**](SignaturePlacement.md) | Signature placement regions configured on this template | 

## Methods

### NewTemplateDetail

`func NewTemplateDetail(id string, name string, createdAt time.Time, updatedAt time.Time, fields []Field, signaturePlacements []SignaturePlacement, ) *TemplateDetail`

NewTemplateDetail instantiates a new TemplateDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDetailWithDefaults

`func NewTemplateDetailWithDefaults() *TemplateDetail`

NewTemplateDetailWithDefaults instantiates a new TemplateDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateDetail) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDetail) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *TemplateDetail) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateDetail) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateDetail) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TemplateDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TemplateDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFields

`func (o *TemplateDetail) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TemplateDetail) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TemplateDetail) SetFields(v []Field)`

SetFields sets Fields field to given value.


### GetSignaturePlacements

`func (o *TemplateDetail) GetSignaturePlacements() []SignaturePlacement`

GetSignaturePlacements returns the SignaturePlacements field if non-nil, zero value otherwise.

### GetSignaturePlacementsOk

`func (o *TemplateDetail) GetSignaturePlacementsOk() (*[]SignaturePlacement, bool)`

GetSignaturePlacementsOk returns a tuple with the SignaturePlacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturePlacements

`func (o *TemplateDetail) SetSignaturePlacements(v []SignaturePlacement)`

SetSignaturePlacements sets SignaturePlacements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


