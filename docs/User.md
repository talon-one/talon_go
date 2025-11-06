# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**Email** | Pointer to **string** | The email address associated with the user profile. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | Name of the user. | 
**State** | Pointer to **string** | State of the user. | 
**InviteToken** | Pointer to **string** | Invitation token of the user.  **Note**: If the user has already accepted their invitation, this is &#x60;null&#x60;.  | 
**IsAdmin** | Pointer to **bool** | Indicates whether the user is an &#x60;admin&#x60;. | [optional] 
**Policy** | Pointer to [**map[string]interface{}**](.md) | Access level of the user. | 
**Roles** | Pointer to **[]int64** | A list of the IDs of the roles assigned to the user. | [optional] 
**AuthMethod** | Pointer to **string** | Authentication method for this user. | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to [**map[string]interface{}**](.md) | Application notifications that the user is subscribed to. | [optional] 
**LastSignedIn** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the user last signed in to Talon.One. | [optional] 
**LastAccessed** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the user&#39;s last activity after signing in to Talon.One. | [optional] 
**LatestFeedTimestamp** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the user was notified for feed. | [optional] 
**AdditionalAttributes** | Pointer to [**map[string]interface{}**](.md) | Additional user attributes, created and used by external identity providers. | [optional] 

## Methods

### NewUser

`func NewUser(id int64, created time.Time, modified time.Time, email string, accountId int64, name string, state string, inviteToken string, policy map[string]interface{}, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *User) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *User) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *User) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *User) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *User) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *User) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAccountId

`func (o *User) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.


### GetInviteToken

`func (o *User) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *User) GetInviteTokenOk() (*string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteToken

`func (o *User) SetInviteToken(v string)`

SetInviteToken sets InviteToken field to given value.


### GetIsAdmin

`func (o *User) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *User) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *User) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *User) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetPolicy

`func (o *User) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *User) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *User) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.


### GetRoles

`func (o *User) GetRoles() []int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []int64)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAuthMethod

`func (o *User) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *User) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *User) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *User) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetApplicationNotificationSubscriptions

`func (o *User) GetApplicationNotificationSubscriptions() map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *User) GetApplicationNotificationSubscriptionsOk() (*map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNotificationSubscriptions

`func (o *User) SetApplicationNotificationSubscriptions(v map[string]interface{})`

SetApplicationNotificationSubscriptions sets ApplicationNotificationSubscriptions field to given value.

### HasApplicationNotificationSubscriptions

`func (o *User) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### GetLastSignedIn

`func (o *User) GetLastSignedIn() time.Time`

GetLastSignedIn returns the LastSignedIn field if non-nil, zero value otherwise.

### GetLastSignedInOk

`func (o *User) GetLastSignedInOk() (*time.Time, bool)`

GetLastSignedInOk returns a tuple with the LastSignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignedIn

`func (o *User) SetLastSignedIn(v time.Time)`

SetLastSignedIn sets LastSignedIn field to given value.

### HasLastSignedIn

`func (o *User) HasLastSignedIn() bool`

HasLastSignedIn returns a boolean if a field has been set.

### GetLastAccessed

`func (o *User) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *User) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *User) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *User) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetLatestFeedTimestamp

`func (o *User) GetLatestFeedTimestamp() time.Time`

GetLatestFeedTimestamp returns the LatestFeedTimestamp field if non-nil, zero value otherwise.

### GetLatestFeedTimestampOk

`func (o *User) GetLatestFeedTimestampOk() (*time.Time, bool)`

GetLatestFeedTimestampOk returns a tuple with the LatestFeedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFeedTimestamp

`func (o *User) SetLatestFeedTimestamp(v time.Time)`

SetLatestFeedTimestamp sets LatestFeedTimestamp field to given value.

### HasLatestFeedTimestamp

`func (o *User) HasLatestFeedTimestamp() bool`

HasLatestFeedTimestamp returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *User) GetAdditionalAttributes() map[string]interface{}`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *User) GetAdditionalAttributesOk() (*map[string]interface{}, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *User) SetAdditionalAttributes(v map[string]interface{})`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *User) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


