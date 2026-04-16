# SignaturePlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SignerLabel** | **string** | Label identifying which signer this placement belongs to | 
**PageNumber** | **int32** |  | 
**X** | **float32** | X coordinate in PDF points from the left edge of the page | 
**Y** | **float32** | Y coordinate in PDF points from the bottom edge of the page | 
**Width** | **float32** |  | 
**Height** | **float32** |  | 

## Methods

### NewSignaturePlacement

`func NewSignaturePlacement(id string, signerLabel string, pageNumber int32, x float32, y float32, width float32, height float32, ) *SignaturePlacement`

NewSignaturePlacement instantiates a new SignaturePlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignaturePlacementWithDefaults

`func NewSignaturePlacementWithDefaults() *SignaturePlacement`

NewSignaturePlacementWithDefaults instantiates a new SignaturePlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignaturePlacement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignaturePlacement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignaturePlacement) SetId(v string)`

SetId sets Id field to given value.


### GetSignerLabel

`func (o *SignaturePlacement) GetSignerLabel() string`

GetSignerLabel returns the SignerLabel field if non-nil, zero value otherwise.

### GetSignerLabelOk

`func (o *SignaturePlacement) GetSignerLabelOk() (*string, bool)`

GetSignerLabelOk returns a tuple with the SignerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerLabel

`func (o *SignaturePlacement) SetSignerLabel(v string)`

SetSignerLabel sets SignerLabel field to given value.


### GetPageNumber

`func (o *SignaturePlacement) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SignaturePlacement) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SignaturePlacement) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetX

`func (o *SignaturePlacement) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SignaturePlacement) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SignaturePlacement) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *SignaturePlacement) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SignaturePlacement) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SignaturePlacement) SetY(v float32)`

SetY sets Y field to given value.


### GetWidth

`func (o *SignaturePlacement) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SignaturePlacement) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SignaturePlacement) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *SignaturePlacement) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SignaturePlacement) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SignaturePlacement) SetHeight(v float32)`

SetHeight sets Height field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


