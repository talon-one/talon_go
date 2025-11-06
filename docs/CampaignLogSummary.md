# CampaignLogSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user that performed the change. | 
**Email** | Pointer to **string** | E-mail of the user that performed the change. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the change was performed. | 
**Action** | Pointer to **string** | Action performed by the user. | 
**Summary** | Pointer to **string** | AI-generated summary of the action performed. | 

## Methods

### NewCampaignLogSummary

`func NewCampaignLogSummary(name string, email string, created time.Time, action string, summary string, ) *CampaignLogSummary`

NewCampaignLogSummary instantiates a new CampaignLogSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignLogSummaryWithDefaults

`func NewCampaignLogSummaryWithDefaults() *CampaignLogSummary`

NewCampaignLogSummaryWithDefaults instantiates a new CampaignLogSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignLogSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignLogSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignLogSummary) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CampaignLogSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CampaignLogSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CampaignLogSummary) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCreated

`func (o *CampaignLogSummary) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignLogSummary) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignLogSummary) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAction

`func (o *CampaignLogSummary) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignLogSummary) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CampaignLogSummary) SetAction(v string)`

SetAction sets Action field to given value.


### GetSummary

`func (o *CampaignLogSummary) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CampaignLogSummary) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CampaignLogSummary) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


