# LoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of loyalty program. Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Title** | Pointer to **string** | The display title for the Loyalty Program. | 
**Description** | Pointer to **string** | Description of our Loyalty Program. | 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | 
**DefaultValidity** | Pointer to **string** | The default duration after which new loyalty points should expire. Can be &#39;unlimited&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | 
**DefaultPending** | Pointer to **string** | The default duration of the pending time after which points should be valid. Can be &#39;immediate&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program. | 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type. | 
**ProgramJoinPolicy** | Pointer to **string** | The policy that defines when the customer joins the loyalty program.   - &#x60;not_join&#x60;: The customer does not join the loyalty program but can still earn and spend loyalty points.       **Note**: The customer does not have a program join date.   - &#x60;points_activated&#x60;: The customer joins the loyalty program only when their earned loyalty points become active for the first time.   - &#x60;points_earned&#x60;: The customer joins the loyalty program when they earn loyalty points for the first time.  | [optional] 
**TiersExpirationPolicy** | Pointer to **string** | The policy that defines how tier expiration, used to reevaluate the customer&#39;s current tier, is determined.  - &#x60;tier_start_date&#x60;: The tier expiration is relative to when the customer joined the current tier.  - &#x60;program_join_date&#x60;: The tier expiration is relative to when the customer joined the loyalty program.  - &#x60;customer_attribute&#x60;: The tier expiration is determined by a custom customer attribute.  - &#x60;absolute_expiration&#x60;: The tier is reevaluated at the start of each tier cycle. For this policy, it is required to provide a &#x60;tierCycleStartDate&#x60;.  | [optional] 
**TierCycleStartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the tier cycle starts for all customers in the loyalty program.  **Note**: This is only required when the tier expiration policy is set to &#x60;absolute_expiration&#x60;.  | [optional] 
**TiersExpireIn** | Pointer to **string** | The amount of time after which the tier expires and is reevaluated.  The time format is an **integer** followed by one letter indicating the time unit. Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | [optional] 
**TiersDowngradePolicy** | Pointer to **string** | The policy that defines how customer tiers are downgraded in the loyalty program after tier reevaluation.  - &#x60;one_down&#x60;: If the customer doesn&#39;t have enough points to stay in the current tier, they are downgraded by one tier.  - &#x60;balance_based&#x60;: The customer&#39;s tier is reevaluated based on the amount of active points they have at the moment.  | [optional] 
**CardCodeSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**AccountID** | Pointer to **int32** | The ID of the Talon.One account that owns this program. | 
**Name** | Pointer to **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]LoyaltyTier**](LoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | Pointer to **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]
**CanUpdateTiers** | Pointer to **bool** | &#x60;True&#x60; if the tier definitions can be updated.  | [optional] [default to false]
**CanUpdateJoinPolicy** | Pointer to **bool** | &#x60;True&#x60; if the program join policy can be updated.  | [optional] 
**CanUpdateTierExpirationPolicy** | Pointer to **bool** | &#x60;True&#x60; if the tier expiration policy can be updated.  | [optional] 
**CanUpgradeToAdvancedTiers** | Pointer to **bool** | &#x60;True&#x60; if the program can be upgraded to use the &#x60;tiersExpireIn&#x60; and &#x60;tiersDowngradePolicy&#x60; properties.  | [optional] [default to false]
**CanUpdateSubledgers** | Pointer to **bool** | &#x60;True&#x60; if the &#x60;allowSubledger&#x60; property can be updated in the loyalty program.  | [optional] [default to false]

## Methods

### GetId

`func (o *LoyaltyProgram) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgram) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyProgram) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyProgram) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *LoyaltyProgram) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyProgram) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LoyaltyProgram) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LoyaltyProgram) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetTitle

`func (o *LoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgram) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *LoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *LoyaltyProgram) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *LoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoyaltyProgram) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *LoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *LoyaltyProgram) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplications

`func (o *LoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *LoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplications

`func (o *LoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### SetSubscribedApplications

`func (o *LoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.

### GetDefaultValidity

`func (o *LoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *LoyaltyProgram) GetDefaultValidityOk() (string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValidity

`func (o *LoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### SetDefaultValidity

`func (o *LoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.

### GetDefaultPending

`func (o *LoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *LoyaltyProgram) GetDefaultPendingOk() (string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultPending

`func (o *LoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### SetDefaultPending

`func (o *LoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.

### GetAllowSubledger

`func (o *LoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *LoyaltyProgram) GetAllowSubledgerOk() (bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowSubledger

`func (o *LoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### SetAllowSubledger

`func (o *LoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.

### GetUsersPerCardLimit

`func (o *LoyaltyProgram) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyProgram) GetUsersPerCardLimitOk() (int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsersPerCardLimit

`func (o *LoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyProgram) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.

### GetSandbox

`func (o *LoyaltyProgram) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *LoyaltyProgram) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *LoyaltyProgram) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *LoyaltyProgram) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetProgramJoinPolicy

`func (o *LoyaltyProgram) GetProgramJoinPolicy() string`

GetProgramJoinPolicy returns the ProgramJoinPolicy field if non-nil, zero value otherwise.

### GetProgramJoinPolicyOk

`func (o *LoyaltyProgram) GetProgramJoinPolicyOk() (string, bool)`

GetProgramJoinPolicyOk returns a tuple with the ProgramJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramJoinPolicy

`func (o *LoyaltyProgram) HasProgramJoinPolicy() bool`

HasProgramJoinPolicy returns a boolean if a field has been set.

### SetProgramJoinPolicy

`func (o *LoyaltyProgram) SetProgramJoinPolicy(v string)`

SetProgramJoinPolicy gets a reference to the given string and assigns it to the ProgramJoinPolicy field.

### GetTiersExpirationPolicy

`func (o *LoyaltyProgram) GetTiersExpirationPolicy() string`

GetTiersExpirationPolicy returns the TiersExpirationPolicy field if non-nil, zero value otherwise.

### GetTiersExpirationPolicyOk

`func (o *LoyaltyProgram) GetTiersExpirationPolicyOk() (string, bool)`

GetTiersExpirationPolicyOk returns a tuple with the TiersExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersExpirationPolicy

`func (o *LoyaltyProgram) HasTiersExpirationPolicy() bool`

HasTiersExpirationPolicy returns a boolean if a field has been set.

### SetTiersExpirationPolicy

`func (o *LoyaltyProgram) SetTiersExpirationPolicy(v string)`

SetTiersExpirationPolicy gets a reference to the given string and assigns it to the TiersExpirationPolicy field.

### GetTierCycleStartDate

`func (o *LoyaltyProgram) GetTierCycleStartDate() time.Time`

GetTierCycleStartDate returns the TierCycleStartDate field if non-nil, zero value otherwise.

### GetTierCycleStartDateOk

`func (o *LoyaltyProgram) GetTierCycleStartDateOk() (time.Time, bool)`

GetTierCycleStartDateOk returns a tuple with the TierCycleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTierCycleStartDate

`func (o *LoyaltyProgram) HasTierCycleStartDate() bool`

HasTierCycleStartDate returns a boolean if a field has been set.

### SetTierCycleStartDate

`func (o *LoyaltyProgram) SetTierCycleStartDate(v time.Time)`

SetTierCycleStartDate gets a reference to the given time.Time and assigns it to the TierCycleStartDate field.

### GetTiersExpireIn

`func (o *LoyaltyProgram) GetTiersExpireIn() string`

GetTiersExpireIn returns the TiersExpireIn field if non-nil, zero value otherwise.

### GetTiersExpireInOk

`func (o *LoyaltyProgram) GetTiersExpireInOk() (string, bool)`

GetTiersExpireInOk returns a tuple with the TiersExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersExpireIn

`func (o *LoyaltyProgram) HasTiersExpireIn() bool`

HasTiersExpireIn returns a boolean if a field has been set.

### SetTiersExpireIn

`func (o *LoyaltyProgram) SetTiersExpireIn(v string)`

SetTiersExpireIn gets a reference to the given string and assigns it to the TiersExpireIn field.

### GetTiersDowngradePolicy

`func (o *LoyaltyProgram) GetTiersDowngradePolicy() string`

GetTiersDowngradePolicy returns the TiersDowngradePolicy field if non-nil, zero value otherwise.

### GetTiersDowngradePolicyOk

`func (o *LoyaltyProgram) GetTiersDowngradePolicyOk() (string, bool)`

GetTiersDowngradePolicyOk returns a tuple with the TiersDowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiersDowngradePolicy

`func (o *LoyaltyProgram) HasTiersDowngradePolicy() bool`

HasTiersDowngradePolicy returns a boolean if a field has been set.

### SetTiersDowngradePolicy

`func (o *LoyaltyProgram) SetTiersDowngradePolicy(v string)`

SetTiersDowngradePolicy gets a reference to the given string and assigns it to the TiersDowngradePolicy field.

### GetCardCodeSettings

`func (o *LoyaltyProgram) GetCardCodeSettings() CodeGeneratorSettings`

GetCardCodeSettings returns the CardCodeSettings field if non-nil, zero value otherwise.

### GetCardCodeSettingsOk

`func (o *LoyaltyProgram) GetCardCodeSettingsOk() (CodeGeneratorSettings, bool)`

GetCardCodeSettingsOk returns a tuple with the CardCodeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardCodeSettings

`func (o *LoyaltyProgram) HasCardCodeSettings() bool`

HasCardCodeSettings returns a boolean if a field has been set.

### SetCardCodeSettings

`func (o *LoyaltyProgram) SetCardCodeSettings(v CodeGeneratorSettings)`

SetCardCodeSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CardCodeSettings field.

### GetAccountID

`func (o *LoyaltyProgram) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *LoyaltyProgram) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *LoyaltyProgram) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *LoyaltyProgram) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetName

`func (o *LoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgram) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyProgram) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyProgram) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTiers

`func (o *LoyaltyProgram) GetTiers() []LoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *LoyaltyProgram) GetTiersOk() ([]LoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *LoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *LoyaltyProgram) SetTiers(v []LoyaltyTier)`

SetTiers gets a reference to the given []LoyaltyTier and assigns it to the Tiers field.

### GetTimezone

`func (o *LoyaltyProgram) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LoyaltyProgram) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *LoyaltyProgram) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *LoyaltyProgram) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCardBased

`func (o *LoyaltyProgram) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *LoyaltyProgram) GetCardBasedOk() (bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardBased

`func (o *LoyaltyProgram) HasCardBased() bool`

HasCardBased returns a boolean if a field has been set.

### SetCardBased

`func (o *LoyaltyProgram) SetCardBased(v bool)`

SetCardBased gets a reference to the given bool and assigns it to the CardBased field.

### GetCanUpdateTiers

`func (o *LoyaltyProgram) GetCanUpdateTiers() bool`

GetCanUpdateTiers returns the CanUpdateTiers field if non-nil, zero value otherwise.

### GetCanUpdateTiersOk

`func (o *LoyaltyProgram) GetCanUpdateTiersOk() (bool, bool)`

GetCanUpdateTiersOk returns a tuple with the CanUpdateTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateTiers

`func (o *LoyaltyProgram) HasCanUpdateTiers() bool`

HasCanUpdateTiers returns a boolean if a field has been set.

### SetCanUpdateTiers

`func (o *LoyaltyProgram) SetCanUpdateTiers(v bool)`

SetCanUpdateTiers gets a reference to the given bool and assigns it to the CanUpdateTiers field.

### GetCanUpdateJoinPolicy

`func (o *LoyaltyProgram) GetCanUpdateJoinPolicy() bool`

GetCanUpdateJoinPolicy returns the CanUpdateJoinPolicy field if non-nil, zero value otherwise.

### GetCanUpdateJoinPolicyOk

`func (o *LoyaltyProgram) GetCanUpdateJoinPolicyOk() (bool, bool)`

GetCanUpdateJoinPolicyOk returns a tuple with the CanUpdateJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateJoinPolicy

`func (o *LoyaltyProgram) HasCanUpdateJoinPolicy() bool`

HasCanUpdateJoinPolicy returns a boolean if a field has been set.

### SetCanUpdateJoinPolicy

`func (o *LoyaltyProgram) SetCanUpdateJoinPolicy(v bool)`

SetCanUpdateJoinPolicy gets a reference to the given bool and assigns it to the CanUpdateJoinPolicy field.

### GetCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgram) GetCanUpdateTierExpirationPolicy() bool`

GetCanUpdateTierExpirationPolicy returns the CanUpdateTierExpirationPolicy field if non-nil, zero value otherwise.

### GetCanUpdateTierExpirationPolicyOk

`func (o *LoyaltyProgram) GetCanUpdateTierExpirationPolicyOk() (bool, bool)`

GetCanUpdateTierExpirationPolicyOk returns a tuple with the CanUpdateTierExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgram) HasCanUpdateTierExpirationPolicy() bool`

HasCanUpdateTierExpirationPolicy returns a boolean if a field has been set.

### SetCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgram) SetCanUpdateTierExpirationPolicy(v bool)`

SetCanUpdateTierExpirationPolicy gets a reference to the given bool and assigns it to the CanUpdateTierExpirationPolicy field.

### GetCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgram) GetCanUpgradeToAdvancedTiers() bool`

GetCanUpgradeToAdvancedTiers returns the CanUpgradeToAdvancedTiers field if non-nil, zero value otherwise.

### GetCanUpgradeToAdvancedTiersOk

`func (o *LoyaltyProgram) GetCanUpgradeToAdvancedTiersOk() (bool, bool)`

GetCanUpgradeToAdvancedTiersOk returns a tuple with the CanUpgradeToAdvancedTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgram) HasCanUpgradeToAdvancedTiers() bool`

HasCanUpgradeToAdvancedTiers returns a boolean if a field has been set.

### SetCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgram) SetCanUpgradeToAdvancedTiers(v bool)`

SetCanUpgradeToAdvancedTiers gets a reference to the given bool and assigns it to the CanUpgradeToAdvancedTiers field.

### GetCanUpdateSubledgers

`func (o *LoyaltyProgram) GetCanUpdateSubledgers() bool`

GetCanUpdateSubledgers returns the CanUpdateSubledgers field if non-nil, zero value otherwise.

### GetCanUpdateSubledgersOk

`func (o *LoyaltyProgram) GetCanUpdateSubledgersOk() (bool, bool)`

GetCanUpdateSubledgersOk returns a tuple with the CanUpdateSubledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateSubledgers

`func (o *LoyaltyProgram) HasCanUpdateSubledgers() bool`

HasCanUpdateSubledgers returns a boolean if a field has been set.

### SetCanUpdateSubledgers

`func (o *LoyaltyProgram) SetCanUpdateSubledgers(v bool)`

SetCanUpdateSubledgers gets a reference to the given bool and assigns it to the CanUpdateSubledgers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


