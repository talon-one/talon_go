# MCPOAuthServerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | The authorization server&#39;s issuer identifier (its base URL). | 
**AuthorizationEndpoint** | Pointer to **string** | URL of the authorization endpoint. | 
**TokenEndpoint** | Pointer to **string** | URL of the token endpoint. | 
**RegistrationEndpoint** | Pointer to **string** | URL of the client registration endpoint. | 
**ResponseTypesSupported** | Pointer to **[]string** | List of supported OAuth2 response types. | 
**GrantTypesSupported** | Pointer to **[]string** | List of supported OAuth2 grant types. | 
**CodeChallengeMethodsSupported** | Pointer to **[]string** | List of supported PKCE code challenge methods. | 

## Methods

### NewMCPOAuthServerMetadata

`func NewMCPOAuthServerMetadata(issuer string, authorizationEndpoint string, tokenEndpoint string, registrationEndpoint string, responseTypesSupported []string, grantTypesSupported []string, codeChallengeMethodsSupported []string, ) *MCPOAuthServerMetadata`

NewMCPOAuthServerMetadata instantiates a new MCPOAuthServerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthServerMetadataWithDefaults

`func NewMCPOAuthServerMetadataWithDefaults() *MCPOAuthServerMetadata`

NewMCPOAuthServerMetadataWithDefaults instantiates a new MCPOAuthServerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *MCPOAuthServerMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MCPOAuthServerMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MCPOAuthServerMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetAuthorizationEndpoint

`func (o *MCPOAuthServerMetadata) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *MCPOAuthServerMetadata) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *MCPOAuthServerMetadata) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.


### GetTokenEndpoint

`func (o *MCPOAuthServerMetadata) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *MCPOAuthServerMetadata) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *MCPOAuthServerMetadata) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetRegistrationEndpoint

`func (o *MCPOAuthServerMetadata) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *MCPOAuthServerMetadata) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *MCPOAuthServerMetadata) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.


### GetResponseTypesSupported

`func (o *MCPOAuthServerMetadata) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *MCPOAuthServerMetadata) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *MCPOAuthServerMetadata) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.


### GetGrantTypesSupported

`func (o *MCPOAuthServerMetadata) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *MCPOAuthServerMetadata) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *MCPOAuthServerMetadata) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.


### GetCodeChallengeMethodsSupported

`func (o *MCPOAuthServerMetadata) GetCodeChallengeMethodsSupported() []string`

GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field if non-nil, zero value otherwise.

### GetCodeChallengeMethodsSupportedOk

`func (o *MCPOAuthServerMetadata) GetCodeChallengeMethodsSupportedOk() (*[]string, bool)`

GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethodsSupported

`func (o *MCPOAuthServerMetadata) SetCodeChallengeMethodsSupported(v []string)`

SetCodeChallengeMethodsSupported sets CodeChallengeMethodsSupported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


