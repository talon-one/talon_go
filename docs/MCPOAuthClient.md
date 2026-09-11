# MCPOAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Unique identifier for the OAuth2 client. | 
**ClientName** | Pointer to **string** | Human-readable name for the OAuth2 client. | 
**RedirectUris** | Pointer to **[]string** | List of allowed redirect URIs for the authorization code flow. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the client was registered. | 

## Methods

### NewMCPOAuthClient

`func NewMCPOAuthClient(clientId string, clientName string, redirectUris []string, createdAt time.Time, ) *MCPOAuthClient`

NewMCPOAuthClient instantiates a new MCPOAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthClientWithDefaults

`func NewMCPOAuthClientWithDefaults() *MCPOAuthClient`

NewMCPOAuthClientWithDefaults instantiates a new MCPOAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MCPOAuthClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MCPOAuthClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MCPOAuthClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientName

`func (o *MCPOAuthClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MCPOAuthClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MCPOAuthClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *MCPOAuthClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *MCPOAuthClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *MCPOAuthClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetCreatedAt

`func (o *MCPOAuthClient) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MCPOAuthClient) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MCPOAuthClient) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


