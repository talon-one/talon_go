# Selector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the selector binding. | 
**Type** | Pointer to **string** | A binding of type &#x60;selector&#x60;. | 
**Source** | Pointer to **string** | The attribute path the pipeline draws items from. | 
**Steps** | Pointer to **[]map[string]interface{}** | Ordered pipeline steps applied to the source items. | 

## Methods

### NewSelector

`func NewSelector(name string, type_ string, source string, steps []map[string]interface{}, ) *Selector`

NewSelector instantiates a new Selector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorWithDefaults

`func NewSelectorWithDefaults() *Selector`

NewSelectorWithDefaults instantiates a new Selector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Selector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Selector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Selector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Selector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Selector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Selector) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *Selector) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Selector) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Selector) SetSource(v string)`

SetSource sets Source field to given value.


### GetSteps

`func (o *Selector) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Selector) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Selector) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


