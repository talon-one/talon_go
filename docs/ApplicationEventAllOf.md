# ApplicationEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **int32** | The globally unique Talon.One ID of the session that contains this event. | [optional] 
**Type** | Pointer to **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Additional JSON serialized data associated with the event. | 
**Effects** | Pointer to [**[]Effect**](Effect.md) | An array containing the effects that were applied as a result of this event. | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) | An array containing the rule failure reasons which happened during this event. | [optional] 

## Methods

### GetSessionId

`func (o *ApplicationEventAllOf) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationEventAllOf) GetSessionIdOk() (int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *ApplicationEventAllOf) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *ApplicationEventAllOf) SetSessionId(v int32)`

SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.

### GetType

`func (o *ApplicationEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationEventAllOf) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ApplicationEventAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ApplicationEventAllOf) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAttributes

`func (o *ApplicationEventAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEventAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationEventAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationEventAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetEffects

`func (o *ApplicationEventAllOf) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ApplicationEventAllOf) GetEffectsOk() ([]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *ApplicationEventAllOf) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *ApplicationEventAllOf) SetEffects(v []Effect)`

SetEffects gets a reference to the given []Effect and assigns it to the Effects field.

### GetRuleFailureReasons

`func (o *ApplicationEventAllOf) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *ApplicationEventAllOf) GetRuleFailureReasonsOk() ([]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleFailureReasons

`func (o *ApplicationEventAllOf) HasRuleFailureReasons() bool`

HasRuleFailureReasons returns a boolean if a field has been set.

### SetRuleFailureReasons

`func (o *ApplicationEventAllOf) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons gets a reference to the given []RuleFailureReason and assigns it to the RuleFailureReasons field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


