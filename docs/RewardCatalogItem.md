# RewardCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique ID of the reward. | 
**Name** | Pointer to **string** | The customer-facing name of the reward. | 
**Description** | Pointer to **string** | The customer-facing description of the reward. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. | [optional] 
**Rule** | Pointer to [**RuleMetadata**](RuleMetadata.md) |  | 
**Eligibility** | Pointer to [**RewardEligibility**](RewardEligibility.md) |  | [optional] 

## Methods

### NewRewardCatalogItem

`func NewRewardCatalogItem(id int64, name string, rule RuleMetadata, ) *RewardCatalogItem`

NewRewardCatalogItem instantiates a new RewardCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardCatalogItemWithDefaults

`func NewRewardCatalogItemWithDefaults() *RewardCatalogItem`

NewRewardCatalogItemWithDefaults instantiates a new RewardCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardCatalogItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardCatalogItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardCatalogItem) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *RewardCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RewardCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RewardCatalogItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RewardCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardCatalogItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPointsRequired

`func (o *RewardCatalogItem) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *RewardCatalogItem) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *RewardCatalogItem) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *RewardCatalogItem) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.

### GetRule

`func (o *RewardCatalogItem) GetRule() RuleMetadata`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RewardCatalogItem) GetRuleOk() (*RuleMetadata, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RewardCatalogItem) SetRule(v RuleMetadata)`

SetRule sets Rule field to given value.


### GetEligibility

`func (o *RewardCatalogItem) GetEligibility() RewardEligibility`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *RewardCatalogItem) GetEligibilityOk() (*RewardEligibility, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *RewardCatalogItem) SetEligibility(v RewardEligibility)`

SetEligibility sets Eligibility field to given value.

### HasEligibility

`func (o *RewardCatalogItem) HasEligibility() bool`

HasEligibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


