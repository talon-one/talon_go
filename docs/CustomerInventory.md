# CustomerInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Referrals** | Pointer to [**[]InventoryReferral**](InventoryReferral.md) |  | [optional] 
**Coupons** | Pointer to [**[]InventoryCoupon**](InventoryCoupon.md) | The coupons reserved by this profile. This array includes hard and soft reservations.  | [optional] 
**Giveaways** | Pointer to [**[]Giveaway**](Giveaway.md) |  | [optional] 
**Achievements** | Pointer to [**[]AchievementProgressWithDefinition**](AchievementProgressWithDefinition.md) |  | [optional] 

## Methods

### NewCustomerInventory

`func NewCustomerInventory() *CustomerInventory`

NewCustomerInventory instantiates a new CustomerInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInventoryWithDefaults

`func NewCustomerInventoryWithDefaults() *CustomerInventory`

NewCustomerInventoryWithDefaults instantiates a new CustomerInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *CustomerInventory) GetProfile() CustomerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CustomerInventory) GetProfileOk() (*CustomerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CustomerInventory) SetProfile(v CustomerProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CustomerInventory) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLoyalty

`func (o *CustomerInventory) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *CustomerInventory) GetLoyaltyOk() (*Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *CustomerInventory) SetLoyalty(v Loyalty)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *CustomerInventory) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetReferrals

`func (o *CustomerInventory) GetReferrals() []InventoryReferral`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *CustomerInventory) GetReferralsOk() (*[]InventoryReferral, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *CustomerInventory) SetReferrals(v []InventoryReferral)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *CustomerInventory) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetCoupons

`func (o *CustomerInventory) GetCoupons() []InventoryCoupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *CustomerInventory) GetCouponsOk() (*[]InventoryCoupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *CustomerInventory) SetCoupons(v []InventoryCoupon)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *CustomerInventory) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetGiveaways

`func (o *CustomerInventory) GetGiveaways() []Giveaway`

GetGiveaways returns the Giveaways field if non-nil, zero value otherwise.

### GetGiveawaysOk

`func (o *CustomerInventory) GetGiveawaysOk() (*[]Giveaway, bool)`

GetGiveawaysOk returns a tuple with the Giveaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveaways

`func (o *CustomerInventory) SetGiveaways(v []Giveaway)`

SetGiveaways sets Giveaways field to given value.

### HasGiveaways

`func (o *CustomerInventory) HasGiveaways() bool`

HasGiveaways returns a boolean if a field has been set.

### GetAchievements

`func (o *CustomerInventory) GetAchievements() []AchievementProgressWithDefinition`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *CustomerInventory) GetAchievementsOk() (*[]AchievementProgressWithDefinition, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *CustomerInventory) SetAchievements(v []AchievementProgressWithDefinition)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *CustomerInventory) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


