# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**CompanyName** | Pointer to **string** |  | 
**DomainName** | Pointer to **string** | Subdomain Name for yourcompany.talon.one. | 
**State** | Pointer to **string** | State of the account (active, deactivated). | 
**BillingEmail** | Pointer to **string** | The billing email address associated with your company account. | 
**PlanName** | Pointer to **string** | The name of your booked plan. | [optional] 
**PlanExpires** | Pointer to [**time.Time**](time.Time.md) | The point in time at which your current plan expires. | [optional] 
**ApplicationLimit** | Pointer to **int64** | The maximum number of Applications covered by your plan. | [optional] 
**UserLimit** | Pointer to **int64** | The maximum number of Campaign Manager Users covered by your plan. | [optional] 
**CampaignLimit** | Pointer to **int64** | The maximum number of Campaigns covered by your plan. | [optional] 
**ApiLimit** | Pointer to **int64** | The maximum number of Integration API calls covered by your plan per billing period. | [optional] 
**ApplicationCount** | Pointer to **int64** | The current number of Applications in your account. | 
**UserCount** | Pointer to **int64** | The current number of Campaign Manager Users in your account. | 
**CampaignsActiveCount** | Pointer to **int64** | The current number of active Campaigns in your account. | 
**CampaignsInactiveCount** | Pointer to **int64** | The current number of inactive Campaigns in your account. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 

## Methods

### NewAccount

`func NewAccount(id int64, created time.Time, modified time.Time, companyName string, domainName string, state string, billingEmail string, applicationCount int64, userCount int64, campaignsActiveCount int64, campaignsInactiveCount int64, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Account) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Account) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Account) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Account) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Account) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetCompanyName

`func (o *Account) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Account) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Account) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetDomainName

`func (o *Account) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Account) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Account) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetState

`func (o *Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Account) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Account) SetState(v string)`

SetState sets State field to given value.


### GetBillingEmail

`func (o *Account) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *Account) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *Account) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.


### GetPlanName

`func (o *Account) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *Account) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *Account) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *Account) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanExpires

`func (o *Account) GetPlanExpires() time.Time`

GetPlanExpires returns the PlanExpires field if non-nil, zero value otherwise.

### GetPlanExpiresOk

`func (o *Account) GetPlanExpiresOk() (*time.Time, bool)`

GetPlanExpiresOk returns a tuple with the PlanExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanExpires

`func (o *Account) SetPlanExpires(v time.Time)`

SetPlanExpires sets PlanExpires field to given value.

### HasPlanExpires

`func (o *Account) HasPlanExpires() bool`

HasPlanExpires returns a boolean if a field has been set.

### GetApplicationLimit

`func (o *Account) GetApplicationLimit() int64`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *Account) GetApplicationLimitOk() (*int64, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLimit

`func (o *Account) SetApplicationLimit(v int64)`

SetApplicationLimit sets ApplicationLimit field to given value.

### HasApplicationLimit

`func (o *Account) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### GetUserLimit

`func (o *Account) GetUserLimit() int64`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *Account) GetUserLimitOk() (*int64, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *Account) SetUserLimit(v int64)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *Account) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### GetCampaignLimit

`func (o *Account) GetCampaignLimit() int64`

GetCampaignLimit returns the CampaignLimit field if non-nil, zero value otherwise.

### GetCampaignLimitOk

`func (o *Account) GetCampaignLimitOk() (*int64, bool)`

GetCampaignLimitOk returns a tuple with the CampaignLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignLimit

`func (o *Account) SetCampaignLimit(v int64)`

SetCampaignLimit sets CampaignLimit field to given value.

### HasCampaignLimit

`func (o *Account) HasCampaignLimit() bool`

HasCampaignLimit returns a boolean if a field has been set.

### GetApiLimit

`func (o *Account) GetApiLimit() int64`

GetApiLimit returns the ApiLimit field if non-nil, zero value otherwise.

### GetApiLimitOk

`func (o *Account) GetApiLimitOk() (*int64, bool)`

GetApiLimitOk returns a tuple with the ApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiLimit

`func (o *Account) SetApiLimit(v int64)`

SetApiLimit sets ApiLimit field to given value.

### HasApiLimit

`func (o *Account) HasApiLimit() bool`

HasApiLimit returns a boolean if a field has been set.

### GetApplicationCount

`func (o *Account) GetApplicationCount() int64`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *Account) GetApplicationCountOk() (*int64, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *Account) SetApplicationCount(v int64)`

SetApplicationCount sets ApplicationCount field to given value.


### GetUserCount

`func (o *Account) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *Account) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *Account) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.


### GetCampaignsActiveCount

`func (o *Account) GetCampaignsActiveCount() int64`

GetCampaignsActiveCount returns the CampaignsActiveCount field if non-nil, zero value otherwise.

### GetCampaignsActiveCountOk

`func (o *Account) GetCampaignsActiveCountOk() (*int64, bool)`

GetCampaignsActiveCountOk returns a tuple with the CampaignsActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignsActiveCount

`func (o *Account) SetCampaignsActiveCount(v int64)`

SetCampaignsActiveCount sets CampaignsActiveCount field to given value.


### GetCampaignsInactiveCount

`func (o *Account) GetCampaignsInactiveCount() int64`

GetCampaignsInactiveCount returns the CampaignsInactiveCount field if non-nil, zero value otherwise.

### GetCampaignsInactiveCountOk

`func (o *Account) GetCampaignsInactiveCountOk() (*int64, bool)`

GetCampaignsInactiveCountOk returns a tuple with the CampaignsInactiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignsInactiveCount

`func (o *Account) SetCampaignsInactiveCount(v int64)`

SetCampaignsInactiveCount sets CampaignsInactiveCount field to given value.


### GetAttributes

`func (o *Account) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Account) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Account) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Account) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


