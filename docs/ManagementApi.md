# \ManagementApi

All URLs are relative to `https://yourbaseurl.talon.one`.

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUserByEmail**](ManagementApi.md#ActivateUserByEmail) | **Post** /v1/users/activate | Enable user by email address
[**AddLoyaltyCardPoints**](ManagementApi.md#AddLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/add_points | Add points to card
[**AddLoyaltyPoints**](ManagementApi.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/add_points | Add points to customer profile
[**CopyCampaignToApplications**](ManagementApi.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into the specified Application
[**CreateAccountCollection**](ManagementApi.md#CreateAccountCollection) | **Post** /v1/collections | Create account-level collection
[**CreateAchievement**](ManagementApi.md#CreateAchievement) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | Create achievement
[**CreateAdditionalCost**](ManagementApi.md#CreateAdditionalCost) | **Post** /v1/additional_costs | Create additional cost
[**CreateAttribute**](ManagementApi.md#CreateAttribute) | **Post** /v1/attributes | Create custom attribute
[**CreateBatchLoyaltyCards**](ManagementApi.md#CreateBatchLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards/batch | Create loyalty cards
[**CreateCampaignFromTemplate**](ManagementApi.md#CreateCampaignFromTemplate) | **Post** /v1/applications/{applicationId}/create_campaign_from_template | Create campaign from campaign template
[**CreateCollection**](ManagementApi.md#CreateCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | Create campaign-level collection
[**CreateCoupons**](ManagementApi.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
[**CreateCouponsAsync**](ManagementApi.md#CreateCouponsAsync) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_async | Create coupons asynchronously
[**CreateCouponsDeletionJob**](ManagementApi.md#CreateCouponsDeletionJob) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_deletion_jobs | Creates a coupon deletion job
[**CreateCouponsForMultipleRecipients**](ManagementApi.md#CreateCouponsForMultipleRecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
[**CreateInviteEmail**](ManagementApi.md#CreateInviteEmail) | **Post** /v1/invite_emails | Resend invitation email
[**CreateInviteV2**](ManagementApi.md#CreateInviteV2) | **Post** /v2/invites | Invite user
[**CreatePasswordRecoveryEmail**](ManagementApi.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateSession**](ManagementApi.md#CreateSession) | **Post** /v1/sessions | Create session
[**CreateStore**](ManagementApi.md#CreateStore) | **Post** /v1/applications/{applicationId}/stores | Create store
[**DeactivateUserByEmail**](ManagementApi.md#DeactivateUserByEmail) | **Post** /v1/users/deactivate | Disable user by email address
[**DeductLoyaltyCardPoints**](ManagementApi.md#DeductLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/deduct_points | Deduct points from card
[**DeleteAccountCollection**](ManagementApi.md#DeleteAccountCollection) | **Delete** /v1/collections/{collectionId} | Delete account-level collection
[**DeleteAchievement**](ManagementApi.md#DeleteAchievement) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Delete achievement
[**DeleteCampaign**](ManagementApi.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
[**DeleteCollection**](ManagementApi.md#DeleteCollection) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Delete campaign-level collection
[**DeleteCoupon**](ManagementApi.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
[**DeleteCoupons**](ManagementApi.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
[**DeleteLoyaltyCard**](ManagementApi.md#DeleteLoyaltyCard) | **Delete** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Delete loyalty card
[**DeleteReferral**](ManagementApi.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
[**DeleteStore**](ManagementApi.md#DeleteStore) | **Delete** /v1/applications/{applicationId}/stores/{storeId} | Delete store
[**DeleteUser**](ManagementApi.md#DeleteUser) | **Delete** /v1/users/{userId} | Delete user
[**DeleteUserByEmail**](ManagementApi.md#DeleteUserByEmail) | **Post** /v1/users/delete | Delete user by email address
[**DestroySession**](ManagementApi.md#DestroySession) | **Delete** /v1/sessions | Destroy session
[**DisconnectCampaignStores**](ManagementApi.md#DisconnectCampaignStores) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores | Disconnect stores
[**ExportAccountCollectionItems**](ManagementApi.md#ExportAccountCollectionItems) | **Get** /v1/collections/{collectionId}/export | Export account-level collection&#39;s items
[**ExportAchievements**](ManagementApi.md#ExportAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId}/export | Export achievement customer data
[**ExportAudiencesMemberships**](ManagementApi.md#ExportAudiencesMemberships) | **Get** /v1/audiences/{audienceId}/memberships/export | Export audience members
[**ExportCampaignStores**](ManagementApi.md#ExportCampaignStores) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/export | Export stores
[**ExportCollectionItems**](ManagementApi.md#ExportCollectionItems) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/export | Export campaign-level collection&#39;s items
[**ExportCoupons**](ManagementApi.md#ExportCoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons
[**ExportCustomerSessions**](ManagementApi.md#ExportCustomerSessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions
[**ExportCustomersTiers**](ManagementApi.md#ExportCustomersTiers) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customers_tiers | Export customers&#39; tier data
[**ExportEffects**](ManagementApi.md#ExportEffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects
[**ExportLoyaltyBalance**](ManagementApi.md#ExportLoyaltyBalance) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balance | Export customer loyalty balance to CSV
[**ExportLoyaltyBalances**](ManagementApi.md#ExportLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balances | Export customer loyalty balances
[**ExportLoyaltyCardBalances**](ManagementApi.md#ExportLoyaltyCardBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_card_balances | Export all card transaction logs
[**ExportLoyaltyCardLedger**](ManagementApi.md#ExportLoyaltyCardLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/export_log | Export card&#39;s ledger log
[**ExportLoyaltyCards**](ManagementApi.md#ExportLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/export | Export loyalty cards
[**ExportLoyaltyLedger**](ManagementApi.md#ExportLoyaltyLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/export_log | Export customer&#39;s transaction logs
[**ExportPoolGiveaways**](ManagementApi.md#ExportPoolGiveaways) | **Get** /v1/giveaways/pools/{poolId}/export | Export giveaway codes of a giveaway pool
[**ExportReferrals**](ManagementApi.md#ExportReferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals
[**GetAccessLogsWithoutTotalCount**](ManagementApi.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
[**GetAccount**](ManagementApi.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get account details
[**GetAccountAnalytics**](ManagementApi.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
[**GetAccountCollection**](ManagementApi.md#GetAccountCollection) | **Get** /v1/collections/{collectionId} | Get account-level collection
[**GetAchievement**](ManagementApi.md#GetAchievement) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Get achievement
[**GetAdditionalCost**](ManagementApi.md#GetAdditionalCost) | **Get** /v1/additional_costs/{additionalCostId} | Get additional cost
[**GetAdditionalCosts**](ManagementApi.md#GetAdditionalCosts) | **Get** /v1/additional_costs | List additional costs
[**GetApplication**](ManagementApi.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementApi.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get Application health
[**GetApplicationCustomer**](ManagementApi.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
[**GetApplicationCustomerFriends**](ManagementApi.md#GetApplicationCustomerFriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
[**GetApplicationCustomers**](ManagementApi.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
[**GetApplicationCustomersByAttributes**](ManagementApi.md#GetApplicationCustomersByAttributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
[**GetApplicationEventTypes**](ManagementApi.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
[**GetApplicationEventsWithoutTotalCount**](ManagementApi.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
[**GetApplicationSession**](ManagementApi.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
[**GetApplicationSessions**](ManagementApi.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
[**GetApplications**](ManagementApi.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementApi.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get custom attribute
[**GetAttributes**](ManagementApi.md#GetAttributes) | **Get** /v1/attributes | List custom attributes
[**GetAudienceMemberships**](ManagementApi.md#GetAudienceMemberships) | **Get** /v1/audiences/{audienceId}/memberships | List audience members
[**GetAudiences**](ManagementApi.md#GetAudiences) | **Get** /v1/audiences | List audiences
[**GetAudiencesAnalytics**](ManagementApi.md#GetAudiencesAnalytics) | **Get** /v1/audiences/analytics | List audience analytics
[**GetCampaign**](ManagementApi.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
[**GetCampaignAnalytics**](ManagementApi.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementApi.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
[**GetCampaignGroup**](ManagementApi.md#GetCampaignGroup) | **Get** /v1/campaign_groups/{campaignGroupId} | Get campaign access group
[**GetCampaignGroups**](ManagementApi.md#GetCampaignGroups) | **Get** /v1/campaign_groups | List campaign access groups
[**GetCampaignTemplates**](ManagementApi.md#GetCampaignTemplates) | **Get** /v1/campaign_templates | List campaign templates
[**GetCampaigns**](ManagementApi.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
[**GetChanges**](ManagementApi.md#GetChanges) | **Get** /v1/changes | Get audit logs for an account
[**GetCollection**](ManagementApi.md#GetCollection) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Get campaign-level collection
[**GetCollectionItems**](ManagementApi.md#GetCollectionItems) | **Get** /v1/collections/{collectionId}/items | Get collection items
[**GetCouponsWithoutTotalCount**](ManagementApi.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
[**GetCustomerActivityReport**](ManagementApi.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementApi.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementApi.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
[**GetCustomerProfile**](ManagementApi.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get customer profile
[**GetCustomerProfileAchievementProgress**](ManagementApi.md#GetCustomerProfileAchievementProgress) | **Get** /v1/applications/{applicationId}/achievement_progress/{integrationId} | List customer achievements
[**GetCustomerProfiles**](ManagementApi.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List customer profiles
[**GetCustomersByAttributes**](ManagementApi.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
[**GetEventTypes**](ManagementApi.md#GetEventTypes) | **Get** /v1/event_types | List event types
[**GetExports**](ManagementApi.md#GetExports) | **Get** /v1/exports | Get exports
[**GetLoyaltyCard**](ManagementApi.md#GetLoyaltyCard) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Get loyalty card
[**GetLoyaltyCardTransactionLogs**](ManagementApi.md#GetLoyaltyCardTransactionLogs) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/logs | List card&#39;s transactions
[**GetLoyaltyCards**](ManagementApi.md#GetLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards | List loyalty cards
[**GetLoyaltyPoints**](ManagementApi.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId} | Get customer&#39;s full loyalty ledger
[**GetLoyaltyProgram**](ManagementApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{loyaltyProgramId} | Get loyalty program
[**GetLoyaltyProgramTransactions**](ManagementApi.md#GetLoyaltyProgramTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/transactions | List loyalty program transactions
[**GetLoyaltyPrograms**](ManagementApi.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List loyalty programs
[**GetLoyaltyStatistics**](ManagementApi.md#GetLoyaltyStatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/statistics | Get loyalty program statistics
[**GetReferralsWithoutTotalCount**](ManagementApi.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
[**GetRoleV2**](ManagementApi.md#GetRoleV2) | **Get** /v2/roles/{roleId} | Get role
[**GetRuleset**](ManagementApi.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
[**GetRulesets**](ManagementApi.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
[**GetStore**](ManagementApi.md#GetStore) | **Get** /v1/applications/{applicationId}/stores/{storeId} | Get store
[**GetUser**](ManagementApi.md#GetUser) | **Get** /v1/users/{userId} | Get user
[**GetUsers**](ManagementApi.md#GetUsers) | **Get** /v1/users | List users in account
[**GetWebhook**](ManagementApi.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get webhook
[**GetWebhookActivationLogs**](ManagementApi.md#GetWebhookActivationLogs) | **Get** /v1/webhook_activation_logs | List webhook activation log entries
[**GetWebhookLogs**](ManagementApi.md#GetWebhookLogs) | **Get** /v1/webhook_logs | List webhook log entries
[**GetWebhooks**](ManagementApi.md#GetWebhooks) | **Get** /v1/webhooks | List webhooks
[**ImportAccountCollection**](ManagementApi.md#ImportAccountCollection) | **Post** /v1/collections/{collectionId}/import | Import data into existing account-level collection
[**ImportAllowedList**](ManagementApi.md#ImportAllowedList) | **Post** /v1/attributes/{attributeId}/allowed_list/import | Import allowed values for attribute
[**ImportAudiencesMemberships**](ManagementApi.md#ImportAudiencesMemberships) | **Post** /v1/audiences/{audienceId}/memberships/import | Import audience members
[**ImportCampaignStores**](ManagementApi.md#ImportCampaignStores) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/import | Import stores
[**ImportCollection**](ManagementApi.md#ImportCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import data into existing campaign-level collection
[**ImportCoupons**](ManagementApi.md#ImportCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons
[**ImportLoyaltyCards**](ManagementApi.md#ImportLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_cards | Import loyalty cards
[**ImportLoyaltyCustomersTiers**](ManagementApi.md#ImportLoyaltyCustomersTiers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_customers_tiers | Import customers into loyalty tiers
[**ImportLoyaltyPoints**](ManagementApi.md#ImportLoyaltyPoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_points | Import loyalty points
[**ImportPoolGiveaways**](ManagementApi.md#ImportPoolGiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
[**ImportReferrals**](ManagementApi.md#ImportReferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals
[**InviteUserExternal**](ManagementApi.md#InviteUserExternal) | **Post** /v1/users/invite | Invite user from identity provider
[**ListAccountCollections**](ManagementApi.md#ListAccountCollections) | **Get** /v1/collections | List collections in account
[**ListAchievements**](ManagementApi.md#ListAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | List achievements
[**ListAllRolesV2**](ManagementApi.md#ListAllRolesV2) | **Get** /v2/roles | List roles
[**ListCatalogItems**](ManagementApi.md#ListCatalogItems) | **Get** /v1/catalogs/{catalogId}/items | List items in a catalog
[**ListCollections**](ManagementApi.md#ListCollections) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | List collections in campaign
[**ListCollectionsInApplication**](ManagementApi.md#ListCollectionsInApplication) | **Get** /v1/applications/{applicationId}/collections | List collections in Application
[**ListStores**](ManagementApi.md#ListStores) | **Get** /v1/applications/{applicationId}/stores | List stores
[**NotificationActivation**](ManagementApi.md#NotificationActivation) | **Put** /v1/notifications/{notificationId}/activation | Activate or deactivate notification
[**OktaEventHandlerChallenge**](ManagementApi.md#OktaEventHandlerChallenge) | **Get** /v1/provisioning/okta | Validate Okta API ownership
[**PostAddedDeductedPointsNotification**](ManagementApi.md#PostAddedDeductedPointsNotification) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/notifications/added_deducted_points | Create notification about added or deducted loyalty points
[**PostCatalogsStrikethroughNotification**](ManagementApi.md#PostCatalogsStrikethroughNotification) | **Post** /v1/applications/{applicationId}/catalogs/notifications/strikethrough | Create strikethrough notification
[**PostPendingPointsNotification**](ManagementApi.md#PostPendingPointsNotification) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/notifications/pending_points | Create notification about pending loyalty points
[**RemoveLoyaltyPoints**](ManagementApi.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/deduct_points | Deduct points from customer profile
[**ResetPassword**](ManagementApi.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**ScimCreateUser**](ManagementApi.md#ScimCreateUser) | **Post** /v1/provisioning/scim/Users | Create SCIM user
[**ScimDeleteUser**](ManagementApi.md#ScimDeleteUser) | **Delete** /v1/provisioning/scim/Users/{userId} | Delete SCIM user
[**ScimGetResourceTypes**](ManagementApi.md#ScimGetResourceTypes) | **Get** /v1/provisioning/scim/ResourceTypes | List supported SCIM resource types
[**ScimGetSchemas**](ManagementApi.md#ScimGetSchemas) | **Get** /v1/provisioning/scim/Schemas | List supported SCIM schemas
[**ScimGetServiceProviderConfig**](ManagementApi.md#ScimGetServiceProviderConfig) | **Get** /v1/provisioning/scim/ServiceProviderConfig | Get SCIM service provider configuration
[**ScimGetUser**](ManagementApi.md#ScimGetUser) | **Get** /v1/provisioning/scim/Users/{userId} | Get SCIM user
[**ScimGetUsers**](ManagementApi.md#ScimGetUsers) | **Get** /v1/provisioning/scim/Users | List SCIM users
[**ScimPatchUser**](ManagementApi.md#ScimPatchUser) | **Patch** /v1/provisioning/scim/Users/{userId} | Update SCIM user attributes
[**ScimReplaceUserAttributes**](ManagementApi.md#ScimReplaceUserAttributes) | **Put** /v1/provisioning/scim/Users/{userId} | Update SCIM user
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes (without total count)
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes in campaign (without total count)
[**TransferLoyaltyCard**](ManagementApi.md#TransferLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transfer | Transfer card data
[**UpdateAccountCollection**](ManagementApi.md#UpdateAccountCollection) | **Put** /v1/collections/{collectionId} | Update account-level collection
[**UpdateAchievement**](ManagementApi.md#UpdateAchievement) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Update achievement
[**UpdateAdditionalCost**](ManagementApi.md#UpdateAdditionalCost) | **Put** /v1/additional_costs/{additionalCostId} | Update additional cost
[**UpdateAttribute**](ManagementApi.md#UpdateAttribute) | **Put** /v1/attributes/{attributeId} | Update custom attribute
[**UpdateCampaign**](ManagementApi.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
[**UpdateCollection**](ManagementApi.md#UpdateCollection) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Update campaign-level collection&#39;s description
[**UpdateCoupon**](ManagementApi.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
[**UpdateCouponBatch**](ManagementApi.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update coupons
[**UpdateLoyaltyCard**](ManagementApi.md#UpdateLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Update loyalty card status
[**UpdateReferral**](ManagementApi.md#UpdateReferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
[**UpdateRoleV2**](ManagementApi.md#UpdateRoleV2) | **Put** /v2/roles/{roleId} | Update role
[**UpdateStore**](ManagementApi.md#UpdateStore) | **Put** /v1/applications/{applicationId}/stores/{storeId} | Update store
[**UpdateUser**](ManagementApi.md#UpdateUser) | **Put** /v1/users/{userId} | Update user



## ActivateUserByEmail

> ActivateUserByEmail(ctx, body)

Enable user by email address

Enable a [disabled user](https://docs.talon.one/docs/product/account/account-settings/managing-users#disabling-a-user) by their email address. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **DeactivateUserRequest**| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyCardPoints

> AddLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId, body)

Add points to card

Add points to the given loyalty card in the specified card-based loyalty program. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**body** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyPoints

> AddLoyaltyPoints(ctx, loyaltyProgramId, integrationId, body)

Add points to customer profile

Add points in the specified loyalty program for the given customer.  To get the `integrationId` of the profile from a `sessionId`, use the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
**integrationId** | **string**| The identifier of the profile. | 
**body** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyCampaignToApplications

> InlineResponse2006 CopyCampaignToApplications(ctx, applicationId, campaignId, body)

Copy the campaign into the specified Application

Copy the campaign into all specified Applications.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**CampaignCopy**](CampaignCopy.md)| body | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCollection

> Collection CreateAccountCollection(ctx, body)

Create account-level collection

Create an account-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewCollection**](NewCollection.md)| body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAchievement

> Achievement CreateAchievement(ctx, applicationId, campaignId, body)

Create achievement

Create a new achievement in a specific campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**CreateAchievement**](CreateAchievement.md)| body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalCost

> AccountAdditionalCost CreateAdditionalCost(ctx, body)

Create additional cost

Create an [additional cost](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs).  These additional costs are shared across all applications in your account, and are never required. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewAdditionalCost**](NewAdditionalCost.md)| body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttribute

> Attribute CreateAttribute(ctx, body)

Create custom attribute

Create a _custom attribute_ in this account. [Custom attributes](https://docs.talon.one/docs/dev/concepts/attributes) allow you to add data to Talon.One domain entities like campaigns, coupons, customers and so on.  These attributes can then be given values when creating/updating these entities, and these values can be used in your campaign rules.  For example, you could define a `zipCode` field for customer sessions, and add a rule to your campaign that only allows certain ZIP codes.  These attributes are shared across all Applications in your account and are never required. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewAttribute**](NewAttribute.md)| body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchLoyaltyCards

> LoyaltyCardBatchResponse CreateBatchLoyaltyCards(ctx, loyaltyProgramId, body)

Create loyalty cards

Create a batch of loyalty cards in a specified [card-based loyalty program](https://docs.talon.one/docs/product/loyalty-programs/overview#loyalty-program-types).  Customers can use loyalty cards to collect and spend loyalty points.  **Important:**  - The specified card-based loyalty program must have a defined card code format that is used to generate the loyalty card codes. - Trying to create more than 20,000 loyalty cards in a single request returns an error message with a `400` status code. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**body** | [**LoyaltyCardBatch**](LoyaltyCardBatch.md)| body | 

### Return type

[**LoyaltyCardBatchResponse**](LoyaltyCardBatchResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignFromTemplate

> CreateTemplateCampaignResponse CreateCampaignFromTemplate(ctx, applicationId, body)

Create campaign from campaign template

Use the campaign template referenced in the request body to create a new campaign in one of the connected Applications.  If the template was created from a campaign with rules referencing [campaign collections](https://docs.talon.one/docs/product/campaigns/managing-collections), the corresponding collections for the new campaign are created automatically. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | [**CreateTemplateCampaign**](CreateTemplateCampaign.md)| body | 

### Return type

[**CreateTemplateCampaignResponse**](CreateTemplateCampaignResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx, applicationId, campaignId, body)

Create campaign-level collection

Create a campaign-level collection in a given campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewCampaignCollection**](NewCampaignCollection.md)| body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupons

> InlineResponse2008 CreateCoupons(ctx, applicationId, campaignId, body, optional)

Create coupons

Create coupons according to some pattern. Up to 20.000 coupons can be created without a unique prefix. When a unique prefix is provided, up to 200.000 coupons can be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewCoupons**](NewCoupons.md)| body | 
 **optional** | ***CreateCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCouponsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **silent** | **optional.**| Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsAsync

> AsyncCouponCreationResponse CreateCouponsAsync(ctx, applicationId, campaignId, body)

Create coupons asynchronously

Create up to 5,000,000 coupons asynchronously. You should typically use this enpdoint when you create at least 20,001 coupons. You receive an email when the creation is complete.  If you want to create less than 20,001 coupons, you can use the [Create coupons](https://docs.talon.one/management-api#tag/Coupons/operation/createCoupons) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewCouponCreationJob**](NewCouponCreationJob.md)| body | 

### Return type

[**AsyncCouponCreationResponse**](AsyncCouponCreationResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsDeletionJob

> AsyncCouponDeletionJobResponse CreateCouponsDeletionJob(ctx, applicationId, campaignId, body)

Creates a coupon deletion job

This endpoint handles creating a job to delete coupons asynchronously. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewCouponDeletionJob**](NewCouponDeletionJob.md)| body | 

### Return type

[**AsyncCouponDeletionJobResponse**](AsyncCouponDeletionJobResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsForMultipleRecipients

> InlineResponse2008 CreateCouponsForMultipleRecipients(ctx, applicationId, campaignId, body, optional)

Create coupons for multiple recipients

Create coupons according to some pattern for up to 1000 recipients.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewCouponsForMultipleRecipients**](NewCouponsForMultipleRecipients.md)| body | 
 **optional** | ***CreateCouponsForMultipleRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCouponsForMultipleRecipientsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **silent** | **optional.**| Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteEmail

> NewInviteEmail CreateInviteEmail(ctx, body)

Resend invitation email

Resend an email invitation to an existing user.  **Note:** The invitation token is valid for 24 hours after the email has been sent. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewInviteEmail**](NewInviteEmail.md)| body | 

### Return type

[**NewInviteEmail**](NewInviteEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteV2

> User CreateInviteV2(ctx, body)

Invite user

Create a new user in the account and send an invitation to their email address.  **Note**: The invitation token is valid for 24 hours after the email has been sent. You can resend an invitation to a user with the [Resend invitation email](https://docs.talon.one/management-api#tag/Accounts-and-users/operation/createInviteEmail) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewInvitation**](NewInvitation.md)| body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePasswordRecoveryEmail

> NewPasswordEmail CreatePasswordRecoveryEmail(ctx, body)

Request a password reset

Send an email with a password recovery link to the email address of an existing account.  **Note:** The password recovery link expires 30 minutes after this endpoint is triggered. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewPasswordEmail**](NewPasswordEmail.md)| body | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> Session CreateSession(ctx, body)

Create session

Create a session to use the Management API endpoints. Use the value of the `token` property provided in the response as bearer token in other API calls.  A token is valid for 3 months. In accordance with best pratices, use your generated token for all your API requests. Do **not** regenerate a token for each request.  This endpoint has a rate limit of 3 to 6 requests per second per account, depending on your setup.  <div class=\"redoc-section\">   <p class=\"title\">Granular API key</p>   Instead of using a session, you can also use the <a href=\"https://docs.talon.one/docs/product/account/dev-tools/managing-mapi-keys\">Management API key feature</a>   in the Campaign Manager to decide which endpoints can be used with a given key. </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LoginParams**](LoginParams.md)| body | 

### Return type

[**Session**](Session.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStore

> Store CreateStore(ctx, applicationId, body)

Create store

Create a new store in a specific Application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewStore**](NewStore.md)| body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUserByEmail

> DeactivateUserByEmail(ctx, body)

Disable user by email address

[Disable a specific user](https://docs.talon.one/docs/product/account/account-settings/managing-users#disabling-a-user) by their email address. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DeactivateUserRequest**](DeactivateUserRequest.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeductLoyaltyCardPoints

> DeductLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId, body)

Deduct points from card

Deduct points from the given loyalty card in the specified card-based loyalty program. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**body** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountCollection

> DeleteAccountCollection(ctx, collectionId)

Delete account-level collection

Delete a given account-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAchievement

> DeleteAchievement(ctx, applicationId, campaignId, achievementId)

Delete achievement

Delete the specified achievement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32**| The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, applicationId, campaignId)

Delete campaign

Delete the given campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, applicationId, campaignId, collectionId)

Delete campaign-level collection

Delete a given campaign-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupon

> DeleteCoupon(ctx, applicationId, campaignId, couponId)

Delete coupon

Delete the specified coupon.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string**| The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupons

> DeleteCoupons(ctx, applicationId, campaignId, optional)

Delete coupons

Deletes all the coupons matching the specified criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***DeleteCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCouponsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **optional.**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| - &#x60;expired&#x60;: Matches coupons in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches coupons in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches coupons in which start date is set and in the future.  | 
 **batchId** | **optional.**| Filter results by batches of coupons | 
 **usable** | **optional.**| - &#x60;true&#x60;: only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **referralId** | **optional.**| Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **optional.**| Filter results by match with a profile ID specified in the coupon&#39;s &#x60;RecipientIntegrationId&#x60; field.  | 
 **exactMatch** | **optional.**| Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoyaltyCard

> DeleteLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId)

Delete loyalty card

Delete the given loyalty card.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReferral

> DeleteReferral(ctx, applicationId, campaignId, referralId)

Delete referral

Delete the specified referral.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string**| The ID of the referral code. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStore(ctx, applicationId, storeId)

Delete store

Delete the specified store.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string**| The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId)

Delete user

Delete a specific user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserByEmail

> DeleteUserByEmail(ctx, body)

Delete user by email address

[Delete a specific user](https://docs.talon.one/docs/product/account/account-settings/managing-users#deleting-a-user) by their email address. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **DeactivateUserRequest**| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySession

> DestroySession(ctx, )

Destroy session

Destroys the session.

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectCampaignStores

> DisconnectCampaignStores(ctx, applicationId, campaignId)

Disconnect stores

Disconnect the stores linked to a specific campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAccountCollectionItems

> string ExportAccountCollectionItems(ctx, collectionId)

Export account-level collection's items

Download a CSV file containing items from a given account-level collection.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAchievements

> string ExportAchievements(ctx, applicationId, campaignId, achievementId)

Export achievement customer data

Download a CSV file containing a list of all the customers who have participated in and are currently participating in the given achievement.  The CSV file contains the following columns: - `profileIntegrationID`: The integration ID of the customer profile participating in the achievement. - `title`: The display name of the achievement in the Campaign Manager. - `target`: The required number of actions or the transactional milestone to complete the achievement. - `progress`: The current progress of the customer in the achievement. - `status`: The status of the achievement. Can be one of: ['inprogress', 'completed', 'expired']. - `startDate`: The date on which the customer profile started the achievement in RFC3339. - `endDate`: The date on which the achievement ends and resets for the customer profile in RFC3339. - `completionDate`: The date on which the customer profile completed the achievement in RFC3339. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32**| The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAudiencesMemberships

> string ExportAudiencesMemberships(ctx, audienceId)

Export audience members

Download a CSV file containing the integration IDs of the members of an audience.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The file contains the following column: - `profileintegrationid`: The integration ID of the customer profile. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32**| The ID of the audience. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCampaignStores

> string ExportCampaignStores(ctx, applicationId, campaignId)

Export stores

Download a CSV file containing the stores linked to a specific campaign.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file contains the following column:  - `store_integration_id`: The identifier of the store. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCollectionItems

> string ExportCollectionItems(ctx, applicationId, campaignId, collectionId)

Export campaign-level collection's items

Download a CSV file containing items from a given campaign-level collection.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCoupons

> string ExportCoupons(ctx, applicationId, optional)

Export coupons

Download a CSV file containing the coupons that match the given properties.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file can contain the following columns:  - `accountid`: The ID of your deployment. - `applicationid`: The ID of the Application this coupon is related to. - `attributes`: A json object describing _custom_ referral attribute names and their values. - `batchid`: The ID of the batch this coupon is part of. - `campaignid`: The ID of the campaign this coupon is related to. - `counter`: The number of times this coupon has been redeemed. - `created`: The creation date in RFC3339 of the coupon code. - `deleted`: Whether the coupon code is deleted. - `deleted_changelogid`: The ID of the delete event in the logs. - `discount_counter`: The amount of discount given by this coupon. - `discount_limitval`: The maximum discount amount that can be given be this coupon. - `expirydate`: The end date in RFC3339 of the code redemption period. - `id`: The internal ID of the coupon code. - `importid`: The ID of the import job that created this coupon. - `is_reservation_mandatory`: Whether this coupon requires a reservation to be redeemed. - `limits`: The limits set on this coupon. - `limitval`: The maximum number of redemptions of this code. - `recipientintegrationid`: The integration ID of the recipient of the coupon.   Only the customer with this integration ID can redeem this code. Available only for personal codes. - `referralid`: The ID of the referral code that triggered the creation of this coupon (create coupon effect). - `reservation`: Whether the coupon can be reserved for multiple customers. - `reservation_counter`: How many times this coupon has been reserved. - `reservation_limitval`: The maximum of number of reservations this coupon can have. - `startdate`: The start date in RFC3339 of the code redemption period. - `value`: The coupon code. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ExportCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportCouponsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **optional.**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.**| Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **optional.**| Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **optional.**| Filter results by batches of coupons | 
 **exactMatch** | **optional.**| Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 
 **campaignState** | **optional.**| Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **valuesOnly** | **optional.**| Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomerSessions

> string ExportCustomerSessions(ctx, applicationId, optional)

Export customer sessions

Download a CSV file containing the customer sessions that match the request.  **Important:** Archived sessions cannot be exported. See the [retention policy](https://docs.talon.one/docs/product/server-infrastructure-and-data-retention#data-retention-policy).  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  - `id`: The internal ID of the session. - `firstsession`: Whether this is a first session. - `integrationid`: The integration ID of the session. - `applicationid`: The ID of the Application. - `profileid`: The internal ID of the customer profile. - `profileintegrationid`: The integration ID of the customer profile. - `created`: The timestamp when the session was created. - `state`: The [state](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions#customer-session-states) of the session. - `cartitems`: The cart items in the session. - `discounts`: The discounts in the session. - `total`: The total value of cart items and additional costs in the session, before any discounts are applied. - `attributes`: The attributes set in the session. - `closedat`: Timestamp when the session was closed. - `cancelledat`: Timestamp when the session was cancelled. - `referral`: The referral code in the session. - `identifiers`: The identifiers in the session. - `additional_costs`: The [additional costs](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs) in the session. - `updated`: Timestamp of the last session update. - `store_integration_id`: The integration ID of the store. - `coupons`: Coupon codes in the session. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ExportCustomerSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportCustomerSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **profileIntegrationId** | **optional.**| Only return sessions for the customer that matches this customer integration ID. | 
 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 
 **customerSessionState** | **optional.**| Filter results by state. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomersTiers

> string ExportCustomersTiers(ctx, loyaltyProgramId, optional)

Export customers' tier data

Download a CSV file containing the tier information for customers of the specified loyalty program.  The generated file contains the following columns:  - `programid`: The identifier of the loyalty program. It is displayed in your Talon.One deployment URL. - `subledgerid`: The ID of the subledger associated with the loyalty program. This column is empty if the loyalty program has no subledger. In this case, refer to the export file name to get the ID of the loyalty program. - `customerprofileid`: The ID used to integrate customer profiles with the loyalty program. - `tiername`: The name of the tier. - `startdate`: The tier start date in RFC3339. - `expirydate`: The tier expiry date in RFC3339.  You can filter the results by providing the following optional input parameters:  - `subledgerIds` (optional): Filter results by an array of subledger IDs. If no value is provided, all subledger data for the specified loyalty program will be exported. - `tierNames` (optional): Filter results by an array of tier names. If no value is provided, all tier data for the specified loyalty program will be exported. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
 **optional** | ***ExportCustomersTiersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportCustomersTiersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subledgerIds** | [**optional.Interface of []string**](string.md)| An array of subledgers IDs to filter the export by. | 
 **tierNames** | [**optional.Interface of []string**](string.md)| An array of tier names to filter the export by. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEffects

> string ExportEffects(ctx, applicationId, optional)

Export triggered effects

Download a CSV file containing the triggered effects that match the given attributes.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The generated file can contain the following columns:  - `applicationid`: The ID of the Application. - `campaignid`: The ID of the campaign. - `couponid`: The ID of the coupon, when applicable to the effect. - `created`: The timestamp of the effect. - `event_type`: The name of the event. See the [docs](https://docs.talon.one/docs/dev/concepts/entities/events). - `eventid`: The internal ID of the effect. - `name`: The effect name. See the [docs](https://docs.talon.one/docs/dev/integration-api/api-effects). - `profileintegrationid`: The ID of the customer profile, when applicable. - `props`: The [properties](https://docs.talon.one/docs/dev/integration-api/api-effects) of the effect. - `ruleindex`: The index of the rule. - `rulesetid`: The ID of the rule set. - `sessionid`: The internal ID of the session that triggered the effect. - `profileid`: The internal ID of the customer profile. - `sessionintegrationid`: The integration ID of the session. - `total_revenue`: The total revenue. - `store_integration_id`: The integration ID of the store. You choose this ID when you create a store. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ExportEffectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportEffectsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalance

> string ExportLoyaltyBalance(ctx, loyaltyProgramId, optional)

Export customer loyalty balance to CSV

⚠️ Deprecation notice: Support for requests to this endpoint will end soon. To export customer loyalty balances to CSV, use the [Export customer loyalty balances to CSV](/management-api#tag/Loyalty/operation/exportLoyaltyBalances) endpoint.  Download a CSV file containing the balance of each customer in the loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
 **optional** | ***ExportLoyaltyBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyBalanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.**| Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalances

> string ExportLoyaltyBalances(ctx, loyaltyProgramId, optional)

Export customer loyalty balances

Download a CSV file containing the balance of each customer in the loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The generated file can contain the following columns:  - `loyaltyProgramID`: The ID of the loyalty program. - `loyaltySubledger`: The name of the subdleger, when applicatble. - `profileIntegrationID`: The integration ID of the customer profile. - `currentBalance`: The current point balance. - `pendingBalance`: The number of pending points. - `expiredBalance`: The number of expired points. - `spentBalance`: The number of spent points. - `currentTier`: The tier that the customer is in at the time of the export. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
 **optional** | ***ExportLoyaltyBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyBalancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.**| Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardBalances

> string ExportLoyaltyCardBalances(ctx, loyaltyProgramId, optional)

Export all card transaction logs

Download a CSV file containing the balances of all cards in the loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file contains the following columns: - `loyaltyProgramID`: The ID of the loyalty program. - `loyaltySubledger`: The name of the subdleger, when applicatble. - `cardIdentifier`: The alphanumeric identifier of the loyalty card. - `cardState`:The state of the loyalty card. It can be `active` or `inactive`. - `currentBalance`: The current point balance. - `pendingBalance`: The number of pending points. - `expiredBalance`: The number of expired points. - `spentBalance`: The number of spent points. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***ExportLoyaltyCardBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyCardBalancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.**| Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardLedger

> string ExportLoyaltyCardLedger(ctx, loyaltyProgramId, loyaltyCardId, rangeStart, rangeEnd, optional)

Export card's ledger log

Download a CSV file containing a loyalty card ledger log of the loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **optional** | ***ExportLoyaltyCardLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyCardLedgerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCards

> string ExportLoyaltyCards(ctx, loyaltyProgramId, optional)

Export loyalty cards

Download a CSV file containing the loyalty cards from a specified loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file contains the following columns: - `identifier`: The unique identifier of the loyalty card. - `created`: The date and time the loyalty card was created. - `status`: The status of the loyalty card. - `userpercardlimit`: The maximum number of customer profiles that can be linked to the card. - `customerprofileids`: Integration IDs of the customer profiles linked to the card. - `blockreason`: The reason for transferring and blocking the loyalty card. - `generated`: An indicator of whether the loyalty card was generated. - `batchid`: The ID of the batch the loyalty card is in. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***ExportLoyaltyCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyCardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchId** | **optional.**| Filter results by loyalty card batch ID. | 
 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyLedger

> string ExportLoyaltyLedger(ctx, rangeStart, rangeEnd, loyaltyProgramId, integrationId, optional)

Export customer's transaction logs

Download a CSV file containing a customer's transaction logs in the loyalty program.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The generated file can contain the following columns:  - `customerprofileid`: The ID of the profile. - `customersessionid`: The ID of the customer session. - `rulesetid`: The ID of the rule set. - `rulename`: The name of the rule. - `programid`: The ID of the loyalty program. - `type`: The transaction type, such as `addition` or `subtraction`. - `name`: The reason for the transaction. - `subledgerid`: The ID of the subledger, when applicable. - `startdate`: The start date of the program. - `expirydate`: The expiration date of the program. - `id`: The ID of the transaction. - `created`: The timestamp of the creation of the loyalty program. - `amount`: The number of points in that transaction. - `archived`: Whether the session related to the transaction is archived. - `campaignid`: The ID of the campaign. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
**integrationId** | **string**| The identifier of the profile. | 
 **optional** | ***ExportLoyaltyLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportLoyaltyLedgerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPoolGiveaways

> string ExportPoolGiveaways(ctx, poolId, optional)

Export giveaway codes of a giveaway pool

Download a CSV file containing the giveaway codes of a specific giveaway pool.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file contains the following columns:  - `id`: The internal ID of the giveaway. - `poolid`: The internal ID of the giveaway pool. - `code`: The giveaway code. - `startdate`: The validity start date in RFC3339 of the giveaway (can be empty). - `enddate`: The validity end date in RFC3339 of the giveaway (can be empty). - `attributes`: Any custom attributes associated with the giveaway code (can be empty). - `used`: An indication of whether the giveaway is already awarded. - `importid`: The ID of the import which created the giveaway. - `created`: The creation time of the giveaway code. - `profileintegrationid`: The third-party integration ID of the customer profile that was awarded the giveaway. Can be empty if the giveaway was not awarded. - `profileid`: The internal ID of the customer profile that was awarded the giveaway. Can be empty if the giveaway was not awarded or an internal ID does not exist. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int32**| The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 
 **optional** | ***ExportPoolGiveawaysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportPoolGiveawaysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **optional.**| Timestamp that filters the results to only contain giveaways created before this date. Must be an RFC3339 timestamp string. | 
 **createdAfter** | **optional.**| Timestamp that filters the results to only contain giveaways created after this date. Must be an RFC3339 timestamp string. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportReferrals

> string ExportReferrals(ctx, applicationId, optional)

Export referrals

Download a CSV file containing the referrals that match the given parameters.  **Tip:** If the exported CSV file is too large to view, you can [split it into multiple files](https://www.makeuseof.com/tag/how-to-split-a-huge-csv-excel-workbook-into-seperate-files/).  The CSV file contains the following columns:  - `code`: The referral code. - `advocateprofileintegrationid`: The profile ID of the advocate. - `startdate`: The start date in RFC3339 of the code redemption period. - `expirydate`: The end date in RFC3339 of the code redemption period. - `limitval`: The maximum number of redemptions of this code. Defaults to `1` when left blank. - `attributes`: A json object describing _custom_ referral attribute names and their values. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ExportReferralsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportReferralsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| - &#x60;expired&#x60;: Matches referrals in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches referrals in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches referrals in which start date is set and in the future.  | 
 **usable** | **optional.**| - &#x60;true&#x60;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;, only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **batchId** | **optional.**| Filter results by batches of referrals | 
 **dateFormat** | **optional.**| Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogsWithoutTotalCount

> InlineResponse20019 GetAccessLogsWithoutTotalCount(ctx, applicationId, rangeStart, rangeEnd, optional)

Get access logs for Application

Retrieve the list of API calls sent to the specified Application. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **optional** | ***GetAccessLogsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccessLogsWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **optional.**| Only return results where the request path matches the given regular expression. | 
 **method** | **optional.**| Only return results where the request method matches the given regular expression. | 
 **status** | **optional.**| Filter results by HTTP status codes. | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId)

Get account details

Return the details of your companies Talon.One account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Return type

[**Account**](Account.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAnalytics

> AccountAnalytics GetAccountAnalytics(ctx, accountId)

Get account analytics

Return the analytics of your Talon.One account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCollection

> Collection GetAccountCollection(ctx, collectionId)

Get account-level collection

Retrieve a given account-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAchievement

> Achievement GetAchievement(ctx, applicationId, campaignId, achievementId)

Get achievement

Get the details of a specific achievement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32**| The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCost

> AccountAdditionalCost GetAdditionalCost(ctx, additionalCostId)

Get additional cost

Returns the additional cost. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32**| The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCosts

> InlineResponse20035 GetAdditionalCosts(ctx, optional)

List additional costs

Returns all the defined additional costs for the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAdditionalCostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAdditionalCostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, applicationId)

Get Application

Get the application specified by the ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Return type

[**Application**](Application.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApiHealth

> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId)

Get Application health

Display the health of the Application and show the last time the Application was used.  You can also find this information in the Campaign Manager. In your Application, click **Settings** > **Integration API Keys**. See the [docs](https://docs.talon.one/docs/dev/tutorials/monitoring-integration-status). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomer

> ApplicationCustomer GetApplicationCustomer(ctx, applicationId, customerId)

Get application's customer

Retrieve the customers of the specified application. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32**| The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomerFriends

> InlineResponse20032 GetApplicationCustomerFriends(ctx, applicationId, integrationId, optional)

List friends referred by customer profile

List the friends referred by the specified customer profile in this Application. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string**| The Integration ID of the Advocate&#39;s Profile. | 
 **optional** | ***GetApplicationCustomerFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationCustomerFriendsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomers

> InlineResponse20021 GetApplicationCustomers(ctx, applicationId, optional)

List application's customers

List all the customers of the specified application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetApplicationCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationCustomersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **optional.**| Filter results performing an exact matching against the profile integration identifier. | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomersByAttributes

> InlineResponse20022 GetApplicationCustomersByAttributes(ctx, applicationId, body, optional)

List application customers matching the given attributes

Get a list of the application customers matching the provided criteria.  The match is successful if all the attributes of the request are found in a profile, even if the profile has more attributes that are not present on the request. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md)| body | 
 **optional** | ***GetApplicationCustomersByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationCustomersByAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventTypes

> InlineResponse20028 GetApplicationEventTypes(ctx, applicationId, optional)

List Applications event types

Get all of the distinct values of the Event `type` property for events recorded in the application.  See also: [Track an event](https://docs.talon.one/integration-api#tag/Events/operation/trackEventV2) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetApplicationEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationEventTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventsWithoutTotalCount

> InlineResponse20027 GetApplicationEventsWithoutTotalCount(ctx, applicationId, optional)

List Applications events

Lists all events recorded for an application. Instead of having the total number of results in the response, this endpoint only mentions whether there are more results. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetApplicationEventsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationEventsWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **type_** | **optional.**| Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **optional.**| Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **session** | **optional.**| Session integration ID filter for events. Must be exact match. | 
 **profile** | **optional.**| Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **optional.**| Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **optional.**| Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **optional.**| Coupon code | 
 **referralCode** | **optional.**| Referral code | 
 **ruleQuery** | **optional.**| Rule name filter for events | 
 **campaignQuery** | **optional.**| Campaign name filter for events | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSession

> ApplicationSession GetApplicationSession(ctx, applicationId, sessionId)

Get Application session

Get the details of the given session. You can list the sessions with the [List Application sessions](https://docs.talon.one/management-api#tag/Customer-data/operation/getApplicationSessions) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**sessionId** | **int32**| The **internal** ID of the session. You can get the ID with the [List Application sessions](https://docs.talon.one/management-api#tag/Customer-data/operation/getApplicationSessions) endpoint.  | 

### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSessions

> InlineResponse20026 GetApplicationSessions(ctx, applicationId, optional)

List Application sessions

List all the sessions of the specified Application. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetApplicationSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **profile** | **optional.**| Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **optional.**| Filter by sessions with this state. Must be exact match. | 
 **createdBefore** | **optional.**| Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **coupon** | **optional.**| Filter by sessions with this coupon. Must be exact match. | 
 **referral** | **optional.**| Filter by sessions with this referral. Must be exact match. | 
 **integrationId** | **optional.**| Filter by sessions with this integration ID. Must be exact match. | 
 **storeIntegrationId** | **optional.**| The integration ID of the store. You choose this ID when you create a store. | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> InlineResponse2005 GetApplications(ctx, optional)

List Applications

List all applications in the current account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttribute

> Attribute GetAttribute(ctx, attributeId)

Get custom attribute

Retrieve the specified custom attribute. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32**| The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributes

> InlineResponse20033 GetAttributes(ctx, optional)

List custom attributes

Return all the custom attributes for the account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **entity** | **optional.**| Returned attributes will be filtered by supplied entity. | 

### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceMemberships

> InlineResponse20031 GetAudienceMemberships(ctx, audienceId, optional)

List audience members

Get a paginated list of the customer profiles in a given audience.  A maximum of 1000 customer profiles per page is allowed. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32**| The ID of the audience. | 
 **optional** | ***GetAudienceMembershipsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAudienceMembershipsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **profileQuery** | **optional.**| The filter to select a profile. | 

### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiences

> InlineResponse20029 GetAudiences(ctx, optional)

List audiences

Get all audiences created in the account. To create an audience, use [Create audience](https://docs.talon.one/integration-api#tag/Audiences/operation/createAudienceV2). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAudiencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAudiencesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiencesAnalytics

> InlineResponse20030 GetAudiencesAnalytics(ctx, audienceIds, optional)

List audience analytics

Get a list of audience IDs and their member count. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceIds** | **string**| The IDs of one or more audiences, separated by commas, by which to filter results. | 
 **optional** | ***GetAudiencesAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAudiencesAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, applicationId, campaignId)

Get campaign

Retrieve the given campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAnalytics

> InlineResponse20020 GetCampaignAnalytics(ctx, applicationId, campaignId, rangeStart, rangeEnd, optional)

Get analytics of campaigns

Retrieve statistical data about the performance of the given campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **optional** | ***GetCampaignAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **granularity** | **optional.**| The time interval between the results in the returned time-series. | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByAttributes

> InlineResponse2006 GetCampaignByAttributes(ctx, applicationId, body, optional)

List campaigns that match the given attributes

Get a list of all the campaigns that match a set of attributes. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | [**CampaignSearch**](CampaignSearch.md)| body | 
 **optional** | ***GetCampaignByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignByAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **campaignState** | **optional.**| Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroup

> CampaignGroup GetCampaignGroup(ctx, campaignGroupId)

Get campaign access group

Get a campaign access group specified by its ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGroupId** | **int32**| The ID of the campaign access group. | 

### Return type

[**CampaignGroup**](CampaignGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroups

> InlineResponse20011 GetCampaignGroups(ctx, optional)

List campaign access groups

List the campaign access groups in the current account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCampaignGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplates

> InlineResponse20012 GetCampaignTemplates(ctx, optional)

List campaign templates

Retrieve a list of campaign templates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCampaignTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **state** | **optional.**| Filter results by the state of the campaign template. | 
 **name** | **optional.**| Filter results performing case-insensitive matching against the name of the campaign template. | 
 **tags** | **optional.**| Filter results performing case-insensitive matching against the tags of the campaign template. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values.  | 
 **userId** | **optional.**| Filter results by user ID. | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> InlineResponse2006 GetCampaigns(ctx, applicationId, optional)

List campaigns

List the campaigns of the specified application that match your filter criteria. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **campaignState** | **optional.**| Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **name** | **optional.**| Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **optional.**| Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **campaignGroupId** | **optional.**| Filter results to campaigns owned by the specified campaign access group ID. | 
 **templateId** | **optional.**| The ID of the campaign template this campaign was created from. | 
 **storeId** | **optional.**| Filter results to campaigns linked to the specified store ID. | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> InlineResponse20041 GetChanges(ctx, optional)

Get audit logs for an account

Retrieve the audit logs displayed in **Accounts > Audit logs**. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChangesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **applicationId** | **optional.**| Filter results by Application ID. | 
 **entityPath** | **optional.**| Filter results on a case insensitive matching of the url path of the entity | 
 **userId** | **optional.**| Filter results by user ID. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **managementKeyId** | **optional.**| Filter results that match the given management key ID. | 
 **includeOld** | **optional.**| When this flag is set to false, the state without the change will not be returned. The default value is true. | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> Collection GetCollection(ctx, applicationId, campaignId, collectionId)

Get campaign-level collection

Retrieve a given campaign-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItems

> InlineResponse20018 GetCollectionItems(ctx, collectionId, optional)

Get collection items

Retrieve items from a given collection.  You can retrieve items from both account-level collections and campaign-level collections using this endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 
 **optional** | ***GetCollectionItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsWithoutTotalCount

> InlineResponse2009 GetCouponsWithoutTotalCount(ctx, applicationId, campaignId, optional)

List coupons

List all the coupons matching the specified criteria. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetCouponsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCouponsWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **optional.**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **redeemed** | **optional.**| - &#x60;true&#x60;: only coupons where &#x60;usageCounter &gt; 0&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &#x3D; 0&#x60; will be returned. - This field cannot be used in conjunction with the &#x60;usable&#x60; query parameter.  | 
 **referralId** | **optional.**| Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **optional.**| Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **optional.**| Filter results by batches of coupons | 
 **exactMatch** | **optional.**| Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **expiresBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valuesOnly** | **optional.**| Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReport

> CustomerActivityReport GetCustomerActivityReport(ctx, rangeStart, rangeEnd, applicationId, customerId, optional)

Get customer's activity report

Fetch the summary report of a given customer in the given application, in a time range.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32**| The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 
 **optional** | ***GetCustomerActivityReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerActivityReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReportsWithoutTotalCount

> InlineResponse20025 GetCustomerActivityReportsWithoutTotalCount(ctx, rangeStart, rangeEnd, applicationId, optional)

Get Activity Reports for Application Customers

Fetch summary reports for all application customers based on a time range. Instead of having the total number of results in the response, this endpoint only mentions whether there are more results. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeStart** | **time.Time**| Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**rangeEnd** | **time.Time**| Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetCustomerActivityReportsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerActivityReportsWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **name** | **optional.**| Only return reports matching the customer name. | 
 **integrationId** | **optional.**| Filter results performing an exact matching against the profile integration identifier. | 
 **campaignName** | **optional.**| Only return reports matching the campaign name. | 
 **advocateName** | **optional.**| Only return reports matching the current customer referrer name. | 

### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAnalytics

> CustomerAnalytics GetCustomerAnalytics(ctx, applicationId, customerId, optional)

Get customer's analytics report

Fetch analytics for a given customer in the given application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32**| The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 
 **optional** | ***GetCustomerAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfile

> CustomerProfile GetCustomerProfile(ctx, customerId)

Get customer profile

Return the details of the specified customer profile.  <div class=\"redoc-section\">   <p class=\"title\">Performance tips</p>    You can retrieve the same information via the Integration API, which can save you extra API requests. consider these options:    - Request the customer profile to be part of the response content using     [Update Customer Session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2).   - Send an empty update with the [Update Customer Profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint with `runRuleEngine=false`. </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32**| The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Return type

[**CustomerProfile**](CustomerProfile.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfileAchievementProgress

> InlineResponse20046 GetCustomerProfileAchievementProgress(ctx, applicationId, integrationId, optional)

List customer achievements

For the given customer profile, list all the achievements that match your filter criteria. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string**| The identifier of the profile. | 
 **optional** | ***GetCustomerProfileAchievementProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerProfileAchievementProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **achievementId** | **optional.**| The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 
 **title** | **optional.**| Filter results by the &#x60;title&#x60; of an achievement. | 

### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfiles

> InlineResponse20024 GetCustomerProfiles(ctx, optional)

List customer profiles

List all customer profiles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCustomerProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerProfilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sandbox** | **optional.**| Indicates whether you are pointing to a sandbox or live customer. | [default to false]

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersByAttributes

> InlineResponse20023 GetCustomersByAttributes(ctx, body, optional)

List customer profiles matching the given attributes

Get a list of the customer profiles matching the provided criteria.  The match is successful if all the attributes of the request are found in a profile, even if the profile has more attributes that are not present on the request. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md)| body | 
 **optional** | ***GetCustomersByAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomersByAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sandbox** | **optional.**| Indicates whether you are pointing to a sandbox or live customer. | [default to false]

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> InlineResponse20039 GetEventTypes(ctx, optional)

List event types

Fetch all event type definitions for your account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.**| Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **optional.**| Include all versions of every event type. | [default to false]
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> InlineResponse20042 GetExports(ctx, optional)

Get exports

List all past exports 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExportsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExportsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **applicationId** | **optional.**| Filter results by Application ID. | 
 **campaignId** | **optional.**| Filter by the campaign ID on which the limit counters are used. | 
 **entity** | **optional.**| The name of the entity type that was exported. | 

### Return type

[**InlineResponse20042**](InlineResponse20042.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCard

> LoyaltyCard GetLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId)

Get loyalty card

Get the given loyalty card.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactionLogs

> InlineResponse20016 GetLoyaltyCardTransactionLogs(ctx, loyaltyProgramId, loyaltyCardId, optional)

List card's transactions

Retrieve the transaction logs for the given [loyalty card](https://docs.talon.one/docs/product/loyalty-programs/card-based/card-based-overview) within the specified [card-based loyalty program](https://docs.talon.one/docs/product/loyalty-programs/overview#loyalty-program-types) with filtering options applied. If no filtering options are applied, the last 50 loyalty transactions for the given loyalty card are returned. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
 **optional** | ***GetLoyaltyCardTransactionLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyCardTransactionLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **optional.**| Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **optional.**| Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **subledgerId** | **optional.**| The ID of the subledger by which we filter the data. | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCards

> InlineResponse20015 GetLoyaltyCards(ctx, loyaltyProgramId, optional)

List loyalty cards

For the given card-based loyalty program, list the loyalty cards that match your filter criteria. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***GetLoyaltyCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyCardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **identifier** | **optional.**| The card code by which to filter loyalty cards in the response. | 
 **profileId** | **optional.**| Filter results by customer profile ID. | 
 **batchId** | **optional.**| Filter results by loyalty card batch ID. | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPoints

> LoyaltyLedger GetLoyaltyPoints(ctx, loyaltyProgramId, integrationId)

Get customer's full loyalty ledger

Get the loyalty ledger for this profile integration ID.  To get the `integrationId` of the profile from a `sessionId`, use the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint.  **Important:** To get loyalty transaction logs for a given Integration ID in a loyalty program, we recommend using the Integration API's [Get customer's loyalty logs](https://docs.talon.one/integration-api#tag/Loyalty/operation/getLoyaltyProgramProfileTransactions). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
**integrationId** | **string**| The identifier of the profile. | 

### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> LoyaltyProgram GetLoyaltyProgram(ctx, loyaltyProgramId)

Get loyalty program

Get the specified [loyalty program](https://docs.talon.one/docs/product/loyalty-programs/overview). To list all loyalty programs in your Application, use [List loyalty programs](#operation/getLoyaltyPrograms).  To list the loyalty programs that a customer profile is part of, use the [List customer data](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/getCustomerInventory) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramTransactions

> InlineResponse20014 GetLoyaltyProgramTransactions(ctx, loyaltyProgramId, optional)

List loyalty program transactions

Retrieve loyalty program transaction logs in a given loyalty program with filtering options applied. Manual and imported transactions are also included. **Note:** If no filters are applied, the last 50 loyalty transactions for the given loyalty program are returned.  **Important:** To get loyalty transaction logs for a given Integration ID in a loyalty program, we recommend using the Integration API's [Get customer's loyalty logs](https://docs.talon.one/integration-api#tag/Loyalty/operation/getLoyaltyProgramProfileTransactions). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***GetLoyaltyProgramTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyProgramTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loyaltyTransactionType** | **optional.**| Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **subledgerId** | **optional.**| The ID of the subledger by which we filter the data. | 
 **startDate** | **optional.**| Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **optional.**| Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPrograms

> InlineResponse20013 GetLoyaltyPrograms(ctx, )

List loyalty programs

List the loyalty programs of the account.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyStatistics

> LoyaltyDashboardData GetLoyaltyStatistics(ctx, loyaltyProgramId)

Get loyalty program statistics

Retrieve the statistics of the specified loyalty program such as the total active points, pending points, spent points, and expired points.  **Important:** The returned data does not include the current day. All statistics are updated daily at 11:59 PM in the loyalty program time zone. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Return type

[**LoyaltyDashboardData**](LoyaltyDashboardData.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralsWithoutTotalCount

> InlineResponse20010 GetReferralsWithoutTotalCount(ctx, applicationId, campaignId, optional)

List referrals

List all referrals of the specified campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetReferralsWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReferralsWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **code** | **optional.**| Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiration date is set and in the past. The second matches referrals in which start date is null or in the past and expiration date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **optional.**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **optional.**| Filter results by match with a profile ID specified in the referral&#39;s AdvocateProfileIntegrationId field. | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleV2

> RoleV2 GetRoleV2(ctx, roleId)

Get role

Get the details of a specific role. To see all the roles, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32**| The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 

### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleset

> Ruleset GetRuleset(ctx, applicationId, campaignId, rulesetId)

Get ruleset

Retrieve the specified ruleset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**rulesetId** | **int32**| The ID of the ruleset. | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesets

> InlineResponse2007 GetRulesets(ctx, applicationId, campaignId, optional)

List campaign rulesets

List all rulesets of this campaign. A ruleset is a revision of the rules of a campaign. **Important:** The response also includes deleted rules. You should only consider the latest revision of the returned rulesets. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***GetRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> Store GetStore(ctx, applicationId, storeId)

Get store

Get store details for a specific store ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string**| The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId)

Get user

Retrieve the data (including an invitation code) for a user. Non-admin users can only get their own profile. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse20040 GetUsers(ctx, optional)

List users in account

Retrieve all users in your account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId)

Get webhook

Returns a webhook by its id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32**| The ID of the webhook. You can find the ID in the Campaign Manager&#39;s URL when you display the details of the webhook in **Account** &gt; **Webhooks**.  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookActivationLogs

> InlineResponse20037 GetWebhookActivationLogs(ctx, optional)

List webhook activation log entries

Webhook activation log entries are created as soon as an integration request triggers a webhook effect. See the [docs](https://docs.talon.one/docs/dev/getting-started/webhooks). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookActivationLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookActivationLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **integrationRequestUuid** | **optional.**| Filter results by integration request UUID. | 
 **webhookId** | **optional.**| Filter results by webhook id. | 
 **applicationId** | **optional.**| Filter results by Application ID. | 
 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **createdBefore** | **optional.**| Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20037**](InlineResponse20037.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogs

> InlineResponse20038 GetWebhookLogs(ctx, optional)

List webhook log entries

Retrieve all webhook log entries.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **status** | **optional.**| Filter results by HTTP status codes. | 
 **webhookId** | **optional.**| Filter results by webhook id. | 
 **applicationId** | **optional.**| Filter results by Application ID. | 
 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **requestUuid** | **optional.**| Filter results by request UUID. | 
 **createdBefore** | **optional.**| Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> InlineResponse20036 GetWebhooks(ctx, optional)

List webhooks

List all webhooks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **optional.**| Checks if the given catalog or its attributes are referenced in the specified Application ID.  **Note**: If no Application ID is provided, we check for all connected Applications.  | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **creationType** | **optional.**| Filter results by creation type. | 
 **visibility** | **optional.**| Filter results by visibility. | 
 **outgoingIntegrationsTypeId** | **optional.**| Filter results by outgoing integration type ID. | 
 **title** | **optional.**| Filter results performing case-insensitive matching against the webhook title. | 

### Return type

[**InlineResponse20036**](InlineResponse20036.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccountCollection

> Import ImportAccountCollection(ctx, collectionId, optional)

Import data into existing account-level collection

Upload a CSV file containing the collection of string values that should be attached as payload for collection. The file should be sent as multipart data.  The import **replaces** the initial content of the collection.  The CSV file **must** only contain the following column:  - `item`: the values in your collection.  A collection is limited to 500,000 items.  Example:  ``` item Addidas Nike Asics ```  **Note:** Before sending a request to this endpoint, ensure the data in the CSV to import is different from the data currently stored in the collection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 
 **optional** | ***ImportAccountCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportAccountCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAllowedList

> Import ImportAllowedList(ctx, attributeId, optional)

Import allowed values for attribute

Upload a CSV file containing a list of [picklist values](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#picklist-values) for the specified attribute.  The file should be sent as multipart data.  The import **replaces** the previous list of allowed values for this attribute, if any.  The CSV file **must** only contain the following column: - `item` (required): the values in your allowed list, for example a list of SKU's.  An allowed list is limited to 500,000 items.  Example:  ```text item CS-VG-04032021-UP-50D-10 CS-DV-04042021-UP-49D-12 CS-DG-02082021-UP-50G-07 ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32**| The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 
 **optional** | ***ImportAllowedListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportAllowedListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAudiencesMemberships

> Import ImportAudiencesMemberships(ctx, audienceId, optional)

Import audience members

Upload a CSV file containing the integration IDs of the members you want to add to an audience.  The file should be sent as multipart data and should contain only the following column (required): - `profileintegrationid`: The integration ID of the customer profile.  The import **replaces** the previous list of audience members.  **Note:** We recommend limiting your file size to 500MB.  Example:  ```text profileintegrationid charles alexa ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32**| The ID of the audience. | 
 **optional** | ***ImportAudiencesMembershipsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportAudiencesMembershipsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCampaignStores

> Import ImportCampaignStores(ctx, applicationId, campaignId, optional)

Import stores

Upload a CSV file containing the stores you want to link to a specific campaign.  Send the file as multipart data.  The CSV file **must** only contain the following column: - `store_integration_id`: The identifier of the store.  The import **replaces** the previous list of stores linked to the campaign. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ImportCampaignStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportCampaignStoresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCollection

> Import ImportCollection(ctx, applicationId, campaignId, collectionId, optional)

Import data into existing campaign-level collection

Upload a CSV file containing the collection of string values that should be attached as payload for collection. The file should be sent as multipart data.  The import **replaces** the initial content of the collection.  The CSV file **must** only contain the following column:  - `item`: the values in your collection.  A collection is limited to 500,000 items.  Example:  ``` item Addidas Nike Asics ```  **Note:** Before sending a request to this endpoint, ensure the data in the CSV to import is different from the data currently stored in the collection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 
 **optional** | ***ImportCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCoupons

> Import ImportCoupons(ctx, applicationId, campaignId, optional)

Import coupons

Upload a CSV file containing the coupons that should be created. The file should be sent as multipart data.  The CSV file contains the following columns:  - `value` (required): The coupon code. - `expirydate`: The end date in RFC3339 of the code redemption period. - `startdate`: The start date in RFC3339 of the code redemption period. - `recipientintegrationid`: The integration ID of the recipient of the coupon.   Only the customer with this integration ID can redeem this code. Available only for personal codes. - `limitval`: The maximum number of redemptions of this code. For unlimited redemptions, use `0`. Defaults to `1` when not provided. - `discountlimit`: The total discount value that the code can give. This is typically used to represent a gift card value. - `attributes`: A JSON object describing _custom_ coupon attribute names and their values, enclosed with double quotation marks.    For example, if you created a [custom attribute](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes)   called `category` associated with the coupon entity, the object in the CSV file, when opened in a text editor, must be: `\"{\"category\": \"10_off\"}\"`.  You can use the time zone of your choice. It is converted to UTC internally by Talon.One.  **Note:** We recommend limiting your file size to 500MB.  **Example:**  ```text \"value\",\"expirydate\",\"startdate\",\"recipientintegrationid\",\"limitval\",\"attributes\",\"discountlimit\" COUP1,2018-07-01T04:00:00Z,2018-05-01T04:00:00Z,cust123,1,\"{\"\"Category\"\": \"\"10_off\"\"}\",2.4 ```  Once imported, you can find the `batchId` in the Campaign Manager or by using [List coupons](#tag/Coupons/operation/getCouponsWithoutTotalCount). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ImportCouponsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportCouponsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipDuplicates** | **optional.**| An indicator of whether to skip duplicate coupon values instead of causing an error. Duplicate values are ignored when &#x60;skipDuplicates&#x3D;true&#x60;.  | 
 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCards

> Import ImportLoyaltyCards(ctx, loyaltyProgramId, optional)

Import loyalty cards

Upload a CSV file containing the loyalty cards that you want to use in your card-based loyalty program. Send the file as multipart data.  It contains the following columns for each card:  - `identifier` (required): The alphanumeric identifier of the loyalty card. - `state` (required): The state of the loyalty card. It can be `active` or `inactive`. - `customerprofileids` (optional): An array of strings representing the identifiers of the customer profiles linked to the loyalty card. The identifiers should be separated with a semicolon (;).  **Note:** We recommend limiting your file size to 500MB.  **Example:**  ```csv identifier,state,customerprofileids 123-456-789AT,active,Alexa001;UserA ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***ImportLoyaltyCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportLoyaltyCardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCustomersTiers

> Import ImportLoyaltyCustomersTiers(ctx, loyaltyProgramId, optional)

Import customers into loyalty tiers

Upload a CSV file containing existing customers to be assigned to existing tiers. Send the file as multipart data.  **Important:** This endpoint only works with loyalty programs with advanced tiers (with expiration and downgrade policy) feature enabled.  The CSV file should contain the following columns: - `subledgerid` (optional): The ID of the subledger. If this field is empty, the main ledger will be used. - `customerprofileid`: The integration ID of the customer profile to whom the tier should be assigned. - `tiername`: The name of an existing tier to assign to the customer. - `expirydate`: The expiration date of the tier when the tier is reevaluated. It should be a future date.  About customer assignment to a tier: - If the customer isn't already in a tier, the customer is assigned to the specified tier during the tier import. - If the customer is already in the tier that's specified in the CSV file, only the expiration date is updated.  **Note:** We recommend not using this endpoint to update the tier of a customer. To update a customer's tier, you can [add](/management-api#tag/Loyalty/operation/addLoyaltyPoints) or [deduct](/management-api#tag/Loyalty/operation/removeLoyaltyPoints) their loyalty points.  You can use the time zone of your choice. It is converted to UTC internally by Talon.One.  **Note:** We recommend limiting your file size to 500MB.  **Example:** ```csv subledgerid,customerprofileid,tiername,expirydate SUB1,alexa,Gold,2024-03-21T07:32:14Z ,george,Silver,2025-04-16T21:12:37Z SUB2,avocado,Bronze,2026-05-03T11:47:01Z ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***ImportLoyaltyCustomersTiersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportLoyaltyCustomersTiersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyPoints

> Import ImportLoyaltyPoints(ctx, loyaltyProgramId, optional)

Import loyalty points

Upload a CSV file containing the loyalty points you want to import into a given loyalty program. Send the file as multipart data.  Depending on the type of loyalty program, you can import points into a given customer profile or loyalty card.  The CSV file contains the following columns:  - `customerprofileid` (optional): For profile-based loyalty programs, the integration ID of the customer profile where the loyalty points are imported.    **Note**: If the customer profile does not exist, it will be created. The profile will not be visible in any Application   until a session or profile update is received for that profile. - `identifier` (optional): For card-based loyalty programs, the identifier of the loyalty card where the loyalty points are imported. - `amount`: The amount of points to award to the customer profile. - `startdate` (optional): The earliest date when the points can be redeemed. The points are `active` from this date until the expiration date.    **Note**: It must be an RFC3339 timestamp string or string `immediate`. Empty or missing values are considered `immediate`. - `expirydate` (optional): The latest date when the points can be redeemed. The points are `expired` after this date.    **Note**: It must be an RFC3339 timestamp string or string `unlimited`. Empty or missing values are considered `unlimited`. - `subledgerid` (optional): The ID of the subledger that should received the points. - `reason` (optional): The reason why these points are awarded.  You can use the time zone of your choice. It is converted to UTC internally by Talon.One.  **Note:** For existing customer profiles and loyalty cards, the imported points are added to any previous active or pending points, depending on the value provided for `startdate`. If `startdate` matches the current date, the imported points are _active_. If it is later, the points are _pending_ until the date provided for `startdate` is reached.  **Note:** We recommend limiting your file size to 500MB.  **Example for profile-based programs:**  ```text customerprofileid,amount,startdate,expirydate,subledgerid,reason URNGV8294NV,100,2009-11-10T23:00:00Z,2009-11-11T23:00:00Z,subledger1,appeasement ```  **Example for card-based programs:**  ```text identifier,amount,startdate,expirydate,subledgerid,reason summer-loyalty-card-0543,100,2009-11-10T23:00:00Z,2009-11-11T23:00:00Z,subledger1,appeasement ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
 **optional** | ***ImportLoyaltyPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportLoyaltyPointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportPoolGiveaways

> Import ImportPoolGiveaways(ctx, poolId, optional)

Import giveaway codes into a giveaway pool

Upload a CSV file containing the giveaway codes that should be created. Send the file as multipart data.  The CSV file contains the following columns: - `code` (required): The code of your giveaway, for instance, a gift card redemption code. - `startdate`:  The start date in RFC3339 of the code redemption period. - `enddate`: The last date in RFC3339 of the code redemption period. - `attributes`: A JSON object describing _custom_ giveaway attribute names and their values, enclosed with double quotation marks.    For example, if you created a [custom attribute](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes)   called `provider` associated with the giveaway entity, the object in the CSV file, when opened in a text editor, must be: `\"{\"provider\": \"myPartnerCompany\"}\"`.  The `startdate` and `enddate` have nothing to do with the _validity_ of the codes. They are only used by the Rule Engine to award the codes or not. You can use the time zone setting of your choice. The values are converted to UTC internally by Talon.One.  **Note:**  - We recommend limiting your file size to 500MB. - You can import the same code multiple times. Duplicate codes are treated and distributed to customers as unique codes.  **Example:**  ```text code,startdate,enddate,attributes GIVEAWAY1,2020-11-10T23:00:00Z,2022-11-11T23:00:00Z,\"{\"\"provider\"\": \"\"Amazon\"\"}\" GIVEAWAY2,2020-11-10T23:00:00Z,2022-11-11T23:00:00Z,\"{\"\"provider\"\": \"\"Amazon\"\"}\" GIVEAWAY3,2021-01-10T23:00:00Z,2022-11-11T23:00:00Z,\"{\"\"provider\"\": \"\"Aliexpress\"\"}\" ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int32**| The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 
 **optional** | ***ImportPoolGiveawaysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportPoolGiveawaysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportReferrals

> Import ImportReferrals(ctx, applicationId, campaignId, optional)

Import referrals

Upload a CSV file containing the referrals that should be created. The file should be sent as multipart data.  The CSV file contains the following columns:  - `code` (required): The referral code. - `advocateprofileintegrationid` (required): The profile ID of the advocate. - `startdate`: The start date in RFC3339 of the code redemption period. - `expirydate`: The end date in RFC3339 of the code redemption period. - `limitval`: The maximum number of redemptions of this code. Defaults to `1` when left blank. - `attributes`: A JSON object describing _custom_ referral attribute names and their values, enclosed with double quotation marks.    For example, if you created a [custom attribute](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes)   called `category` associated with the referral entity, the object in the CSV file, when opened in a text editor, must be: `\"{\"category\": \"10_off\"}\"`.  You can use the time zone of your choice. It is converted to UTC internally by Talon.One.  **Important:** When you import a CSV file with referrals, a [customer profile](https://docs.talon.one/docs/dev/concepts/entities/customer-profiles) is **not** automatically created for each `advocateprofileintegrationid` column value. Use the [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint or the [Update multiple customer profiles](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfilesV2) endpoint to create the customer profiles.  **Note:** We recommend limiting your file size to 500MB.  **Example:**  ```text code,startdate,expirydate,advocateprofileintegrationid,limitval,attributes REFERRAL_CODE1,2020-11-10T23:00:00Z,2021-11-11T23:00:00Z,integid_4,1,\"{\"\"my_attribute\"\": \"\"10_off\"\"}\" REFERRAL_CODE2,2020-11-10T23:00:00Z,2021-11-11T23:00:00Z,integid1,1,\"{\"\"my_attribute\"\": \"\"20_off\"\"}\" ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ImportReferralsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportReferralsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **optional.**| The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserExternal

> InviteUserExternal(ctx, body)

Invite user from identity provider

[Invite a user](https://docs.talon.one/docs/product/account/account-settings/managing-users#inviting-a-user) from an external identity provider to Talon.One by sending an invitation to their email address. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewExternalInvitation**](NewExternalInvitation.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountCollections

> InlineResponse20017 ListAccountCollections(ctx, optional)

List collections in account

List account-level collections in the account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAccountCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccountCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **optional.**| Filter by collection name. | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAchievements

> InlineResponse20045 ListAchievements(ctx, applicationId, campaignId, optional)

List achievements

List all the achievements for a specific campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ListAchievementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAchievementsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **title** | **optional.**| Filter by the display name for the achievement in the campaign manager.  **Note**: If no &#x60;title&#x60; is provided, all the achievements from the campaign are returned.  | 

### Return type

[**InlineResponse20045**](InlineResponse20045.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRolesV2

> InlineResponse20043 ListAllRolesV2(ctx, )

List roles

List all roles.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20043**](InlineResponse20043.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItems

> InlineResponse20034 ListCatalogItems(ctx, catalogId, optional)

List items in a catalog

Return a paginated list of cart items in the given catalog. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int32**| The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 
 **optional** | ***ListCatalogItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCatalogItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **sku** | [**optional.Interface of []string**](string.md)| Filter results by one or more SKUs. Must be exact match. | 
 **productNames** | [**optional.Interface of []string**](string.md)| Filter results by one or more product names. Must be exact match. | 

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> InlineResponse20017 ListCollections(ctx, applicationId, campaignId, optional)

List collections in campaign

List collections in a given campaign.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ListCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **optional.**| Filter by collection name. | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsInApplication

> InlineResponse20017 ListCollectionsInApplication(ctx, applicationId, optional)

List collections in Application

List campaign-level collections from all campaigns in a given Application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ListCollectionsInApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollectionsInApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **optional.**| Filter by collection name. | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> InlineResponse20044 ListStores(ctx, applicationId, optional)

List stores

List all stores for a specific Application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
 **optional** | ***ListStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListStoresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **optional.**| When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **campaignId** | **optional.**| Filter results by campaign ID. | 
 **name** | **optional.**| The name of the store. | 
 **integrationId** | **optional.**| The integration ID of the store. | 
 **query** | **optional.**| Filter results by &#x60;name&#x60; or &#x60;integrationId&#x60;. | 

### Return type

[**InlineResponse20044**](InlineResponse20044.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationActivation

> NotificationActivation(ctx, notificationId, body)

Activate or deactivate notification

Activate or deactivate the given notification. When `enabled` is false, updates will no longer be sent for the given notification. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int32**| The ID of the notification. Get it with the appropriate _List notifications_ endpoint. | 
**body** | [**NotificationActivation**](NotificationActivation.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OktaEventHandlerChallenge

> OktaEventHandlerChallenge(ctx, )

Validate Okta API ownership

Validate the ownership of the API through a challenge-response mechanism.  This challenger endpoint is used by Okta to confirm that communication between Talon.One and Okta is correctly configured and accessible for provisioning and deprovisioning of Talon.One users, and that only Talon.One can receive and respond to events from Okta. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddedDeductedPointsNotification

> BaseNotification PostAddedDeductedPointsNotification(ctx, loyaltyProgramId, body)

Create notification about added or deducted loyalty points

Create a notification about added or deducted loyalty points in a given profile-based loyalty program. A notification for added or deducted loyalty points is different from regular webhooks in that it is loyalty program-scoped and has a predefined payload.  For more information, see [Managing loyalty notifications](https://docs.talon.one/docs/product/loyalty-programs/managing-loyalty-notifications). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**body** | [**NewBaseNotification**](NewBaseNotification.md)| body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCatalogsStrikethroughNotification

> BaseNotification PostCatalogsStrikethroughNotification(ctx, applicationId, body)

Create strikethrough notification

Create a notification for the in the given Application. For more information, see [Managing notifications](https://docs.talon.one/docs/product/applications/outbound-notifications).  See the [payload](https://docs.talon.one/outbound-notifications) you will receive. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | [**NewBaseNotification**](NewBaseNotification.md)| body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPendingPointsNotification

> BaseNotification PostPendingPointsNotification(ctx, loyaltyProgramId, body)

Create notification about pending loyalty points

Create a notification about pending loyalty points for a given profile-based loyalty program. For more information, see [Managing loyalty notifications](https://docs.talon.one/docs/product/loyalty-programs/managing-loyalty-notifications). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**body** | [**NewBaseNotification**](NewBaseNotification.md)| body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoyaltyPoints

> RemoveLoyaltyPoints(ctx, loyaltyProgramId, integrationId, body)

Deduct points from customer profile

Deduct points from the specified loyalty program and specified customer profile.  **Important:** - Only active points can be deducted. - Only pending points are rolled back when a session is cancelled or reopened.  To get the `integrationId` of the profile from a `sessionId`, use the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string**| The identifier for the loyalty program. | 
**integrationId** | **string**| The identifier of the profile. | 
**body** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> NewPassword ResetPassword(ctx, body)

Reset password

Consumes the supplied password reset token and updates the password for the associated account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewPassword**](NewPassword.md)| body | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimCreateUser

> ScimUser ScimCreateUser(ctx, body)

Create SCIM user

Create a new Talon.One user using the SCIM provisioning protocol with an identity provider, for example, Microsoft Entra ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ScimNewUser**](ScimNewUser.md)| body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimDeleteUser

> ScimDeleteUser(ctx, userId)

Delete SCIM user

Delete a specific Talon.One user created using the SCIM provisioning protocol with an identity provider, for example, Microsoft Entra ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetResourceTypes

> ScimResourceTypesListResponse ScimGetResourceTypes(ctx, )

List supported SCIM resource types

Retrieve a list of resource types supported by the SCIM provisioning protocol.  Resource types define the various kinds of resources that can be managed via the SCIM API, such as users, groups, or custom-defined resources. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ScimResourceTypesListResponse**](ScimResourceTypesListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetSchemas

> ScimSchemasListResponse ScimGetSchemas(ctx, )

List supported SCIM schemas

Retrieve a list of schemas supported by the SCIM provisioning protocol.  Schemas define the structure and attributes of the different resources that can be managed via the SCIM API, such as users, groups, and any custom-defined resources. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ScimSchemasListResponse**](ScimSchemasListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetServiceProviderConfig

> ScimServiceProviderConfigResponse ScimGetServiceProviderConfig(ctx, )

Get SCIM service provider configuration

Retrieve the configuration settings of the SCIM service provider. It provides details about the features and capabilities supported by the SCIM API, such as the different operation settings. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ScimServiceProviderConfigResponse**](ScimServiceProviderConfigResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUser

> ScimUser ScimGetUser(ctx, userId)

Get SCIM user

Retrieve data for a specific Talon.One user created using the SCIM provisioning protocol with an identity provider, for example, Microsoft Entra ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUsers

> ScimUsersListResponse ScimGetUsers(ctx, )

List SCIM users

Retrieve a paginated list of users that have been provisioned using the SCIM protocol with an identity provider, for example, Microsoft Entra ID.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ScimUsersListResponse**](ScimUsersListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimPatchUser

> ScimUser ScimPatchUser(ctx, userId, body)

Update SCIM user attributes

Update certain attributes of a specific Talon.One user created using the SCIM provisioning protocol with an identity provider, for example, Microsoft Entra ID.  This endpoint allows for selective adding, removing, or replacing specific attributes while leaving other attributes unchanged. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 
**body** | [**ScimPatchRequest**](ScimPatchRequest.md)| body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimReplaceUserAttributes

> ScimUser ScimReplaceUserAttributes(ctx, userId, body)

Update SCIM user

Update the details of a specific Talon.One user created using the SCIM provisioning protocol with an identity provider, for example, Microsoft Entra ID.  This endpoint replaces all attributes of the specific user with the attributes provided in the request payload. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 
**body** | [**ScimNewUser**](ScimNewUser.md)| body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWideWithoutTotalCount

> InlineResponse2009 SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId, body, optional)

List coupons that match the given attributes (without total count)

List the coupons whose attributes match the query criteria in all the campaigns of the given Application.  The match is successful if all the attributes of the request are found in a coupon, even if the coupon has more attributes that are not present on the request.  **Note:** The total count is not included in the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**body** | **map[string]interface{}**| body | 
 **optional** | ***SearchCouponsAdvancedApplicationWideWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchCouponsAdvancedApplicationWideWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **optional.**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.**| Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **optional.**| Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **optional.**| Filter results by batches of coupons | 
 **exactMatch** | **optional.**| Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **campaignState** | **optional.**| Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedWithoutTotalCount

> InlineResponse2009 SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId, body, optional)

List coupons that match the given attributes in campaign (without total count)

List the coupons whose attributes match the query criteria in the given campaign.  The match is successful if all the attributes of the request are found in a coupon, even if the coupon has more attributes that are not present on the request.  **Note:** The total count is not included in the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | **map[string]interface{}**| body | 
 **optional** | ***SearchCouponsAdvancedWithoutTotalCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchCouponsAdvancedWithoutTotalCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 
 **sort** | **optional.**| The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **optional.**| Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **optional.**| Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **optional.**| Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **optional.**| Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **optional.**| Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **optional.**| Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **exactMatch** | **optional.**| Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **batchId** | **optional.**| Filter results by batches of coupons | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferLoyaltyCard

> TransferLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId, body)

Transfer card data

Transfer loyalty card data, such as linked customers, loyalty balances and transactions, from a given loyalty card to a new, automatically created loyalty card.  **Important:**  - The original card is automatically blocked once the new card is created, and it cannot be activated again. - The default status of the new card is _active_. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**body** | [**TransferLoyaltyCard**](TransferLoyaltyCard.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountCollection

> Collection UpdateAccountCollection(ctx, collectionId, body)

Update account-level collection

Edit the description of a given account-level collection and enable or disable the collection in the specified Applications.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 
**body** | [**UpdateCollection**](UpdateCollection.md)| body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAchievement

> Achievement UpdateAchievement(ctx, applicationId, campaignId, achievementId, body)

Update achievement

Update the details of a specific achievement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32**| The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 
**body** | [**UpdateAchievement**](UpdateAchievement.md)| body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalCost

> AccountAdditionalCost UpdateAdditionalCost(ctx, additionalCostId, body)

Update additional cost

Updates an existing additional cost. Once created, the only property of an additional cost that cannot be changed is the `name` property (or **API name** in the Campaign Manager). This restriction is in place to prevent accidentally breaking live integrations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32**| The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 
**body** | [**NewAdditionalCost**](NewAdditionalCost.md)| body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> Attribute UpdateAttribute(ctx, attributeId, body)

Update custom attribute

Update an existing custom attribute. Once created, the only property of a custom attribute that can be changed is the description.  To change the `type` or `name` property of a custom attribute, create a new attribute and update any relevant integrations and rules to use the new attribute. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32**| The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 
**body** | [**NewAttribute**](NewAttribute.md)| body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, applicationId, campaignId, body)

Update campaign

Update the given campaign.  **Important:** You cannot use this endpoint to update campaigns if [campaign staging and revisions](https://docs.talon.one/docs/product/applications/managing-general-settings#campaign-staging-and-revisions) is enabled for your Application. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**UpdateCampaign**](UpdateCampaign.md)| body | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, applicationId, campaignId, collectionId, body)

Update campaign-level collection's description

Edit the description of a given campaign-level collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32**| The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 
**body** | [**UpdateCampaignCollection**](UpdateCampaignCollection.md)| body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> Coupon UpdateCoupon(ctx, applicationId, campaignId, couponId, body)

Update coupon

Update the specified coupon.  <div class=\"redoc-section\">   <p class=\"title\">Important</p>    <p>With this <code>PUT</code> endpoint, if you do not explicitly set a value for the <code>startDate</code>, <code>expiryDate</code>, and <code>recipientIntegrationId</code> properties in your request, it is automatically set to <code>null</code>.</p>  </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string**| The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 
**body** | [**UpdateCoupon**](UpdateCoupon.md)| body | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponBatch

> UpdateCouponBatch(ctx, applicationId, campaignId, body)

Update coupons

Update all coupons or a specific batch of coupons in the given campaign. You can find the `batchId` on the **Coupons** page of your campaign in the Campaign Manager, or you can use [List coupons](#operation/getCouponsWithoutTotalCount).  <div class=\"redoc-section\">   <p class=\"title\">Important</p>    <ul>     <li>Only send sequential requests to this endpoint.</li>     <li>Requests to this endpoint time out after 30 minutes. If you hit a timeout, contact our support team.</li>     <li>With this <code>PUT</code> endpoint, if you do not explicitly set a value for the <code>startDate</code> and <code>expiryDate</code> properties in your request, it is automatically set to <code>null</code>.</li>   </ul>  </div>  To update a specific coupon, use [Update coupon](#operation/updateCoupon). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**body** | [**UpdateCouponBatch**](UpdateCouponBatch.md)| body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoyaltyCard

> LoyaltyCard UpdateLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId, body)

Update loyalty card status

Update the status of the given loyalty card. A card can be _active_ or _inactive_.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**body** | [**UpdateLoyaltyCard**](UpdateLoyaltyCard.md)| body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferral

> Referral UpdateReferral(ctx, applicationId, campaignId, referralId, body)

Update referral

Update the specified referral.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32**| The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string**| The ID of the referral code. | 
**body** | [**UpdateReferral**](UpdateReferral.md)| body | 

### Return type

[**Referral**](Referral.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleV2

> RoleV2 UpdateRoleV2(ctx, roleId, body)

Update role

Update a specific role.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32**| The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 
**body** | **RoleV2Base**| body | 

### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStore

> Store UpdateStore(ctx, applicationId, storeId, body)

Update store

Update store details for a specific store ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32**| The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string**| The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 
**body** | [**NewStore**](NewStore.md)| body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId, body)

Update user

Update the details of a specific user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The ID of the user. | 
**body** | [**UpdateUser**](UpdateUser.md)| body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

