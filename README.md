# Installation

```
go get github.com/ssbanjo/whop-client
```

# Introduction

With this client you can interact with the [Whop API](https://dev.whop.com/reference/home) without needing to prepare the requests yourself. It provides a clean and easy way to build client calls.

# Example

```golang
package main

import (
	"fmt"

	"github.com/ssbanjo/whop-client/client"
)

func main() {

	// Initialize the Whop client.

	apiKey := "YOUR_API_KEY"

	whopClient := client.NewClient(apiKey)


	// Validate a license key.

	licenseKey := "LICENSE_KEY"

	metadata := map[string]interface{}{
		"metadata": map[string]string{
			"hwid": "HARDWARE_ID",
		},
	}

	membership, err := whopClient.ValidateLicenseKey(
		licenseKey,
		client.ValidateLicenseKeyParams{
			Metadata: metadata,
		},
	)

	if err != nil {
		panic(fmt.Errorf("Invalid or already in use license key: %v", err))
	}

	fmt.Printf("Member %s has a valid license key!", membership.Discord.Username)


	// List your customers.

	customers, err := whopClient.ListCustomers(
		client.ListCustomersParams{
			PageNumber: 1,
			PageSize:   10,
		},
	)

	if err != nil {
		panic(fmt.Errorf("Failed to get customers: %v", err))
	}

	for _, customer := range customers.Data {

		fmt.Printf("Found customer: %s\n", customer.Username)
	}

}
```
