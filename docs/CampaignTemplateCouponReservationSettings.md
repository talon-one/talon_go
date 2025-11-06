# CampaignTemplateCouponReservationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]

## Methods

### NewCampaignTemplateCouponReservationSettings

`func NewCampaignTemplateCouponReservationSettings() *CampaignTemplateCouponReservationSettings`

NewCampaignTemplateCouponReservationSettings instantiates a new CampaignTemplateCouponReservationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignTemplateCouponReservationSettingsWithDefaults

`func NewCampaignTemplateCouponReservationSettingsWithDefaults() *CampaignTemplateCouponReservationSettings`

NewCampaignTemplateCouponReservationSettingsWithDefaults instantiates a new CampaignTemplateCouponReservationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *CampaignTemplateCouponReservationSettings) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *CampaignTemplateCouponReservationSettings) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *CampaignTemplateCouponReservationSettings) GetIsReservationMandatoryOk() (*bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory sets IsReservationMandatory field to given value.

### HasIsReservationMandatory

`func (o *CampaignTemplateCouponReservationSettings) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


