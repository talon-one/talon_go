# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the tier. | 
**Name** | Pointer to **string** | The name of the tier. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date and time when the customer moved to this tier. This value uses the loyalty program&#39;s time zone setting. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date when tier level expires in the RFC3339 format (in the Loyalty Program&#39;s timezone). | [optional] 
**DowngradePolicy** | Pointer to **string** | The policy that defines how customer tiers are downgraded in the loyalty program after tier reevaluation.  - &#x60;one_down&#x60;: If the customer doesn&#39;t have enough points to stay in the current tier, they are downgraded by one tier.  - &#x60;balance_based&#x60;: The customer&#39;s tier is reevaluated based on the amount of active points they have at the moment.  | [optional] 

## Methods

### NewTier

`func NewTier(id int64, name string, ) *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tier) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *Tier) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Tier) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Tier) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Tier) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Tier) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Tier) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Tier) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Tier) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetDowngradePolicy

`func (o *Tier) GetDowngradePolicy() string`

GetDowngradePolicy returns the DowngradePolicy field if non-nil, zero value otherwise.

### GetDowngradePolicyOk

`func (o *Tier) GetDowngradePolicyOk() (*string, bool)`

GetDowngradePolicyOk returns a tuple with the DowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradePolicy

`func (o *Tier) SetDowngradePolicy(v string)`

SetDowngradePolicy sets DowngradePolicy field to given value.

### HasDowngradePolicy

`func (o *Tier) HasDowngradePolicy() bool`

HasDowngradePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


