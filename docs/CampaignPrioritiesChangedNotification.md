# CampaignPrioritiesChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Application**](Application.md) |  | 
**OldPriorities** | Pointer to [**map[string][]int32**](array.md) | Campaign IDs for each priority. The priority can be one of: [&#39;universal&#39;, &#39;stackable&#39;, &#39;exclusive&#39;]  | [optional] 
**Priorities** | Pointer to [**map[string][]int32**](array.md) | Campaign IDs for each priority. The priority can be one of: [&#39;universal&#39;, &#39;stackable&#39;, &#39;exclusive&#39;]  | 

## Methods

### GetApplication

`func (o *CampaignPrioritiesChangedNotification) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CampaignPrioritiesChangedNotification) GetApplicationOk() (Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplication

`func (o *CampaignPrioritiesChangedNotification) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplication

`func (o *CampaignPrioritiesChangedNotification) SetApplication(v Application)`

SetApplication gets a reference to the given Application and assigns it to the Application field.

### GetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) GetOldPriorities() map[string][]int32`

GetOldPriorities returns the OldPriorities field if non-nil, zero value otherwise.

### GetOldPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetOldPrioritiesOk() (map[string][]int32, bool)`

GetOldPrioritiesOk returns a tuple with the OldPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldPriorities

`func (o *CampaignPrioritiesChangedNotification) HasOldPriorities() bool`

HasOldPriorities returns a boolean if a field has been set.

### SetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) SetOldPriorities(v map[string][]int32)`

SetOldPriorities gets a reference to the given map[string][]int32 and assigns it to the OldPriorities field.

### GetPriorities

`func (o *CampaignPrioritiesChangedNotification) GetPriorities() map[string][]int32`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetPrioritiesOk() (map[string][]int32, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriorities

`func (o *CampaignPrioritiesChangedNotification) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPriorities

`func (o *CampaignPrioritiesChangedNotification) SetPriorities(v map[string][]int32)`

SetPriorities gets a reference to the given map[string][]int32 and assigns it to the Priorities field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


