# ExpiringPointsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to **string** | The expiration date of loyalty points. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**CustomerProfileID** | Pointer to **string** | The integration ID of the customer profile that has expiring points. | 
**AmountOfExpiringPoints** | Pointer to **float32** | The amount of loyalty points that will be expired soon. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 

## Methods

### NewExpiringPointsData

`func NewExpiringPointsData(expiryDate string, loyaltyProgramID int64, customerProfileID string, amountOfExpiringPoints float32, subledgerID string, ) *ExpiringPointsData`

NewExpiringPointsData instantiates a new ExpiringPointsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringPointsDataWithDefaults

`func NewExpiringPointsDataWithDefaults() *ExpiringPointsData`

NewExpiringPointsDataWithDefaults instantiates a new ExpiringPointsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *ExpiringPointsData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ExpiringPointsData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ExpiringPointsData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetLoyaltyProgramID

`func (o *ExpiringPointsData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *ExpiringPointsData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *ExpiringPointsData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetCustomerProfileID

`func (o *ExpiringPointsData) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *ExpiringPointsData) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *ExpiringPointsData) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.


### GetAmountOfExpiringPoints

`func (o *ExpiringPointsData) GetAmountOfExpiringPoints() float32`

GetAmountOfExpiringPoints returns the AmountOfExpiringPoints field if non-nil, zero value otherwise.

### GetAmountOfExpiringPointsOk

`func (o *ExpiringPointsData) GetAmountOfExpiringPointsOk() (*float32, bool)`

GetAmountOfExpiringPointsOk returns a tuple with the AmountOfExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfExpiringPoints

`func (o *ExpiringPointsData) SetAmountOfExpiringPoints(v float32)`

SetAmountOfExpiringPoints sets AmountOfExpiringPoints field to given value.


### GetSubledgerID

`func (o *ExpiringPointsData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *ExpiringPointsData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *ExpiringPointsData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


