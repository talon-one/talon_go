# CampaignReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the campaign that references this achievement. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application the campaign belongs to. | 

## Methods

### NewCampaignReference

`func NewCampaignReference(id int64, applicationId int64, ) *CampaignReference`

NewCampaignReference instantiates a new CampaignReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignReferenceWithDefaults

`func NewCampaignReferenceWithDefaults() *CampaignReference`

NewCampaignReferenceWithDefaults instantiates a new CampaignReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignReference) SetId(v int64)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *CampaignReference) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignReference) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignReference) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


