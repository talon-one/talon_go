# OutgoingIntegrationCleverTapPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The base URL that is based on the region key of your CleverTap account. | 
**AccountId** | Pointer to **string** | The CleverTap Project ID. | 
**Passcode** | Pointer to **string** | The CleverTap Project passcode. | 

## Methods

### NewOutgoingIntegrationCleverTapPolicy

`func NewOutgoingIntegrationCleverTapPolicy(baseUrl string, accountId string, passcode string, ) *OutgoingIntegrationCleverTapPolicy`

NewOutgoingIntegrationCleverTapPolicy instantiates a new OutgoingIntegrationCleverTapPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationCleverTapPolicyWithDefaults

`func NewOutgoingIntegrationCleverTapPolicyWithDefaults() *OutgoingIntegrationCleverTapPolicy`

NewOutgoingIntegrationCleverTapPolicyWithDefaults instantiates a new OutgoingIntegrationCleverTapPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *OutgoingIntegrationCleverTapPolicy) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *OutgoingIntegrationCleverTapPolicy) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *OutgoingIntegrationCleverTapPolicy) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetAccountId

`func (o *OutgoingIntegrationCleverTapPolicy) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OutgoingIntegrationCleverTapPolicy) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OutgoingIntegrationCleverTapPolicy) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPasscode

`func (o *OutgoingIntegrationCleverTapPolicy) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *OutgoingIntegrationCleverTapPolicy) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *OutgoingIntegrationCleverTapPolicy) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


