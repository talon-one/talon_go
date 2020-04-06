# \ManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLoyaltyPoints**](ManagementApi.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/add_points | Add points in a certain loyalty program for the specified customer
[**CopyCampaignToApplications**](ManagementApi.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into every specified application
[**CreateAdditionalCost**](ManagementApi.md#CreateAdditionalCost) | **Post** /v1/additional_costs | Define a new additional cost
[**CreateAttribute**](ManagementApi.md#CreateAttribute) | **Post** /v1/attributes | Define a new custom attribute
[**CreateCampaign**](ManagementApi.md#CreateCampaign) | **Post** /v1/applications/{applicationId}/campaigns | Create a Campaign
[**CreateCoupons**](ManagementApi.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create Coupons
[**CreatePasswordRecoveryEmail**](ManagementApi.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateRuleset**](ManagementApi.md#CreateRuleset) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | Create a Ruleset
[**CreateSession**](ManagementApi.md#CreateSession) | **Post** /v1/sessions | Create a Session
[**DeleteCampaign**](ManagementApi.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete a Campaign
[**DeleteCoupon**](ManagementApi.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete one Coupon
[**DeleteCoupons**](ManagementApi.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete Coupons
[**DeleteReferral**](ManagementApi.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete one Referral
[**DeleteRuleset**](ManagementApi.md#DeleteRuleset) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Delete a Ruleset
[**GetAccessLogs**](ManagementApi.md#GetAccessLogs) | **Get** /v1/applications/{applicationId}/access_logs | Get access logs for application
[**GetAccessLogsWithoutTotalCount**](ManagementApi.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for application
[**GetAccount**](ManagementApi.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get Account Details
[**GetAccountAnalytics**](ManagementApi.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get Account Analytics
[**GetAdditionalCost**](ManagementApi.md#GetAdditionalCost) | **Get** /v1/additional_costs/{additionalCostId} | Get an additional cost
[**GetAdditionalCosts**](ManagementApi.md#GetAdditionalCosts) | **Get** /v1/additional_costs | List additional costs
[**GetAllAccessLogs**](ManagementApi.md#GetAllAccessLogs) | **Get** /v1/access_logs | Get all access logs
[**GetAllRoles**](ManagementApi.md#GetAllRoles) | **Get** /v1/roles | Get all roles.
[**GetApplication**](ManagementApi.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementApi.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get report of health of application API
[**GetApplicationCustomer**](ManagementApi.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get Application Customer
[**GetApplicationCustomers**](ManagementApi.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List Application Customers
[**GetApplicationCustomersByAttributes**](ManagementApi.md#GetApplicationCustomersByAttributes) | **Post** /v1/application_customer_search | Get a list of the customer profiles that match the given attributes
[**GetApplicationEventTypes**](ManagementApi.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications Event Types
[**GetApplicationEvents**](ManagementApi.md#GetApplicationEvents) | **Get** /v1/applications/{applicationId}/events | List Applications Events
[**GetApplicationEventsWithoutTotalCount**](ManagementApi.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications Events
[**GetApplicationSession**](ManagementApi.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application Session
[**GetApplicationSessions**](ManagementApi.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application Sessions
[**GetApplications**](ManagementApi.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementApi.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get a custom attribute
[**GetAttributes**](ManagementApi.md#GetAttributes) | **Get** /v1/attributes | List custom attributes
[**GetCampaign**](ManagementApi.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get a Campaign
[**GetCampaignAnalytics**](ManagementApi.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementApi.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | Get a list of all campaigns that match the given attributes
[**GetCampaignSet**](ManagementApi.md#GetCampaignSet) | **Get** /v1/applications/{applicationId}/campaign_set | List CampaignSet
[**GetCampaigns**](ManagementApi.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List your Campaigns
[**GetChanges**](ManagementApi.md#GetChanges) | **Get** /v1/changes | Get audit log for an account
[**GetCoupons**](ManagementApi.md#GetCoupons) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | List Coupons
[**GetCouponsByAttributes**](ManagementApi.md#GetCouponsByAttributes) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search | Get a list of the coupons that match the given attributes
[**GetCouponsByAttributesApplicationWide**](ManagementApi.md#GetCouponsByAttributesApplicationWide) | **Post** /v1/applications/{applicationId}/coupons_search | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**GetCouponsWithoutTotalCount**](ManagementApi.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List Coupons
[**GetCustomerActivityReport**](ManagementApi.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get Activity Report for Single Customer
[**GetCustomerActivityReports**](ManagementApi.md#GetCustomerActivityReports) | **Get** /v1/applications/{applicationId}/customer_activity_reports | Get Activity Reports for Application Customers
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementApi.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementApi.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get Analytics Report for a Customer
[**GetCustomerProfile**](ManagementApi.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get Customer Profile
[**GetCustomerProfiles**](ManagementApi.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List Customer Profiles
[**GetCustomersByAttributes**](ManagementApi.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | Get a list of the customer profiles that match the given attributes
[**GetEventTypes**](ManagementApi.md#GetEventTypes) | **Get** /v1/event_types | List Event Types
[**GetExports**](ManagementApi.md#GetExports) | **Get** /v1/exports | Get Exports
[**GetImports**](ManagementApi.md#GetImports) | **Get** /v1/imports | Get Imports
[**GetLoyaltyPoints**](ManagementApi.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{programID}/profile/{integrationID} | get the Loyalty Ledger for this integrationID
[**GetLoyaltyProgram**](ManagementApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{programID} | Get a loyalty program
[**GetLoyaltyPrograms**](ManagementApi.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List all loyalty Programs
[**GetReferrals**](ManagementApi.md#GetReferrals) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals | List Referrals
[**GetReferralsWithoutTotalCount**](ManagementApi.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List Referrals
[**GetRole**](ManagementApi.md#GetRole) | **Get** /v1/roles/{roleId} | Get information for the specified role.
[**GetRuleset**](ManagementApi.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get a Ruleset
[**GetRulesets**](ManagementApi.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List Campaign Rulesets
[**GetUser**](ManagementApi.md#GetUser) | **Get** /v1/users/{userId} | Get a single User
[**GetUsers**](ManagementApi.md#GetUsers) | **Get** /v1/users | List Users in your account
[**GetWebhook**](ManagementApi.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get Webhook
[**GetWebhookActivationLogs**](ManagementApi.md#GetWebhookActivationLogs) | **Get** /v1/webhook_activation_logs | List Webhook activation Log Entries
[**GetWebhookLogs**](ManagementApi.md#GetWebhookLogs) | **Get** /v1/webhook_logs | List Webhook Log Entries
[**GetWebhooks**](ManagementApi.md#GetWebhooks) | **Get** /v1/webhooks | List Webhooks
[**RemoveLoyaltyPoints**](ManagementApi.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/deduct_points | Deduct points in a certain loyalty program for the specified customer
[**ResetPassword**](ManagementApi.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**SearchCouponsAdvanced**](ManagementApi.md#SearchCouponsAdvanced) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced | Get a list of the coupons that match the given attributes
[**SearchCouponsAdvancedApplicationWide**](ManagementApi.md#SearchCouponsAdvancedApplicationWide) | **Post** /v1/applications/{applicationId}/coupons_search_advanced | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | Get a list of the coupons that match the given attributes in all active campaigns of an application
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | Get a list of the coupons that match the given attributes
[**UpdateAdditionalCost**](ManagementApi.md#UpdateAdditionalCost) | **Put** /v1/additional_costs/{additionalCostId} | Update an additional cost
[**UpdateAttribute**](ManagementApi.md#UpdateAttribute) | **Put** /v1/attributes/{attributeId} | Update a custom attribute
[**UpdateCampaign**](ManagementApi.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update a Campaign
[**UpdateCampaignSet**](ManagementApi.md#UpdateCampaignSet) | **Put** /v1/applications/{applicationId}/campaign_set | Update a Campaign Set
[**UpdateCoupon**](ManagementApi.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update a Coupon
[**UpdateCouponBatch**](ManagementApi.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update a Batch of Coupons
[**UpdateRuleset**](ManagementApi.md#UpdateRuleset) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Update a Ruleset



## AddLoyaltyPoints

> AddLoyaltyPoints(ctx, programID, integrationID).Body(body).Execute()

Add points in a certain loyalty program for the specified customer

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**LoyaltyPoints**](LoyaltyPoints.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyCampaignToApplications

> InlineResponse2003 CopyCampaignToApplications(ctx, applicationId, campaignId).Body(body).Execute()

Copy the campaign into every specified application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyCampaignToApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CampaignCopy**](CampaignCopy.md) |  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalCost

> AccountAdditionalCost CreateAdditionalCost(ctx).Body(body).Execute()

Define a new additional cost



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) |  | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttribute

> Attribute CreateAttribute(ctx).Body(body).Execute()

Define a new custom attribute



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAttribute**](NewAttribute.md) |  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaign

> Campaign CreateCampaign(ctx, applicationId).Body(body).Execute()

Create a Campaign

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewCampaign**](NewCampaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupons

> InlineResponse2001 CreateCoupons(ctx, applicationId, campaignId).Body(body).Silent(silent).Execute()

Create Coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCoupons**](NewCoupons.md) |  | 
 **silent** | **string** | If set to &#39;yes&#39;, response will be an empty 204, otherwise a list of the coupons generated (to to 1000). | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePasswordRecoveryEmail

> NewPasswordEmail CreatePasswordRecoveryEmail(ctx).Body(body).Execute()

Request a password reset



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordRecoveryEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewPasswordEmail**](NewPasswordEmail.md) |  | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleset

> Ruleset CreateRuleset(ctx, applicationId, campaignId).Body(body).Execute()

Create a Ruleset

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewRuleset**](NewRuleset.md) |  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> Session CreateSession(ctx).Body(body).Execute()

Create a Session

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LoginParams**](LoginParams.md) |  | 

### Return type

[**Session**](Session.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, applicationId, campaignId).Execute()

Delete a Campaign

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupon

> DeleteCoupon(ctx, applicationId, campaignId, couponId).Execute()

Delete one Coupon

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**couponId** | **string** | The ID of the coupon code to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupons

> DeleteCoupons(ctx, applicationId, campaignId).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartsAfter(startsAfter).StartsBefore(startsBefore).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).Valid(valid).BatchId(batchId).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).Execute()

Delete Coupons

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReferral

> DeleteReferral(ctx, applicationId, campaignId, referralId).Execute()

Delete one Referral

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**referralId** | **string** | The ID of the referral code to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleset

> DeleteRuleset(ctx, applicationId, campaignId, rulesetId).Execute()

Delete a Ruleset

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**rulesetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogs

> InlineResponse2009 GetAccessLogs(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get access logs for application

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogsWithoutTotalCount

> InlineResponse20010 GetAccessLogsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get access logs for application

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get Account Details



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAnalytics

> AccountAnalytics GetAccountAnalytics(ctx, accountId).Execute()

Get Account Analytics



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCost

> AccountAdditionalCost GetAdditionalCost(ctx, additionalCostId).Execute()

Get an additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCosts

> InlineResponse20021 GetAdditionalCosts(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List additional costs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessLogs

> InlineResponse2009 GetAllAccessLogs(ctx).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get all access logs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoles

> InlineResponse20030 GetAllRoles(ctx).Execute()

Get all roles.

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesRequest struct via the builder pattern


### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, applicationId).Execute()

Get Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApiHealth

> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId).Execute()

Get report of health of application API

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApiHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomer

> ApplicationCustomer GetApplicationCustomer(ctx, applicationId, customerId).Execute()

Get Application Customer

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomers

> InlineResponse20012 GetApplicationCustomers(ctx, applicationId).Execute()

List Application Customers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomersByAttributes

> InlineResponse20013 GetApplicationCustomersByAttributes(ctx).Body(body).Execute()

Get a list of the customer profiles that match the given attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationCustomerSearch**](ApplicationCustomerSearch.md) |  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventTypes

> InlineResponse20019 GetApplicationEventTypes(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications Event Types



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEvents

> InlineResponse20017 GetApplicationEvents(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).Execute()

List Applications Events



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **type_** | **string** | Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **time.Time** | Only return events created before this date | 
 **createdAfter** | **time.Time** | Only return events created after this date | 
 **session** | **string** | Session integration ID filter for events. Must be exact match. | 
 **profile** | **string** | Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **string** | Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **string** | Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **string** | Coupon code | 
 **referralCode** | **string** | Referral code | 
 **ruleQuery** | **string** | Rule name filter for events | 
 **campaignQuery** | **string** | Campaign name filter for events | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventsWithoutTotalCount

> InlineResponse20018 GetApplicationEventsWithoutTotalCount(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).Execute()

List Applications Events



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **type_** | **string** | Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **time.Time** | Only return events created before this date | 
 **createdAfter** | **time.Time** | Only return events created after this date | 
 **session** | **string** | Session integration ID filter for events. Must be exact match. | 
 **profile** | **string** | Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **string** | Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **string** | Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **string** | Coupon code | 
 **referralCode** | **string** | Referral code | 
 **ruleQuery** | **string** | Rule name filter for events | 
 **campaignQuery** | **string** | Campaign name filter for events | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSession

> ApplicationSession GetApplicationSession(ctx, applicationId, sessionId).Execute()

Get Application Session

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**sessionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSessions

> InlineResponse20016 GetApplicationSessions(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Profile(profile).State(state).Coupon(coupon).Referral(referral).IntegrationId(integrationId).CustomerId(customerId).Execute()

List Application Sessions

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **profile** | **string** | Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **string** | Filter by sessions with this state. Must be exact match. | 
 **coupon** | **string** | Filter by sessions with this coupon. Must be exact match. | 
 **referral** | **string** | Filter by sessions with this referral. Must be exact match. | 
 **integrationId** | **string** | Filter by sessions with this integrationId. Must be exact match. | 
 **customerId** | **string** | Filter by integration ID of the customer for the session | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> InlineResponse2002 GetApplications(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttribute

> Attribute GetAttribute(ctx, attributeId).Execute()

Get a custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributes

> InlineResponse20020 GetAttributes(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List custom attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, applicationId, campaignId).Execute()

Get a Campaign

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAnalytics

> InlineResponse20011 GetCampaignAnalytics(ctx, applicationId, campaignId).RangeStart(rangeStart).RangeEnd(rangeEnd).Granularity(granularity).Execute()

Get analytics of campaigns

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **granularity** | **string** | The time interval between the results in the returned time-series. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByAttributes

> InlineResponse2003 GetCampaignByAttributes(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Execute()

Get a list of all campaigns that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignSearch**](CampaignSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **string** | Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSet

> CampaignSet GetCampaignSet(ctx, applicationId).Execute()

List CampaignSet

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignSet**](CampaignSet.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> InlineResponse2003 GetCampaigns(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Name(name).Tags(tags).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List your Campaigns

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **string** | Filter results by the state of the campaign. | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> InlineResponse20027 GetChanges(ctx).PageSize(pageSize).Skip(skip).Sort(sort).ApplicationId(applicationId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).WithTotalResultSize(withTotalResultSize).IncludeOld(includeOld).Execute()

Get audit log for an account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **applicationId** | **int32** |  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, hasMore will be be true whenever there is a next page. totalResultSize will always be zero. With this flag set to false, hasMore will always be set to false. totalResultSize will contain the total number of results for this query.  | 
 **includeOld** | **bool** | When this flag is set to false, the state without the change will not be returned. The default value is true. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoupons

> InlineResponse2001 GetCoupons(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartsAfter(startsAfter).StartsBefore(startsBefore).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).Valid(valid).BatchId(batchId).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).Execute()

List Coupons

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsByAttributes

> InlineResponse2001 GetCouponsByAttributes(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

Get a list of the coupons that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CouponSearch**](CouponSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsByAttributesApplicationWide

> InlineResponse2001 GetCouponsByAttributesApplicationWide(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

Get a list of the coupons that match the given attributes in all active campaigns of an application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsByAttributesApplicationWideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponSearch**](CouponSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsWithoutTotalCount

> InlineResponse2005 GetCouponsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).Execute()

List Coupons

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReport

> CustomerActivityReport GetCustomerActivityReport(ctx, applicationId, customerId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Execute()

Get Activity Report for Single Customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReports

> InlineResponse20014 GetCustomerActivityReports(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()

Get Activity Reports for Application Customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **name** | **string** | Only return reports matching the customer name | 
 **integrationId** | **string** | Only return reports matching the integrationId | 
 **campaignName** | **string** | Only return reports matching the campaignName | 
 **advocateName** | **string** | Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReportsWithoutTotalCount

> InlineResponse20015 GetCustomerActivityReportsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()

Get Activity Reports for Application Customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **name** | **string** | Only return reports matching the customer name | 
 **integrationId** | **string** | Only return reports matching the integrationId | 
 **campaignName** | **string** | Only return reports matching the campaignName | 
 **advocateName** | **string** | Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAnalytics

> CustomerAnalytics GetCustomerAnalytics(ctx, applicationId, customerId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get Analytics Report for a Customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfile

> ApplicationCustomer GetCustomerProfile(ctx, applicationId, customerId).Execute()

Get Customer Profile

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfiles

> InlineResponse20013 GetCustomerProfiles(ctx).PageSize(pageSize).Skip(skip).Execute()

List Customer Profiles

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersByAttributes

> InlineResponse20013 GetCustomersByAttributes(ctx).Body(body).PageSize(pageSize).Skip(skip).Execute()

Get a list of the customer profiles that match the given attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationCustomerSearch**](ApplicationCustomerSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> InlineResponse20025 GetEventTypes(ctx).ApplicationIds(applicationIds).Name(name).IncludeOldVersions(includeOldVersions).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Event Types



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Filter by one or more application ids separated by comma | 
 **name** | **string** | Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **bool** | Include all versions of every event type. | [default to false]
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> InlineResponse20028 GetExports(ctx).PageSize(pageSize).Skip(skip).ApplicationId(applicationId).CampaignId(campaignId).Entity(entity).Execute()

Get Exports



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **applicationId** | **int32** |  | 
 **campaignId** | **int32** |  | 
 **entity** | **string** | The name of the entity type that was exported. | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImports

> InlineResponse20029 GetImports(ctx).PageSize(pageSize).Skip(skip).Execute()

Get Imports



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPoints

> LoyaltyLedger GetLoyaltyPoints(ctx, programID, integrationID).Execute()

get the Loyalty Ledger for this integrationID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** | The identifier for the application, must be unique within the account. | 
**integrationID** | **string** | The identifier for the application, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> LoyaltyProgram GetLoyaltyProgram(ctx, programID).Execute()

Get a loyalty program

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPrograms

> InlineResponse2008 GetLoyaltyPrograms(ctx).Execute()

List all loyalty Programs

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramsRequest struct via the builder pattern


### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferrals

> InlineResponse2006 GetReferrals(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()

List Referrals

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **code** | **string** | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **string** | Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralsWithoutTotalCount

> InlineResponse2007 GetReferralsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()

List Referrals

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **code** | **string** | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **string** | Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleId).Execute()

Get information for the specified role.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleset

> Ruleset GetRuleset(ctx, applicationId, campaignId, rulesetId).Execute()

Get a Ruleset

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**rulesetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesets

> InlineResponse2004 GetRulesets(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Campaign Rulesets

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get a single User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse20026 GetUsers(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Users in your account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get Webhook



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookActivationLogs

> InlineResponse20023 GetWebhookActivationLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).IntegrationRequestUuid(integrationRequestUuid).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List Webhook activation Log Entries



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookActivationLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **integrationRequestUuid** | **string** | Filter results by integration request UUID. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** |  | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Only return events created before this date. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogs

> InlineResponse20024 GetWebhookLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Status(status).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).RequestUuid(requestUuid).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List Webhook Log Entries

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** |  | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **requestUuid** | **string** | Filter results by request UUID. | 
 **createdBefore** | **time.Time** | Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> InlineResponse20022 GetWebhooks(ctx).ApplicationIds(applicationIds).Sort(sort).PageSize(pageSize).Skip(skip).Execute()

List Webhooks

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Filter by one or more application ids separated by comma | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoyaltyPoints

> RemoveLoyaltyPoints(ctx, programID, integrationID).Body(body).Execute()

Deduct points in a certain loyalty program for the specified customer

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**LoyaltyPoints**](LoyaltyPoints.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> NewPassword ResetPassword(ctx).Body(body).Execute()

Reset password



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewPassword**](NewPassword.md) |  | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvanced

> InlineResponse2001 SearchCouponsAdvanced(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

Get a list of the coupons that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWide

> InlineResponse2001 SearchCouponsAdvancedApplicationWide(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

Get a list of the coupons that match the given attributes in all active campaigns of an application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedApplicationWideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWideWithoutTotalCount

> InlineResponse2005 SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

Get a list of the coupons that match the given attributes in all active campaigns of an application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedApplicationWideWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedWithoutTotalCount

> InlineResponse2005 SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

Get a list of the coupons that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalCost

> AccountAdditionalCost UpdateAdditionalCost(ctx, additionalCostId).Body(body).Execute()

Update an additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) |  | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> Attribute UpdateAttribute(ctx, attributeId).Body(body).Execute()

Update a custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAttribute**](NewAttribute.md) |  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, applicationId, campaignId).Body(body).Execute()

Update a Campaign

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCampaign**](UpdateCampaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignSet

> CampaignSet UpdateCampaignSet(ctx, applicationId).Body(body).Execute()

Update a Campaign Set

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewCampaignSet**](NewCampaignSet.md) |  | 

### Return type

[**CampaignSet**](CampaignSet.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> Coupon UpdateCoupon(ctx, applicationId, campaignId, couponId).Body(body).Execute()

Update a Coupon

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**couponId** | **string** | The ID of the coupon code to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateCoupon**](UpdateCoupon.md) |  | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponBatch

> UpdateCouponBatch(ctx, applicationId, campaignId).Body(body).Execute()

Update a Batch of Coupons

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCouponBatch**](UpdateCouponBatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleset

> Ruleset UpdateRuleset(ctx, applicationId, campaignId, rulesetId).Body(body).Execute()

Update a Ruleset

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**rulesetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NewRuleset**](NewRuleset.md) |  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

