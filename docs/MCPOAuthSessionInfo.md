# MCPOAuthSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | The identifier of the authorization session. | 
**ExpiresAt** | Pointer to [**time.Time**](time.Time.md) | The date and time at which the session expires. Date and time. Follows RFC3339 format. | [optional] 
**Client** | Pointer to [**MCPOAuthClient**](MCPOAuthClient.md) |  | 

## Methods

### NewMCPOAuthSessionInfo

`func NewMCPOAuthSessionInfo(sessionId string, client MCPOAuthClient, ) *MCPOAuthSessionInfo`

NewMCPOAuthSessionInfo instantiates a new MCPOAuthSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthSessionInfoWithDefaults

`func NewMCPOAuthSessionInfoWithDefaults() *MCPOAuthSessionInfo`

NewMCPOAuthSessionInfoWithDefaults instantiates a new MCPOAuthSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *MCPOAuthSessionInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MCPOAuthSessionInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MCPOAuthSessionInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetExpiresAt

`func (o *MCPOAuthSessionInfo) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MCPOAuthSessionInfo) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MCPOAuthSessionInfo) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *MCPOAuthSessionInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetClient

`func (o *MCPOAuthSessionInfo) GetClient() MCPOAuthClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *MCPOAuthSessionInfo) GetClientOk() (*MCPOAuthClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *MCPOAuthSessionInfo) SetClient(v MCPOAuthClient)`

SetClient sets Client field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


