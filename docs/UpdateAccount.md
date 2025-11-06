# UpdateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**CompanyName** | Pointer to **string** | Name of your company. | 
**BillingEmail** | Pointer to **string** | The billing email address associated with your company account. | 
**State** | Pointer to **string** | State of the account (active, deactivated). | [optional] 
**PlanExpires** | Pointer to [**time.Time**](time.Time.md) | The point in time at which your current plan expires. | [optional] 

## Methods

### NewUpdateAccount

`func NewUpdateAccount(companyName string, billingEmail string, ) *UpdateAccount`

NewUpdateAccount instantiates a new UpdateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountWithDefaults

`func NewUpdateAccountWithDefaults() *UpdateAccount`

NewUpdateAccountWithDefaults instantiates a new UpdateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateAccount) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateAccount) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateAccount) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateAccount) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCompanyName

`func (o *UpdateAccount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateAccount) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateAccount) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetBillingEmail

`func (o *UpdateAccount) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *UpdateAccount) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *UpdateAccount) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.


### GetState

`func (o *UpdateAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateAccount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateAccount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPlanExpires

`func (o *UpdateAccount) GetPlanExpires() time.Time`

GetPlanExpires returns the PlanExpires field if non-nil, zero value otherwise.

### GetPlanExpiresOk

`func (o *UpdateAccount) GetPlanExpiresOk() (*time.Time, bool)`

GetPlanExpiresOk returns a tuple with the PlanExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanExpires

`func (o *UpdateAccount) SetPlanExpires(v time.Time)`

SetPlanExpires sets PlanExpires field to given value.

### HasPlanExpires

`func (o *UpdateAccount) HasPlanExpires() bool`

HasPlanExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


