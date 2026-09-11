# MCPOAuthProtectedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** | The URL of the protected resource (the MCP entrypoint). | 
**AuthorizationServers** | Pointer to **[]string** | List of authorization server base URLs that can issue tokens for this resource. | 

## Methods

### NewMCPOAuthProtectedResource

`func NewMCPOAuthProtectedResource(resource string, authorizationServers []string, ) *MCPOAuthProtectedResource`

NewMCPOAuthProtectedResource instantiates a new MCPOAuthProtectedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthProtectedResourceWithDefaults

`func NewMCPOAuthProtectedResourceWithDefaults() *MCPOAuthProtectedResource`

NewMCPOAuthProtectedResourceWithDefaults instantiates a new MCPOAuthProtectedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *MCPOAuthProtectedResource) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MCPOAuthProtectedResource) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MCPOAuthProtectedResource) SetResource(v string)`

SetResource sets Resource field to given value.


### GetAuthorizationServers

`func (o *MCPOAuthProtectedResource) GetAuthorizationServers() []string`

GetAuthorizationServers returns the AuthorizationServers field if non-nil, zero value otherwise.

### GetAuthorizationServersOk

`func (o *MCPOAuthProtectedResource) GetAuthorizationServersOk() (*[]string, bool)`

GetAuthorizationServersOk returns a tuple with the AuthorizationServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServers

`func (o *MCPOAuthProtectedResource) SetAuthorizationServers(v []string)`

SetAuthorizationServers sets AuthorizationServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


