# ExperimentBlueprintVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this variant. | 
**Weight** | Pointer to **int64** | The percentage split of this variant. For &#x60;random&#x60; assignment, the split must be between 1 and 99 and the sum across all variants must equal 100. Ignored for &#x60;audience&#x60; and &#x60;external&#x60; assignment.  | 
**Ruleset** | Pointer to [**CatalogRuleset**](CatalogRuleset.md) |  | 
**IsPrimary** | Pointer to **bool** | Indicates whether this variant is the primary variant of the experiment. | 

## Methods

### NewExperimentBlueprintVariant

`func NewExperimentBlueprintVariant(name string, weight int64, ruleset CatalogRuleset, isPrimary bool, ) *ExperimentBlueprintVariant`

NewExperimentBlueprintVariant instantiates a new ExperimentBlueprintVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentBlueprintVariantWithDefaults

`func NewExperimentBlueprintVariantWithDefaults() *ExperimentBlueprintVariant`

NewExperimentBlueprintVariantWithDefaults instantiates a new ExperimentBlueprintVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExperimentBlueprintVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentBlueprintVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentBlueprintVariant) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *ExperimentBlueprintVariant) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExperimentBlueprintVariant) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExperimentBlueprintVariant) SetWeight(v int64)`

SetWeight sets Weight field to given value.


### GetRuleset

`func (o *ExperimentBlueprintVariant) GetRuleset() CatalogRuleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ExperimentBlueprintVariant) GetRulesetOk() (*CatalogRuleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ExperimentBlueprintVariant) SetRuleset(v CatalogRuleset)`

SetRuleset sets Ruleset field to given value.


### GetIsPrimary

`func (o *ExperimentBlueprintVariant) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ExperimentBlueprintVariant) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ExperimentBlueprintVariant) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


