# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | 
**CompanyName** | Pointer to **string** |  | 
**DomainName** | Pointer to **string** | Subdomain Name for yourcompany.talon.one | 
**State** | Pointer to **string** | State of the account (active, deactivated) | 
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
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign | [optional] 

## Methods

### GetId

`func (o *Account) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Account) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Account) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Account) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Account) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *Account) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Account) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *Account) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *Account) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetCompanyName

`func (o *Account) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Account) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *Account) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *Account) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### GetDomainName

`func (o *Account) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Account) GetDomainNameOk() (string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainName

`func (o *Account) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainName

`func (o *Account) SetDomainName(v string)`

SetDomainName gets a reference to the given string and assigns it to the DomainName field.

### GetState

`func (o *Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Account) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Account) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Account) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetBillingEmail

`func (o *Account) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *Account) GetBillingEmailOk() (string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBillingEmail

`func (o *Account) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmail

`func (o *Account) SetBillingEmail(v string)`

SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.

### GetPlanName

`func (o *Account) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *Account) GetPlanNameOk() (string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanName

`func (o *Account) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### SetPlanName

`func (o *Account) SetPlanName(v string)`

SetPlanName gets a reference to the given string and assigns it to the PlanName field.

### GetPlanExpires

`func (o *Account) GetPlanExpires() time.Time`

GetPlanExpires returns the PlanExpires field if non-nil, zero value otherwise.

### GetPlanExpiresOk

`func (o *Account) GetPlanExpiresOk() (time.Time, bool)`

GetPlanExpiresOk returns a tuple with the PlanExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanExpires

`func (o *Account) HasPlanExpires() bool`

HasPlanExpires returns a boolean if a field has been set.

### SetPlanExpires

`func (o *Account) SetPlanExpires(v time.Time)`

SetPlanExpires gets a reference to the given time.Time and assigns it to the PlanExpires field.

### GetApplicationLimit

`func (o *Account) GetApplicationLimit() int32`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *Account) GetApplicationLimitOk() (int32, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationLimit

`func (o *Account) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### SetApplicationLimit

`func (o *Account) SetApplicationLimit(v int32)`

SetApplicationLimit gets a reference to the given int32 and assigns it to the ApplicationLimit field.

### GetUserLimit

`func (o *Account) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *Account) GetUserLimitOk() (int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserLimit

`func (o *Account) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimit

`func (o *Account) SetUserLimit(v int32)`

SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.

### GetCampaignLimit

`func (o *Account) GetCampaignLimit() int32`

GetCampaignLimit returns the CampaignLimit field if non-nil, zero value otherwise.

### GetCampaignLimitOk

`func (o *Account) GetCampaignLimitOk() (int32, bool)`

GetCampaignLimitOk returns a tuple with the CampaignLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignLimit

`func (o *Account) HasCampaignLimit() bool`

HasCampaignLimit returns a boolean if a field has been set.

### SetCampaignLimit

`func (o *Account) SetCampaignLimit(v int32)`

SetCampaignLimit gets a reference to the given int32 and assigns it to the CampaignLimit field.

### GetApiLimit

`func (o *Account) GetApiLimit() int32`

GetApiLimit returns the ApiLimit field if non-nil, zero value otherwise.

### GetApiLimitOk

`func (o *Account) GetApiLimitOk() (int32, bool)`

GetApiLimitOk returns a tuple with the ApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApiLimit

`func (o *Account) HasApiLimit() bool`

HasApiLimit returns a boolean if a field has been set.

### SetApiLimit

`func (o *Account) SetApiLimit(v int32)`

SetApiLimit gets a reference to the given int32 and assigns it to the ApiLimit field.

### GetApplicationCount

`func (o *Account) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *Account) GetApplicationCountOk() (int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationCount

`func (o *Account) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### SetApplicationCount

`func (o *Account) SetApplicationCount(v int32)`

SetApplicationCount gets a reference to the given int32 and assigns it to the ApplicationCount field.

### GetUserCount

`func (o *Account) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *Account) GetUserCountOk() (int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserCount

`func (o *Account) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCount

`func (o *Account) SetUserCount(v int32)`

SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.

### GetCampaignsActiveCount

`func (o *Account) GetCampaignsActiveCount() int32`

GetCampaignsActiveCount returns the CampaignsActiveCount field if non-nil, zero value otherwise.

### GetCampaignsActiveCountOk

`func (o *Account) GetCampaignsActiveCountOk() (int32, bool)`

GetCampaignsActiveCountOk returns a tuple with the CampaignsActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsActiveCount

`func (o *Account) HasCampaignsActiveCount() bool`

HasCampaignsActiveCount returns a boolean if a field has been set.

### SetCampaignsActiveCount

`func (o *Account) SetCampaignsActiveCount(v int32)`

SetCampaignsActiveCount gets a reference to the given int32 and assigns it to the CampaignsActiveCount field.

### GetCampaignsInactiveCount

`func (o *Account) GetCampaignsInactiveCount() int32`

GetCampaignsInactiveCount returns the CampaignsInactiveCount field if non-nil, zero value otherwise.

### GetCampaignsInactiveCountOk

`func (o *Account) GetCampaignsInactiveCountOk() (int32, bool)`

GetCampaignsInactiveCountOk returns a tuple with the CampaignsInactiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsInactiveCount

`func (o *Account) HasCampaignsInactiveCount() bool`

HasCampaignsInactiveCount returns a boolean if a field has been set.

### SetCampaignsInactiveCount

`func (o *Account) SetCampaignsInactiveCount(v int32)`

SetCampaignsInactiveCount gets a reference to the given int32 and assigns it to the CampaignsInactiveCount field.

### GetAttributes

`func (o *Account) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Account) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Account) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Account) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


