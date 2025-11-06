# BaseNotificationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**map[string]interface{}**](.md) | Indicates which notification properties to apply. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]

## Methods

### NewBaseNotificationEntity

`func NewBaseNotificationEntity(policy map[string]interface{}, ) *BaseNotificationEntity`

NewBaseNotificationEntity instantiates a new BaseNotificationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotificationEntityWithDefaults

`func NewBaseNotificationEntityWithDefaults() *BaseNotificationEntity`

NewBaseNotificationEntityWithDefaults instantiates a new BaseNotificationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *BaseNotificationEntity) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BaseNotificationEntity) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BaseNotificationEntity) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.


### GetEnabled

`func (o *BaseNotificationEntity) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseNotificationEntity) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseNotificationEntity) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseNotificationEntity) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


