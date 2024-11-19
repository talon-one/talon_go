# UpdateLoyaltyProgram

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
**ProgramJoinPolicy** | Pointer to **string** | The policy that defines when the customer joins the loyalty program.   - &#x60;not_join&#x60;: The customer does not join the loyalty program but can still earn and spend loyalty points.       **Note**: The customer does not have a program join date.   - &#x60;points_activated&#x60;: The customer joins the loyalty program only when their earned loyalty points become active for the first time.   - &#x60;points_earned&#x60;: The customer joins the loyalty program when they earn loyalty points for the first time.  | [optional] 
**TiersExpirationPolicy** | Pointer to **string** | The policy that defines how tier expiration, used to reevaluate the customer&#39;s current tier, is determined.  - &#x60;tier_start_date&#x60;: The tier expiration is relative to when the customer joined the current tier.  - &#x60;program_join_date&#x60;: The tier expiration is relative to when the customer joined the loyalty program.  - &#x60;customer_attribute&#x60;: The tier expiration is determined by a custom customer attribute.  - &#x60;absolute_expiration&#x60;: The tier is reevaluated at the start of each tier cycle. For this policy, it is required to provide a &#x60;tierCycleStartDate&#x60;.  | [optional] 
**TierCycleStartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the tier cycle starts for all customers in the loyalty program.  **Note**: This is only required when the tier expiration policy is set to &#x60;absolute_expiration&#x60;.  | [optional] 
**TiersExpireIn** | Pointer to **string** | The amount of time after which the tier expires and is reevaluated.  The time format is an **integer** followed by one letter indicating the time unit. Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | [optional] 
**TiersDowngradePolicy** | Pointer to **string** | The policy that defines how customer tiers are downgraded in the loyalty program after tier reevaluation.  - &#x60;one_down&#x60;: If the customer doesn&#39;t have enough points to stay in the current tier, they are downgraded by one tier.  - &#x60;balance_based&#x60;: The customer&#39;s tier is reevaluated based on the amount of active points they have at the moment.  | [optional] 
**CardCodeSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program. | [optional] 

## Methods

### GetTitle

`func (o *UpdateLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateLoyaltyProgram) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *UpdateLoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *UpdateLoyaltyProgram) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *UpdateLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoyaltyProgram) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateLoyaltyProgram) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplications

`func (o *UpdateLoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *UpdateLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplications

`func (o *UpdateLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### SetSubscribedApplications

`func (o *UpdateLoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.

### GetDefaultValidity

`func (o *UpdateLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *UpdateLoyaltyProgram) GetDefaultValidityOk() (string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValidity

`func (o *UpdateLoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### SetDefaultValidity

`func (o *UpdateLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.

### GetDefaultPending

`func (o *UpdateLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *UpdateLoyaltyProgram) GetDefaultPendingOk() (string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultPending

`func (o *UpdateLoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### SetDefaultPending

`func (o *UpdateLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.

### GetAllowSubledger

`func (o *UpdateLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *UpdateLoyaltyProgram) GetAllowSubledgerOk() (bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowSubledger

`func (o *UpdateLoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### SetAllowSubledger

`func (o *UpdateLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.

### GetUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *UpdateLoyaltyProgram) GetUsersPerCardLimitOk() (int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### SetUsersPerCardLimit

`func (o *UpdateLoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.

### GetSandbox

`func (o *UpdateLoyaltyProgram) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *UpdateLoyaltyProgram) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *UpdateLoyaltyProgram) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *UpdateLoyaltyProgram) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetProgramJoinPolicy

`func (o *UpdateLoyaltyProgram) GetProgramJoinPolicy() string`

GetProgramJoinPolicy returns the ProgramJoinPolicy field if non-nil, zero value otherwise.

### GetProgramJoinPolicyOk

`func (o *UpdateLoyaltyProgram) GetProgramJoinPolicyOk() (string, bool)`

GetProgramJoinPolicyOk returns a tuple with the ProgramJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramJoinPolicy

`func (o *UpdateLoyaltyProgram) HasProgramJoinPolicy() bool`

HasProgramJoinPolicy returns a boolean if a field has been set.

### SetProgramJoinPolicy

`func (o *UpdateLoyaltyProgram) SetProgramJoinPolicy(v string)`

SetProgramJoinPolicy gets a reference to the given string and assigns it to the ProgramJoinPolicy field.

### GetTiersExpirationPolicy

`func (o *UpdateLoyaltyProgram) GetTiersExpirationPolicy() string`

GetTiersExpirationPolicy returns the TiersExpirationPolicy field if non-nil, zero value otherwise.

### GetTiersExpirationPolicyOk

`func (o *UpdateLoyaltyProgram) GetTiersExpirationPolicyOk() (string, bool)`

GetTiersExpirationPolicyOk returns a tuple with the TiersExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersExpirationPolicy

`func (o *UpdateLoyaltyProgram) HasTiersExpirationPolicy() bool`

HasTiersExpirationPolicy returns a boolean if a field has been set.

### SetTiersExpirationPolicy

`func (o *UpdateLoyaltyProgram) SetTiersExpirationPolicy(v string)`

SetTiersExpirationPolicy gets a reference to the given string and assigns it to the TiersExpirationPolicy field.

### GetTierCycleStartDate

`func (o *UpdateLoyaltyProgram) GetTierCycleStartDate() time.Time`

GetTierCycleStartDate returns the TierCycleStartDate field if non-nil, zero value otherwise.

### GetTierCycleStartDateOk

`func (o *UpdateLoyaltyProgram) GetTierCycleStartDateOk() (time.Time, bool)`

GetTierCycleStartDateOk returns a tuple with the TierCycleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTierCycleStartDate

`func (o *UpdateLoyaltyProgram) HasTierCycleStartDate() bool`

HasTierCycleStartDate returns a boolean if a field has been set.

### SetTierCycleStartDate

`func (o *UpdateLoyaltyProgram) SetTierCycleStartDate(v time.Time)`

SetTierCycleStartDate gets a reference to the given time.Time and assigns it to the TierCycleStartDate field.

### GetTiersExpireIn

`func (o *UpdateLoyaltyProgram) GetTiersExpireIn() string`

GetTiersExpireIn returns the TiersExpireIn field if non-nil, zero value otherwise.

### GetTiersExpireInOk

`func (o *UpdateLoyaltyProgram) GetTiersExpireInOk() (string, bool)`

GetTiersExpireInOk returns a tuple with the TiersExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersExpireIn

`func (o *UpdateLoyaltyProgram) HasTiersExpireIn() bool`

HasTiersExpireIn returns a boolean if a field has been set.

### SetTiersExpireIn

`func (o *UpdateLoyaltyProgram) SetTiersExpireIn(v string)`

SetTiersExpireIn gets a reference to the given string and assigns it to the TiersExpireIn field.

### GetTiersDowngradePolicy

`func (o *UpdateLoyaltyProgram) GetTiersDowngradePolicy() string`

GetTiersDowngradePolicy returns the TiersDowngradePolicy field if non-nil, zero value otherwise.

### GetTiersDowngradePolicyOk

`func (o *UpdateLoyaltyProgram) GetTiersDowngradePolicyOk() (string, bool)`

GetTiersDowngradePolicyOk returns a tuple with the TiersDowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersDowngradePolicy

`func (o *UpdateLoyaltyProgram) HasTiersDowngradePolicy() bool`

HasTiersDowngradePolicy returns a boolean if a field has been set.

### SetTiersDowngradePolicy

`func (o *UpdateLoyaltyProgram) SetTiersDowngradePolicy(v string)`

SetTiersDowngradePolicy gets a reference to the given string and assigns it to the TiersDowngradePolicy field.

### GetCardCodeSettings

`func (o *UpdateLoyaltyProgram) GetCardCodeSettings() CodeGeneratorSettings`

GetCardCodeSettings returns the CardCodeSettings field if non-nil, zero value otherwise.

### GetCardCodeSettingsOk

`func (o *UpdateLoyaltyProgram) GetCardCodeSettingsOk() (CodeGeneratorSettings, bool)`

GetCardCodeSettingsOk returns a tuple with the CardCodeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardCodeSettings

`func (o *UpdateLoyaltyProgram) HasCardCodeSettings() bool`

HasCardCodeSettings returns a boolean if a field has been set.

### SetCardCodeSettings

`func (o *UpdateLoyaltyProgram) SetCardCodeSettings(v CodeGeneratorSettings)`

SetCardCodeSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CardCodeSettings field.

### GetTiers

`func (o *UpdateLoyaltyProgram) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *UpdateLoyaltyProgram) GetTiersOk() ([]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *UpdateLoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *UpdateLoyaltyProgram) SetTiers(v []NewLoyaltyTier)`

SetTiers gets a reference to the given []NewLoyaltyTier and assigns it to the Tiers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


