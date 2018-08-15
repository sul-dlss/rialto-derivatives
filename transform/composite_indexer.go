package transform

import (
	"log"

	"github.com/sul-dlss-labs/rialto-derivatives/models"
	"github.com/sul-dlss-labs/rialto-derivatives/repository"
	"github.com/vanng822/go-solr/solr"
)

// CompositeIndexer delegates to subindexers to transform resources to solr Documents
type CompositeIndexer struct {
	conceptIndexer      Indexer
	grantIndexer        Indexer
	publicationIndexer  Indexer
	personIndexer       Indexer
	organizationIndexer Indexer
	projectIndexer      Indexer
	defaultIndexer      Indexer
}

// NewCompositeIndexer creates a new CompositeIndexer instance
func NewCompositeIndexer(service *repository.Service) *CompositeIndexer {
	return &CompositeIndexer{
		conceptIndexer:      &ConceptIndexer{},
		grantIndexer:        &GrantIndexer{},
		publicationIndexer:  &PublicationIndexer{},
		personIndexer:       NewPersonIndexer(service),
		organizationIndexer: &OrganizationIndexer{},
		projectIndexer:      &ProjectIndexer{},
		defaultIndexer:      &DefaultIndexer{},
	}
}

// Map transforms a collection of resources into a collection of Solr Documents
func (m *CompositeIndexer) Map(resources []models.Resource) []solr.Document {
	docs := make([]solr.Document, len(resources))
	for i, v := range resources {
		docs[i] = m.mapOne(v)
	}
	return docs
}

// mapOne sets the id and type and then delegates to the type specific indexer
func (m *CompositeIndexer) mapOne(resource models.Resource) solr.Document {
	doc := make(solr.Document)
	doc.Set("id", resource.Subject())
	types := resource.ValueOf("type")
	if types != nil {
		doc.Set("type_ssi", types[0].String())
	} else {
		log.Printf("No resource types exist for %s", resource)
	}
	var indexer Indexer
	if resource.IsPublication() {
		indexer = m.publicationIndexer
	} else if resource.IsPerson() {
		indexer = m.personIndexer
	} else if resource.IsOrganization() {
		indexer = m.organizationIndexer
	} else if resource.IsGrant() {
		indexer = m.grantIndexer
	} else if resource.IsProject() {
		indexer = m.projectIndexer
	} else if resource.IsConcept() {
		indexer = m.conceptIndexer
	} else {
		indexer = m.defaultIndexer
	}
	return indexer.Index(resource, doc)
}
