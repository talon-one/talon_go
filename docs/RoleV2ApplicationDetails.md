# RoleV2ApplicationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Name of the Application-related permission set for the given Application. | [optional] 
**Campaign** | Pointer to **string** | Name of the campaign-related permission set for the given Application. | [optional] 
**DraftCampaign** | Pointer to **string** | Name of the draft campaign-related permission set for the given Application. | [optional] 
**Tools** | Pointer to **string** | Name of the tools-related permission set. | [optional] 

## Methods

### NewRoleV2ApplicationDetails

`func NewRoleV2ApplicationDetails() *RoleV2ApplicationDetails`

NewRoleV2ApplicationDetails instantiates a new RoleV2ApplicationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2ApplicationDetailsWithDefaults

`func NewRoleV2ApplicationDetailsWithDefaults() *RoleV2ApplicationDetails`

NewRoleV2ApplicationDetailsWithDefaults instantiates a new RoleV2ApplicationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *RoleV2ApplicationDetails) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RoleV2ApplicationDetails) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RoleV2ApplicationDetails) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *RoleV2ApplicationDetails) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCampaign

`func (o *RoleV2ApplicationDetails) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *RoleV2ApplicationDetails) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *RoleV2ApplicationDetails) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *RoleV2ApplicationDetails) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetDraftCampaign

`func (o *RoleV2ApplicationDetails) GetDraftCampaign() string`

GetDraftCampaign returns the DraftCampaign field if non-nil, zero value otherwise.

### GetDraftCampaignOk

`func (o *RoleV2ApplicationDetails) GetDraftCampaignOk() (*string, bool)`

GetDraftCampaignOk returns a tuple with the DraftCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftCampaign

`func (o *RoleV2ApplicationDetails) SetDraftCampaign(v string)`

SetDraftCampaign sets DraftCampaign field to given value.

### HasDraftCampaign

`func (o *RoleV2ApplicationDetails) HasDraftCampaign() bool`

HasDraftCampaign returns a boolean if a field has been set.

### GetTools

`func (o *RoleV2ApplicationDetails) GetTools() string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *RoleV2ApplicationDetails) GetToolsOk() (*string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *RoleV2ApplicationDetails) SetTools(v string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *RoleV2ApplicationDetails) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


