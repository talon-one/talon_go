# MCPOAuthTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | Pointer to **string** | OAuth2 grant type. | 
**Code** | Pointer to **string** | Authorization code. Required for &#x60;authorization_code&#x60; grant. | [optional] 
**ClientId** | Pointer to **string** | Client ID. Required for &#x60;authorization_code&#x60; grant. | [optional] 
**RedirectUri** | Pointer to **string** | Redirect URI. Required for &#x60;authorization_code&#x60; grant. | [optional] 
**CodeVerifier** | Pointer to **string** | PKCE code verifier. Required for &#x60;authorization_code&#x60; grant. | [optional] 
**RefreshToken** | Pointer to **string** | Refresh token. Required for &#x60;refresh_token&#x60; grant. | [optional] 

## Methods

### NewMCPOAuthTokenRequest

`func NewMCPOAuthTokenRequest(grantType string, ) *MCPOAuthTokenRequest`

NewMCPOAuthTokenRequest instantiates a new MCPOAuthTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthTokenRequestWithDefaults

`func NewMCPOAuthTokenRequestWithDefaults() *MCPOAuthTokenRequest`

NewMCPOAuthTokenRequestWithDefaults instantiates a new MCPOAuthTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *MCPOAuthTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *MCPOAuthTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *MCPOAuthTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetCode

`func (o *MCPOAuthTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MCPOAuthTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MCPOAuthTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MCPOAuthTokenRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetClientId

`func (o *MCPOAuthTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MCPOAuthTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MCPOAuthTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MCPOAuthTokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRedirectUri

`func (o *MCPOAuthTokenRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *MCPOAuthTokenRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *MCPOAuthTokenRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *MCPOAuthTokenRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetCodeVerifier

`func (o *MCPOAuthTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *MCPOAuthTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *MCPOAuthTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.

### HasCodeVerifier

`func (o *MCPOAuthTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.

### GetRefreshToken

`func (o *MCPOAuthTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *MCPOAuthTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *MCPOAuthTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *MCPOAuthTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


