# GetIntegrationCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignIds** | Pointer to **[]int64** | A list of IDs of the campaigns to get coupons from. | 
**Limit** | Pointer to **int64** | The maximum number of coupons included in the response. | [default to 10]

## Methods

### NewGetIntegrationCouponRequest

`func NewGetIntegrationCouponRequest(campaignIds []int64, limit int64, ) *GetIntegrationCouponRequest`

NewGetIntegrationCouponRequest instantiates a new GetIntegrationCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationCouponRequestWithDefaults

`func NewGetIntegrationCouponRequestWithDefaults() *GetIntegrationCouponRequest`

NewGetIntegrationCouponRequestWithDefaults instantiates a new GetIntegrationCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignIds

`func (o *GetIntegrationCouponRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GetIntegrationCouponRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GetIntegrationCouponRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.


### GetLimit

`func (o *GetIntegrationCouponRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetIntegrationCouponRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetIntegrationCouponRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


