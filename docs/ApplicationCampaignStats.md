# ApplicationCampaignStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **int64** | Number of disabled campaigns. | 
**Staged** | Pointer to **int64** | Number of staged campaigns. | 
**Scheduled** | Pointer to **int64** | Number of scheduled campaigns. | 
**Running** | Pointer to **int64** | Number of running campaigns. | 
**Expired** | Pointer to **int64** | Number of expired campaigns. | 
**Archived** | Pointer to **int64** | Number of archived campaigns. | 

## Methods

### NewApplicationCampaignStats

`func NewApplicationCampaignStats(disabled int64, staged int64, scheduled int64, running int64, expired int64, archived int64, ) *ApplicationCampaignStats`

NewApplicationCampaignStats instantiates a new ApplicationCampaignStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCampaignStatsWithDefaults

`func NewApplicationCampaignStatsWithDefaults() *ApplicationCampaignStats`

NewApplicationCampaignStatsWithDefaults instantiates a new ApplicationCampaignStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *ApplicationCampaignStats) GetDisabled() int64`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApplicationCampaignStats) GetDisabledOk() (*int64, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApplicationCampaignStats) SetDisabled(v int64)`

SetDisabled sets Disabled field to given value.


### GetStaged

`func (o *ApplicationCampaignStats) GetStaged() int64`

GetStaged returns the Staged field if non-nil, zero value otherwise.

### GetStagedOk

`func (o *ApplicationCampaignStats) GetStagedOk() (*int64, bool)`

GetStagedOk returns a tuple with the Staged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaged

`func (o *ApplicationCampaignStats) SetStaged(v int64)`

SetStaged sets Staged field to given value.


### GetScheduled

`func (o *ApplicationCampaignStats) GetScheduled() int64`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *ApplicationCampaignStats) GetScheduledOk() (*int64, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *ApplicationCampaignStats) SetScheduled(v int64)`

SetScheduled sets Scheduled field to given value.


### GetRunning

`func (o *ApplicationCampaignStats) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ApplicationCampaignStats) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ApplicationCampaignStats) SetRunning(v int64)`

SetRunning sets Running field to given value.


### GetExpired

`func (o *ApplicationCampaignStats) GetExpired() int64`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApplicationCampaignStats) GetExpiredOk() (*int64, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApplicationCampaignStats) SetExpired(v int64)`

SetExpired sets Expired field to given value.


### GetArchived

`func (o *ApplicationCampaignStats) GetArchived() int64`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ApplicationCampaignStats) GetArchivedOk() (*int64, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ApplicationCampaignStats) SetArchived(v int64)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


