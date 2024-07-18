# AccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Subdomain Name for yourcompany.talon.one. | 
**State** | Pointer to **string** | State of the account (active, deactivated). | 
**BillingEmail** | Pointer to **string** | The billing email address associated with your company account. | 
**PlanName** | Pointer to **string** | The name of your booked plan. | [optional] 
**PlanExpires** | Pointer to [**time.Time**](time.Time.md) | The point in time at which your current plan expires. | [optional] 
**ApplicationLimit** | Pointer to **int32** | The maximum number of Applications covered by your plan. | [optional] 
**UserLimit** | Pointer to **int32** | The maximum number of Campaign Manager Users covered by your plan. | [optional] 
**CampaignLimit** | Pointer to **int32** | The maximum number of Campaigns covered by your plan. | [optional] 
**ApiLimit** | Pointer to **int32** | The maximum number of Integration API calls covered by your plan per billing period. | [optional] 
**ApplicationCount** | Pointer to **int32** | The current number of Applications in your account. | 
**UserCount** | Pointer to **int32** | The current number of Campaign Manager Users in your account. | 
**CampaignsActiveCount** | Pointer to **int32** | The current number of active Campaigns in your account. | 
**CampaignsInactiveCount** | Pointer to **int32** | The current number of inactive Campaigns in your account. | 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this campaign. | [optional] 

## Methods

### GetDomainName

`func (o *AccountAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AccountAllOf) GetDomainNameOk() (string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainName

`func (o *AccountAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainName

`func (o *AccountAllOf) SetDomainName(v string)`

SetDomainName gets a reference to the given string and assigns it to the DomainName field.

### GetState

`func (o *AccountAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountAllOf) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *AccountAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *AccountAllOf) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetBillingEmail

`func (o *AccountAllOf) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *AccountAllOf) GetBillingEmailOk() (string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBillingEmail

`func (o *AccountAllOf) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmail

`func (o *AccountAllOf) SetBillingEmail(v string)`

SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.

### GetPlanName

`func (o *AccountAllOf) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *AccountAllOf) GetPlanNameOk() (string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanName

`func (o *AccountAllOf) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### SetPlanName

`func (o *AccountAllOf) SetPlanName(v string)`

SetPlanName gets a reference to the given string and assigns it to the PlanName field.

### GetPlanExpires

`func (o *AccountAllOf) GetPlanExpires() time.Time`

GetPlanExpires returns the PlanExpires field if non-nil, zero value otherwise.

### GetPlanExpiresOk

`func (o *AccountAllOf) GetPlanExpiresOk() (time.Time, bool)`

GetPlanExpiresOk returns a tuple with the PlanExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanExpires

`func (o *AccountAllOf) HasPlanExpires() bool`

HasPlanExpires returns a boolean if a field has been set.

### SetPlanExpires

`func (o *AccountAllOf) SetPlanExpires(v time.Time)`

SetPlanExpires gets a reference to the given time.Time and assigns it to the PlanExpires field.

### GetApplicationLimit

`func (o *AccountAllOf) GetApplicationLimit() int32`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *AccountAllOf) GetApplicationLimitOk() (int32, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationLimit

`func (o *AccountAllOf) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### SetApplicationLimit

`func (o *AccountAllOf) SetApplicationLimit(v int32)`

SetApplicationLimit gets a reference to the given int32 and assigns it to the ApplicationLimit field.

### GetUserLimit

`func (o *AccountAllOf) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *AccountAllOf) GetUserLimitOk() (int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserLimit

`func (o *AccountAllOf) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimit

`func (o *AccountAllOf) SetUserLimit(v int32)`

SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.

### GetCampaignLimit

`func (o *AccountAllOf) GetCampaignLimit() int32`

GetCampaignLimit returns the CampaignLimit field if non-nil, zero value otherwise.

### GetCampaignLimitOk

`func (o *AccountAllOf) GetCampaignLimitOk() (int32, bool)`

GetCampaignLimitOk returns a tuple with the CampaignLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignLimit

`func (o *AccountAllOf) HasCampaignLimit() bool`

HasCampaignLimit returns a boolean if a field has been set.

### SetCampaignLimit

`func (o *AccountAllOf) SetCampaignLimit(v int32)`

SetCampaignLimit gets a reference to the given int32 and assigns it to the CampaignLimit field.

### GetApiLimit

`func (o *AccountAllOf) GetApiLimit() int32`

GetApiLimit returns the ApiLimit field if non-nil, zero value otherwise.

### GetApiLimitOk

`func (o *AccountAllOf) GetApiLimitOk() (int32, bool)`

GetApiLimitOk returns a tuple with the ApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApiLimit

`func (o *AccountAllOf) HasApiLimit() bool`

HasApiLimit returns a boolean if a field has been set.

### SetApiLimit

`func (o *AccountAllOf) SetApiLimit(v int32)`

SetApiLimit gets a reference to the given int32 and assigns it to the ApiLimit field.

### GetApplicationCount

`func (o *AccountAllOf) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *AccountAllOf) GetApplicationCountOk() (int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationCount

`func (o *AccountAllOf) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### SetApplicationCount

`func (o *AccountAllOf) SetApplicationCount(v int32)`

SetApplicationCount gets a reference to the given int32 and assigns it to the ApplicationCount field.

### GetUserCount

`func (o *AccountAllOf) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *AccountAllOf) GetUserCountOk() (int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserCount

`func (o *AccountAllOf) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCount

`func (o *AccountAllOf) SetUserCount(v int32)`

SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.

### GetCampaignsActiveCount

`func (o *AccountAllOf) GetCampaignsActiveCount() int32`

GetCampaignsActiveCount returns the CampaignsActiveCount field if non-nil, zero value otherwise.

### GetCampaignsActiveCountOk

`func (o *AccountAllOf) GetCampaignsActiveCountOk() (int32, bool)`

GetCampaignsActiveCountOk returns a tuple with the CampaignsActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsActiveCount

`func (o *AccountAllOf) HasCampaignsActiveCount() bool`

HasCampaignsActiveCount returns a boolean if a field has been set.

### SetCampaignsActiveCount

`func (o *AccountAllOf) SetCampaignsActiveCount(v int32)`

SetCampaignsActiveCount gets a reference to the given int32 and assigns it to the CampaignsActiveCount field.

### GetCampaignsInactiveCount

`func (o *AccountAllOf) GetCampaignsInactiveCount() int32`

GetCampaignsInactiveCount returns the CampaignsInactiveCount field if non-nil, zero value otherwise.

### GetCampaignsInactiveCountOk

`func (o *AccountAllOf) GetCampaignsInactiveCountOk() (int32, bool)`

GetCampaignsInactiveCountOk returns a tuple with the CampaignsInactiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsInactiveCount

`func (o *AccountAllOf) HasCampaignsInactiveCount() bool`

HasCampaignsInactiveCount returns a boolean if a field has been set.

### SetCampaignsInactiveCount

`func (o *AccountAllOf) SetCampaignsInactiveCount(v int32)`

SetCampaignsInactiveCount gets a reference to the given int32 and assigns it to the CampaignsInactiveCount field.

### GetAttributes

`func (o *AccountAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AccountAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *AccountAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *AccountAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


