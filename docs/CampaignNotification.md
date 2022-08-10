# CampaignNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 

## Methods

### GetEvent

`func (o *CampaignNotification) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignNotification) GetEventOk() (string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *CampaignNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *CampaignNotification) SetEvent(v string)`

SetEvent gets a reference to the given string and assigns it to the Event field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


