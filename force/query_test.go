package force

import (
	"fmt"
	"testing"

	"github.com/Vivino/go-force/sobjects"
)

const (
	queryAll = "SELECT %v FROM Account WHERE Id = '%v'"
)

type AccountQueryResponse struct {
	sobjects.BaseQuery
	Records []sobjects.Account `json:"Records" force:"records"`
}

func TestQuery(t *testing.T) {
	forceAPI := createTest()
	desc, err := forceAPI.DescribeSObject(&sobjects.Account{})
	if err != nil {
		t.Fatalf("Failed to retrieve description of sobject: %v", err)
	}

	list := &AccountQueryResponse{}
	err = forceAPI.Query(buildQuery(desc.AllFields, desc.Name, nil), list)
	if err != nil {
		t.Fatalf("Failed to query: %v", err)
	}

	t.Logf("%#v", list)
}

func TestQueryAll(t *testing.T) {
	forceAPI := createTest()
	// First Insert and Delete an Account
	newID, err := insertSObject(forceAPI, t)
	if err != nil {
		t.Fatal(err)
	}
	deleteSObject(forceAPI, t, newID)

	// Then look for it.
	desc, err := forceAPI.DescribeSObject(&sobjects.Account{})
	if err != nil {
		t.Fatalf("Failed to retrieve description of sobject: %v", err)
	}

	list := &AccountQueryResponse{}
	err = forceAPI.QueryAll(fmt.Sprintf(queryAll, desc.AllFields, newID), list)
	if err != nil {
		t.Fatalf("Failed to queryAll: %v", err)
	}

	if len(list.Records) == 0 {
		t.Fatal("Failed to retrieve deleted record using queryAll")
	}

	t.Logf("%#v", list)
}

func TestQueryNext(t *testing.T) {
	// TODO
}
