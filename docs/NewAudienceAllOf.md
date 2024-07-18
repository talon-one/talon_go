# NewAudienceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | Pointer to **string** | The Talon.One-supported [3rd-party platform](https://docs.talon.one/docs/dev/technology-partners/overview) that this audience was created in.  For example, &#x60;mParticle&#x60;, &#x60;Segment&#x60;, &#x60;Selligent&#x60;, &#x60;Braze&#x60;, or &#x60;Iterable&#x60;.  **Note:** If you do not integrate with any of these platforms, do not use this property.  | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration.  **Note:** To create an audience that doesn&#39;t come from a 3rd party platform, do not use this property.  | [optional] 
**CreatedIn3rdParty** | Pointer to **bool** | Determines if this audience is a 3rd party audience or not. | [optional] 
**LastUpdate** | Pointer to [**time.Time**](time.Time.md) | The last time that the audience memberships changed. | [optional] 

## Methods

### GetIntegration

`func (o *NewAudienceAllOf) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *NewAudienceAllOf) GetIntegrationOk() (string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegration

`func (o *NewAudienceAllOf) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegration

`func (o *NewAudienceAllOf) SetIntegration(v string)`

SetIntegration gets a reference to the given string and assigns it to the Integration field.

### GetIntegrationId

`func (o *NewAudienceAllOf) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewAudienceAllOf) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *NewAudienceAllOf) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *NewAudienceAllOf) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetCreatedIn3rdParty

`func (o *NewAudienceAllOf) GetCreatedIn3rdParty() bool`

GetCreatedIn3rdParty returns the CreatedIn3rdParty field if non-nil, zero value otherwise.

### GetCreatedIn3rdPartyOk

`func (o *NewAudienceAllOf) GetCreatedIn3rdPartyOk() (bool, bool)`

GetCreatedIn3rdPartyOk returns a tuple with the CreatedIn3rdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedIn3rdParty

`func (o *NewAudienceAllOf) HasCreatedIn3rdParty() bool`

HasCreatedIn3rdParty returns a boolean if a field has been set.

### SetCreatedIn3rdParty

`func (o *NewAudienceAllOf) SetCreatedIn3rdParty(v bool)`

SetCreatedIn3rdParty gets a reference to the given bool and assigns it to the CreatedIn3rdParty field.

### GetLastUpdate

`func (o *NewAudienceAllOf) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *NewAudienceAllOf) GetLastUpdateOk() (time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdate

`func (o *NewAudienceAllOf) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdate

`func (o *NewAudienceAllOf) SetLastUpdate(v time.Time)`

SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


