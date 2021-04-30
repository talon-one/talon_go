# CustomerInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Referrals** | Pointer to [**[]Referral**](Referral.md) |  | [optional] 
**Coupons** | Pointer to [**[]InventoryCoupon**](InventoryCoupon.md) |  | [optional] 
**Giveaways** | Pointer to [**[]Giveaway**](Giveaway.md) |  | [optional] 

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

### GetLoyalty

`func (o *CustomerInventory) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *CustomerInventory) GetLoyaltyOk() (Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyalty

`func (o *CustomerInventory) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### SetLoyalty

`func (o *CustomerInventory) SetLoyalty(v Loyalty)`

SetLoyalty gets a reference to the given Loyalty and assigns it to the Loyalty field.

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

### GetCoupons

`func (o *CustomerInventory) GetCoupons() []InventoryCoupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *CustomerInventory) GetCouponsOk() ([]InventoryCoupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupons

`func (o *CustomerInventory) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### SetCoupons

`func (o *CustomerInventory) SetCoupons(v []InventoryCoupon)`

SetCoupons gets a reference to the given []InventoryCoupon and assigns it to the Coupons field.

### GetGiveaways

`func (o *CustomerInventory) GetGiveaways() []Giveaway`

GetGiveaways returns the Giveaways field if non-nil, zero value otherwise.

### GetGiveawaysOk

`func (o *CustomerInventory) GetGiveawaysOk() ([]Giveaway, bool)`

GetGiveawaysOk returns a tuple with the Giveaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGiveaways

`func (o *CustomerInventory) HasGiveaways() bool`

HasGiveaways returns a boolean if a field has been set.

### SetGiveaways

`func (o *CustomerInventory) SetGiveaways(v []Giveaway)`

SetGiveaways gets a reference to the given []Giveaway and assigns it to the Giveaways field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


