# CustomerProfileAudienceRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defines the action to perform: - &#x60;add&#x60;: Adds the customer profile to the audience.    **Note**: If the customer profile does not exist, it will be created. The profile will not be visible in any Application   until a session or profile update is received for that profile. - &#x60;delete&#x60;: Removes the customer profile from the audience.  | 
**ProfileIntegrationId** | Pointer to **string** | The ID of this customer profile in the third-party integration. | 
**AudienceId** | Pointer to **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

## Methods

### GetAction

`func (o *CustomerProfileAudienceRequestItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CustomerProfileAudienceRequestItem) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *CustomerProfileAudienceRequestItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *CustomerProfileAudienceRequestItem) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetProfileIntegrationId

`func (o *CustomerProfileAudienceRequestItem) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *CustomerProfileAudienceRequestItem) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *CustomerProfileAudienceRequestItem) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *CustomerProfileAudienceRequestItem) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetAudienceId

`func (o *CustomerProfileAudienceRequestItem) GetAudienceId() int32`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *CustomerProfileAudienceRequestItem) GetAudienceIdOk() (int32, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceId

`func (o *CustomerProfileAudienceRequestItem) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### SetAudienceId

`func (o *CustomerProfileAudienceRequestItem) SetAudienceId(v int32)`

SetAudienceId gets a reference to the given int32 and assigns it to the AudienceId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


