# UpdateCampaignGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the campaign access group. | 
**Description** | Pointer to **string** | A longer description of the campaign access group. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of IDs of the Applications that this campaign access group is enabled for. | [optional] 
**CampaignIds** | Pointer to **[]int64** | A list of IDs of the campaigns that are part of the campaign access group. | [optional] 

## Methods

### NewUpdateCampaignGroup

`func NewUpdateCampaignGroup(name string, ) *UpdateCampaignGroup`

NewUpdateCampaignGroup instantiates a new UpdateCampaignGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignGroupWithDefaults

`func NewUpdateCampaignGroupWithDefaults() *UpdateCampaignGroup`

NewUpdateCampaignGroupWithDefaults instantiates a new UpdateCampaignGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCampaignGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaignGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCampaignGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaignGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCampaignGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCampaignGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *UpdateCampaignGroup) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetCampaignIds

`func (o *UpdateCampaignGroup) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *UpdateCampaignGroup) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *UpdateCampaignGroup) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *UpdateCampaignGroup) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


