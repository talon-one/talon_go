# IntegrationStateV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | Pointer to [**CustomerSessionV2**](CustomerSessionV2.md) |  | [optional] 
**CustomerProfile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Event** | Pointer to [**Event**](Event.md) |  | [optional] 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Referral** | Pointer to [**Referral**](Referral.md) |  | [optional] 
**Coupons** | Pointer to [**[]Coupon**](Coupon.md) |  | [optional] 
**TriggeredCampaigns** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 
**Effects** | Pointer to [**[]Effect**](Effect.md) |  | 
**CreatedCoupons** | Pointer to [**[]Coupon**](Coupon.md) |  | 
**CreatedReferrals** | Pointer to [**[]Referral**](Referral.md) |  | 

## Methods

### GetCustomerSession

`func (o *IntegrationStateV2) GetCustomerSession() CustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationStateV2) GetCustomerSessionOk() (CustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSession

`func (o *IntegrationStateV2) HasCustomerSession() bool`

HasCustomerSession returns a boolean if a field has been set.

### SetCustomerSession

`func (o *IntegrationStateV2) SetCustomerSession(v CustomerSessionV2)`

SetCustomerSession gets a reference to the given CustomerSessionV2 and assigns it to the CustomerSession field.

### GetCustomerProfile

`func (o *IntegrationStateV2) GetCustomerProfile() CustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *IntegrationStateV2) GetCustomerProfileOk() (CustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerProfile

`func (o *IntegrationStateV2) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### SetCustomerProfile

`func (o *IntegrationStateV2) SetCustomerProfile(v CustomerProfile)`

SetCustomerProfile gets a reference to the given CustomerProfile and assigns it to the CustomerProfile field.

### GetEvent

`func (o *IntegrationStateV2) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IntegrationStateV2) GetEventOk() (Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *IntegrationStateV2) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *IntegrationStateV2) SetEvent(v Event)`

SetEvent gets a reference to the given Event and assigns it to the Event field.

### GetLoyalty

`func (o *IntegrationStateV2) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *IntegrationStateV2) GetLoyaltyOk() (Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyalty

`func (o *IntegrationStateV2) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### SetLoyalty

`func (o *IntegrationStateV2) SetLoyalty(v Loyalty)`

SetLoyalty gets a reference to the given Loyalty and assigns it to the Loyalty field.

### GetReferral

`func (o *IntegrationStateV2) GetReferral() Referral`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *IntegrationStateV2) GetReferralOk() (Referral, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferral

`func (o *IntegrationStateV2) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### SetReferral

`func (o *IntegrationStateV2) SetReferral(v Referral)`

SetReferral gets a reference to the given Referral and assigns it to the Referral field.

### GetCoupons

`func (o *IntegrationStateV2) GetCoupons() []Coupon`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *IntegrationStateV2) GetCouponsOk() ([]Coupon, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupons

`func (o *IntegrationStateV2) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### SetCoupons

`func (o *IntegrationStateV2) SetCoupons(v []Coupon)`

SetCoupons gets a reference to the given []Coupon and assigns it to the Coupons field.

### GetTriggeredCampaigns

`func (o *IntegrationStateV2) GetTriggeredCampaigns() []Campaign`

GetTriggeredCampaigns returns the TriggeredCampaigns field if non-nil, zero value otherwise.

### GetTriggeredCampaignsOk

`func (o *IntegrationStateV2) GetTriggeredCampaignsOk() ([]Campaign, bool)`

GetTriggeredCampaignsOk returns a tuple with the TriggeredCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggeredCampaigns

`func (o *IntegrationStateV2) HasTriggeredCampaigns() bool`

HasTriggeredCampaigns returns a boolean if a field has been set.

### SetTriggeredCampaigns

`func (o *IntegrationStateV2) SetTriggeredCampaigns(v []Campaign)`

SetTriggeredCampaigns gets a reference to the given []Campaign and assigns it to the TriggeredCampaigns field.

### GetEffects

`func (o *IntegrationStateV2) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *IntegrationStateV2) GetEffectsOk() ([]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *IntegrationStateV2) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *IntegrationStateV2) SetEffects(v []Effect)`

SetEffects gets a reference to the given []Effect and assigns it to the Effects field.

### GetCreatedCoupons

`func (o *IntegrationStateV2) GetCreatedCoupons() []Coupon`

GetCreatedCoupons returns the CreatedCoupons field if non-nil, zero value otherwise.

### GetCreatedCouponsOk

`func (o *IntegrationStateV2) GetCreatedCouponsOk() ([]Coupon, bool)`

GetCreatedCouponsOk returns a tuple with the CreatedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedCoupons

`func (o *IntegrationStateV2) HasCreatedCoupons() bool`

HasCreatedCoupons returns a boolean if a field has been set.

### SetCreatedCoupons

`func (o *IntegrationStateV2) SetCreatedCoupons(v []Coupon)`

SetCreatedCoupons gets a reference to the given []Coupon and assigns it to the CreatedCoupons field.

### GetCreatedReferrals

`func (o *IntegrationStateV2) GetCreatedReferrals() []Referral`

GetCreatedReferrals returns the CreatedReferrals field if non-nil, zero value otherwise.

### GetCreatedReferralsOk

`func (o *IntegrationStateV2) GetCreatedReferralsOk() ([]Referral, bool)`

GetCreatedReferralsOk returns a tuple with the CreatedReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedReferrals

`func (o *IntegrationStateV2) HasCreatedReferrals() bool`

HasCreatedReferrals returns a boolean if a field has been set.

### SetCreatedReferrals

`func (o *IntegrationStateV2) SetCreatedReferrals(v []Referral)`

SetCreatedReferrals gets a reference to the given []Referral and assigns it to the CreatedReferrals field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


