# CampaignSetLeafNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates the node type. | 
**CampaignId** | Pointer to **int64** | ID of the campaign | 

## Methods

### NewCampaignSetLeafNode

`func NewCampaignSetLeafNode(type_ string, campaignId int64, ) *CampaignSetLeafNode`

NewCampaignSetLeafNode instantiates a new CampaignSetLeafNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetLeafNodeWithDefaults

`func NewCampaignSetLeafNodeWithDefaults() *CampaignSetLeafNode`

NewCampaignSetLeafNodeWithDefaults instantiates a new CampaignSetLeafNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSetLeafNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetLeafNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSetLeafNode) SetType(v string)`

SetType sets Type field to given value.


### GetCampaignId

`func (o *CampaignSetLeafNode) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignSetLeafNode) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignSetLeafNode) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


