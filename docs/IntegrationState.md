# IntegrationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**CustomerSession**](CustomerSession.md) |  | 
**Profile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | 
**Event** | Pointer to [**Event**](Event.md) |  | 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Coupon** | Pointer to [**Coupon**](Coupon.md) |  | [optional] 

## Methods

### NewIntegrationState

`func NewIntegrationState(session CustomerSession, profile CustomerProfile, event Event, ) *IntegrationState`

NewIntegrationState instantiates a new IntegrationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationStateWithDefaults

`func NewIntegrationStateWithDefaults() *IntegrationState`

NewIntegrationStateWithDefaults instantiates a new IntegrationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *IntegrationState) GetSession() CustomerSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *IntegrationState) GetSessionOk() (*CustomerSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *IntegrationState) SetSession(v CustomerSession)`

SetSession sets Session field to given value.


### GetProfile

`func (o *IntegrationState) GetProfile() CustomerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IntegrationState) GetProfileOk() (*CustomerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IntegrationState) SetProfile(v CustomerProfile)`

SetProfile sets Profile field to given value.


### GetEvent

`func (o *IntegrationState) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IntegrationState) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IntegrationState) SetEvent(v Event)`

SetEvent sets Event field to given value.


### GetLoyalty

`func (o *IntegrationState) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *IntegrationState) GetLoyaltyOk() (*Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *IntegrationState) SetLoyalty(v Loyalty)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *IntegrationState) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetCoupon

`func (o *IntegrationState) GetCoupon() Coupon`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *IntegrationState) GetCouponOk() (*Coupon, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *IntegrationState) SetCoupon(v Coupon)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *IntegrationState) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


