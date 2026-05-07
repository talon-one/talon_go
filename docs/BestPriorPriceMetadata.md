# BestPriorPriceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfluencingCampaignDetails** | Pointer to [**[]InfluencingCampaignDetails**](InfluencingCampaignDetails.md) | Details about campaigns that influenced the final price. | 
**AdjustmentDetails** | Pointer to [**AdjustmentDetails**](AdjustmentDetails.md) |  | [optional] 

## Methods

### NewBestPriorPriceMetadata

`func NewBestPriorPriceMetadata(influencingCampaignDetails []InfluencingCampaignDetails, ) *BestPriorPriceMetadata`

NewBestPriorPriceMetadata instantiates a new BestPriorPriceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorPriceMetadataWithDefaults

`func NewBestPriorPriceMetadataWithDefaults() *BestPriorPriceMetadata`

NewBestPriorPriceMetadataWithDefaults instantiates a new BestPriorPriceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfluencingCampaignDetails

`func (o *BestPriorPriceMetadata) GetInfluencingCampaignDetails() []InfluencingCampaignDetails`

GetInfluencingCampaignDetails returns the InfluencingCampaignDetails field if non-nil, zero value otherwise.

### GetInfluencingCampaignDetailsOk

`func (o *BestPriorPriceMetadata) GetInfluencingCampaignDetailsOk() (*[]InfluencingCampaignDetails, bool)`

GetInfluencingCampaignDetailsOk returns a tuple with the InfluencingCampaignDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingCampaignDetails

`func (o *BestPriorPriceMetadata) SetInfluencingCampaignDetails(v []InfluencingCampaignDetails)`

SetInfluencingCampaignDetails sets InfluencingCampaignDetails field to given value.


### GetAdjustmentDetails

`func (o *BestPriorPriceMetadata) GetAdjustmentDetails() AdjustmentDetails`

GetAdjustmentDetails returns the AdjustmentDetails field if non-nil, zero value otherwise.

### GetAdjustmentDetailsOk

`func (o *BestPriorPriceMetadata) GetAdjustmentDetailsOk() (*AdjustmentDetails, bool)`

GetAdjustmentDetailsOk returns a tuple with the AdjustmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDetails

`func (o *BestPriorPriceMetadata) SetAdjustmentDetails(v AdjustmentDetails)`

SetAdjustmentDetails sets AdjustmentDetails field to given value.

### HasAdjustmentDetails

`func (o *BestPriorPriceMetadata) HasAdjustmentDetails() bool`

HasAdjustmentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


