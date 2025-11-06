# CampaignCollectionWithoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**Description** | Pointer to **string** | A short description of the purpose of this collection. | [optional] 
**Name** | Pointer to **string** | The name of this collection. | 
**ModifiedBy** | Pointer to **int64** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | Pointer to **int64** | ID of the user who created this effect. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | [optional] 
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | [optional] 

## Methods

### NewCampaignCollectionWithoutPayload

`func NewCampaignCollectionWithoutPayload(id int64, created time.Time, accountId int64, modified time.Time, name string, createdBy int64, ) *CampaignCollectionWithoutPayload`

NewCampaignCollectionWithoutPayload instantiates a new CampaignCollectionWithoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCollectionWithoutPayloadWithDefaults

`func NewCampaignCollectionWithoutPayloadWithDefaults() *CampaignCollectionWithoutPayload`

NewCampaignCollectionWithoutPayloadWithDefaults instantiates a new CampaignCollectionWithoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignCollectionWithoutPayload) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignCollectionWithoutPayload) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignCollectionWithoutPayload) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignCollectionWithoutPayload) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignCollectionWithoutPayload) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignCollectionWithoutPayload) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *CampaignCollectionWithoutPayload) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignCollectionWithoutPayload) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CampaignCollectionWithoutPayload) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetModified

`func (o *CampaignCollectionWithoutPayload) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignCollectionWithoutPayload) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CampaignCollectionWithoutPayload) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetDescription

`func (o *CampaignCollectionWithoutPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignCollectionWithoutPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignCollectionWithoutPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignCollectionWithoutPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CampaignCollectionWithoutPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCollectionWithoutPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCollectionWithoutPayload) SetName(v string)`

SetName sets Name field to given value.


### GetModifiedBy

`func (o *CampaignCollectionWithoutPayload) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CampaignCollectionWithoutPayload) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CampaignCollectionWithoutPayload) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CampaignCollectionWithoutPayload) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CampaignCollectionWithoutPayload) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignCollectionWithoutPayload) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignCollectionWithoutPayload) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetApplicationId

`func (o *CampaignCollectionWithoutPayload) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignCollectionWithoutPayload) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignCollectionWithoutPayload) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CampaignCollectionWithoutPayload) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCampaignId

`func (o *CampaignCollectionWithoutPayload) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignCollectionWithoutPayload) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignCollectionWithoutPayload) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CampaignCollectionWithoutPayload) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


