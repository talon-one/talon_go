# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**Email** | Pointer to **string** | The email address associated with your account. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**InviteToken** | Pointer to **string** | Invite token, empty if the user as already accepted their invite. | 
**State** | Pointer to **string** | Current user state. | 
**Name** | Pointer to **string** | Full name | 
**Policy** | Pointer to [**map[string]interface{}**](.md) | User ACL Policy | 
**LatestFeedTimestamp** | Pointer to [**time.Time**](time.Time.md) | Latest timestamp the user has been notified for feed. | [optional] 
**Roles** | Pointer to **[]int32** | Contains a list of all roles the user is a member of. | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**AuthMethod** | Pointer to **string** | The Authentication method for this user. | [optional] 

## Methods

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *User) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *User) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *User) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *User) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *User) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *User) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *User) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *User) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetAccountId

`func (o *User) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *User) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *User) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetInviteToken

`func (o *User) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *User) GetInviteTokenOk() (string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteToken

`func (o *User) HasInviteToken() bool`

HasInviteToken returns a boolean if a field has been set.

### SetInviteToken

`func (o *User) SetInviteToken(v string)`

SetInviteToken gets a reference to the given string and assigns it to the InviteToken field.

### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *User) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *User) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPolicy

`func (o *User) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *User) GetPolicyOk() (map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *User) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *User) SetPolicy(v map[string]interface{})`

SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.

### GetLatestFeedTimestamp

`func (o *User) GetLatestFeedTimestamp() time.Time`

GetLatestFeedTimestamp returns the LatestFeedTimestamp field if non-nil, zero value otherwise.

### GetLatestFeedTimestampOk

`func (o *User) GetLatestFeedTimestampOk() (time.Time, bool)`

GetLatestFeedTimestampOk returns a tuple with the LatestFeedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLatestFeedTimestamp

`func (o *User) HasLatestFeedTimestamp() bool`

HasLatestFeedTimestamp returns a boolean if a field has been set.

### SetLatestFeedTimestamp

`func (o *User) SetLatestFeedTimestamp(v time.Time)`

SetLatestFeedTimestamp gets a reference to the given time.Time and assigns it to the LatestFeedTimestamp field.

### GetRoles

`func (o *User) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() ([]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *User) SetRoles(v []int32)`

SetRoles gets a reference to the given []int32 and assigns it to the Roles field.

### GetApplicationNotificationSubscriptions

`func (o *User) GetApplicationNotificationSubscriptions() map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *User) GetApplicationNotificationSubscriptionsOk() (map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationNotificationSubscriptions

`func (o *User) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### SetApplicationNotificationSubscriptions

`func (o *User) SetApplicationNotificationSubscriptions(v map[string]interface{})`

SetApplicationNotificationSubscriptions gets a reference to the given map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.

### GetAuthMethod

`func (o *User) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *User) GetAuthMethodOk() (string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthMethod

`func (o *User) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethod

`func (o *User) SetAuthMethod(v string)`

SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


