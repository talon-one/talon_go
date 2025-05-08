# NewAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The human-friendly display name for this audience. | 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**Description** | Pointer to **string** | A description of the audience. | [optional] 
**Integration** | Pointer to **string** | The Talon.One-supported [3rd-party platform](https://docs.talon.one/docs/dev/technology-partners/overview) that this audience was created in.  For example, &#x60;mParticle&#x60;, &#x60;Segment&#x60;, &#x60;Shopify&#x60;, &#x60;Braze&#x60;, or &#x60;Iterable&#x60;.  **Note:** If you do not integrate with any of these platforms, do not use this property.  | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration.  **Note:** To create an audience that doesn&#39;t come from a 3rd party platform, do not use this property.  | [optional] 
**CreatedIn3rdParty** | Pointer to **bool** | Determines if this audience is a 3rd party audience or not. | [optional] 
**LastUpdate** | Pointer to [**time.Time**](time.Time.md) | The last time that the audience memberships changed. | [optional] 

## Methods

### GetName

`func (o *NewAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAudience) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewAudience) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewAudience) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSandbox

`func (o *NewAudience) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewAudience) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *NewAudience) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *NewAudience) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetDescription

`func (o *NewAudience) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewAudience) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewAudience) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewAudience) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetIntegration

`func (o *NewAudience) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *NewAudience) GetIntegrationOk() (string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegration

`func (o *NewAudience) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegration

`func (o *NewAudience) SetIntegration(v string)`

SetIntegration gets a reference to the given string and assigns it to the Integration field.

### GetIntegrationId

`func (o *NewAudience) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewAudience) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *NewAudience) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *NewAudience) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetCreatedIn3rdParty

`func (o *NewAudience) GetCreatedIn3rdParty() bool`

GetCreatedIn3rdParty returns the CreatedIn3rdParty field if non-nil, zero value otherwise.

### GetCreatedIn3rdPartyOk

`func (o *NewAudience) GetCreatedIn3rdPartyOk() (bool, bool)`

GetCreatedIn3rdPartyOk returns a tuple with the CreatedIn3rdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedIn3rdParty

`func (o *NewAudience) HasCreatedIn3rdParty() bool`

HasCreatedIn3rdParty returns a boolean if a field has been set.

### SetCreatedIn3rdParty

`func (o *NewAudience) SetCreatedIn3rdParty(v bool)`

SetCreatedIn3rdParty gets a reference to the given bool and assigns it to the CreatedIn3rdParty field.

### GetLastUpdate

`func (o *NewAudience) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *NewAudience) GetLastUpdateOk() (time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdate

`func (o *NewAudience) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdate

`func (o *NewAudience) SetLastUpdate(v time.Time)`

SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


