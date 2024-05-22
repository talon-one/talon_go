# AudienceCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. The time this entity was created. | 
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | 
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. | 
**ClosedSessions** | Pointer to **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | 
**TotalSales** | Pointer to **float32** | The total amount of money spent by the customer **before** discounts are applied.  The total sales amount excludes the following: - Cancelled or reopened sessions. - Returned items.  | 
**LoyaltyMemberships** | Pointer to [**[]LoyaltyMembership**](LoyaltyMembership.md) | **DEPRECATED** A list of loyalty programs joined by the customer.  | [optional] 
**AudienceMemberships** | Pointer to [**[]AudienceMembership**](AudienceMembership.md) | The audiences the customer belongs to. | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received from this customer. This field is updated on calls that trigger the Rule Engine and that are not [dry requests](https://docs.talon.one/docs/dev/integration-api/dry-requests/#overlay).  For example, [reserving a coupon](https://docs.talon.one/integration-api#operation/createCouponReservation) for a customer doesn&#39;t impact this field.  | 
**Sandbox** | Pointer to **bool** | An indicator of whether the customer is part of a sandbox or live Application. See the [docs](https://docs.talon.one/docs/product/applications/overview#application-environments).  | [optional] 
**ConnectedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the Applications that are connected to this customer profile. | [optional] 
**ConnectedAudiences** | Pointer to **[]int32** | A list of the IDs of the audiences that are connected to this customer profile. | [optional] 

## Methods

### GetId

`func (o *AudienceCustomer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceCustomer) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *AudienceCustomer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *AudienceCustomer) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *AudienceCustomer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AudienceCustomer) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *AudienceCustomer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *AudienceCustomer) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetIntegrationId

`func (o *AudienceCustomer) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *AudienceCustomer) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *AudienceCustomer) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *AudienceCustomer) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetAttributes

`func (o *AudienceCustomer) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AudienceCustomer) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *AudienceCustomer) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *AudienceCustomer) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetAccountId

`func (o *AudienceCustomer) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AudienceCustomer) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *AudienceCustomer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *AudienceCustomer) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetClosedSessions

`func (o *AudienceCustomer) GetClosedSessions() int32`

GetClosedSessions returns the ClosedSessions field if non-nil, zero value otherwise.

### GetClosedSessionsOk

`func (o *AudienceCustomer) GetClosedSessionsOk() (int32, bool)`

GetClosedSessionsOk returns a tuple with the ClosedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClosedSessions

`func (o *AudienceCustomer) HasClosedSessions() bool`

HasClosedSessions returns a boolean if a field has been set.

### SetClosedSessions

`func (o *AudienceCustomer) SetClosedSessions(v int32)`

SetClosedSessions gets a reference to the given int32 and assigns it to the ClosedSessions field.

### GetTotalSales

`func (o *AudienceCustomer) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *AudienceCustomer) GetTotalSalesOk() (float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSales

`func (o *AudienceCustomer) HasTotalSales() bool`

HasTotalSales returns a boolean if a field has been set.

### SetTotalSales

`func (o *AudienceCustomer) SetTotalSales(v float32)`

SetTotalSales gets a reference to the given float32 and assigns it to the TotalSales field.

### GetLoyaltyMemberships

`func (o *AudienceCustomer) GetLoyaltyMemberships() []LoyaltyMembership`

GetLoyaltyMemberships returns the LoyaltyMemberships field if non-nil, zero value otherwise.

### GetLoyaltyMembershipsOk

`func (o *AudienceCustomer) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool)`

GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyMemberships

`func (o *AudienceCustomer) HasLoyaltyMemberships() bool`

HasLoyaltyMemberships returns a boolean if a field has been set.

### SetLoyaltyMemberships

`func (o *AudienceCustomer) SetLoyaltyMemberships(v []LoyaltyMembership)`

SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.

### GetAudienceMemberships

`func (o *AudienceCustomer) GetAudienceMemberships() []AudienceMembership`

GetAudienceMemberships returns the AudienceMemberships field if non-nil, zero value otherwise.

### GetAudienceMembershipsOk

`func (o *AudienceCustomer) GetAudienceMembershipsOk() ([]AudienceMembership, bool)`

GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceMemberships

`func (o *AudienceCustomer) HasAudienceMemberships() bool`

HasAudienceMemberships returns a boolean if a field has been set.

### SetAudienceMemberships

`func (o *AudienceCustomer) SetAudienceMemberships(v []AudienceMembership)`

SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.

### GetLastActivity

`func (o *AudienceCustomer) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *AudienceCustomer) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *AudienceCustomer) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *AudienceCustomer) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetSandbox

`func (o *AudienceCustomer) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *AudienceCustomer) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *AudienceCustomer) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *AudienceCustomer) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetConnectedApplicationsIds

`func (o *AudienceCustomer) GetConnectedApplicationsIds() []int32`

GetConnectedApplicationsIds returns the ConnectedApplicationsIds field if non-nil, zero value otherwise.

### GetConnectedApplicationsIdsOk

`func (o *AudienceCustomer) GetConnectedApplicationsIdsOk() ([]int32, bool)`

GetConnectedApplicationsIdsOk returns a tuple with the ConnectedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectedApplicationsIds

`func (o *AudienceCustomer) HasConnectedApplicationsIds() bool`

HasConnectedApplicationsIds returns a boolean if a field has been set.

### SetConnectedApplicationsIds

`func (o *AudienceCustomer) SetConnectedApplicationsIds(v []int32)`

SetConnectedApplicationsIds gets a reference to the given []int32 and assigns it to the ConnectedApplicationsIds field.

### GetConnectedAudiences

`func (o *AudienceCustomer) GetConnectedAudiences() []int32`

GetConnectedAudiences returns the ConnectedAudiences field if non-nil, zero value otherwise.

### GetConnectedAudiencesOk

`func (o *AudienceCustomer) GetConnectedAudiencesOk() ([]int32, bool)`

GetConnectedAudiencesOk returns a tuple with the ConnectedAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectedAudiences

`func (o *AudienceCustomer) HasConnectedAudiences() bool`

HasConnectedAudiences returns a boolean if a field has been set.

### SetConnectedAudiences

`func (o *AudienceCustomer) SetConnectedAudiences(v []int32)`

SetConnectedAudiences gets a reference to the given []int32 and assigns it to the ConnectedAudiences field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


