# BaseLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display title for the Loyalty Program. | [optional] 
**Description** | Pointer to **string** | Description of our Loyalty Program. | [optional] 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] 
**DefaultValidity** | Pointer to **string** | The default duration after which new loyalty points should expire. Can be &#39;unlimited&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | [optional] 
**DefaultPending** | Pointer to **string** | The default duration of the pending time after which points should be valid. Can be &#39;immediate&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | [optional] 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program. | [optional] 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type. | [optional] 

## Methods

### GetTitle

`func (o *BaseLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BaseLoyaltyProgram) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *BaseLoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *BaseLoyaltyProgram) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *BaseLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseLoyaltyProgram) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *BaseLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *BaseLoyaltyProgram) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplications

`func (o *BaseLoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *BaseLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplications

`func (o *BaseLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### SetSubscribedApplications

`func (o *BaseLoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.

### GetDefaultValidity

`func (o *BaseLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *BaseLoyaltyProgram) GetDefaultValidityOk() (string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValidity

`func (o *BaseLoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### SetDefaultValidity

`func (o *BaseLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.

### GetDefaultPending

`func (o *BaseLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *BaseLoyaltyProgram) GetDefaultPendingOk() (string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultPending

`func (o *BaseLoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### SetDefaultPending

`func (o *BaseLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.

### GetAllowSubledger

`func (o *BaseLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *BaseLoyaltyProgram) GetAllowSubledgerOk() (bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowSubledger

`func (o *BaseLoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### SetAllowSubledger

`func (o *BaseLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.

### GetUsersPerCardLimit

`func (o *BaseLoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *BaseLoyaltyProgram) GetUsersPerCardLimitOk() (int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsersPerCardLimit

`func (o *BaseLoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### SetUsersPerCardLimit

`func (o *BaseLoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.

### GetSandbox

`func (o *BaseLoyaltyProgram) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *BaseLoyaltyProgram) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *BaseLoyaltyProgram) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *BaseLoyaltyProgram) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


