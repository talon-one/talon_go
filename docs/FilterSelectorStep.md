# FilterSelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A step discriminator of type &#x60;filter&#x60;. | 
**Predicate** | Pointer to [**map[string]interface{}**](.md) | Describes a part of the logic of the rule. | 

## Methods

### NewFilterSelectorStep

`func NewFilterSelectorStep(type_ string, predicate map[string]interface{}, ) *FilterSelectorStep`

NewFilterSelectorStep instantiates a new FilterSelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterSelectorStepWithDefaults

`func NewFilterSelectorStepWithDefaults() *FilterSelectorStep`

NewFilterSelectorStepWithDefaults instantiates a new FilterSelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FilterSelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilterSelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilterSelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetPredicate

`func (o *FilterSelectorStep) GetPredicate() map[string]interface{}`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *FilterSelectorStep) GetPredicateOk() (*map[string]interface{}, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *FilterSelectorStep) SetPredicate(v map[string]interface{})`

SetPredicate sets Predicate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


