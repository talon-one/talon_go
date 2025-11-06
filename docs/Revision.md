# Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | 
**ApplicationId** | Pointer to **int64** |  | 
**CampaignId** | Pointer to **int64** |  | 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**CreatedBy** | Pointer to **int64** |  | 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ActivatedBy** | Pointer to **int64** |  | [optional] 
**CurrentVersion** | Pointer to [**RevisionVersion**](RevisionVersion.md) |  | [optional] 

## Methods

### NewRevision

`func NewRevision(id int64, accountId int64, applicationId int64, campaignId int64, created time.Time, createdBy int64, ) *Revision`

NewRevision instantiates a new Revision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionWithDefaults

`func NewRevisionWithDefaults() *Revision`

NewRevisionWithDefaults instantiates a new Revision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Revision) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Revision) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Revision) SetId(v int64)`

SetId sets Id field to given value.


### GetActivateAt

`func (o *Revision) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *Revision) GetActivateAtOk() (*time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateAt

`func (o *Revision) SetActivateAt(v time.Time)`

SetActivateAt sets ActivateAt field to given value.

### HasActivateAt

`func (o *Revision) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.

### GetAccountId

`func (o *Revision) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Revision) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Revision) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetApplicationId

`func (o *Revision) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Revision) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Revision) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *Revision) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Revision) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Revision) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetCreated

`func (o *Revision) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Revision) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Revision) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *Revision) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Revision) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Revision) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetActivatedAt

`func (o *Revision) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Revision) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *Revision) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *Revision) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetActivatedBy

`func (o *Revision) GetActivatedBy() int64`

GetActivatedBy returns the ActivatedBy field if non-nil, zero value otherwise.

### GetActivatedByOk

`func (o *Revision) GetActivatedByOk() (*int64, bool)`

GetActivatedByOk returns a tuple with the ActivatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedBy

`func (o *Revision) SetActivatedBy(v int64)`

SetActivatedBy sets ActivatedBy field to given value.

### HasActivatedBy

`func (o *Revision) HasActivatedBy() bool`

HasActivatedBy returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *Revision) GetCurrentVersion() RevisionVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Revision) GetCurrentVersionOk() (*RevisionVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *Revision) SetCurrentVersion(v RevisionVersion)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *Revision) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


