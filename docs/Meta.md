# Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaigns** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Maps each evaluated campaign ID to a key-value list of that campaigns attributes. Campaigns without attributes will be omitted. | [optional] 
**Coupons** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Maps the coupon value to a key-value list of that coupons attributes. | [optional] 
**CouponRejectionReason** | Pointer to [**CouponRejectionReason**](CouponRejectionReason.md) |  | [optional] 
**ReferralRejectionReason** | Pointer to [**ReferralRejectionReason**](ReferralRejectionReason.md) |  | [optional] 
**Warnings** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Contains warnings about possible misuse. | [optional] 

## Methods

### GetCampaigns

`func (o *Meta) GetCampaigns() map[string]map[string]interface{}`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *Meta) GetCampaignsOk() (map[string]map[string]interface{}, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaigns

`func (o *Meta) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaigns

`func (o *Meta) SetCampaigns(v map[string]map[string]interface{})`

SetCampaigns gets a reference to the given map[string]map[string]interface{} and assigns it to the Campaigns field.

### GetCoupons

`func (o *Meta) GetCoupons() map[string]map[string]interface{}`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *Meta) GetCouponsOk() (map[string]map[string]interface{}, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupons

`func (o *Meta) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### SetCoupons

`func (o *Meta) SetCoupons(v map[string]map[string]interface{})`

SetCoupons gets a reference to the given map[string]map[string]interface{} and assigns it to the Coupons field.

### GetCouponRejectionReason

`func (o *Meta) GetCouponRejectionReason() CouponRejectionReason`

GetCouponRejectionReason returns the CouponRejectionReason field if non-nil, zero value otherwise.

### GetCouponRejectionReasonOk

`func (o *Meta) GetCouponRejectionReasonOk() (CouponRejectionReason, bool)`

GetCouponRejectionReasonOk returns a tuple with the CouponRejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRejectionReason

`func (o *Meta) HasCouponRejectionReason() bool`

HasCouponRejectionReason returns a boolean if a field has been set.

### SetCouponRejectionReason

`func (o *Meta) SetCouponRejectionReason(v CouponRejectionReason)`

SetCouponRejectionReason gets a reference to the given CouponRejectionReason and assigns it to the CouponRejectionReason field.

### GetReferralRejectionReason

`func (o *Meta) GetReferralRejectionReason() ReferralRejectionReason`

GetReferralRejectionReason returns the ReferralRejectionReason field if non-nil, zero value otherwise.

### GetReferralRejectionReasonOk

`func (o *Meta) GetReferralRejectionReasonOk() (ReferralRejectionReason, bool)`

GetReferralRejectionReasonOk returns a tuple with the ReferralRejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRejectionReason

`func (o *Meta) HasReferralRejectionReason() bool`

HasReferralRejectionReason returns a boolean if a field has been set.

### SetReferralRejectionReason

`func (o *Meta) SetReferralRejectionReason(v ReferralRejectionReason)`

SetReferralRejectionReason gets a reference to the given ReferralRejectionReason and assigns it to the ReferralRejectionReason field.

### GetWarnings

`func (o *Meta) GetWarnings() map[string]map[string]interface{}`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Meta) GetWarningsOk() (map[string]map[string]interface{}, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWarnings

`func (o *Meta) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarnings

`func (o *Meta) SetWarnings(v map[string]map[string]interface{})`

SetWarnings gets a reference to the given map[string]map[string]interface{} and assigns it to the Warnings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


