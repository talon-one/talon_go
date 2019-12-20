# NewSamlConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X509certificate** | **string** | X.509 Certificate. | [default to null]
**AccountId** | **int32** | The ID of the account that owns this entity. | [default to null]
**Name** | **string** | ID of the SAML service. | [default to null]
**Enabled** | **bool** | Determines if this SAML connection active. | [default to null]
**Issuer** | **string** | Identity Provider Entity ID. | [default to null]
**SignOnURL** | **string** | Single Sign-On URL. | [default to null]
**SignOutURL** | **string** | Single Sign-Out URL. | [optional] [default to null]
**MetadataURL** | **string** | Metadata URL. | [optional] [default to null]
**AudienceURI** | **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


