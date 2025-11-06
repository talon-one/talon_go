# ExpiringCardPointsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to **string** | The expiration date of loyalty points. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**AmountOfExpiringPoints** | Pointer to **float32** | The amount of loyalty points that will be expired soon. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card. | 
**UsersPerCardLimit** | Pointer to **int64** | The maximum number of customer profiles with which a card can be shared. This can be set to &#x60;0&#x60; for no limit.  | 
**Profiles** | Pointer to **[]string** | The integration IDs of the customer profiles linked to the card. | 

## Methods

### NewExpiringCardPointsData

`func NewExpiringCardPointsData(expiryDate string, loyaltyProgramID int64, amountOfExpiringPoints float32, subledgerID string, cardIdentifier string, usersPerCardLimit int64, profiles []string, ) *ExpiringCardPointsData`

NewExpiringCardPointsData instantiates a new ExpiringCardPointsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringCardPointsDataWithDefaults

`func NewExpiringCardPointsDataWithDefaults() *ExpiringCardPointsData`

NewExpiringCardPointsDataWithDefaults instantiates a new ExpiringCardPointsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *ExpiringCardPointsData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ExpiringCardPointsData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ExpiringCardPointsData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetLoyaltyProgramID

`func (o *ExpiringCardPointsData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *ExpiringCardPointsData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *ExpiringCardPointsData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetAmountOfExpiringPoints

`func (o *ExpiringCardPointsData) GetAmountOfExpiringPoints() float32`

GetAmountOfExpiringPoints returns the AmountOfExpiringPoints field if non-nil, zero value otherwise.

### GetAmountOfExpiringPointsOk

`func (o *ExpiringCardPointsData) GetAmountOfExpiringPointsOk() (*float32, bool)`

GetAmountOfExpiringPointsOk returns a tuple with the AmountOfExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfExpiringPoints

`func (o *ExpiringCardPointsData) SetAmountOfExpiringPoints(v float32)`

SetAmountOfExpiringPoints sets AmountOfExpiringPoints field to given value.


### GetSubledgerID

`func (o *ExpiringCardPointsData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *ExpiringCardPointsData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *ExpiringCardPointsData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetCardIdentifier

`func (o *ExpiringCardPointsData) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *ExpiringCardPointsData) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *ExpiringCardPointsData) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetUsersPerCardLimit

`func (o *ExpiringCardPointsData) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *ExpiringCardPointsData) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *ExpiringCardPointsData) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetProfiles

`func (o *ExpiringCardPointsData) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ExpiringCardPointsData) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ExpiringCardPointsData) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


