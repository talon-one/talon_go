# CampaignVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionFrontendState** | Pointer to **string** | The campaign revision state displayed in the Campaign Manager. | [optional] 
**ActiveRevisionId** | Pointer to **int64** | ID of the revision that was last activated on this campaign.  | [optional] 
**ActiveRevisionVersionId** | Pointer to **int64** | ID of the revision version that is active on the campaign.  | [optional] 
**Version** | Pointer to **int64** | Incrementing number representing how many revisions have been activated on this campaign, starts from 0 for a new campaign.  | [optional] 
**CurrentRevisionId** | Pointer to **int64** | ID of the revision currently being modified for the campaign.  | [optional] 
**CurrentRevisionVersionId** | Pointer to **int64** | ID of the latest version applied on the current revision.  | [optional] 
**StageRevision** | Pointer to **bool** | Flag for determining whether we use current revision when sending requests with staging API key.  | [optional] [default to false]

## Methods

### NewCampaignVersions

`func NewCampaignVersions() *CampaignVersions`

NewCampaignVersions instantiates a new CampaignVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignVersionsWithDefaults

`func NewCampaignVersionsWithDefaults() *CampaignVersions`

NewCampaignVersionsWithDefaults instantiates a new CampaignVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionFrontendState

`func (o *CampaignVersions) GetRevisionFrontendState() string`

GetRevisionFrontendState returns the RevisionFrontendState field if non-nil, zero value otherwise.

### GetRevisionFrontendStateOk

`func (o *CampaignVersions) GetRevisionFrontendStateOk() (*string, bool)`

GetRevisionFrontendStateOk returns a tuple with the RevisionFrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionFrontendState

`func (o *CampaignVersions) SetRevisionFrontendState(v string)`

SetRevisionFrontendState sets RevisionFrontendState field to given value.

### HasRevisionFrontendState

`func (o *CampaignVersions) HasRevisionFrontendState() bool`

HasRevisionFrontendState returns a boolean if a field has been set.

### GetActiveRevisionId

`func (o *CampaignVersions) GetActiveRevisionId() int64`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *CampaignVersions) GetActiveRevisionIdOk() (*int64, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionId

`func (o *CampaignVersions) SetActiveRevisionId(v int64)`

SetActiveRevisionId sets ActiveRevisionId field to given value.

### HasActiveRevisionId

`func (o *CampaignVersions) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.

### GetActiveRevisionVersionId

`func (o *CampaignVersions) GetActiveRevisionVersionId() int64`

GetActiveRevisionVersionId returns the ActiveRevisionVersionId field if non-nil, zero value otherwise.

### GetActiveRevisionVersionIdOk

`func (o *CampaignVersions) GetActiveRevisionVersionIdOk() (*int64, bool)`

GetActiveRevisionVersionIdOk returns a tuple with the ActiveRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionVersionId

`func (o *CampaignVersions) SetActiveRevisionVersionId(v int64)`

SetActiveRevisionVersionId sets ActiveRevisionVersionId field to given value.

### HasActiveRevisionVersionId

`func (o *CampaignVersions) HasActiveRevisionVersionId() bool`

HasActiveRevisionVersionId returns a boolean if a field has been set.

### GetVersion

`func (o *CampaignVersions) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignVersions) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignVersions) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CampaignVersions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCurrentRevisionId

`func (o *CampaignVersions) GetCurrentRevisionId() int64`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *CampaignVersions) GetCurrentRevisionIdOk() (*int64, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionId

`func (o *CampaignVersions) SetCurrentRevisionId(v int64)`

SetCurrentRevisionId sets CurrentRevisionId field to given value.

### HasCurrentRevisionId

`func (o *CampaignVersions) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### GetCurrentRevisionVersionId

`func (o *CampaignVersions) GetCurrentRevisionVersionId() int64`

GetCurrentRevisionVersionId returns the CurrentRevisionVersionId field if non-nil, zero value otherwise.

### GetCurrentRevisionVersionIdOk

`func (o *CampaignVersions) GetCurrentRevisionVersionIdOk() (*int64, bool)`

GetCurrentRevisionVersionIdOk returns a tuple with the CurrentRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionVersionId

`func (o *CampaignVersions) SetCurrentRevisionVersionId(v int64)`

SetCurrentRevisionVersionId sets CurrentRevisionVersionId field to given value.

### HasCurrentRevisionVersionId

`func (o *CampaignVersions) HasCurrentRevisionVersionId() bool`

HasCurrentRevisionVersionId returns a boolean if a field has been set.

### GetStageRevision

`func (o *CampaignVersions) GetStageRevision() bool`

GetStageRevision returns the StageRevision field if non-nil, zero value otherwise.

### GetStageRevisionOk

`func (o *CampaignVersions) GetStageRevisionOk() (*bool, bool)`

GetStageRevisionOk returns a tuple with the StageRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageRevision

`func (o *CampaignVersions) SetStageRevision(v bool)`

SetStageRevision sets StageRevision field to given value.

### HasStageRevision

`func (o *CampaignVersions) HasStageRevision() bool`

HasStageRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


