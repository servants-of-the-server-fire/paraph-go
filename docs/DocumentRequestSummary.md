# DocumentRequestSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TemplateId** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**HasSigning** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDocumentRequestSummary

`func NewDocumentRequestSummary(id string, templateId string, status RequestStatus, hasSigning bool, createdAt time.Time, updatedAt time.Time, ) *DocumentRequestSummary`

NewDocumentRequestSummary instantiates a new DocumentRequestSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentRequestSummaryWithDefaults

`func NewDocumentRequestSummaryWithDefaults() *DocumentRequestSummary`

NewDocumentRequestSummaryWithDefaults instantiates a new DocumentRequestSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentRequestSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentRequestSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentRequestSummary) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *DocumentRequestSummary) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DocumentRequestSummary) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DocumentRequestSummary) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetTitle

`func (o *DocumentRequestSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentRequestSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentRequestSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentRequestSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentRequestSummary) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentRequestSummary) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentRequestSummary) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetHasSigning

`func (o *DocumentRequestSummary) GetHasSigning() bool`

GetHasSigning returns the HasSigning field if non-nil, zero value otherwise.

### GetHasSigningOk

`func (o *DocumentRequestSummary) GetHasSigningOk() (*bool, bool)`

GetHasSigningOk returns a tuple with the HasSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSigning

`func (o *DocumentRequestSummary) SetHasSigning(v bool)`

SetHasSigning sets HasSigning field to given value.


### GetCreatedAt

`func (o *DocumentRequestSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentRequestSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentRequestSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DocumentRequestSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentRequestSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentRequestSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


