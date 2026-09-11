# UpdateRiskNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The entity type to analyze within the given time frame. | 
**Activity** | Pointer to **string** | The activity metric to analyze within the given entity. | 
**TimeFrame** | Pointer to **string** | The rolling time window for risk evaluation. | 
**Active** | Pointer to **bool** | Indicates whether this risk notification is active. | 

## Methods

### NewUpdateRiskNotification

`func NewUpdateRiskNotification(entity string, activity string, timeFrame string, active bool, ) *UpdateRiskNotification`

NewUpdateRiskNotification instantiates a new UpdateRiskNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRiskNotificationWithDefaults

`func NewUpdateRiskNotificationWithDefaults() *UpdateRiskNotification`

NewUpdateRiskNotificationWithDefaults instantiates a new UpdateRiskNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *UpdateRiskNotification) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *UpdateRiskNotification) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *UpdateRiskNotification) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetActivity

`func (o *UpdateRiskNotification) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *UpdateRiskNotification) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *UpdateRiskNotification) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetTimeFrame

`func (o *UpdateRiskNotification) GetTimeFrame() string`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *UpdateRiskNotification) GetTimeFrameOk() (*string, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *UpdateRiskNotification) SetTimeFrame(v string)`

SetTimeFrame sets TimeFrame field to given value.


### GetActive

`func (o *UpdateRiskNotification) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateRiskNotification) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateRiskNotification) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


