# IntegrationCustomerProfileAudienceRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defines the action to perform: - &#x60;add&#x60;: Adds the customer profile to the audience.    **Note**: If the customer profile does not exist, it will be created. The profile will not be visible in any Application   until a session or profile update is received for that profile. - &#x60;delete&#x60;: Removes the customer profile from the audience.  | 
**ProfileIntegrationId** | Pointer to **string** | The ID of this customer profile in the third-party integration. | 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration. | 

## Methods

### GetAction

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *IntegrationCustomerProfileAudienceRequestItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetProfileIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


