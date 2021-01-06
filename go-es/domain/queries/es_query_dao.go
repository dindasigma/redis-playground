package queries

import (
	"github.com/olivere/elastic"
)

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQueries...)

	notEqualsQueries := make([]elastic.Query, 0)
	for _, eq := range q.NotEquals {
		notEqualsQueries = append(notEqualsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.MustNot(notEqualsQueries...)

	likeQueries := make([]elastic.Query, 0)
	for _, eq := range q.Like {
		likeQueries = append(likeQueries, elastic.NewFuzzyQuery(eq.Field, eq.Value))
	}
	query.Must(likeQueries...)

	notLikeQueries := make([]elastic.Query, 0)
	for _, eq := range q.NotLike {
		notLikeQueries = append(notLikeQueries, elastic.NewFuzzyQuery(eq.Field, eq.Value))
	}
	query.MustNot(notLikeQueries...)

	return query
}
