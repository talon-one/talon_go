# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Slots** | Pointer to [**[]SlotDef**](SlotDef.md) | The slots defined for this application. | 
**Functions** | Pointer to [**[]FunctionDef**](FunctionDef.md) | The functions defined for this application. | 
**Templates** | Pointer to [**[]TemplateDef**](TemplateDef.md) | The templates defined for this application. | 
**Variables** | Pointer to **string** |  | 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


