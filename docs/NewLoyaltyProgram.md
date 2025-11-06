# NewLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display title for the Loyalty Program. | 
**Description** | Pointer to **string** | Description of our Loyalty Program. | [optional] 
**SubscribedApplications** | Pointer to **[]int64** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] 
**DefaultValidity** | Pointer to **string** | The default duration after which new loyalty points should expire. Can be &#39;unlimited&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | 
**DefaultPending** | Pointer to **string** | The default duration of the pending time after which points should be valid. Accepted values: &#39;immediate&#39;, &#39;on_action&#39; or a specific time. The time format is a number followed by one letter indicating the time unit, like &#39;30s&#39;, &#39;40m&#39;, &#39;1h&#39;, &#39;5D&#39;, &#39;7W&#39;, or 10M&#39;. These rounding suffixes are also supported: - &#39;_D&#39; for rounding down. Can be used as a suffix after &#39;D&#39;, and signifies the start of the day. - &#39;_U&#39; for rounding up. Can be used as a suffix after &#39;D&#39;, &#39;W&#39;, and &#39;M&#39;, and signifies the end of the day, week, and month.  | 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program. | 
**UsersPerCardLimit** | Pointer to **int64** | The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when &#x60;cardBased&#x60; is &#x60;true&#x60;.  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type. | 
**ProgramJoinPolicy** | Pointer to **string** | The policy that defines when the customer joins the loyalty program.   - &#x60;not_join&#x60;: The customer does not join the loyalty program but can still earn and spend loyalty points.       **Note**: The customer does not have a program join date.   - &#x60;points_activated&#x60;: The customer joins the loyalty program only when their earned loyalty points become active for the first time.   - &#x60;points_earned&#x60;: The customer joins the loyalty program when they earn loyalty points for the first time.  | [optional] 
**TiersExpirationPolicy** | Pointer to **string** | The policy that defines how tier expiration, used to reevaluate the customer&#39;s current tier, is determined.  - &#x60;tier_start_date&#x60;: The tier expiration is relative to when the customer joined the current tier.  - &#x60;program_join_date&#x60;: The tier expiration is relative to when the customer joined the loyalty program.  - &#x60;customer_attribute&#x60;: The tier expiration is determined by a custom customer attribute.  - &#x60;absolute_expiration&#x60;: The tier is reevaluated at the start of each tier cycle. For this policy, it is required to provide a &#x60;tierCycleStartDate&#x60;.  | [optional] 
**TierCycleStartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the tier cycle starts for all customers in the loyalty program.  **Note**: This is only required when the tier expiration policy is set to &#x60;absolute_expiration&#x60;.  | [optional] 
**TiersExpireIn** | Pointer to **string** | The amount of time after which the tier expires and is reevaluated.  The time format is an **integer** followed by one letter indicating the time unit. Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | [optional] 
**TiersDowngradePolicy** | Pointer to **string** | The policy that defines how customer tiers are downgraded in the loyalty program after tier reevaluation.  - &#x60;one_down&#x60;: If the customer doesn&#39;t have enough points to stay in the current tier, they are downgraded by one tier.  - &#x60;balance_based&#x60;: The customer&#39;s tier is reevaluated based on the amount of active points they have at the moment.  | [optional] 
**CardCodeSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReturnPolicy** | Pointer to **string** | The policy that defines the rollback of points in case of a partially returned, cancelled, or reopened [customer session](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions). - &#x60;only_pending&#x60;: Only pending points can be rolled back. - &#x60;within_balance&#x60;: Available active points can be rolled back if there aren&#39;t enough pending points. The active balance of the customer cannot be negative. - &#x60;unlimited&#x60;: Allows negative balance without any limit.  | [optional] 
**Name** | Pointer to **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | Pointer to **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]

## Methods

### NewNewLoyaltyProgram

`func NewNewLoyaltyProgram(title string, defaultValidity string, defaultPending string, allowSubledger bool, sandbox bool, name string, timezone string, cardBased bool, ) *NewLoyaltyProgram`

NewNewLoyaltyProgram instantiates a new NewLoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewLoyaltyProgramWithDefaults

`func NewNewLoyaltyProgramWithDefaults() *NewLoyaltyProgram`

NewNewLoyaltyProgramWithDefaults instantiates a new NewLoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewLoyaltyProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewLoyaltyProgram) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewLoyaltyProgram) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewLoyaltyProgram) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplications

`func (o *NewLoyaltyProgram) GetSubscribedApplications() []int64`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *NewLoyaltyProgram) GetSubscribedApplicationsOk() (*[]int64, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *NewLoyaltyProgram) SetSubscribedApplications(v []int64)`

SetSubscribedApplications sets SubscribedApplications field to given value.

### HasSubscribedApplications

`func (o *NewLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### GetDefaultValidity

`func (o *NewLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *NewLoyaltyProgram) GetDefaultValidityOk() (*string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValidity

`func (o *NewLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity sets DefaultValidity field to given value.


### GetDefaultPending

`func (o *NewLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *NewLoyaltyProgram) GetDefaultPendingOk() (*string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPending

`func (o *NewLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending sets DefaultPending field to given value.


### GetAllowSubledger

`func (o *NewLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *NewLoyaltyProgram) GetAllowSubledgerOk() (*bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubledger

`func (o *NewLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger sets AllowSubledger field to given value.


### GetUsersPerCardLimit

`func (o *NewLoyaltyProgram) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *NewLoyaltyProgram) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *NewLoyaltyProgram) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.

### HasUsersPerCardLimit

`func (o *NewLoyaltyProgram) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### GetSandbox

`func (o *NewLoyaltyProgram) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewLoyaltyProgram) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewLoyaltyProgram) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetProgramJoinPolicy

`func (o *NewLoyaltyProgram) GetProgramJoinPolicy() string`

GetProgramJoinPolicy returns the ProgramJoinPolicy field if non-nil, zero value otherwise.

### GetProgramJoinPolicyOk

`func (o *NewLoyaltyProgram) GetProgramJoinPolicyOk() (*string, bool)`

GetProgramJoinPolicyOk returns a tuple with the ProgramJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramJoinPolicy

`func (o *NewLoyaltyProgram) SetProgramJoinPolicy(v string)`

SetProgramJoinPolicy sets ProgramJoinPolicy field to given value.

### HasProgramJoinPolicy

`func (o *NewLoyaltyProgram) HasProgramJoinPolicy() bool`

HasProgramJoinPolicy returns a boolean if a field has been set.

### GetTiersExpirationPolicy

`func (o *NewLoyaltyProgram) GetTiersExpirationPolicy() string`

GetTiersExpirationPolicy returns the TiersExpirationPolicy field if non-nil, zero value otherwise.

### GetTiersExpirationPolicyOk

`func (o *NewLoyaltyProgram) GetTiersExpirationPolicyOk() (*string, bool)`

GetTiersExpirationPolicyOk returns a tuple with the TiersExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersExpirationPolicy

`func (o *NewLoyaltyProgram) SetTiersExpirationPolicy(v string)`

SetTiersExpirationPolicy sets TiersExpirationPolicy field to given value.

### HasTiersExpirationPolicy

`func (o *NewLoyaltyProgram) HasTiersExpirationPolicy() bool`

HasTiersExpirationPolicy returns a boolean if a field has been set.

### GetTierCycleStartDate

`func (o *NewLoyaltyProgram) GetTierCycleStartDate() time.Time`

GetTierCycleStartDate returns the TierCycleStartDate field if non-nil, zero value otherwise.

### GetTierCycleStartDateOk

`func (o *NewLoyaltyProgram) GetTierCycleStartDateOk() (*time.Time, bool)`

GetTierCycleStartDateOk returns a tuple with the TierCycleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierCycleStartDate

`func (o *NewLoyaltyProgram) SetTierCycleStartDate(v time.Time)`

SetTierCycleStartDate sets TierCycleStartDate field to given value.

### HasTierCycleStartDate

`func (o *NewLoyaltyProgram) HasTierCycleStartDate() bool`

HasTierCycleStartDate returns a boolean if a field has been set.

### GetTiersExpireIn

`func (o *NewLoyaltyProgram) GetTiersExpireIn() string`

GetTiersExpireIn returns the TiersExpireIn field if non-nil, zero value otherwise.

### GetTiersExpireInOk

`func (o *NewLoyaltyProgram) GetTiersExpireInOk() (*string, bool)`

GetTiersExpireInOk returns a tuple with the TiersExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersExpireIn

`func (o *NewLoyaltyProgram) SetTiersExpireIn(v string)`

SetTiersExpireIn sets TiersExpireIn field to given value.

### HasTiersExpireIn

`func (o *NewLoyaltyProgram) HasTiersExpireIn() bool`

HasTiersExpireIn returns a boolean if a field has been set.

### GetTiersDowngradePolicy

`func (o *NewLoyaltyProgram) GetTiersDowngradePolicy() string`

GetTiersDowngradePolicy returns the TiersDowngradePolicy field if non-nil, zero value otherwise.

### GetTiersDowngradePolicyOk

`func (o *NewLoyaltyProgram) GetTiersDowngradePolicyOk() (*string, bool)`

GetTiersDowngradePolicyOk returns a tuple with the TiersDowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersDowngradePolicy

`func (o *NewLoyaltyProgram) SetTiersDowngradePolicy(v string)`

SetTiersDowngradePolicy sets TiersDowngradePolicy field to given value.

### HasTiersDowngradePolicy

`func (o *NewLoyaltyProgram) HasTiersDowngradePolicy() bool`

HasTiersDowngradePolicy returns a boolean if a field has been set.

### GetCardCodeSettings

`func (o *NewLoyaltyProgram) GetCardCodeSettings() CodeGeneratorSettings`

GetCardCodeSettings returns the CardCodeSettings field if non-nil, zero value otherwise.

### GetCardCodeSettingsOk

`func (o *NewLoyaltyProgram) GetCardCodeSettingsOk() (*CodeGeneratorSettings, bool)`

GetCardCodeSettingsOk returns a tuple with the CardCodeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCodeSettings

`func (o *NewLoyaltyProgram) SetCardCodeSettings(v CodeGeneratorSettings)`

SetCardCodeSettings sets CardCodeSettings field to given value.

### HasCardCodeSettings

`func (o *NewLoyaltyProgram) HasCardCodeSettings() bool`

HasCardCodeSettings returns a boolean if a field has been set.

### GetReturnPolicy

`func (o *NewLoyaltyProgram) GetReturnPolicy() string`

GetReturnPolicy returns the ReturnPolicy field if non-nil, zero value otherwise.

### GetReturnPolicyOk

`func (o *NewLoyaltyProgram) GetReturnPolicyOk() (*string, bool)`

GetReturnPolicyOk returns a tuple with the ReturnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicy

`func (o *NewLoyaltyProgram) SetReturnPolicy(v string)`

SetReturnPolicy sets ReturnPolicy field to given value.

### HasReturnPolicy

`func (o *NewLoyaltyProgram) HasReturnPolicy() bool`

HasReturnPolicy returns a boolean if a field has been set.

### GetName

`func (o *NewLoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewLoyaltyProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewLoyaltyProgram) SetName(v string)`

SetName sets Name field to given value.


### GetTiers

`func (o *NewLoyaltyProgram) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *NewLoyaltyProgram) GetTiersOk() (*[]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *NewLoyaltyProgram) SetTiers(v []NewLoyaltyTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *NewLoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTimezone

`func (o *NewLoyaltyProgram) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewLoyaltyProgram) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NewLoyaltyProgram) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCardBased

`func (o *NewLoyaltyProgram) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *NewLoyaltyProgram) GetCardBasedOk() (*bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBased

`func (o *NewLoyaltyProgram) SetCardBased(v bool)`

SetCardBased sets CardBased field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


