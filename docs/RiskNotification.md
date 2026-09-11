# RiskNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Entity** | Pointer to **string** | The entity type to analyze within the given time frame. | 
**Activity** | Pointer to **string** | The activity metric to analyze within the given entity. | 
**TimeFrame** | Pointer to **string** | The rolling time window for risk evaluation. | 
**Active** | Pointer to **bool** | Indicates whether this risk notification is active. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update. | 

## Methods

### NewRiskNotification

`func NewRiskNotification(id int64, created time.Time, entity string, activity string, timeFrame string, active bool, modified time.Time, ) *RiskNotification`

NewRiskNotification instantiates a new RiskNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskNotificationWithDefaults

`func NewRiskNotificationWithDefaults() *RiskNotification`

NewRiskNotificationWithDefaults instantiates a new RiskNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskNotification) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskNotification) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskNotification) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RiskNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskNotification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskNotification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEntity

`func (o *RiskNotification) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RiskNotification) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RiskNotification) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetActivity

`func (o *RiskNotification) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *RiskNotification) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *RiskNotification) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetTimeFrame

`func (o *RiskNotification) GetTimeFrame() string`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *RiskNotification) GetTimeFrameOk() (*string, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *RiskNotification) SetTimeFrame(v string)`

SetTimeFrame sets TimeFrame field to given value.


### GetActive

`func (o *RiskNotification) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RiskNotification) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RiskNotification) SetActive(v bool)`

SetActive sets Active field to given value.


### GetModified

`func (o *RiskNotification) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RiskNotification) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RiskNotification) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


