# JWT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Access token used to authenticate a user in Talon.One. | 
**ExpiresIn** | Pointer to **int64** | Time until the token expires (in seconds). | 

## Methods

### NewJWT

`func NewJWT(accessToken string, expiresIn int64, ) *JWT`

NewJWT instantiates a new JWT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTWithDefaults

`func NewJWTWithDefaults() *JWT`

NewJWTWithDefaults instantiates a new JWT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *JWT) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *JWT) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *JWT) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *JWT) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *JWT) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *JWT) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


