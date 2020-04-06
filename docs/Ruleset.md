# Ruleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**UserId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Rules** | Pointer to [**[]Rule**](Rule.md) | Set of rules to apply. | 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | 
**RbVersion** | Pointer to **string** | A string indicating which version of the rulebuilder was used to create this ruleset. | [optional] 
**Activate** | Pointer to **bool** | A boolean indicating whether this newly created ruleset should also be activated for the campaign owns it | [optional] 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp indicating when this Ruleset was activated. | [optional] 

## Methods

### GetId

`func (o *Ruleset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ruleset) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Ruleset) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Ruleset) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Ruleset) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Ruleset) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Ruleset) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Ruleset) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *Ruleset) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Ruleset) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Ruleset) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Ruleset) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetUserId

`func (o *Ruleset) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Ruleset) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *Ruleset) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *Ruleset) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetRules

`func (o *Ruleset) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Ruleset) GetRulesOk() ([]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRules

`func (o *Ruleset) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRules

`func (o *Ruleset) SetRules(v []Rule)`

SetRules gets a reference to the given []Rule and assigns it to the Rules field.

### GetBindings

`func (o *Ruleset) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *Ruleset) GetBindingsOk() ([]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBindings

`func (o *Ruleset) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindings

`func (o *Ruleset) SetBindings(v []Binding)`

SetBindings gets a reference to the given []Binding and assigns it to the Bindings field.

### GetRbVersion

`func (o *Ruleset) GetRbVersion() string`

GetRbVersion returns the RbVersion field if non-nil, zero value otherwise.

### GetRbVersionOk

`func (o *Ruleset) GetRbVersionOk() (string, bool)`

GetRbVersionOk returns a tuple with the RbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRbVersion

`func (o *Ruleset) HasRbVersion() bool`

HasRbVersion returns a boolean if a field has been set.

### SetRbVersion

`func (o *Ruleset) SetRbVersion(v string)`

SetRbVersion gets a reference to the given string and assigns it to the RbVersion field.

### GetActivate

`func (o *Ruleset) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *Ruleset) GetActivateOk() (bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivate

`func (o *Ruleset) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivate

`func (o *Ruleset) SetActivate(v bool)`

SetActivate gets a reference to the given bool and assigns it to the Activate field.

### GetActivatedAt

`func (o *Ruleset) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Ruleset) GetActivatedAtOk() (time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivatedAt

`func (o *Ruleset) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAt

`func (o *Ruleset) SetActivatedAt(v time.Time)`

SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


