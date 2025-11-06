# BestPriorPriceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfluencingCampaignIDs** | Pointer to **[]int64** |  | [optional] 
**AdjustmentReferenceID** | Pointer to **string** | Identifier related to the &#x60;referenceId&#x60; used during a &#x60;ADD_PRICE_ADJUSTMENT&#x60; action  using the [Sync cart item catalog endpoint](https://docs.talon.one/integration-api#tag/Catalogs/operation/syncCatalog). | [optional] 

## Methods

### NewBestPriorPriceMetadata

`func NewBestPriorPriceMetadata() *BestPriorPriceMetadata`

NewBestPriorPriceMetadata instantiates a new BestPriorPriceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorPriceMetadataWithDefaults

`func NewBestPriorPriceMetadataWithDefaults() *BestPriorPriceMetadata`

NewBestPriorPriceMetadataWithDefaults instantiates a new BestPriorPriceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfluencingCampaignIDs

`func (o *BestPriorPriceMetadata) GetInfluencingCampaignIDs() []int64`

GetInfluencingCampaignIDs returns the InfluencingCampaignIDs field if non-nil, zero value otherwise.

### GetInfluencingCampaignIDsOk

`func (o *BestPriorPriceMetadata) GetInfluencingCampaignIDsOk() (*[]int64, bool)`

GetInfluencingCampaignIDsOk returns a tuple with the InfluencingCampaignIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingCampaignIDs

`func (o *BestPriorPriceMetadata) SetInfluencingCampaignIDs(v []int64)`

SetInfluencingCampaignIDs sets InfluencingCampaignIDs field to given value.

### HasInfluencingCampaignIDs

`func (o *BestPriorPriceMetadata) HasInfluencingCampaignIDs() bool`

HasInfluencingCampaignIDs returns a boolean if a field has been set.

### GetAdjustmentReferenceID

`func (o *BestPriorPriceMetadata) GetAdjustmentReferenceID() string`

GetAdjustmentReferenceID returns the AdjustmentReferenceID field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIDOk

`func (o *BestPriorPriceMetadata) GetAdjustmentReferenceIDOk() (*string, bool)`

GetAdjustmentReferenceIDOk returns a tuple with the AdjustmentReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReferenceID

`func (o *BestPriorPriceMetadata) SetAdjustmentReferenceID(v string)`

SetAdjustmentReferenceID sets AdjustmentReferenceID field to given value.

### HasAdjustmentReferenceID

`func (o *BestPriorPriceMetadata) HasAdjustmentReferenceID() bool`

HasAdjustmentReferenceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


