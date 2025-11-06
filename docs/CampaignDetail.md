# CampaignDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | The ID of the campaign that references the application cart item filter. | [optional] 
**CampaignName** | Pointer to **string** | A user-facing name for this campaign. | [optional] 

## Methods

### NewCampaignDetail

`func NewCampaignDetail() *CampaignDetail`

NewCampaignDetail instantiates a new CampaignDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailWithDefaults

`func NewCampaignDetailWithDefaults() *CampaignDetail`

NewCampaignDetailWithDefaults instantiates a new CampaignDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *CampaignDetail) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignDetail) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignDetail) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CampaignDetail) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetCampaignName

`func (o *CampaignDetail) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CampaignDetail) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *CampaignDetail) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.

### HasCampaignName

`func (o *CampaignDetail) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


