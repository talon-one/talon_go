# ApplicationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**ProfileId** | Pointer to **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**SessionId** | Pointer to **int32** | The globally unique Talon.One ID of the session that contains this event. | [optional] 
**Type** | Pointer to **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Additional JSON serialized data associated with the event. | 
**Effects** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | An array containing the effects that were applied as a result of this event. | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) | An array containing the rule failure reasons which happened during this event. | [optional] 

## Methods

### GetId

`func (o *ApplicationEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationEvent) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationEvent) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *ApplicationEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationEvent) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationEvent) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationEvent) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *ApplicationEvent) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationEvent) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationEvent) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationEvent) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetProfileId

`func (o *ApplicationEvent) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationEvent) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *ApplicationEvent) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *ApplicationEvent) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetSessionId

`func (o *ApplicationEvent) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationEvent) GetSessionIdOk() (int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *ApplicationEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *ApplicationEvent) SetSessionId(v int32)`

SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.

### GetType

`func (o *ApplicationEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationEvent) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ApplicationEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ApplicationEvent) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAttributes

`func (o *ApplicationEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEvent) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationEvent) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetEffects

`func (o *ApplicationEvent) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ApplicationEvent) GetEffectsOk() ([]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *ApplicationEvent) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *ApplicationEvent) SetEffects(v []map[string]interface{})`

SetEffects gets a reference to the given []map[string]interface{} and assigns it to the Effects field.

### GetRuleFailureReasons

`func (o *ApplicationEvent) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *ApplicationEvent) GetRuleFailureReasonsOk() ([]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleFailureReasons

`func (o *ApplicationEvent) HasRuleFailureReasons() bool`

HasRuleFailureReasons returns a boolean if a field has been set.

### SetRuleFailureReasons

`func (o *ApplicationEvent) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons gets a reference to the given []RuleFailureReason and assigns it to the RuleFailureReasons field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


