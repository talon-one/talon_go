# NewExperimentVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this variant. | 
**Weight** | Pointer to **int64** | The percentage split of this variant. The sum of all variant percentages must be 100. | 
**Ruleset** | Pointer to [**NewRuleset**](NewRuleset.md) |  | 
**IsPrimary** | Pointer to **bool** |  | 

## Methods

### NewNewExperimentVariant

`func NewNewExperimentVariant(name string, weight int64, ruleset NewRuleset, isPrimary bool, ) *NewExperimentVariant`

NewNewExperimentVariant instantiates a new NewExperimentVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExperimentVariantWithDefaults

`func NewNewExperimentVariantWithDefaults() *NewExperimentVariant`

NewNewExperimentVariantWithDefaults instantiates a new NewExperimentVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewExperimentVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewExperimentVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewExperimentVariant) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *NewExperimentVariant) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *NewExperimentVariant) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *NewExperimentVariant) SetWeight(v int64)`

SetWeight sets Weight field to given value.


### GetRuleset

`func (o *NewExperimentVariant) GetRuleset() NewRuleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *NewExperimentVariant) GetRulesetOk() (*NewRuleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *NewExperimentVariant) SetRuleset(v NewRuleset)`

SetRuleset sets Ruleset field to given value.


### GetIsPrimary

`func (o *NewExperimentVariant) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *NewExperimentVariant) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *NewExperimentVariant) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


