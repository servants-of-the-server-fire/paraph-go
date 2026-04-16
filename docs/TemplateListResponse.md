# TemplateListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to [**[]Template**](Template.md) |  | [optional] 
**ListInfo** | Pointer to [**ListInfo**](ListInfo.md) |  | [optional] 

## Methods

### NewTemplateListResponse

`func NewTemplateListResponse() *TemplateListResponse`

NewTemplateListResponse instantiates a new TemplateListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateListResponseWithDefaults

`func NewTemplateListResponseWithDefaults() *TemplateListResponse`

NewTemplateListResponseWithDefaults instantiates a new TemplateListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *TemplateListResponse) GetTemplates() []Template`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *TemplateListResponse) GetTemplatesOk() (*[]Template, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *TemplateListResponse) SetTemplates(v []Template)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *TemplateListResponse) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetListInfo

`func (o *TemplateListResponse) GetListInfo() ListInfo`

GetListInfo returns the ListInfo field if non-nil, zero value otherwise.

### GetListInfoOk

`func (o *TemplateListResponse) GetListInfoOk() (*ListInfo, bool)`

GetListInfoOk returns a tuple with the ListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListInfo

`func (o *TemplateListResponse) SetListInfo(v ListInfo)`

SetListInfo sets ListInfo field to given value.

### HasListInfo

`func (o *TemplateListResponse) HasListInfo() bool`

HasListInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


