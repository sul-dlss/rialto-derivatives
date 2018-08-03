package models

// Predicates stores the predicates we use in rialto
var Predicates = map[string]map[string]string{
	"rdf": map[string]string{"type": "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"},
	"bibo": map[string]string{
		"abstract":   "http://purl.org/ontology/bibo/abstract",
		"cites":      "http://purl.org/ontology/bibo/abstract",
		"doi":        "http://purl.org/ontology/bibo/doi",
		"identifier": "http://purl.org/ontology/bibo/identifier",
		"uri":        "http://purl.org/ontology/bibo/uri",
	},
	"dct": map[string]string{
		"title":       "http://purl.org/dc/terms/title",
		"created":     "http://purl.org/dc/terms/created",
		"hasPart":     "http://purl.org/dc/terms/hasPart",
		"subject":     "http://purl.org/dc/terms/subject",
		"alternative": "http://purl.org/dc/terms/alternative",
	},
	"vivo": map[string]string{
		"description":                    "http://vivoweb.org/ontology/core#description",
		"hasFundingVehicle":              "http://vivoweb.org/ontology/core#hasFundingVehicle",
		"publisher":                      "http://vivoweb.org/ontology/core#publisher",
		"informationResourceSupportedBy": "http://vivoweb.org/ontology/core#informationResourceSupportedBy",
	},
	"gcis": map[string]string{
		"hasInstrument": "http://data.globalchange.gov/gcis.owl#hasInstrument",
	},
	"owl": map[string]string{
		"sameAs": "http://www.w3.org/2002/07/owl#sameAs",
	},
}

var publicationTypes = []string{
	"http://purl.org/ontology/bibo/Document",
	"http://vivoweb.org/ontology/core#Abstract",
	"http://purl.org/ontology/bibo/Article",
	"http://purl.org/ontology/bibo/Book",
	"http://vivoweb.org/ontology/core#CaseStudy",
	"http://vivoweb.org/ontology/core#Catalog",
	"http://purl.org/spar/fabio/ClinicalGuideline",
	"http://vivoweb.org/ontology/core#ConferencePoster",
	"http://purl.org/ontology/bibo/Manual",
	"http://purl.org/ontology/bibo/Manuscript",
	"http://purl.org/ontology/bibo/Patent",
	"http://purl.org/ontology/bibo/Report",
	"http://vivoweb.org/ontology/core#ResearchProposal",
	"http://vivoweb.org/ontology/core#Score",
	"http://vivoweb.org/ontology/core#Screenplay",
	"http://purl.org/ontology/bibo/Slideshow",
	"http://vivoweb.org/ontology/core#Speech",
	"http://purl.org/ontology/bibo/Standard",
	"http://purl.org/ontology/bibo/Thesis",
	"http://vivoweb.org/ontology/core#Translation",
	"http://purl.org/ontology/bibo/Webpage",
	"http://vivoweb.org/ontology/core#WorkingPaper",
}