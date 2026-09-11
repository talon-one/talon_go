# CouponEligibilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns the coupon. | 
**CampaignName** | Pointer to **string** | The name of the campaign that owns the coupon. | 
**FailureReason** | Pointer to **string** | The reason the coupon is not eligible, if applicable. | [optional] 

## Methods

### NewCouponEligibilityInfo

`func NewCouponEligibilityInfo(campaignId int64, campaignName string, ) *CouponEligibilityInfo`

NewCouponEligibilityInfo instantiates a new CouponEligibilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponEligibilityInfoWithDefaults

`func NewCouponEligibilityInfoWithDefaults() *CouponEligibilityInfo`

NewCouponEligibilityInfoWithDefaults instantiates a new CouponEligibilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *CouponEligibilityInfo) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CouponEligibilityInfo) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CouponEligibilityInfo) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetCampaignName

`func (o *CouponEligibilityInfo) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CouponEligibilityInfo) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *CouponEligibilityInfo) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.


### GetFailureReason

`func (o *CouponEligibilityInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *CouponEligibilityInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *CouponEligibilityInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *CouponEligibilityInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


