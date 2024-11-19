# CampaignVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveRevisionId** | Pointer to **int32** | ID of the revision that was last activated on this campaign.  | [optional] 
**ActiveRevisionVersionId** | Pointer to **int32** | ID of the revision version that is active on the campaign.  | [optional] 
**Version** | Pointer to **int32** | Incrementing number representing how many revisions have been activated on this campaign, starts from 0 for a new campaign.  | [optional] 
**CurrentRevisionId** | Pointer to **int32** | ID of the revision currently being modified for the campaign.  | [optional] 
**CurrentRevisionVersionId** | Pointer to **int32** | ID of the latest version applied on the current revision.  | [optional] 
**StageRevision** | Pointer to **bool** | Flag for determining whether we use current revision when sending requests with staging API key.  | [optional] [default to false]

## Methods

### GetActiveRevisionId

`func (o *CampaignVersions) GetActiveRevisionId() int32`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *CampaignVersions) GetActiveRevisionIdOk() (int32, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRevisionId

`func (o *CampaignVersions) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.

### SetActiveRevisionId

`func (o *CampaignVersions) SetActiveRevisionId(v int32)`

SetActiveRevisionId gets a reference to the given int32 and assigns it to the ActiveRevisionId field.

### GetActiveRevisionVersionId

`func (o *CampaignVersions) GetActiveRevisionVersionId() int32`

GetActiveRevisionVersionId returns the ActiveRevisionVersionId field if non-nil, zero value otherwise.

### GetActiveRevisionVersionIdOk

`func (o *CampaignVersions) GetActiveRevisionVersionIdOk() (int32, bool)`

GetActiveRevisionVersionIdOk returns a tuple with the ActiveRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRevisionVersionId

`func (o *CampaignVersions) HasActiveRevisionVersionId() bool`

HasActiveRevisionVersionId returns a boolean if a field has been set.

### SetActiveRevisionVersionId

`func (o *CampaignVersions) SetActiveRevisionVersionId(v int32)`

SetActiveRevisionVersionId gets a reference to the given int32 and assigns it to the ActiveRevisionVersionId field.

### GetVersion

`func (o *CampaignVersions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignVersions) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CampaignVersions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CampaignVersions) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetCurrentRevisionId

`func (o *CampaignVersions) GetCurrentRevisionId() int32`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *CampaignVersions) GetCurrentRevisionIdOk() (int32, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentRevisionId

`func (o *CampaignVersions) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### SetCurrentRevisionId

`func (o *CampaignVersions) SetCurrentRevisionId(v int32)`

SetCurrentRevisionId gets a reference to the given int32 and assigns it to the CurrentRevisionId field.

### GetCurrentRevisionVersionId

`func (o *CampaignVersions) GetCurrentRevisionVersionId() int32`

GetCurrentRevisionVersionId returns the CurrentRevisionVersionId field if non-nil, zero value otherwise.

### GetCurrentRevisionVersionIdOk

`func (o *CampaignVersions) GetCurrentRevisionVersionIdOk() (int32, bool)`

GetCurrentRevisionVersionIdOk returns a tuple with the CurrentRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentRevisionVersionId

`func (o *CampaignVersions) HasCurrentRevisionVersionId() bool`

HasCurrentRevisionVersionId returns a boolean if a field has been set.

### SetCurrentRevisionVersionId

`func (o *CampaignVersions) SetCurrentRevisionVersionId(v int32)`

SetCurrentRevisionVersionId gets a reference to the given int32 and assigns it to the CurrentRevisionVersionId field.

### GetStageRevision

`func (o *CampaignVersions) GetStageRevision() bool`

GetStageRevision returns the StageRevision field if non-nil, zero value otherwise.

### GetStageRevisionOk

`func (o *CampaignVersions) GetStageRevisionOk() (bool, bool)`

GetStageRevisionOk returns a tuple with the StageRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStageRevision

`func (o *CampaignVersions) HasStageRevision() bool`

HasStageRevision returns a boolean if a field has been set.

### SetStageRevision

`func (o *CampaignVersions) SetStageRevision(v bool)`

SetStageRevision gets a reference to the given bool and assigns it to the StageRevision field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


