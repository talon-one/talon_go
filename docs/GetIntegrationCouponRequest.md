# GetIntegrationCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignIds** | Pointer to **[]int32** | A list of IDs of the campaigns to get coupons from. | 
**Limit** | Pointer to **int32** | The maximum number of coupons included in the response. | [default to 10]

## Methods

### GetCampaignIds

`func (o *GetIntegrationCouponRequest) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GetIntegrationCouponRequest) GetCampaignIdsOk() ([]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignIds

`func (o *GetIntegrationCouponRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIds

`func (o *GetIntegrationCouponRequest) SetCampaignIds(v []int32)`

SetCampaignIds gets a reference to the given []int32 and assigns it to the CampaignIds field.

### GetLimit

`func (o *GetIntegrationCouponRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetIntegrationCouponRequest) GetLimitOk() (int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *GetIntegrationCouponRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *GetIntegrationCouponRequest) SetLimit(v int32)`

SetLimit gets a reference to the given int32 and assigns it to the Limit field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


