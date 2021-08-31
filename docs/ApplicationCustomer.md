# ApplicationCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. The exact moment this entity was created. The exact moment this entity was created. The exact moment this entity was created. | 
**IntegrationId** | Pointer to **string** | The integration ID for this entity sent to and used in the Talon.One system. The integration ID for this entity sent to and used in the Talon.One system. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | 
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. The ID of the Talon.One account that owns this profile. | 
**ClosedSessions** | Pointer to **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | 
**TotalSales** | Pointer to **float32** | Sum of all purchases made by this customer | 
**LoyaltyMemberships** | Pointer to [**[]LoyaltyMembership**](LoyaltyMembership.md) | A list of loyalty programs joined by the customer | [optional] 
**AudienceMemberships** | Pointer to [**[]AudienceMembership**](AudienceMembership.md) | A list of audiences the customer belongs to | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received from this customer | 
**AdvocateIntegrationId** | Pointer to **string** | The Integration ID of the Customer Profile that referred this Customer in the Application. | [optional] 

## Methods

### GetId

`func (o *ApplicationCustomer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCustomer) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationCustomer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationCustomer) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *ApplicationCustomer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationCustomer) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationCustomer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationCustomer) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetIntegrationId

`func (o *ApplicationCustomer) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ApplicationCustomer) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *ApplicationCustomer) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *ApplicationCustomer) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetAttributes

`func (o *ApplicationCustomer) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationCustomer) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationCustomer) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationCustomer) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetAccountId

`func (o *ApplicationCustomer) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApplicationCustomer) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *ApplicationCustomer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *ApplicationCustomer) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetClosedSessions

`func (o *ApplicationCustomer) GetClosedSessions() int32`

GetClosedSessions returns the ClosedSessions field if non-nil, zero value otherwise.

### GetClosedSessionsOk

`func (o *ApplicationCustomer) GetClosedSessionsOk() (int32, bool)`

GetClosedSessionsOk returns a tuple with the ClosedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClosedSessions

`func (o *ApplicationCustomer) HasClosedSessions() bool`

HasClosedSessions returns a boolean if a field has been set.

### SetClosedSessions

`func (o *ApplicationCustomer) SetClosedSessions(v int32)`

SetClosedSessions gets a reference to the given int32 and assigns it to the ClosedSessions field.

### GetTotalSales

`func (o *ApplicationCustomer) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *ApplicationCustomer) GetTotalSalesOk() (float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSales

`func (o *ApplicationCustomer) HasTotalSales() bool`

HasTotalSales returns a boolean if a field has been set.

### SetTotalSales

`func (o *ApplicationCustomer) SetTotalSales(v float32)`

SetTotalSales gets a reference to the given float32 and assigns it to the TotalSales field.

### GetLoyaltyMemberships

`func (o *ApplicationCustomer) GetLoyaltyMemberships() []LoyaltyMembership`

GetLoyaltyMemberships returns the LoyaltyMemberships field if non-nil, zero value otherwise.

### GetLoyaltyMembershipsOk

`func (o *ApplicationCustomer) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool)`

GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyMemberships

`func (o *ApplicationCustomer) HasLoyaltyMemberships() bool`

HasLoyaltyMemberships returns a boolean if a field has been set.

### SetLoyaltyMemberships

`func (o *ApplicationCustomer) SetLoyaltyMemberships(v []LoyaltyMembership)`

SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.

### GetAudienceMemberships

`func (o *ApplicationCustomer) GetAudienceMemberships() []AudienceMembership`

GetAudienceMemberships returns the AudienceMemberships field if non-nil, zero value otherwise.

### GetAudienceMembershipsOk

`func (o *ApplicationCustomer) GetAudienceMembershipsOk() ([]AudienceMembership, bool)`

GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceMemberships

`func (o *ApplicationCustomer) HasAudienceMemberships() bool`

HasAudienceMemberships returns a boolean if a field has been set.

### SetAudienceMemberships

`func (o *ApplicationCustomer) SetAudienceMemberships(v []AudienceMembership)`

SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.

### GetLastActivity

`func (o *ApplicationCustomer) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *ApplicationCustomer) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *ApplicationCustomer) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *ApplicationCustomer) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetAdvocateIntegrationId

`func (o *ApplicationCustomer) GetAdvocateIntegrationId() string`

GetAdvocateIntegrationId returns the AdvocateIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateIntegrationIdOk

`func (o *ApplicationCustomer) GetAdvocateIntegrationIdOk() (string, bool)`

GetAdvocateIntegrationIdOk returns a tuple with the AdvocateIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdvocateIntegrationId

`func (o *ApplicationCustomer) HasAdvocateIntegrationId() bool`

HasAdvocateIntegrationId returns a boolean if a field has been set.

### SetAdvocateIntegrationId

`func (o *ApplicationCustomer) SetAdvocateIntegrationId(v string)`

SetAdvocateIntegrationId gets a reference to the given string and assigns it to the AdvocateIntegrationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


