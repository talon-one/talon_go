# WebhookLogEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID reference of the webhook request | [default to null]
**IntegrationRequestUuid** | **string** | UUID reference of the integration request linked to this webhook request | [default to null]
**WebhookId** | **int32** | ID of the webhook that triggered the request | [default to null]
**ApplicationId** | **int32** | ID of the application that triggered the webhook | [optional] [default to null]
**Url** | **string** | Target url of request | [default to null]
**Request** | **string** | Request message | [default to null]
**Response** | **string** | Response message | [optional] [default to null]
**Status** | **int32** | HTTP status code of response | [optional] [default to null]
**RequestTime** | [**time.Time**](time.Time.md) | Timestamp of request | [default to null]
**ResponseTime** | [**time.Time**](time.Time.md) | Timestamp of response | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


