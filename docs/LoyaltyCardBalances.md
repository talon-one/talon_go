# LoyaltyCardBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**LoyaltyBalance**](LoyaltyBalance.md) |  | [optional] 
**SubledgerBalances** | Pointer to [**map[string]LoyaltyBalance**](LoyaltyBalance.md) | Map of the loyalty balances of the subledgers of a ledger. | [optional] 
**Profiles** | Pointer to [**[]LoyaltyCardProfileRegistration**](LoyaltyCardProfileRegistration.md) | Customer profiles linked to the loyalty card. | [optional] 

## Methods

### NewLoyaltyCardBalances

`func NewLoyaltyCardBalances() *LoyaltyCardBalances`

NewLoyaltyCardBalances instantiates a new LoyaltyCardBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardBalancesWithDefaults

`func NewLoyaltyCardBalancesWithDefaults() *LoyaltyCardBalances`

NewLoyaltyCardBalancesWithDefaults instantiates a new LoyaltyCardBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *LoyaltyCardBalances) GetBalance() LoyaltyBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LoyaltyCardBalances) GetBalanceOk() (*LoyaltyBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LoyaltyCardBalances) SetBalance(v LoyaltyBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *LoyaltyCardBalances) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetSubledgerBalances

`func (o *LoyaltyCardBalances) GetSubledgerBalances() map[string]LoyaltyBalance`

GetSubledgerBalances returns the SubledgerBalances field if non-nil, zero value otherwise.

### GetSubledgerBalancesOk

`func (o *LoyaltyCardBalances) GetSubledgerBalancesOk() (*map[string]LoyaltyBalance, bool)`

GetSubledgerBalancesOk returns a tuple with the SubledgerBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerBalances

`func (o *LoyaltyCardBalances) SetSubledgerBalances(v map[string]LoyaltyBalance)`

SetSubledgerBalances sets SubledgerBalances field to given value.

### HasSubledgerBalances

`func (o *LoyaltyCardBalances) HasSubledgerBalances() bool`

HasSubledgerBalances returns a boolean if a field has been set.

### GetProfiles

`func (o *LoyaltyCardBalances) GetProfiles() []LoyaltyCardProfileRegistration`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *LoyaltyCardBalances) GetProfilesOk() (*[]LoyaltyCardProfileRegistration, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *LoyaltyCardBalances) SetProfiles(v []LoyaltyCardProfileRegistration)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *LoyaltyCardBalances) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


