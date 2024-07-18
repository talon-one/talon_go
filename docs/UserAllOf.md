# UserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. | 
**State** | Pointer to **string** | State of the user. | 
**InviteToken** | Pointer to **string** | Invitation token of the user.  **Note**: If the user has already accepted their invitation, this is &#x60;null&#x60;.  | 
**IsAdmin** | Pointer to **bool** | Indicates whether the user is an &#x60;admin&#x60;. | [optional] 
**Policy** | Pointer to [**map[string]interface{}**](.md) | Access level of the user. | 
**Roles** | Pointer to **[]int32** | A list of the IDs of the roles assigned to the user. | [optional] 
**AuthMethod** | Pointer to **string** | Authentication method for this user. | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Application notifications that the user is subscribed to. | [optional] 
**LastSignedIn** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the user last signed in to Talon.One. | [optional] 
**LastAccessed** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the user&#39;s last activity after signing in to Talon.One. | [optional] 
**LatestFeedTimestamp** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the user was notified for feed. | [optional] 
**AdditionalAttributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Additional user attributes, created and used by external identity providers. | [optional] 

## Methods

### GetName

`func (o *UserAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UserAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UserAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetState

`func (o *UserAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserAllOf) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *UserAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *UserAllOf) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetInviteToken

`func (o *UserAllOf) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *UserAllOf) GetInviteTokenOk() (string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteToken

`func (o *UserAllOf) HasInviteToken() bool`

HasInviteToken returns a boolean if a field has been set.

### SetInviteToken

`func (o *UserAllOf) SetInviteToken(v string)`

SetInviteToken gets a reference to the given string and assigns it to the InviteToken field.

### GetIsAdmin

`func (o *UserAllOf) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserAllOf) GetIsAdminOk() (bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdmin

`func (o *UserAllOf) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdmin

`func (o *UserAllOf) SetIsAdmin(v bool)`

SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.

### GetPolicy

`func (o *UserAllOf) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UserAllOf) GetPolicyOk() (map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *UserAllOf) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *UserAllOf) SetPolicy(v map[string]interface{})`

SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.

### GetRoles

`func (o *UserAllOf) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAllOf) GetRolesOk() ([]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *UserAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *UserAllOf) SetRoles(v []int32)`

SetRoles gets a reference to the given []int32 and assigns it to the Roles field.

### GetAuthMethod

`func (o *UserAllOf) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *UserAllOf) GetAuthMethodOk() (string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthMethod

`func (o *UserAllOf) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethod

`func (o *UserAllOf) SetAuthMethod(v string)`

SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.

### GetApplicationNotificationSubscriptions

`func (o *UserAllOf) GetApplicationNotificationSubscriptions() map[string]map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *UserAllOf) GetApplicationNotificationSubscriptionsOk() (map[string]map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationNotificationSubscriptions

`func (o *UserAllOf) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### SetApplicationNotificationSubscriptions

`func (o *UserAllOf) SetApplicationNotificationSubscriptions(v map[string]map[string]interface{})`

SetApplicationNotificationSubscriptions gets a reference to the given map[string]map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.

### GetLastSignedIn

`func (o *UserAllOf) GetLastSignedIn() time.Time`

GetLastSignedIn returns the LastSignedIn field if non-nil, zero value otherwise.

### GetLastSignedInOk

`func (o *UserAllOf) GetLastSignedInOk() (time.Time, bool)`

GetLastSignedInOk returns a tuple with the LastSignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSignedIn

`func (o *UserAllOf) HasLastSignedIn() bool`

HasLastSignedIn returns a boolean if a field has been set.

### SetLastSignedIn

`func (o *UserAllOf) SetLastSignedIn(v time.Time)`

SetLastSignedIn gets a reference to the given time.Time and assigns it to the LastSignedIn field.

### GetLastAccessed

`func (o *UserAllOf) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *UserAllOf) GetLastAccessedOk() (time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAccessed

`func (o *UserAllOf) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### SetLastAccessed

`func (o *UserAllOf) SetLastAccessed(v time.Time)`

SetLastAccessed gets a reference to the given time.Time and assigns it to the LastAccessed field.

### GetLatestFeedTimestamp

`func (o *UserAllOf) GetLatestFeedTimestamp() time.Time`

GetLatestFeedTimestamp returns the LatestFeedTimestamp field if non-nil, zero value otherwise.

### GetLatestFeedTimestampOk

`func (o *UserAllOf) GetLatestFeedTimestampOk() (time.Time, bool)`

GetLatestFeedTimestampOk returns a tuple with the LatestFeedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLatestFeedTimestamp

`func (o *UserAllOf) HasLatestFeedTimestamp() bool`

HasLatestFeedTimestamp returns a boolean if a field has been set.

### SetLatestFeedTimestamp

`func (o *UserAllOf) SetLatestFeedTimestamp(v time.Time)`

SetLatestFeedTimestamp gets a reference to the given time.Time and assigns it to the LatestFeedTimestamp field.

### GetAdditionalAttributes

`func (o *UserAllOf) GetAdditionalAttributes() map[string]map[string]interface{}`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *UserAllOf) GetAdditionalAttributesOk() (map[string]map[string]interface{}, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalAttributes

`func (o *UserAllOf) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributes

`func (o *UserAllOf) SetAdditionalAttributes(v map[string]map[string]interface{})`

SetAdditionalAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the AdditionalAttributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


