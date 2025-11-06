# OutgoingIntegrationIterablePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The base URL that is based on the region key of your Iterable account. | 
**ApiKey** | Pointer to **string** | The API key generated from your Iterable account. See [Iterable API Key Guide](https://support.iterable.com/hc/en-us/articles/360043464871-API-Keys-) | 

## Methods

### NewOutgoingIntegrationIterablePolicy

`func NewOutgoingIntegrationIterablePolicy(baseUrl string, apiKey string, ) *OutgoingIntegrationIterablePolicy`

NewOutgoingIntegrationIterablePolicy instantiates a new OutgoingIntegrationIterablePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationIterablePolicyWithDefaults

`func NewOutgoingIntegrationIterablePolicyWithDefaults() *OutgoingIntegrationIterablePolicy`

NewOutgoingIntegrationIterablePolicyWithDefaults instantiates a new OutgoingIntegrationIterablePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *OutgoingIntegrationIterablePolicy) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *OutgoingIntegrationIterablePolicy) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *OutgoingIntegrationIterablePolicy) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *OutgoingIntegrationIterablePolicy) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OutgoingIntegrationIterablePolicy) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OutgoingIntegrationIterablePolicy) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


