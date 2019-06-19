# Role

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the role corresponding to the DB row | [default to null]
**AccountID** | **int32** | The ID of the Talon.One account that owns this role. | [default to null]
**Name** | **string** | Name of the role | [optional] [default to null]
**Description** | **string** | Description of the role | [optional] [default to null]
**Members** | **[]int32** | A list of userid in this role | [optional] [default to null]
**Acl** | **string** | Role Policy this should be a stringified blob of json | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


