# IntegrationHubConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationHubUrl** | Pointer to **string** | The url used to integrate the IntegrationHub Marketplace. | 
**AccessToken** | Pointer to **string** | Access token used to authenticate a user in Talon.One. | 

## Methods

### NewIntegrationHubConfig

`func NewIntegrationHubConfig(integrationHubUrl string, accessToken string, ) *IntegrationHubConfig`

NewIntegrationHubConfig instantiates a new IntegrationHubConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubConfigWithDefaults

`func NewIntegrationHubConfigWithDefaults() *IntegrationHubConfig`

NewIntegrationHubConfigWithDefaults instantiates a new IntegrationHubConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationHubUrl

`func (o *IntegrationHubConfig) GetIntegrationHubUrl() string`

GetIntegrationHubUrl returns the IntegrationHubUrl field if non-nil, zero value otherwise.

### GetIntegrationHubUrlOk

`func (o *IntegrationHubConfig) GetIntegrationHubUrlOk() (*string, bool)`

GetIntegrationHubUrlOk returns a tuple with the IntegrationHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationHubUrl

`func (o *IntegrationHubConfig) SetIntegrationHubUrl(v string)`

SetIntegrationHubUrl sets IntegrationHubUrl field to given value.


### GetAccessToken

`func (o *IntegrationHubConfig) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IntegrationHubConfig) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IntegrationHubConfig) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


