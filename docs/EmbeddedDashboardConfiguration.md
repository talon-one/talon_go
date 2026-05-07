# EmbeddedDashboardConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | Pointer to **string** | The ID of the workspace that contains dashboards. | 
**DashboardId** | Pointer to **string** | The ID of the dashboard that contains metrics. | 

## Methods

### NewEmbeddedDashboardConfiguration

`func NewEmbeddedDashboardConfiguration(workspaceId string, dashboardId string, ) *EmbeddedDashboardConfiguration`

NewEmbeddedDashboardConfiguration instantiates a new EmbeddedDashboardConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedDashboardConfigurationWithDefaults

`func NewEmbeddedDashboardConfigurationWithDefaults() *EmbeddedDashboardConfiguration`

NewEmbeddedDashboardConfigurationWithDefaults instantiates a new EmbeddedDashboardConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *EmbeddedDashboardConfiguration) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *EmbeddedDashboardConfiguration) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *EmbeddedDashboardConfiguration) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetDashboardId

`func (o *EmbeddedDashboardConfiguration) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *EmbeddedDashboardConfiguration) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *EmbeddedDashboardConfiguration) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


