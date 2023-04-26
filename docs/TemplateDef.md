# TemplateDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Title** | Pointer to **string** | Campaigner-friendly name for the template that will be shown in the rule editor. | 
**Description** | Pointer to **string** | A short description of the template that will be shown in the rule editor. | 
**Help** | Pointer to **string** | Extended help text for the template. | 
**Category** | Pointer to **string** | Used for grouping templates in the rule editor sidebar. | 
**Expr** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | A Talang expression that contains variable bindings referring to args. | 
**Args** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | An array of argument definitions. | 
**Expose** | Pointer to **bool** | A flag to control exposure in Rule Builder. | [optional] [default to false]
**Name** | Pointer to **string** | The template name used in Talang. | 

## Methods

### GetId

`func (o *TemplateDef) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDef) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TemplateDef) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TemplateDef) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *TemplateDef) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TemplateDef) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *TemplateDef) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *TemplateDef) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *TemplateDef) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *TemplateDef) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *TemplateDef) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *TemplateDef) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetTitle

`func (o *TemplateDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateDef) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *TemplateDef) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *TemplateDef) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *TemplateDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateDef) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TemplateDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TemplateDef) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetHelp

`func (o *TemplateDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *TemplateDef) GetHelpOk() (string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHelp

`func (o *TemplateDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### SetHelp

`func (o *TemplateDef) SetHelp(v string)`

SetHelp gets a reference to the given string and assigns it to the Help field.

### GetCategory

`func (o *TemplateDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TemplateDef) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *TemplateDef) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *TemplateDef) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### GetExpr

`func (o *TemplateDef) GetExpr() []map[string]interface{}`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *TemplateDef) GetExprOk() ([]map[string]interface{}, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpr

`func (o *TemplateDef) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### SetExpr

`func (o *TemplateDef) SetExpr(v []map[string]interface{})`

SetExpr gets a reference to the given []map[string]interface{} and assigns it to the Expr field.

### GetArgs

`func (o *TemplateDef) GetArgs() []TemplateArgDef`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *TemplateDef) GetArgsOk() ([]TemplateArgDef, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArgs

`func (o *TemplateDef) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgs

`func (o *TemplateDef) SetArgs(v []TemplateArgDef)`

SetArgs gets a reference to the given []TemplateArgDef and assigns it to the Args field.

### GetExpose

`func (o *TemplateDef) GetExpose() bool`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *TemplateDef) GetExposeOk() (bool, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpose

`func (o *TemplateDef) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### SetExpose

`func (o *TemplateDef) SetExpose(v bool)`

SetExpose gets a reference to the given bool and assigns it to the Expose field.

### GetName

`func (o *TemplateDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDef) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TemplateDef) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TemplateDef) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


