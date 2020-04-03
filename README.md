# go-force

[Golang](http://golang.org/) API wrapper for [Force.com](http://www.force.com/), [Salesforce.com](http://www.salesforce.com/)

## Installation

	go get github.com/nimajalali/go-force/force

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/talkinjitsu/go-force/force"
	"github.com/talkinjitsu/go-force/sobjects"
)

type SomeCustomSObject struct {
	sobjects.BaseSObject
	
	Active    bool   `force:"Active__c"`
	AccountId string `force:"Account__c"`
}

func (t *SomeCustomSObject) APIName() string {
	return "SomeCustomObject__c"
}

type SomeCustomSObjectQueryResponse struct {
	sobjects.BaseQuery

	Records []*SomeCustomSObject `force:"records"`
}

func main() {
	// Init the force
	forceAPI, err := force.Create(
		"YOUR-API-VERSION",
		"YOUR-CLIENT-ID",
		"YOUR-CLIENT-SECRET",
		"YOUR-USERNAME",
		"YOUR-PASSWORD",
		"YOUR-SECURITY-TOKEN",
		"YOUR-ENVIRONMENT",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Get somCustomSObject by ID
	someCustomSObject := &SomeCustomSObject{}
	err = forceAPI.GetSObject("Your-Object-ID", nil, someCustomSObject)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", someCustomSObject)

	// Query
	someCustomSObjects := &SomeCustomSObjectQueryResponse{}
	err = forceAPI.Query("SELECT Id FROM SomeCustomSObject__c LIMIT 10", someCustomSObjects)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", someCustomSObjects)
}
```

## Documentation

* [Package Reference](http://godoc.org/github.com/talkinjitsu/go-force/force)
* [Force.com API Reference](http://www.salesforce.com/us/developer/docs/api_rest/)
