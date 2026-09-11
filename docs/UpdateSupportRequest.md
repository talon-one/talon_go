# UpdateSupportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestStatus** | Pointer to **string** | Current status of the support request. | 
**ProcessingNote** | Pointer to **string** | Notes attached by the admin when rejecting or approving a request. | [optional] 

## Methods

### NewUpdateSupportRequest

`func NewUpdateSupportRequest(requestStatus string, ) *UpdateSupportRequest`

NewUpdateSupportRequest instantiates a new UpdateSupportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSupportRequestWithDefaults

`func NewUpdateSupportRequestWithDefaults() *UpdateSupportRequest`

NewUpdateSupportRequestWithDefaults instantiates a new UpdateSupportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestStatus

`func (o *UpdateSupportRequest) GetRequestStatus() string`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *UpdateSupportRequest) GetRequestStatusOk() (*string, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *UpdateSupportRequest) SetRequestStatus(v string)`

SetRequestStatus sets RequestStatus field to given value.


### GetProcessingNote

`func (o *UpdateSupportRequest) GetProcessingNote() string`

GetProcessingNote returns the ProcessingNote field if non-nil, zero value otherwise.

### GetProcessingNoteOk

`func (o *UpdateSupportRequest) GetProcessingNoteOk() (*string, bool)`

GetProcessingNoteOk returns a tuple with the ProcessingNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingNote

`func (o *UpdateSupportRequest) SetProcessingNote(v string)`

SetProcessingNote sets ProcessingNote field to given value.

### HasProcessingNote

`func (o *UpdateSupportRequest) HasProcessingNote() bool`

HasProcessingNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


