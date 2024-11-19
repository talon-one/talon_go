# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
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
**ApplicationCartItemFilters** | Pointer to [**[]ApplicationCif**](ApplicationCIF.md) | The cart item filters belonging to the Application. | [optional] 

## Methods

### GetId

`func (o *Environment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Environment) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Environment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Environment) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Environment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Environment) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *Environment) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Environment) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Environment) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Environment) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetSlots

`func (o *Environment) GetSlots() []SlotDef`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *Environment) GetSlotsOk() ([]SlotDef, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSlots

`func (o *Environment) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### SetSlots

`func (o *Environment) SetSlots(v []SlotDef)`

SetSlots gets a reference to the given []SlotDef and assigns it to the Slots field.

### GetFunctions

`func (o *Environment) GetFunctions() []FunctionDef`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Environment) GetFunctionsOk() ([]FunctionDef, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFunctions

`func (o *Environment) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctions

`func (o *Environment) SetFunctions(v []FunctionDef)`

SetFunctions gets a reference to the given []FunctionDef and assigns it to the Functions field.

### GetTemplates

`func (o *Environment) GetTemplates() []TemplateDef`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Environment) GetTemplatesOk() ([]TemplateDef, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplates

`func (o *Environment) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplates

`func (o *Environment) SetTemplates(v []TemplateDef)`

SetTemplates gets a reference to the given []TemplateDef and assigns it to the Templates field.

### GetVariables

`func (o *Environment) GetVariables() string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Environment) GetVariablesOk() (string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariables

`func (o *Environment) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariables

`func (o *Environment) SetVariables(v string)`

SetVariables gets a reference to the given string and assigns it to the Variables field.

### GetGiveawaysPools

`func (o *Environment) GetGiveawaysPools() []GiveawaysPool`

GetGiveawaysPools returns the GiveawaysPools field if non-nil, zero value otherwise.

### GetGiveawaysPoolsOk

`func (o *Environment) GetGiveawaysPoolsOk() ([]GiveawaysPool, bool)`

GetGiveawaysPoolsOk returns a tuple with the GiveawaysPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGiveawaysPools

`func (o *Environment) HasGiveawaysPools() bool`

HasGiveawaysPools returns a boolean if a field has been set.

### SetGiveawaysPools

`func (o *Environment) SetGiveawaysPools(v []GiveawaysPool)`

SetGiveawaysPools gets a reference to the given []GiveawaysPool and assigns it to the GiveawaysPools field.

### GetLoyaltyPrograms

`func (o *Environment) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *Environment) GetLoyaltyProgramsOk() ([]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyPrograms

`func (o *Environment) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### SetLoyaltyPrograms

`func (o *Environment) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms gets a reference to the given []LoyaltyProgram and assigns it to the LoyaltyPrograms field.

### GetAchievements

`func (o *Environment) GetAchievements() []Achievement`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *Environment) GetAchievementsOk() ([]Achievement, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievements

`func (o *Environment) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### SetAchievements

`func (o *Environment) SetAchievements(v []Achievement)`

SetAchievements gets a reference to the given []Achievement and assigns it to the Achievements field.

### GetAttributes

`func (o *Environment) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Environment) GetAttributesOk() ([]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Environment) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Environment) SetAttributes(v []Attribute)`

SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.

### GetAdditionalCosts

`func (o *Environment) GetAdditionalCosts() []AccountAdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *Environment) GetAdditionalCostsOk() ([]AccountAdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCosts

`func (o *Environment) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### SetAdditionalCosts

`func (o *Environment) SetAdditionalCosts(v []AccountAdditionalCost)`

SetAdditionalCosts gets a reference to the given []AccountAdditionalCost and assigns it to the AdditionalCosts field.

### GetAudiences

`func (o *Environment) GetAudiences() []Audience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *Environment) GetAudiencesOk() ([]Audience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudiences

`func (o *Environment) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### SetAudiences

`func (o *Environment) SetAudiences(v []Audience)`

SetAudiences gets a reference to the given []Audience and assigns it to the Audiences field.

### GetCollections

`func (o *Environment) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Environment) GetCollectionsOk() ([]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCollections

`func (o *Environment) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### SetCollections

`func (o *Environment) SetCollections(v []Collection)`

SetCollections gets a reference to the given []Collection and assigns it to the Collections field.

### GetApplicationCartItemFilters

`func (o *Environment) GetApplicationCartItemFilters() []ApplicationCif`

GetApplicationCartItemFilters returns the ApplicationCartItemFilters field if non-nil, zero value otherwise.

### GetApplicationCartItemFiltersOk

`func (o *Environment) GetApplicationCartItemFiltersOk() ([]ApplicationCif, bool)`

GetApplicationCartItemFiltersOk returns a tuple with the ApplicationCartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationCartItemFilters

`func (o *Environment) HasApplicationCartItemFilters() bool`

HasApplicationCartItemFilters returns a boolean if a field has been set.

### SetApplicationCartItemFilters

`func (o *Environment) SetApplicationCartItemFilters(v []ApplicationCif)`

SetApplicationCartItemFilters gets a reference to the given []ApplicationCif and assigns it to the ApplicationCartItemFilters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


