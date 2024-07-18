# CustomerProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. | 
**ClosedSessions** | Pointer to **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | 
**TotalSales** | Pointer to **float32** | The total amount of money spent by the customer **before** discounts are applied.  The total sales amount excludes the following: - Cancelled or reopened sessions. - Returned items.  | 
**LoyaltyMemberships** | Pointer to [**[]LoyaltyMembership**](LoyaltyMembership.md) | **DEPRECATED** A list of loyalty programs joined by the customer.  | [optional] 
**AudienceMemberships** | Pointer to [**[]AudienceMembership**](AudienceMembership.md) | The audiences the customer belongs to. | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received from this customer. This field is updated on calls that trigger the Rule Engine and that are not [dry requests](https://docs.talon.one/docs/dev/integration-api/dry-requests/#overlay).  For example, [reserving a coupon](https://docs.talon.one/integration-api#operation/createCouponReservation) for a customer doesn&#39;t impact this field.  | 
**Sandbox** | Pointer to **bool** | An indicator of whether the customer is part of a sandbox or live Application. See the [docs](https://docs.talon.one/docs/product/applications/overview#application-environments).  | [optional] 

## Methods

### GetAccountId

`func (o *CustomerProfileAllOf) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomerProfileAllOf) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CustomerProfileAllOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CustomerProfileAllOf) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetClosedSessions

`func (o *CustomerProfileAllOf) GetClosedSessions() int32`

GetClosedSessions returns the ClosedSessions field if non-nil, zero value otherwise.

### GetClosedSessionsOk

`func (o *CustomerProfileAllOf) GetClosedSessionsOk() (int32, bool)`

GetClosedSessionsOk returns a tuple with the ClosedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClosedSessions

`func (o *CustomerProfileAllOf) HasClosedSessions() bool`

HasClosedSessions returns a boolean if a field has been set.

### SetClosedSessions

`func (o *CustomerProfileAllOf) SetClosedSessions(v int32)`

SetClosedSessions gets a reference to the given int32 and assigns it to the ClosedSessions field.

### GetTotalSales

`func (o *CustomerProfileAllOf) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *CustomerProfileAllOf) GetTotalSalesOk() (float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSales

`func (o *CustomerProfileAllOf) HasTotalSales() bool`

HasTotalSales returns a boolean if a field has been set.

### SetTotalSales

`func (o *CustomerProfileAllOf) SetTotalSales(v float32)`

SetTotalSales gets a reference to the given float32 and assigns it to the TotalSales field.

### GetLoyaltyMemberships

`func (o *CustomerProfileAllOf) GetLoyaltyMemberships() []LoyaltyMembership`

GetLoyaltyMemberships returns the LoyaltyMemberships field if non-nil, zero value otherwise.

### GetLoyaltyMembershipsOk

`func (o *CustomerProfileAllOf) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool)`

GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyMemberships

`func (o *CustomerProfileAllOf) HasLoyaltyMemberships() bool`

HasLoyaltyMemberships returns a boolean if a field has been set.

### SetLoyaltyMemberships

`func (o *CustomerProfileAllOf) SetLoyaltyMemberships(v []LoyaltyMembership)`

SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.

### GetAudienceMemberships

`func (o *CustomerProfileAllOf) GetAudienceMemberships() []AudienceMembership`

GetAudienceMemberships returns the AudienceMemberships field if non-nil, zero value otherwise.

### GetAudienceMembershipsOk

`func (o *CustomerProfileAllOf) GetAudienceMembershipsOk() ([]AudienceMembership, bool)`

GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceMemberships

`func (o *CustomerProfileAllOf) HasAudienceMemberships() bool`

HasAudienceMemberships returns a boolean if a field has been set.

### SetAudienceMemberships

`func (o *CustomerProfileAllOf) SetAudienceMemberships(v []AudienceMembership)`

SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.

### GetLastActivity

`func (o *CustomerProfileAllOf) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerProfileAllOf) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *CustomerProfileAllOf) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *CustomerProfileAllOf) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetSandbox

`func (o *CustomerProfileAllOf) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *CustomerProfileAllOf) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *CustomerProfileAllOf) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *CustomerProfileAllOf) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


