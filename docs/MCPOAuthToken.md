# MCPOAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Bearer access token. | 
**TokenType** | Pointer to **string** | Token type. Always \&quot;Bearer\&quot;. | 
**ExpiresIn** | Pointer to **int64** | Seconds until the access token expires. | 
**RefreshToken** | Pointer to **string** | Refresh token for obtaining a new access token. | 
**RefreshTokenExpiresIn** | Pointer to **int64** | Seconds until the refresh token expires. | 

## Methods

### NewMCPOAuthToken

`func NewMCPOAuthToken(accessToken string, tokenType string, expiresIn int64, refreshToken string, refreshTokenExpiresIn int64, ) *MCPOAuthToken`

NewMCPOAuthToken instantiates a new MCPOAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthTokenWithDefaults

`func NewMCPOAuthTokenWithDefaults() *MCPOAuthToken`

NewMCPOAuthTokenWithDefaults instantiates a new MCPOAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MCPOAuthToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MCPOAuthToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MCPOAuthToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *MCPOAuthToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *MCPOAuthToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *MCPOAuthToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *MCPOAuthToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *MCPOAuthToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *MCPOAuthToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.


### GetRefreshToken

`func (o *MCPOAuthToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *MCPOAuthToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *MCPOAuthToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetRefreshTokenExpiresIn

`func (o *MCPOAuthToken) GetRefreshTokenExpiresIn() int64`

GetRefreshTokenExpiresIn returns the RefreshTokenExpiresIn field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresInOk

`func (o *MCPOAuthToken) GetRefreshTokenExpiresInOk() (*int64, bool)`

GetRefreshTokenExpiresInOk returns a tuple with the RefreshTokenExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresIn

`func (o *MCPOAuthToken) SetRefreshTokenExpiresIn(v int64)`

SetRefreshTokenExpiresIn sets RefreshTokenExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


