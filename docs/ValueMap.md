# ValueMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CreatedBy** | Pointer to **int64** | The ID of the user who created the value map. | [optional] 
**CampaignId** | Pointer to **int64** |  | 

## Methods

### NewValueMap

`func NewValueMap(id int64, campaignId int64, ) *ValueMap`

NewValueMap instantiates a new ValueMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueMapWithDefaults

`func NewValueMapWithDefaults() *ValueMap`

NewValueMapWithDefaults instantiates a new ValueMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValueMap) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValueMap) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValueMap) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ValueMap) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ValueMap) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ValueMap) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ValueMap) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ValueMap) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ValueMap) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ValueMap) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ValueMap) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCampaignId

`func (o *ValueMap) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ValueMap) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ValueMap) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


