package models

import (
	"strings"
	"testing"

	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
	"github.com/stretchr/testify/assert"
)

func TestNewPersonMinimalFields(t *testing.T) {
	data := make(map[string]rdf.Term)
	id, _ := rdf.NewIRI("http://example.com/record1")
	faculty, _ := rdf.NewIRI("http://vivoweb.org/ontology/core#FacultyMember")

	data["id"] = id
	data["subtype"] = faculty

	resource := NewPerson(data)
	assert.IsType(t, &Person{}, resource)
	assert.Equal(t, id.String(), resource.Subject())
}

func TestNewPersonAllFields(t *testing.T) {
	data := make(map[string]rdf.Term)
	id, _ := rdf.NewIRI("http://example.com/record1")
	faculty, _ := rdf.NewIRI("http://vivoweb.org/ontology/core#FacultyMember")
	organization, _ := rdf.NewIRI("http://sul.stanford.edu/rialto/agents/orgs/school-of-engineering/electrical-engineering")

	fname, _ := rdf.NewLiteral("Justin")
	lname, _ := rdf.NewLiteral("Coyne")
	data["id"] = id
	data["subtype"] = faculty
	data["lastname"] = lname
	data["firstname"] = fname
	data["organization"] = organization

	resource := NewPerson(data)
	assert.IsType(t, &Person{}, resource)
	assert.Equal(t, faculty.String(), resource.Subtype)
	assert.Equal(t, resource.Firstname, fname.String())
	assert.Equal(t, resource.Lastname, lname.String())
	assert.Equal(t, *resource.Organization, organization.String())

}

func TestSetOrganizationInfo(t *testing.T) {
	data := make(map[string]rdf.Term)
	id, _ := rdf.NewIRI("http://example.com/record1")
	faculty, _ := rdf.NewIRI("http://vivoweb.org/ontology/core#FacultyMember")
	organization, _ := rdf.NewIRI("http://sul.stanford.edu/rialto/agents/orgs/school-of-engineering/electrical-engineering")

	data["id"] = id
	data["subtype"] = faculty
	data["organization"] = organization

	resource := NewPerson(data)

	organizationJSON := strings.NewReader(`{
    "head" : {
  "vars" : [ "org", "type", "name" ]
},
"results" : {
  "bindings" : [ {
    "org" : {
      "type" : "uri",
      "value" : "http://sul.stanford.edu/rialto/agents/orgs/school-of-engineering/electrical-engineering"
    },
    "name" : {
      "type" : "literal",
      "value" : "Electrical Engineering"
    },
    "type" : {
      "type" : "uri",
      "value" : "http://vivoweb.org/ontology/core#Department"
    }
  }, {
    "org" : {
      "type" : "uri",
      "value" : "http://sul.stanford.edu/rialto/agents/orgs/stanford"
    },
    "name" : {
      "type" : "literal",
      "value" : "Stanford University"
    },
    "type" : {
      "type" : "uri",
      "value" : "http://vivoweb.org/ontology/core#University"
    }
  }, {
    "org" : {
      "type" : "uri",
      "value" : "http://sul.stanford.edu/rialto/agents/orgs/school-of-engineering"
    },
    "name" : {
      "type" : "literal",
      "value" : "School of Engineering"
    },
    "type" : {
      "type" : "uri",
      "value" : "http://vivoweb.org/ontology/core#School"
    }
  } ]
}
	  }`)
	results, _ := sparql.ParseJSON(organizationJSON)
	resource.SetOrganizationInfo(results)
	assert.IsType(t, &Person{}, resource)
	assert.Equal(t, "Electrical Engineering", *resource.DepartmentLabel)
	assert.Equal(t, "School of Engineering", *resource.SchoolLabel)
	assert.Equal(t, "Stanford University", *resource.InstitutionLabel)

}