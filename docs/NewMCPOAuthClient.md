# NewMCPOAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | Human-readable name for the OAuth2 client. | 
**RedirectUris** | Pointer to **[]string** | List of allowed redirect URIs for the authorization code flow. At least one URI is required. | 

## Methods

### NewNewMCPOAuthClient

`func NewNewMCPOAuthClient(clientName string, redirectUris []string, ) *NewMCPOAuthClient`

NewNewMCPOAuthClient instantiates a new NewMCPOAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMCPOAuthClientWithDefaults

`func NewNewMCPOAuthClientWithDefaults() *NewMCPOAuthClient`

NewNewMCPOAuthClientWithDefaults instantiates a new NewMCPOAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *NewMCPOAuthClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *NewMCPOAuthClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *NewMCPOAuthClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *NewMCPOAuthClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *NewMCPOAuthClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *NewMCPOAuthClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


