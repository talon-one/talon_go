# ApplicationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Event type. It can be one of the following: [&#39;campaign_evaluation_tree_changed&#39;]  | 

## Methods

### NewApplicationNotification

`func NewApplicationNotification(event string, ) *ApplicationNotification`

NewApplicationNotification instantiates a new ApplicationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationNotificationWithDefaults

`func NewApplicationNotificationWithDefaults() *ApplicationNotification`

NewApplicationNotificationWithDefaults instantiates a new ApplicationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ApplicationNotification) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ApplicationNotification) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ApplicationNotification) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


