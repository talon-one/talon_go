# AddLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this loyalty point addition. | 
**ProgramId** | Pointer to **int64** | The ID of the loyalty program where these points were added. | 
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
**BundleIndex** | Pointer to **int64** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 

## Methods

### NewAddLoyaltyPointsEffectProps

`func NewAddLoyaltyPointsEffectProps(name string, programId int64, subLedgerId string, value float32, recipientIntegrationId string, transactionUUID string, ) *AddLoyaltyPointsEffectProps`

NewAddLoyaltyPointsEffectProps instantiates a new AddLoyaltyPointsEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLoyaltyPointsEffectPropsWithDefaults

`func NewAddLoyaltyPointsEffectPropsWithDefaults() *AddLoyaltyPointsEffectProps`

NewAddLoyaltyPointsEffectPropsWithDefaults instantiates a new AddLoyaltyPointsEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddLoyaltyPointsEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLoyaltyPointsEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLoyaltyPointsEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetProgramId

`func (o *AddLoyaltyPointsEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *AddLoyaltyPointsEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *AddLoyaltyPointsEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *AddLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *AddLoyaltyPointsEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *AddLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetValue

`func (o *AddLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddLoyaltyPointsEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetDesiredValue

`func (o *AddLoyaltyPointsEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *AddLoyaltyPointsEffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *AddLoyaltyPointsEffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *AddLoyaltyPointsEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *AddLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.


### GetStartDate

`func (o *AddLoyaltyPointsEffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddLoyaltyPointsEffectProps) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddLoyaltyPointsEffectProps) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddLoyaltyPointsEffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *AddLoyaltyPointsEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AddLoyaltyPointsEffectProps) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *AddLoyaltyPointsEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *AddLoyaltyPointsEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *AddLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *AddLoyaltyPointsEffectProps) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *AddLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *AddLoyaltyPointsEffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.

### HasCartItemPosition

`func (o *AddLoyaltyPointsEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### GetCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *AddLoyaltyPointsEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *AddLoyaltyPointsEffectProps) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *AddLoyaltyPointsEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### GetBundleIndex

`func (o *AddLoyaltyPointsEffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *AddLoyaltyPointsEffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *AddLoyaltyPointsEffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *AddLoyaltyPointsEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *AddLoyaltyPointsEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *AddLoyaltyPointsEffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *AddLoyaltyPointsEffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *AddLoyaltyPointsEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


