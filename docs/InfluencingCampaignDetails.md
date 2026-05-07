# InfluencingCampaignDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | Identifier of the campaign that influenced the final price. | 
**DiscountValue** | Pointer to **float32** | Discount value applied by the campaign. | 

## Methods

### NewInfluencingCampaignDetails

`func NewInfluencingCampaignDetails(campaignId int64, discountValue float32, ) *InfluencingCampaignDetails`

NewInfluencingCampaignDetails instantiates a new InfluencingCampaignDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfluencingCampaignDetailsWithDefaults

`func NewInfluencingCampaignDetailsWithDefaults() *InfluencingCampaignDetails`

NewInfluencingCampaignDetailsWithDefaults instantiates a new InfluencingCampaignDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *InfluencingCampaignDetails) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *InfluencingCampaignDetails) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *InfluencingCampaignDetails) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetDiscountValue

`func (o *InfluencingCampaignDetails) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *InfluencingCampaignDetails) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *InfluencingCampaignDetails) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


