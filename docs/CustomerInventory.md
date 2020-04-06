# CustomerInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Referrals** | Pointer to [**[]Referral**](Referral.md) |  | [optional] 

## Methods

### GetProfile

`func (o *CustomerInventory) GetProfile() CustomerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CustomerInventory) GetProfileOk() (CustomerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfile

`func (o *CustomerInventory) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfile

`func (o *CustomerInventory) SetProfile(v CustomerProfile)`

SetProfile gets a reference to the given CustomerProfile and assigns it to the Profile field.

### GetReferrals

`func (o *CustomerInventory) GetReferrals() []Referral`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *CustomerInventory) GetReferralsOk() ([]Referral, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferrals

`func (o *CustomerInventory) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### SetReferrals

`func (o *CustomerInventory) SetReferrals(v []Referral)`

SetReferrals gets a reference to the given []Referral and assigns it to the Referrals field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


