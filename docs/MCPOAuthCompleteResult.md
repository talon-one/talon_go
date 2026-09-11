# MCPOAuthCompleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **string** | The full redirect URL the browser should be sent to, containing the authorization code and state as query parameters. | 

## Methods

### NewMCPOAuthCompleteResult

`func NewMCPOAuthCompleteResult(redirectUrl string, ) *MCPOAuthCompleteResult`

NewMCPOAuthCompleteResult instantiates a new MCPOAuthCompleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthCompleteResultWithDefaults

`func NewMCPOAuthCompleteResultWithDefaults() *MCPOAuthCompleteResult`

NewMCPOAuthCompleteResultWithDefaults instantiates a new MCPOAuthCompleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *MCPOAuthCompleteResult) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *MCPOAuthCompleteResult) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *MCPOAuthCompleteResult) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


