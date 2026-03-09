# UpdateExperimentVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | 
**Name** | Pointer to **string** | The name of this variant. | 
**Ruleset** | Pointer to [**NewRuleset**](NewRuleset.md) |  | 
**Weight** | Pointer to **int64** | The percentage split of this variant. The sum of all variant percentages must be 100. | 

## Methods

### NewUpdateExperimentVariant

`func NewUpdateExperimentVariant(id int64, name string, ruleset NewRuleset, weight int64, ) *UpdateExperimentVariant`

NewUpdateExperimentVariant instantiates a new UpdateExperimentVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExperimentVariantWithDefaults

`func NewUpdateExperimentVariantWithDefaults() *UpdateExperimentVariant`

NewUpdateExperimentVariantWithDefaults instantiates a new UpdateExperimentVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateExperimentVariant) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateExperimentVariant) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateExperimentVariant) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateExperimentVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExperimentVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExperimentVariant) SetName(v string)`

SetName sets Name field to given value.


### GetRuleset

`func (o *UpdateExperimentVariant) GetRuleset() NewRuleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *UpdateExperimentVariant) GetRulesetOk() (*NewRuleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *UpdateExperimentVariant) SetRuleset(v NewRuleset)`

SetRuleset sets Ruleset field to given value.


### GetWeight

`func (o *UpdateExperimentVariant) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateExperimentVariant) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateExperimentVariant) SetWeight(v int64)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


