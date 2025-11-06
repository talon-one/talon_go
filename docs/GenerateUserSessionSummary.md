# GenerateUserSessionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionID** | Pointer to **string** | The ID of the session. | 
**ApplicationID** | Pointer to **float32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

## Methods

### NewGenerateUserSessionSummary

`func NewGenerateUserSessionSummary(sessionID string, applicationID float32, ) *GenerateUserSessionSummary`

NewGenerateUserSessionSummary instantiates a new GenerateUserSessionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateUserSessionSummaryWithDefaults

`func NewGenerateUserSessionSummaryWithDefaults() *GenerateUserSessionSummary`

NewGenerateUserSessionSummaryWithDefaults instantiates a new GenerateUserSessionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionID

`func (o *GenerateUserSessionSummary) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *GenerateUserSessionSummary) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *GenerateUserSessionSummary) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetApplicationID

`func (o *GenerateUserSessionSummary) GetApplicationID() float32`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *GenerateUserSessionSummary) GetApplicationIDOk() (*float32, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *GenerateUserSessionSummary) SetApplicationID(v float32)`

SetApplicationID sets ApplicationID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


