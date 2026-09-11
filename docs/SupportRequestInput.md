# SupportRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | Identifier of the Application connected to the loyalty program or the campaign. It is displayed in your Talon.One deployment URL. | 
**CampaignId** | Pointer to **int64** | Identifier of the campaign where the coupon or gift card is created. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. | [optional] 
**SubledgerId** | Pointer to **int64** | Identifier of the subledger the points are added to or deducted from. If there is no existing subledger with this ID, the subledger is created automatically. | [optional] 
**CustomerProfileId** | Pointer to **string** | Integration ID of the customer profile linked to the support request. | 
**RequestType** | Pointer to **string** | Type of reward requested, including gift cards, personal coupons, and loyalty point additions or deductions. | 
**RequestValue** | Pointer to **float32** | Requested monetary balance of the gift card or the number of loyalty points to be added or deducted. | [optional] 
**RequestNote** | Pointer to **string** | Notes attached to the support request. | 

## Methods

### NewSupportRequestInput

`func NewSupportRequestInput(applicationId int64, customerProfileId string, requestType string, requestNote string, ) *SupportRequestInput`

NewSupportRequestInput instantiates a new SupportRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestInputWithDefaults

`func NewSupportRequestInputWithDefaults() *SupportRequestInput`

NewSupportRequestInputWithDefaults instantiates a new SupportRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *SupportRequestInput) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SupportRequestInput) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SupportRequestInput) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *SupportRequestInput) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *SupportRequestInput) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *SupportRequestInput) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *SupportRequestInput) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *SupportRequestInput) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *SupportRequestInput) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *SupportRequestInput) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *SupportRequestInput) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetSubledgerId

`func (o *SupportRequestInput) GetSubledgerId() int64`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *SupportRequestInput) GetSubledgerIdOk() (*int64, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *SupportRequestInput) SetSubledgerId(v int64)`

SetSubledgerId sets SubledgerId field to given value.

### HasSubledgerId

`func (o *SupportRequestInput) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### GetCustomerProfileId

`func (o *SupportRequestInput) GetCustomerProfileId() string`

GetCustomerProfileId returns the CustomerProfileId field if non-nil, zero value otherwise.

### GetCustomerProfileIdOk

`func (o *SupportRequestInput) GetCustomerProfileIdOk() (*string, bool)`

GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileId

`func (o *SupportRequestInput) SetCustomerProfileId(v string)`

SetCustomerProfileId sets CustomerProfileId field to given value.


### GetRequestType

`func (o *SupportRequestInput) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SupportRequestInput) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SupportRequestInput) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetRequestValue

`func (o *SupportRequestInput) GetRequestValue() float32`

GetRequestValue returns the RequestValue field if non-nil, zero value otherwise.

### GetRequestValueOk

`func (o *SupportRequestInput) GetRequestValueOk() (*float32, bool)`

GetRequestValueOk returns a tuple with the RequestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestValue

`func (o *SupportRequestInput) SetRequestValue(v float32)`

SetRequestValue sets RequestValue field to given value.

### HasRequestValue

`func (o *SupportRequestInput) HasRequestValue() bool`

HasRequestValue returns a boolean if a field has been set.

### GetRequestNote

`func (o *SupportRequestInput) GetRequestNote() string`

GetRequestNote returns the RequestNote field if non-nil, zero value otherwise.

### GetRequestNoteOk

`func (o *SupportRequestInput) GetRequestNoteOk() (*string, bool)`

GetRequestNoteOk returns a tuple with the RequestNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestNote

`func (o *SupportRequestInput) SetRequestNote(v string)`

SetRequestNote sets RequestNote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


