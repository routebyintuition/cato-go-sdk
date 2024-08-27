# Cato Networks Go Client SDK

This client is built to support the Cato Networks GraphQL API as outlined here:

https://www.catonetworks.com/platform/cato-api/

### Installation

```
go install github.com/routebyintuition/cato-go-sdk
```

### Client Initialization
URL: The URL to the Cato Networks GraphQL endpoint (https://api.catonetworks.com/api/v1/graphql2)
Token: Your API access token
HTTP Client: Use either the default HTTP client or leverage more advanced configuration options by passing in a client.

```
catoClient, _ := cato.New(url, token, *http.DefaultClient)
```

### Client Queries
In this example, we are looking up Entity information regarding Cato Site data from the API. We pass in the EntityID as a string slice ([]string{}) which can leverage multiple EntityIDs.

To read the returned data, we can either leverage a JSON Marshall call to read the JSON directly or perform GetItems()/GetItem() on the elements.

```
	queryResult, err := catoClient.EntityLookup(ctx, accountId, cato_models.EntityType("site"), nil, nil, nil, nil, entityIds, nil, nil, nil)
	if err != nil {
		fmt.Println("error in EntityLookup: ", err)
		os.Exit(1)
	}

	for _, val := range queryResult.EntityLookup.GetItems() {
		fmt.Println("entity item: ", *val.Entity.Name)
	}

	queryResultJson, err := json.Marshal(queryResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(queryResultJson))

```

### Client Mutations
Mustations are used in GraphQL to perform a change. This can include a create/update/delete operation.

