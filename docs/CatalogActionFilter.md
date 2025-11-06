# CatalogActionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to **string** | The name of the attribute to filter on. | 
**Op** | Pointer to **string** | The filtering operator. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | The value to filter for. | 

## Methods

### NewCatalogActionFilter

`func NewCatalogActionFilter(attr string, op string, value map[string]interface{}, ) *CatalogActionFilter`

NewCatalogActionFilter instantiates a new CatalogActionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionFilterWithDefaults

`func NewCatalogActionFilterWithDefaults() *CatalogActionFilter`

NewCatalogActionFilterWithDefaults instantiates a new CatalogActionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *CatalogActionFilter) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *CatalogActionFilter) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *CatalogActionFilter) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetOp

`func (o *CatalogActionFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CatalogActionFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CatalogActionFilter) SetOp(v string)`

SetOp sets Op field to given value.


### GetValue

`func (o *CatalogActionFilter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CatalogActionFilter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CatalogActionFilter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


