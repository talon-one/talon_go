# CampaignTemplateCouponReservationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReservationLimit** | Pointer to **int32** | The number of reservations that can be made with this coupon code.  | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]

## Methods

### GetReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) GetReservationLimit() int32`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *CampaignTemplateCouponReservationSettings) GetReservationLimitOk() (int32, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### SetReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) SetReservationLimit(v int32)`

SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.

### GetIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *CampaignTemplateCouponReservationSettings) GetIsReservationMandatoryOk() (bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### SetIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


