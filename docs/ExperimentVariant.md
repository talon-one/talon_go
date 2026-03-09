# ExperimentVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** |  | 
**ExperimentId** | Pointer to **int64** |  | [optional] 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | 

## Methods

### NewExperimentVariant

`func NewExperimentVariant(id int64, created time.Time, name string, isPrimary bool, ) *ExperimentVariant`

NewExperimentVariant instantiates a new ExperimentVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantWithDefaults

`func NewExperimentVariantWithDefaults() *ExperimentVariant`

NewExperimentVariantWithDefaults instantiates a new ExperimentVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExperimentVariant) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentVariant) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentVariant) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ExperimentVariant) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExperimentVariant) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExperimentVariant) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *ExperimentVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentVariant) SetName(v string)`

SetName sets Name field to given value.


### GetExperimentId

`func (o *ExperimentVariant) GetExperimentId() int64`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *ExperimentVariant) GetExperimentIdOk() (*int64, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *ExperimentVariant) SetExperimentId(v int64)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *ExperimentVariant) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetRuleset

`func (o *ExperimentVariant) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ExperimentVariant) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ExperimentVariant) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *ExperimentVariant) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetWeight

`func (o *ExperimentVariant) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExperimentVariant) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExperimentVariant) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ExperimentVariant) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetIsPrimary

`func (o *ExperimentVariant) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ExperimentVariant) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ExperimentVariant) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


