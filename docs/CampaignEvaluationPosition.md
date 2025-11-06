# CampaignEvaluationPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int64** | The ID of the campaign evaluation group the campaign belongs to. | 
**GroupName** | Pointer to **string** | The name of the campaign evaluation group the campaign belongs to. | 
**Position** | Pointer to **int64** | The position of the campaign node in its parent group. | 

## Methods

### NewCampaignEvaluationPosition

`func NewCampaignEvaluationPosition(groupId int64, groupName string, position int64, ) *CampaignEvaluationPosition`

NewCampaignEvaluationPosition instantiates a new CampaignEvaluationPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEvaluationPositionWithDefaults

`func NewCampaignEvaluationPositionWithDefaults() *CampaignEvaluationPosition`

NewCampaignEvaluationPositionWithDefaults instantiates a new CampaignEvaluationPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CampaignEvaluationPosition) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CampaignEvaluationPosition) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CampaignEvaluationPosition) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *CampaignEvaluationPosition) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CampaignEvaluationPosition) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CampaignEvaluationPosition) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetPosition

`func (o *CampaignEvaluationPosition) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CampaignEvaluationPosition) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CampaignEvaluationPosition) SetPosition(v int64)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


