# NewTemplateDef

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Campaigner-friendly name for the template that will be shown in the rule editor. | [default to null]
**Description** | **string** | A short description of the template that will be shown in the rule editor. | [optional] [default to null]
**Help** | **string** | Extended help text for the template. | [optional] [default to null]
**Category** | **string** | Used for grouping templates in the rule editor sidebar. | [default to null]
**Expr** | [**[]interface{}**](interface{}.md) | A Talang expression that contains variable bindings referring to args. | [default to null]
**Args** | [**[]TemplateArgDef**](TemplateArgDef.md) | An array of argument definitions. | [default to null]
**Expose** | **bool** | A flag to control exposure in Rule Builder. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


