# EnvironmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slots** | Pointer to [**[]SlotDef**](SlotDef.md) | The slots defined for this application. | 
**Functions** | Pointer to [**[]FunctionDef**](FunctionDef.md) | The functions defined for this application. | 
**Templates** | Pointer to [**[]TemplateDef**](TemplateDef.md) | The templates defined for this application. | 
**Variables** | Pointer to **string** | A stringified version of the environment&#39;s Talang variables scope. | 
**GiveawaysPools** | Pointer to [**[]GiveawaysPool**](GiveawaysPool.md) | The giveaways pools that the application is subscribed to. | [optional] 
**LoyaltyPrograms** | Pointer to [**[]LoyaltyProgram**](LoyaltyProgram.md) | The loyalty programs that the application is subscribed to. | [optional] 
**Achievements** | Pointer to [**[]Achievement**](Achievement.md) | The achievements, linked to the campaigns, belonging to the application. | [optional] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) | The attributes that the application is subscribed to. | [optional] 
**AdditionalCosts** | Pointer to [**[]AccountAdditionalCost**](AccountAdditionalCost.md) | The additional costs that the application is subscribed to. | [optional] 
**Audiences** | Pointer to [**[]Audience**](Audience.md) | The audiences contained in the account which the application belongs to. | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | The account-level collections that the application is subscribed to. | [optional] 

## Methods

### GetSlots

`func (o *EnvironmentAllOf) GetSlots() []SlotDef`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *EnvironmentAllOf) GetSlotsOk() ([]SlotDef, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSlots

`func (o *EnvironmentAllOf) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### SetSlots

`func (o *EnvironmentAllOf) SetSlots(v []SlotDef)`

SetSlots gets a reference to the given []SlotDef and assigns it to the Slots field.

### GetFunctions

`func (o *EnvironmentAllOf) GetFunctions() []FunctionDef`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *EnvironmentAllOf) GetFunctionsOk() ([]FunctionDef, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFunctions

`func (o *EnvironmentAllOf) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctions

`func (o *EnvironmentAllOf) SetFunctions(v []FunctionDef)`

SetFunctions gets a reference to the given []FunctionDef and assigns it to the Functions field.

### GetTemplates

`func (o *EnvironmentAllOf) GetTemplates() []TemplateDef`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *EnvironmentAllOf) GetTemplatesOk() ([]TemplateDef, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplates

`func (o *EnvironmentAllOf) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplates

`func (o *EnvironmentAllOf) SetTemplates(v []TemplateDef)`

SetTemplates gets a reference to the given []TemplateDef and assigns it to the Templates field.

### GetVariables

`func (o *EnvironmentAllOf) GetVariables() string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *EnvironmentAllOf) GetVariablesOk() (string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariables

`func (o *EnvironmentAllOf) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariables

`func (o *EnvironmentAllOf) SetVariables(v string)`

SetVariables gets a reference to the given string and assigns it to the Variables field.

### GetGiveawaysPools

`func (o *EnvironmentAllOf) GetGiveawaysPools() []GiveawaysPool`

GetGiveawaysPools returns the GiveawaysPools field if non-nil, zero value otherwise.

### GetGiveawaysPoolsOk

`func (o *EnvironmentAllOf) GetGiveawaysPoolsOk() ([]GiveawaysPool, bool)`

GetGiveawaysPoolsOk returns a tuple with the GiveawaysPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGiveawaysPools

`func (o *EnvironmentAllOf) HasGiveawaysPools() bool`

HasGiveawaysPools returns a boolean if a field has been set.

### SetGiveawaysPools

`func (o *EnvironmentAllOf) SetGiveawaysPools(v []GiveawaysPool)`

SetGiveawaysPools gets a reference to the given []GiveawaysPool and assigns it to the GiveawaysPools field.

### GetLoyaltyPrograms

`func (o *EnvironmentAllOf) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *EnvironmentAllOf) GetLoyaltyProgramsOk() ([]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyPrograms

`func (o *EnvironmentAllOf) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### SetLoyaltyPrograms

`func (o *EnvironmentAllOf) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms gets a reference to the given []LoyaltyProgram and assigns it to the LoyaltyPrograms field.

### GetAchievements

`func (o *EnvironmentAllOf) GetAchievements() []Achievement`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *EnvironmentAllOf) GetAchievementsOk() ([]Achievement, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievements

`func (o *EnvironmentAllOf) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### SetAchievements

`func (o *EnvironmentAllOf) SetAchievements(v []Achievement)`

SetAchievements gets a reference to the given []Achievement and assigns it to the Achievements field.

### GetAttributes

`func (o *EnvironmentAllOf) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EnvironmentAllOf) GetAttributesOk() ([]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *EnvironmentAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *EnvironmentAllOf) SetAttributes(v []Attribute)`

SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.

### GetAdditionalCosts

`func (o *EnvironmentAllOf) GetAdditionalCosts() []AccountAdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *EnvironmentAllOf) GetAdditionalCostsOk() ([]AccountAdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCosts

`func (o *EnvironmentAllOf) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### SetAdditionalCosts

`func (o *EnvironmentAllOf) SetAdditionalCosts(v []AccountAdditionalCost)`

SetAdditionalCosts gets a reference to the given []AccountAdditionalCost and assigns it to the AdditionalCosts field.

### GetAudiences

`func (o *EnvironmentAllOf) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *EnvironmentAllOf) GetAudiencesOk() ([]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudiences

`func (o *EnvironmentAllOf) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### SetAudiences

`func (o *EnvironmentAllOf) SetAudiences(v []Audience)`

SetAudiences gets a reference to the given []Audience and assigns it to the Audiences field.

### GetCollections

`func (o *EnvironmentAllOf) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *EnvironmentAllOf) GetCollectionsOk() ([]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCollections

`func (o *EnvironmentAllOf) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### SetCollections

`func (o *EnvironmentAllOf) SetCollections(v []Collection)`

SetCollections gets a reference to the given []Collection and assigns it to the Collections field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


