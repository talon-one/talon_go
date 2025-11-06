# OutgoingIntegrationMoEngagePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The base URL of your MoEngage deployment, containing the MoEngage data center number (represented by &#x60;0X&#x60;). | 
**AppId** | Pointer to **string** | MoEngage APP ID. See [MoEngage Developer Guide](https://developers.moengage.com/hc/en-us/articles/4404674776724-Overview). | 
**DataApiId** | Pointer to **string** | MoEngage DATA API ID. See [MoEngage Developer Guide](https://developers.moengage.com/hc/en-us/articles/4404674776724-Overview). | 
**DataApiKey** | Pointer to **string** | MoEngage DATA API Key. See [MoEngage Developer Guide](https://developers.moengage.com/hc/en-us/articles/4404674776724-Overview). | 

## Methods

### NewOutgoingIntegrationMoEngagePolicy

`func NewOutgoingIntegrationMoEngagePolicy(baseUrl string, appId string, dataApiId string, dataApiKey string, ) *OutgoingIntegrationMoEngagePolicy`

NewOutgoingIntegrationMoEngagePolicy instantiates a new OutgoingIntegrationMoEngagePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationMoEngagePolicyWithDefaults

`func NewOutgoingIntegrationMoEngagePolicyWithDefaults() *OutgoingIntegrationMoEngagePolicy`

NewOutgoingIntegrationMoEngagePolicyWithDefaults instantiates a new OutgoingIntegrationMoEngagePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *OutgoingIntegrationMoEngagePolicy) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *OutgoingIntegrationMoEngagePolicy) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *OutgoingIntegrationMoEngagePolicy) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetAppId

`func (o *OutgoingIntegrationMoEngagePolicy) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *OutgoingIntegrationMoEngagePolicy) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *OutgoingIntegrationMoEngagePolicy) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetDataApiId

`func (o *OutgoingIntegrationMoEngagePolicy) GetDataApiId() string`

GetDataApiId returns the DataApiId field if non-nil, zero value otherwise.

### GetDataApiIdOk

`func (o *OutgoingIntegrationMoEngagePolicy) GetDataApiIdOk() (*string, bool)`

GetDataApiIdOk returns a tuple with the DataApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataApiId

`func (o *OutgoingIntegrationMoEngagePolicy) SetDataApiId(v string)`

SetDataApiId sets DataApiId field to given value.


### GetDataApiKey

`func (o *OutgoingIntegrationMoEngagePolicy) GetDataApiKey() string`

GetDataApiKey returns the DataApiKey field if non-nil, zero value otherwise.

### GetDataApiKeyOk

`func (o *OutgoingIntegrationMoEngagePolicy) GetDataApiKeyOk() (*string, bool)`

GetDataApiKeyOk returns a tuple with the DataApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataApiKey

`func (o *OutgoingIntegrationMoEngagePolicy) SetDataApiKey(v string)`

SetDataApiKey sets DataApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


