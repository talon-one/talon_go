# NewTemplateDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Campaigner-friendly name for the template that will be shown in the rule editor. | 
**Description** | Pointer to **string** | A short description of the template that will be shown in the rule editor. | [optional] 
**Help** | Pointer to **string** | Extended help text for the template. | [optional] 
**Category** | Pointer to **string** | Used for grouping templates in the rule editor sidebar. | 
**Expr** | Pointer to **[]map[string]interface{}** | A Talang expression that contains variable bindings referring to args. | 
**Args** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | An array of argument definitions. | 
**Expose** | Pointer to **bool** | A flag to control exposure in Rule Builder. | [optional] [default to false]

## Methods

### NewNewTemplateDef

`func NewNewTemplateDef(title string, category string, expr []map[string]interface{}, args []TemplateArgDef, ) *NewTemplateDef`

NewNewTemplateDef instantiates a new NewTemplateDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTemplateDefWithDefaults

`func NewNewTemplateDefWithDefaults() *NewTemplateDef`

NewNewTemplateDefWithDefaults instantiates a new NewTemplateDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewTemplateDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewTemplateDef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewTemplateDef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewTemplateDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewTemplateDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewTemplateDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewTemplateDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHelp

`func (o *NewTemplateDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *NewTemplateDef) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *NewTemplateDef) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *NewTemplateDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetCategory

`func (o *NewTemplateDef) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NewTemplateDef) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NewTemplateDef) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetExpr

`func (o *NewTemplateDef) GetExpr() []interface{}`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *NewTemplateDef) GetExprOk() (*[]map[string]interface{}, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *NewTemplateDef) SetExpr(v []interface{})`

SetExpr sets Expr field to given value.


### GetArgs

`func (o *NewTemplateDef) GetArgs() []TemplateArgDef`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *NewTemplateDef) GetArgsOk() (*[]TemplateArgDef, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *NewTemplateDef) SetArgs(v []TemplateArgDef)`

SetArgs sets Args field to given value.


### GetExpose

`func (o *NewTemplateDef) GetExpose() bool`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *NewTemplateDef) GetExposeOk() (*bool, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *NewTemplateDef) SetExpose(v bool)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *NewTemplateDef) HasExpose() bool`

HasExpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


