package get_meta

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

func intPropertyFields(class *models.SemanticSchemaClass, property *models.SemanticSchemaClassProperty) *graphql.Object {
	getMetaIntFields := graphql.Fields{
		"sum": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sSum", class.Class, property.Name),
			Description: descriptions.GetMetaPropertySumDesc,
			Type:        graphql.Float,
		},
		"type": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sType", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyTypeDesc,
			Type:        graphql.String,
		},
		"lowest": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sLowest", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyLowestDesc,
			Type:        graphql.Float,
		},
		"highest": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sHighest", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyHighestDesc,
			Type:        graphql.Float,
		},
		"average": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sAverage", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyAverageDesc,
			Type:        graphql.Float,
		},
		"count": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sCount", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyCountDesc,
			Type:        graphql.Int,
		},
	}

	getMetaIntProperty := graphql.ObjectConfig{
		Name:        fmt.Sprintf("Meta%s%sObj", class.Class, property.Name),
		Fields:      getMetaIntFields,
		Description: descriptions.GetMetaPropertyObjectDesc,
	}

	return graphql.NewObject(getMetaIntProperty)
}