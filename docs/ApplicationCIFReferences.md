# ApplicationCIFReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCartItemFilterId** | Pointer to **int64** | The ID of the Application Cart Item Filter that is referenced by a campaign. | [optional] 
**Campaigns** | Pointer to [**[]CampaignDetail**](CampaignDetail.md) | Campaigns that reference a speciifc Application Cart Item Filter. | [optional] 

## Methods

### NewApplicationCIFReferences

`func NewApplicationCIFReferences() *ApplicationCIFReferences`

NewApplicationCIFReferences instantiates a new ApplicationCIFReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCIFReferencesWithDefaults

`func NewApplicationCIFReferencesWithDefaults() *ApplicationCIFReferences`

NewApplicationCIFReferencesWithDefaults instantiates a new ApplicationCIFReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCartItemFilterId

`func (o *ApplicationCIFReferences) GetApplicationCartItemFilterId() int64`

GetApplicationCartItemFilterId returns the ApplicationCartItemFilterId field if non-nil, zero value otherwise.

### GetApplicationCartItemFilterIdOk

`func (o *ApplicationCIFReferences) GetApplicationCartItemFilterIdOk() (*int64, bool)`

GetApplicationCartItemFilterIdOk returns a tuple with the ApplicationCartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCartItemFilterId

`func (o *ApplicationCIFReferences) SetApplicationCartItemFilterId(v int64)`

SetApplicationCartItemFilterId sets ApplicationCartItemFilterId field to given value.

### HasApplicationCartItemFilterId

`func (o *ApplicationCIFReferences) HasApplicationCartItemFilterId() bool`

HasApplicationCartItemFilterId returns a boolean if a field has been set.

### GetCampaigns

`func (o *ApplicationCIFReferences) GetCampaigns() []CampaignDetail`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *ApplicationCIFReferences) GetCampaignsOk() (*[]CampaignDetail, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *ApplicationCIFReferences) SetCampaigns(v []CampaignDetail)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *ApplicationCIFReferences) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


