# SupportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identifier of the support request. | 
**ApplicationId** | Pointer to **int64** | Identifier of the Application connected to the loyalty program or the campaign. It is displayed in your Talon.One deployment URL. | 
**CampaignId** | Pointer to **int64** | Identifier of the campaign where the coupon or gift card is created. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | Identifier of the loyalty program where the points are added or deducted. | [optional] 
**SubledgerId** | Pointer to **int64** | Identifier of the subledger the points are added to or deducted from. If there is no existing subledger with this ID, the subledger is created automatically. | [optional] 
**CreatedByUser** | Pointer to **string** | Email address of the customer support agent who created the support request. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the request was made. | 
**CustomerProfileId** | Pointer to **string** | Integration ID of the customer profile linked to the support request. | 
**RequestType** | Pointer to **string** | Type of reward requested, including gift cards, personal coupons, and loyalty point additions or deductions. | 
**RequestValue** | Pointer to **float32** | Requested monetary balance of the gift card or the number of loyalty points to be added or deducted. | [optional] 
**RequestNote** | Pointer to **string** | Notes attached to the support request. | 
**RequestStatus** | Pointer to **string** | Current status of the support request. | 
**ProcessedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the request was approved or rejected. | [optional] 
**ProcessingNote** | Pointer to **string** | Notes attached by the admin when rejecting or approving a request. | [optional] 
**ProcessedByUser** | Pointer to **string** | Email address of the admin who approved or rejected the support request. | [optional] 
**CouponCode** | Pointer to **string** | Coupon code associated with the approved support request. | [optional] 

## Methods

### NewSupportRequest

`func NewSupportRequest(id int64, applicationId int64, createdByUser string, createdAt time.Time, customerProfileId string, requestType string, requestNote string, requestStatus string, ) *SupportRequest`

NewSupportRequest instantiates a new SupportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestWithDefaults

`func NewSupportRequestWithDefaults() *SupportRequest`

NewSupportRequestWithDefaults instantiates a new SupportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *SupportRequest) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SupportRequest) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SupportRequest) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *SupportRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *SupportRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *SupportRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *SupportRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *SupportRequest) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *SupportRequest) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *SupportRequest) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *SupportRequest) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetSubledgerId

`func (o *SupportRequest) GetSubledgerId() int64`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *SupportRequest) GetSubledgerIdOk() (*int64, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *SupportRequest) SetSubledgerId(v int64)`

SetSubledgerId sets SubledgerId field to given value.

### HasSubledgerId

`func (o *SupportRequest) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *SupportRequest) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *SupportRequest) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *SupportRequest) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetCreatedAt

`func (o *SupportRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SupportRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SupportRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomerProfileId

`func (o *SupportRequest) GetCustomerProfileId() string`

GetCustomerProfileId returns the CustomerProfileId field if non-nil, zero value otherwise.

### GetCustomerProfileIdOk

`func (o *SupportRequest) GetCustomerProfileIdOk() (*string, bool)`

GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileId

`func (o *SupportRequest) SetCustomerProfileId(v string)`

SetCustomerProfileId sets CustomerProfileId field to given value.


### GetRequestType

`func (o *SupportRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SupportRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SupportRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetRequestValue

`func (o *SupportRequest) GetRequestValue() float32`

GetRequestValue returns the RequestValue field if non-nil, zero value otherwise.

### GetRequestValueOk

`func (o *SupportRequest) GetRequestValueOk() (*float32, bool)`

GetRequestValueOk returns a tuple with the RequestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestValue

`func (o *SupportRequest) SetRequestValue(v float32)`

SetRequestValue sets RequestValue field to given value.

### HasRequestValue

`func (o *SupportRequest) HasRequestValue() bool`

HasRequestValue returns a boolean if a field has been set.

### GetRequestNote

`func (o *SupportRequest) GetRequestNote() string`

GetRequestNote returns the RequestNote field if non-nil, zero value otherwise.

### GetRequestNoteOk

`func (o *SupportRequest) GetRequestNoteOk() (*string, bool)`

GetRequestNoteOk returns a tuple with the RequestNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestNote

`func (o *SupportRequest) SetRequestNote(v string)`

SetRequestNote sets RequestNote field to given value.


### GetRequestStatus

`func (o *SupportRequest) GetRequestStatus() string`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *SupportRequest) GetRequestStatusOk() (*string, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *SupportRequest) SetRequestStatus(v string)`

SetRequestStatus sets RequestStatus field to given value.


### GetProcessedAt

`func (o *SupportRequest) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *SupportRequest) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *SupportRequest) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *SupportRequest) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetProcessingNote

`func (o *SupportRequest) GetProcessingNote() string`

GetProcessingNote returns the ProcessingNote field if non-nil, zero value otherwise.

### GetProcessingNoteOk

`func (o *SupportRequest) GetProcessingNoteOk() (*string, bool)`

GetProcessingNoteOk returns a tuple with the ProcessingNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingNote

`func (o *SupportRequest) SetProcessingNote(v string)`

SetProcessingNote sets ProcessingNote field to given value.

### HasProcessingNote

`func (o *SupportRequest) HasProcessingNote() bool`

HasProcessingNote returns a boolean if a field has been set.

### GetProcessedByUser

`func (o *SupportRequest) GetProcessedByUser() string`

GetProcessedByUser returns the ProcessedByUser field if non-nil, zero value otherwise.

### GetProcessedByUserOk

`func (o *SupportRequest) GetProcessedByUserOk() (*string, bool)`

GetProcessedByUserOk returns a tuple with the ProcessedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedByUser

`func (o *SupportRequest) SetProcessedByUser(v string)`

SetProcessedByUser sets ProcessedByUser field to given value.

### HasProcessedByUser

`func (o *SupportRequest) HasProcessedByUser() bool`

HasProcessedByUser returns a boolean if a field has been set.

### GetCouponCode

`func (o *SupportRequest) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *SupportRequest) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *SupportRequest) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *SupportRequest) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


