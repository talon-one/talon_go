# MultipleCustomerProfileIntegrationRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**IntegrationId** | Pointer to **string** | The identifier of this profile, set by your integration layer. It must be unique within the account.  To get the &#x60;integrationId&#x60; of the profile from a &#x60;sessionId&#x60;, use the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2).  | 

## Methods

### NewMultipleCustomerProfileIntegrationRequestItem

`func NewMultipleCustomerProfileIntegrationRequestItem(integrationId string, ) *MultipleCustomerProfileIntegrationRequestItem`

NewMultipleCustomerProfileIntegrationRequestItem instantiates a new MultipleCustomerProfileIntegrationRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleCustomerProfileIntegrationRequestItemWithDefaults

`func NewMultipleCustomerProfileIntegrationRequestItemWithDefaults() *MultipleCustomerProfileIntegrationRequestItem`

NewMultipleCustomerProfileIntegrationRequestItemWithDefaults instantiates a new MultipleCustomerProfileIntegrationRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MultipleCustomerProfileIntegrationRequestItem) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MultipleCustomerProfileIntegrationRequestItem) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MultipleCustomerProfileIntegrationRequestItem) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MultipleCustomerProfileIntegrationRequestItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIntegrationId

`func (o *MultipleCustomerProfileIntegrationRequestItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *MultipleCustomerProfileIntegrationRequestItem) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *MultipleCustomerProfileIntegrationRequestItem) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


