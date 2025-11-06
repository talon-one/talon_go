# CampaignEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | 

## Methods

### NewCampaignEntity

`func NewCampaignEntity(campaignId int64, ) *CampaignEntity`

NewCampaignEntity instantiates a new CampaignEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEntityWithDefaults

`func NewCampaignEntityWithDefaults() *CampaignEntity`

NewCampaignEntityWithDefaults instantiates a new CampaignEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *CampaignEntity) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignEntity) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignEntity) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


