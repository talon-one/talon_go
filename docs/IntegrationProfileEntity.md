# IntegrationProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 

## Methods

### GetProfileId

`func (o *IntegrationProfileEntity) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationProfileEntity) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *IntegrationProfileEntity) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *IntegrationProfileEntity) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


