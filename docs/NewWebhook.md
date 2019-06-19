# NewWebhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | **[]int32** | The IDs of the applications that are related to this entity. | [default to null]
**Title** | **string** | Friendly title for this webhook | [default to null]
**Verb** | **string** | API method for this webhook | [default to null]
**Url** | **string** | API url (supports templating using parameters) for this webhook | [default to null]
**Headers** | **[]string** | List of API HTTP headers for this webhook | [default to null]
**Payload** | **string** | API payload (supports templating using parameters) for this webhook | [optional] [default to null]
**Params** | [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions | [default to null]
**Enabled** | **bool** | Enables or disables webhook from showing in rule builder | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


