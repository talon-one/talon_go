# SamlConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [default to null]
**Created** | [**time.Time**](time.Time.md) | Unix timestamp indicating when the session was first created. | [default to null]
**AccountId** | **int32** | The ID of the account that owns this SAML Service. | [default to null]
**AssertionConsumerServiceURL** | **string** | The location where the SAML assertion is sent with a HTTP POST. | [default to null]
**Name** | **string** | ID of the SAML service. | [default to null]
**Enabled** | **bool** | Determines if this SAML connection active. | [default to null]
**Issuer** | **string** | Identity Provider Entity ID. | [default to null]
**SignOnURL** | **string** | Single Sign-On URL. | [default to null]
**SignOutURL** | **string** | Single Sign-Out URL. | [optional] [default to null]
**MetadataURL** | **string** | Metadata URL. | [optional] [default to null]
**X509certificate** | **string** | X.509 Certificate. | [default to null]
**Audience** | **string** | The application-defined unique identifier that is the intended audience of the SAML assertion.  This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


