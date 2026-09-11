# FilterAndMapValuesSelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A step discriminator of type &#x60;filterAndMapValues&#x60;. | 
**ValueMap** | Pointer to [**SelectorValueMapRef**](SelectorValueMapRef.md) |  | 

## Methods

### NewFilterAndMapValuesSelectorStep

`func NewFilterAndMapValuesSelectorStep(type_ string, valueMap SelectorValueMapRef, ) *FilterAndMapValuesSelectorStep`

NewFilterAndMapValuesSelectorStep instantiates a new FilterAndMapValuesSelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterAndMapValuesSelectorStepWithDefaults

`func NewFilterAndMapValuesSelectorStepWithDefaults() *FilterAndMapValuesSelectorStep`

NewFilterAndMapValuesSelectorStepWithDefaults instantiates a new FilterAndMapValuesSelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FilterAndMapValuesSelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilterAndMapValuesSelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilterAndMapValuesSelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetValueMap

`func (o *FilterAndMapValuesSelectorStep) GetValueMap() SelectorValueMapRef`

GetValueMap returns the ValueMap field if non-nil, zero value otherwise.

### GetValueMapOk

`func (o *FilterAndMapValuesSelectorStep) GetValueMapOk() (*SelectorValueMapRef, bool)`

GetValueMapOk returns a tuple with the ValueMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMap

`func (o *FilterAndMapValuesSelectorStep) SetValueMap(v SelectorValueMapRef)`

SetValueMap sets ValueMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


