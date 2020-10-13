# AddLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name/description of this loyalty point addition | 
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were added | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added | 
**Value** | Pointer to **float32** | The amount of points that were added | 
**RecipientIntegrationId** | Pointer to **string** | The user for whom these points were added | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date after which points will be valid | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date after which points will expire | [optional] 

## Methods

### GetName

`func (o *AddLoyaltyPointsEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLoyaltyPointsEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AddLoyaltyPointsEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AddLoyaltyPointsEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetProgramId

`func (o *AddLoyaltyPointsEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *AddLoyaltyPointsEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *AddLoyaltyPointsEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *AddLoyaltyPointsEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *AddLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *AddLoyaltyPointsEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *AddLoyaltyPointsEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *AddLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetValue

`func (o *AddLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddLoyaltyPointsEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *AddLoyaltyPointsEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *AddLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetRecipientIntegrationId

`func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *AddLoyaltyPointsEffectProps) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *AddLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetStartDate

`func (o *AddLoyaltyPointsEffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddLoyaltyPointsEffectProps) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *AddLoyaltyPointsEffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *AddLoyaltyPointsEffectProps) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *AddLoyaltyPointsEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AddLoyaltyPointsEffectProps) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *AddLoyaltyPointsEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *AddLoyaltyPointsEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


