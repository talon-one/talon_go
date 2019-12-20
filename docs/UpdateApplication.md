# UpdateApplication

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this application. | [default to null]
**Description** | **string** | A longer description of the application. | [optional] [default to null]
**Timezone** | **string** | A string containing an IANA timezone descriptor. | [default to null]
**Currency** | **string** | A string describing a default currency for new customer sessions. | [default to null]
**CaseSensitivity** | **string** | A string indicating how should campaigns in this application deal with case sensitivity on coupon codes. | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this campaign | [optional] [default to null]
**Limits** | [**[]LimitConfig**](LimitConfig.md) | Default limits for campaigns created in this application | [optional] [default to null]
**AttributesSettings** | [***AttributesSettings**](AttributesSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

