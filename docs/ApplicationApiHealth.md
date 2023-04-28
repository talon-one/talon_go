# ApplicationApiHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **string** | One-word summary of the health of the API connection of an application. Possible values are: - &#x60;OK&#x60;: The Application has received only successful API requests in the last 5 minutes. - &#x60;WARNING&#x60;: The Application received at least one failed request in the last 50 minutes. - &#x60;ERROR&#x60;: More than 50% of received requests failed. - &#x60;CRITICAL&#x60;: All received requests failed. - &#x60;NONE&#x60;: During the last 5 minutes, the Application hasn&#39;t recorded any integration API requests.  | 
**LastUsed** | Pointer to [**time.Time**](time.Time.md) | time of last request relevant to the API health test. | 

## Methods

### GetSummary

`func (o *ApplicationApiHealth) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ApplicationApiHealth) GetSummaryOk() (string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSummary

`func (o *ApplicationApiHealth) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummary

`func (o *ApplicationApiHealth) SetSummary(v string)`

SetSummary gets a reference to the given string and assigns it to the Summary field.

### GetLastUsed

`func (o *ApplicationApiHealth) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ApplicationApiHealth) GetLastUsedOk() (time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUsed

`func (o *ApplicationApiHealth) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsed

`func (o *ApplicationApiHealth) SetLastUsed(v time.Time)`

SetLastUsed gets a reference to the given time.Time and assigns it to the LastUsed field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


