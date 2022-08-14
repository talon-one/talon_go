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

### GetAttributes

`func (o *UpdateAccount) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateAccount) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *UpdateAccount) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *UpdateAccount) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetCompanyName

`func (o *UpdateAccount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateAccount) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *UpdateAccount) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *UpdateAccount) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### GetBillingEmail

`func (o *UpdateAccount) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *UpdateAccount) GetBillingEmailOk() (string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBillingEmail

`func (o *UpdateAccount) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmail

`func (o *UpdateAccount) SetBillingEmail(v string)`

SetBillingEmail gets a reference to the given string and assigns it to the BillingEmail field.

### GetState

`func (o *UpdateAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateAccount) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *UpdateAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *UpdateAccount) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetPlanExpires

`func (o *UpdateAccount) GetPlanExpires() time.Time`

GetPlanExpires returns the PlanExpires field if non-nil, zero value otherwise.

### GetPlanExpiresOk

`func (o *UpdateAccount) GetPlanExpiresOk() (time.Time, bool)`

GetPlanExpiresOk returns a tuple with the PlanExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanExpires

`func (o *UpdateAccount) HasPlanExpires() bool`

HasPlanExpires returns a boolean if a field has been set.

### SetPlanExpires

`func (o *UpdateAccount) SetPlanExpires(v time.Time)`

SetPlanExpires gets a reference to the given time.Time and assigns it to the PlanExpires field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


