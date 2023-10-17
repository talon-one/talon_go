# AddLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this loyalty point addition. | 
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were added. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**Value** | Pointer to **float32** | The amount of points that were added. | 
**DesiredValue** | Pointer to **float32** | The original amount of loyalty points to be awarded. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The user for whom these points were added. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date after which points will be valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date after which points will expire. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of this addition in the loyalty ledger. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart items list on which the loyal points addition should be applied. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the sub position indicates to which item the loyalty points addition is applied.  | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**BundleIndex** | Pointer to **int32** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 

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

### GetDesiredValue

`func (o *AddLoyaltyPointsEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *AddLoyaltyPointsEffectProps) GetDesiredValueOk() (float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredValue

`func (o *AddLoyaltyPointsEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### SetDesiredValue

`func (o *AddLoyaltyPointsEffectProps) SetDesiredValue(v float32)`

SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.

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

### GetTransactionUUID

`func (o *AddLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *AddLoyaltyPointsEffectProps) GetTransactionUUIDOk() (string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactionUUID

`func (o *AddLoyaltyPointsEffectProps) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.

### SetTransactionUUID

`func (o *AddLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID gets a reference to the given string and assigns it to the TransactionUUID field.

### GetCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *AddLoyaltyPointsEffectProps) GetCartItemPositionOk() (float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### SetCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.

### GetCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPositionOk() (float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### SetCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.

### GetCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *AddLoyaltyPointsEffectProps) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.

### GetBundleIndex

`func (o *AddLoyaltyPointsEffectProps) GetBundleIndex() int32`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *AddLoyaltyPointsEffectProps) GetBundleIndexOk() (int32, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleIndex

`func (o *AddLoyaltyPointsEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### SetBundleIndex

`func (o *AddLoyaltyPointsEffectProps) SetBundleIndex(v int32)`

SetBundleIndex gets a reference to the given int32 and assigns it to the BundleIndex field.

### GetBundleName

`func (o *AddLoyaltyPointsEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *AddLoyaltyPointsEffectProps) GetBundleNameOk() (string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleName

`func (o *AddLoyaltyPointsEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### SetBundleName

`func (o *AddLoyaltyPointsEffectProps) SetBundleName(v string)`

SetBundleName gets a reference to the given string and assigns it to the BundleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


