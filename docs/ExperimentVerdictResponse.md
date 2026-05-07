# ExperimentVerdictResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | Pointer to [**ExperimentVerdict**](ExperimentVerdict.md) |  | 
**Generated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment when the verdict was generated. | 

## Methods

### NewExperimentVerdictResponse

`func NewExperimentVerdictResponse(verdict ExperimentVerdict, generated time.Time, ) *ExperimentVerdictResponse`

NewExperimentVerdictResponse instantiates a new ExperimentVerdictResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVerdictResponseWithDefaults

`func NewExperimentVerdictResponseWithDefaults() *ExperimentVerdictResponse`

NewExperimentVerdictResponseWithDefaults instantiates a new ExperimentVerdictResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *ExperimentVerdictResponse) GetVerdict() ExperimentVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ExperimentVerdictResponse) GetVerdictOk() (*ExperimentVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ExperimentVerdictResponse) SetVerdict(v ExperimentVerdict)`

SetVerdict sets Verdict field to given value.


### GetGenerated

`func (o *ExperimentVerdictResponse) GetGenerated() time.Time`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *ExperimentVerdictResponse) GetGeneratedOk() (*time.Time, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *ExperimentVerdictResponse) SetGenerated(v time.Time)`

SetGenerated sets Generated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


