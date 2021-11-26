# Installation

To install `b1ddi-go-client` use `go get` 

```bash
go get github.com/infobloxopen/b1ddi-go-client
```

# Usage Guide

## Examples

Following program will print subnet mask of each subnet in the B1DDI API:
```go
package main

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	b1cli "github.com/infobloxopen/b1ddi-go-client/client"
)

func main() {
	// Create new go-swagger runtime client
	transport := httptransport.New(
		os.Getenv("B1DDI_HOST"), "api/ddi/v1", nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1cli.B1DDIToken(os.Getenv("B1DDI_TOKEN"))
	transport.DefaultAuthentication = tokenAuth

	// Create the BloxOne API client
	client := b1cli.NewClient(transport, strfmt.Default)

	// List all subnets using IPAM API client
	subnetList, err := client.IPAddressManagementAPI.Subnet.SubnetList(nil, nil)
	if err != nil {
		panic(err)
	}

	// Print subnet mask for each subnet
	for _, subnet := range subnetList.Payload.Results {
		fmt.Printf("%s/%d\n", *subnet.Address, subnet.Cidr)
	}
}
```
