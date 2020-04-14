# Migration to v2.0.0
As part of migrating our code generation to a new, more modernized and more modular one, a few _minor_ changes needed to take place in the provided API of the library.
This guide will explore these new feature and walk you through those in order to migrate any existing code from versions prior to [v2.0.0](https://github.com/talon-one/talon_go/releases/tag/v2.0.0)

# Table of Contents
1. [Configuration Initialization](#configuration-initialization)
2. [Authentication Contexts Definitions](#authentication-contexts-definitions)
3. [Models](#models)
4. [Creating, Preparing and Executing operations (API calls)](#creating-preparing-and-executing-operations-API-calls)

## Configuration Initialization

The definition of the server destination has moved to a new property and now is set in an array of [`ServerConfiguration`](https://github.com/talon-one/talon_go/blob/master/configuration.go#L75-L79):
```diff
-configuration.BasePath = "https://mycompany.talon.one"
+configuration.Servers = talon.ServerConfigurations{
+	{
+		// Notice that there is no trailing '/'
+		URL:         "http://localhost:9000",
+		Description: "Talon.One's API base URL",
+	},
+}
```

Note that:
 1. There is no trailing `/` at the end of the URL
 2. The `Description` property is optional

###### [^ Back to Top](#table-of-contents "Back to Top")

## Authentication Contexts Definitions

The definition of API keys and their corresponding Prefixes for authentication are made using the same method `context.WithValue(...)` though two changes must take place:
 1. The key of the value has changed from `talon.ContextAPIKey` to `talon.ContextAPIKeys` (plural declension)
 2. The required type of value now accepts multiple keys and therefore has been changed to a `map[string]talon.APIKey`. The default supported key is `Authorization` as the header name

```diff
-integrationContext := context.WithValue(context.Background(), talon.ContextAPIKey, talon.APIKey{
	Prefix: "ApiKey-v1",
	Key:    "fd1fd219b1e953a6b2700e8034de5bfc877462ae106127311ddd710978654312",
})
// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
+ integrationContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
+	"Authorization": talon.APIKey{
		Prefix: "ApiKey-v1",
		Key:    "fd1fd219b1e953a6b2700e8034de5bfc877462ae106127311ddd710978654312",
+	},
})
```

## Models

Models now support the getter/setter paradigm. That allows you to set their properties not only upon creation but also at later stage in the code execution flow.
As well, this paradigm helps when properties are required vs. optional - you will not be allowed to drop required ones upon creation of a struct but you will be able to set optional ones later.

### Examples
1. [`LoginParams`](https://github.com/talon-one/talon_go/blob/master/model_login_params.go#L18-L23)
As we can see in the definition below, both the `Email` and `Password` properties are required:
```go
type LoginParams struct {
	// The email address associated with your account.
	Email string `json:"email"`
	// The password for your account.
	Password string `json:"password"`
}
```

Therefore, when instantiating a new struct of the `LoginParams` type you must provide them:
```go
loginParams := talon.LoginParams{
	Email:    "admin@talon.one",
	Password: "50meSecureVeryPa$$w0rd!",
}
```

2. [`NewCustomerSession`](https://github.com/talon-one/talon_go/blob/master/model_new_customer_session.go#L18-L35)
At the below definition, we can see that none of the properties is required. Therefore you can use both struct literals combined with the property-setters to achieve the state of model you wish:
```go
type NewCustomerSession struct {
	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId *string `json:"profileId,omitempty"`
	// Any coupon code entered.
	Coupon *string `json:"coupon,omitempty"`
	// Any referral code entered.
	Referral *string `json:"referral,omitempty"`
	// Indicates the current state of the session. All sessions must start in the \"open\" state, after which valid transitions are...  1. open -> closed 2. open -> cancelled 3. closed -> cancelled
	State *string `json:"state,omitempty"`
	// Serialized JSON representation.
	CartItems *[]CartItem `json:"cartItems,omitempty"`
	// Identifiers for the customer, this can be used for limits on values such as device ID.
	Identifiers *[]string `json:"identifiers,omitempty"`
	// The total sum of the cart in one session.
	Total *float32 `json:"total,omitempty"`
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}
```

Usage example:
```go
customerSession := talon.NewCustomerSession{
	// You can use both struct literals
	ProfileId: talon.PtrString("DEADBEEF"),
	State:     talon.PtrString("open"),
}

// Or alternatively, using the relevant setter in a later stage in the code
customerSession.SetTotal(42.0)
```

The library provides you also with utility functions to streamline the required usage of pointers. You can see the provided functions in the [utils.go](https://github.com/talon-one/talon_go/blob/master/utils.go) file.

Accessing values of models could be done normally via the struct literal, but also using the getter functions which will help in case of pointers handling and value normalizations.

###### [^ Back to Top](#table-of-contents "Back to Top")

## Creating, Preparing and Executing operations (API calls)

The new generated code's approach consists of a more modular way of instantiating and executing the actual API calls.
Each operation creates a `Request` struct which abstracts the interaction with it.
These `Request` structs implements two methods: `Body` & `Execute`.
The first to set the request's body when applicable and the second to actually fire the request.
This allows you to seperate the initialization and preparation of requests from their execution.

For the most simple flow, the minimal changes would look something like this:
```diff
session, response, err := managementClient.ManagementApi.
	CreateSession(
-		context.Background(),
		talon.LoginParams{
			Email:    "admin@talon.one",
			Password: "50meSecureVeryPa$$w0rd!",
		}
-	)
// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
session, response, err := managementClient.ManagementApi.
	CreateSession(
+		context.Background()
+	).
+	Body(
		talon.LoginParams{
			Email:    "admin@talon.one",
			Password: "50meSecureVeryPa$$w0rd!",
		}
+	).
+	Execute()
```

Please note that there are a few operations that support direct setting of URL params on the request object itself, such as [`GetCustomerInventory`](https://github.com/talon-one/talon_go/blob/8868f0141255093a7376118a5813c99f241bc9e0/api_integration.go#L574-L580). For those you can use the direct setter methods which are defined on the request struct:
```go
customerInventory, response, err := integrationClient.IntegrationApi.
 	GetCustomerInventory(integrationAuthContext, "Unique_Profile_ID").
 	Profile(true). // setting the `profile` url param to true directly on the request
 	Execute()
```

As can be seen, all requests methods support piping (return the Request struct) for streamlining of usage.

###### [^ Back to Top](#table-of-contents "Back to Top")
