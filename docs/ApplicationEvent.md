# ApplicationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**ProfileId** | Pointer to **int64** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**StoreId** | Pointer to **int64** | The ID of the store. | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**SessionId** | Pointer to **int64** | The globally unique Talon.One ID of the session that contains this event. | [optional] 
**Type** | Pointer to **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Additional JSON serialized data associated with the event. | 
**Effects** | Pointer to [**[]Effect**](Effect.md) | An array containing the effects that were applied as a result of this event. | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) | An array containing the rule failure reasons which happened during this event. | [optional] 

## Methods

### NewApplicationEvent

`func NewApplicationEvent(id int64, created time.Time, applicationId int64, type_ string, attributes map[string]interface{}, effects []Effect, ) *ApplicationEvent`

NewApplicationEvent instantiates a new ApplicationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEventWithDefaults

`func NewApplicationEventWithDefaults() *ApplicationEvent`

NewApplicationEventWithDefaults instantiates a new ApplicationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationEvent) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ApplicationEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *ApplicationEvent) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationEvent) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationEvent) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *ApplicationEvent) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationEvent) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ApplicationEvent) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ApplicationEvent) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetStoreId

`func (o *ApplicationEvent) GetStoreId() int64`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ApplicationEvent) GetStoreIdOk() (*int64, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ApplicationEvent) SetStoreId(v int64)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ApplicationEvent) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetStoreIntegrationId

`func (o *ApplicationEvent) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *ApplicationEvent) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *ApplicationEvent) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *ApplicationEvent) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetSessionId

`func (o *ApplicationEvent) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationEvent) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApplicationEvent) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ApplicationEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationEvent) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEvent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationEvent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetEffects

`func (o *ApplicationEvent) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ApplicationEvent) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *ApplicationEvent) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.


### GetRuleFailureReasons

`func (o *ApplicationEvent) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *ApplicationEvent) GetRuleFailureReasonsOk() (*[]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleFailureReasons

`func (o *ApplicationEvent) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons sets RuleFailureReasons field to given value.

### HasRuleFailureReasons

`func (o *ApplicationEvent) HasRuleFailureReasons() bool`

HasRuleFailureReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


