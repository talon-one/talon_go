# CampaignSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Version** | Pointer to **int64** | Version of the campaign set. | 
**Set** | Pointer to [**CampaignSetBranchNode**](CampaignSetBranchNode.md) |  | 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign set, if available. | [optional] 

## Methods

### NewCampaignSet

`func NewCampaignSet(applicationId int64, id int64, version int64, set CampaignSetBranchNode, ) *CampaignSet`

NewCampaignSet instantiates a new CampaignSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetWithDefaults

`func NewCampaignSetWithDefaults() *CampaignSet`

NewCampaignSetWithDefaults instantiates a new CampaignSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CampaignSet) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignSet) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignSet) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetId

`func (o *CampaignSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSet) SetId(v int64)`

SetId sets Id field to given value.


### GetVersion

`func (o *CampaignSet) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignSet) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignSet) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetSet

`func (o *CampaignSet) GetSet() CampaignSetBranchNode`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *CampaignSet) GetSetOk() (*CampaignSetBranchNode, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *CampaignSet) SetSet(v CampaignSetBranchNode)`

SetSet sets Set field to given value.


### GetUpdatedBy

`func (o *CampaignSet) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignSet) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CampaignSet) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CampaignSet) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


