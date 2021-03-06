package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss/rialto-derivatives/models"
	"github.com/sul-dlss/rialto-derivatives/repository"
)

func TestSerializePersonResource(t *testing.T) {
	fakeSparql := new(MockedReader)

	indexer := NewPersonSerializer(repository.NewService(fakeSparql))

	resource := &models.Person{
		Firstname: "Harry",
		Lastname:  "Potter",
		URI:       "http://example.com/record1",
		DepartmentOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/department1",
			Label: "Department 1"}},
		InstitutionOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/institution1",
			Label: "Institution 1"}},
		InstituteOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/institute1",
			Label: "Institute 1"}},
		Countries: []*models.Labeled{&models.Labeled{
			URI:   "http://sws.geonames.org/1814991/",
			Label: "United States"}},
		SchoolOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/school1",
			Label: "School 1"}},
	}

	doc := indexer.Serialize(resource)

	assert.Equal(t, `{"departments":["http://example.com/department1"],"department_labels":["Department 1"],"schools":["http://example.com/school1"],"school_labels":["School 1"],"institutions":["http://example.com/institution1"],"institution_labels":["Institution 1"],"institutes":["http://example.com/institute1"],"institute_labels":["Institute 1"],"country_labels":["United States"]}`, doc)
}

func TestToSQLPersonResource(t *testing.T) {
	fakeSparql := new(MockedReader)

	indexer := NewPersonSerializer(repository.NewService(fakeSparql))

	resource := &models.Person{
		Firstname: "Harry",
		Lastname:  "Potter",
		URI:       "http://example.com/record1",
		DepartmentOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/department1",
			Label: "Department 1"}},
		InstitutionOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/institution1",
			Label: "Institution 1"}},
		InstituteOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/institute1",
			Label: "Institute 1"}},
		Countries: []*models.Labeled{&models.Labeled{
			URI:   "http://sws.geonames.org/1814991/",
			Label: "United States"}},
		SchoolOrgs: []*models.Labeled{&models.Labeled{
			URI:   "http://example.com/school1",
			Label: "School 1"}},
	}

	sql, values := indexer.SQLForInsert(resource)

	assert.Equal(t, `INSERT INTO "people" ("uri", "name", "metadata", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (uri) DO UPDATE SET name=$2, metadata=$3, updated_at=$5 WHERE people.uri=$1`, sql)
	assert.Equal(t, "Harry Potter", values[1])
	assert.Equal(t, `{"departments":["http://example.com/department1"],"department_labels":["Department 1"],"schools":["http://example.com/school1"],"school_labels":["School 1"],"institutions":["http://example.com/institution1"],"institution_labels":["Institution 1"],"institutes":["http://example.com/institute1"],"institute_labels":["Institute 1"],"country_labels":["United States"]}`, values[2])
}
