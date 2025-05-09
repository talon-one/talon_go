# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | 
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**Type** | Pointer to **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON data associated with the event. | 
**SessionId** | Pointer to **string** | The ID of the session that this event occurred in. | [optional] 
**Effects** | Pointer to [**[]interface{}**]([]interface{}.md) | An array of effects generated by the rules of the enabled campaigns of the Application.  You decide how to apply them in your system. See the list of [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects).  | 
**LedgerEntries** | Pointer to [**[]LedgerEntry**](LedgerEntry.md) | Ledger entries for the event. | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### GetId

`func (o *Event) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Event) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Event) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Event) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Event) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Event) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *Event) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Event) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Event) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Event) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetProfileId

`func (o *Event) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Event) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *Event) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *Event) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetStoreIntegrationId

`func (o *Event) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *Event) GetStoreIntegrationIdOk() (string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoreIntegrationId

`func (o *Event) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### SetStoreIntegrationId

`func (o *Event) SetStoreIntegrationId(v string)`

SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.

### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAttributes

`func (o *Event) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Event) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Event) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Event) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetSessionId

`func (o *Event) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Event) GetSessionIdOk() (string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *Event) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *Event) SetSessionId(v string)`

SetSessionId gets a reference to the given string and assigns it to the SessionId field.

### GetEffects

`func (o *Event) GetEffects() []interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Event) GetEffectsOk() ([]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *Event) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *Event) SetEffects(v []interface{})`

SetEffects gets a reference to the given []interface{} and assigns it to the Effects field.

### GetLedgerEntries

`func (o *Event) GetLedgerEntries() []LedgerEntry`

GetLedgerEntries returns the LedgerEntries field if non-nil, zero value otherwise.

### GetLedgerEntriesOk

`func (o *Event) GetLedgerEntriesOk() ([]LedgerEntry, bool)`

GetLedgerEntriesOk returns a tuple with the LedgerEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLedgerEntries

`func (o *Event) HasLedgerEntries() bool`

HasLedgerEntries returns a boolean if a field has been set.

### SetLedgerEntries

`func (o *Event) SetLedgerEntries(v []LedgerEntry)`

SetLedgerEntries gets a reference to the given []LedgerEntry and assigns it to the LedgerEntries field.

### GetMeta

`func (o *Event) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Event) GetMetaOk() (Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeta

`func (o *Event) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMeta

`func (o *Event) SetMeta(v Meta)`

SetMeta gets a reference to the given Meta and assigns it to the Meta field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


