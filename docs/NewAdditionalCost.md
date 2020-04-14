# NewAdditionalCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The additional cost name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;shipping\&quot;&#x60; then you would set the shipping additional cost by including an &#x60;additionalCosts.shipping&#x60; property in your request payload. | 
**Title** | Pointer to **string** | The human-readable name for the additional cost that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Description** | Pointer to **string** | A description of this additional cost. | 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that are subscribed to this additional cost | [optional] 

## Methods

### GetName

`func (o *NewAdditionalCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAdditionalCost) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewAdditionalCost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewAdditionalCost) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *NewAdditionalCost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAdditionalCost) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewAdditionalCost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewAdditionalCost) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *NewAdditionalCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewAdditionalCost) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewAdditionalCost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewAdditionalCost) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplicationsIds

`func (o *NewAdditionalCost) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewAdditionalCost) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *NewAdditionalCost) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *NewAdditionalCost) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


