# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of the tier. | 
**Name** | Pointer to **string** | The name of the tier. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date and time when the customer moved to this tier. This value uses the loyalty program&#39;s time zone setting. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date when tier level expires in the RFC3339 format (in the Loyalty Program&#39;s timezone). | [optional] 
**DowngradePolicy** | Pointer to **string** | The policy that defines how customer tiers are downgraded in the loyalty program after tier reevaluation.  - &#x60;one_down&#x60;: If the customer doesn&#39;t have enough points to stay in the current tier, they are downgraded by one tier.  - &#x60;balance_based&#x60;: The customer&#39;s tier is reevaluated based on the amount of active points they have at the moment.  | [optional] 

## Methods

### GetId

`func (o *Tier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tier) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Tier) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Tier) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Tier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartDate

`func (o *Tier) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Tier) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Tier) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Tier) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *Tier) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Tier) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *Tier) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *Tier) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetDowngradePolicy

`func (o *Tier) GetDowngradePolicy() string`

GetDowngradePolicy returns the DowngradePolicy field if non-nil, zero value otherwise.

### GetDowngradePolicyOk

`func (o *Tier) GetDowngradePolicyOk() (string, bool)`

GetDowngradePolicyOk returns a tuple with the DowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDowngradePolicy

`func (o *Tier) HasDowngradePolicy() bool`

HasDowngradePolicy returns a boolean if a field has been set.

### SetDowngradePolicy

`func (o *Tier) SetDowngradePolicy(v string)`

SetDowngradePolicy gets a reference to the given string and assigns it to the DowngradePolicy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


