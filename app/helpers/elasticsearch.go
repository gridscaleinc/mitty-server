package helpers

import (
	"context"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

// const
const (
	esURL = "http://13.113.101.120:9200"
)

// User ...
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	user := new(User)
	user.ID = 1
	user.Name = "yuan jin"
	user.Message = "dongri こんにちは、わたしはどんぐりと申します。"
	//ESAddIndex("dongri", "user", user, "1")
	//ESSearchMatchQuery()
}

// ESIndex ...
func ESIndex(esIndex string, esType string, body interface{}, esID string) error {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		// Handle error
		return err
	}

	exists, err := client.IndexExists(esIndex).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, errCreateIndex := client.CreateIndex(esIndex).Do(ctx)
		if errCreateIndex != nil {
			// Handle error
			return errCreateIndex
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Add a document to the index
	//tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index(esIndex).
		Type(esType).
		Id(esID).
		BodyJson(body).
		Refresh("true").
		Do(ctx)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

// ESDelete ...
func ESDelete(esIndex, esType, esID string) error {
	// Create a context
	ctx := context.Background()
	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		// Handle error
		return err
	}

	// Delete tweet with specified ID
	res, err := client.Delete().
		Index(esIndex).
		Type(esType).
		Id(esID).
		Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	if res.Found {
		fmt.Print("Document deleted from from index\n")
	}
	return nil
}

// ESSearchMatchQuery ...
func ESSearchMatchQuery(esIndex, esType, esSort string, from, size int, matchQuery *elastic.MatchQuery) (*elastic.SearchResult, error) {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		// Handle error
		return nil, err
	}

	// Search with a term query
	//matchQuery := elastic.NewMatchQuery("message", "申す")
	searchResult, err := client.Search().
		Index(esIndex).        // search in index "twitter"
		Query(matchQuery).     // specify the query
		Type(esType).          // search in type
		Sort(esSort, true).    // sort by "user" field, ascending
		From(from).Size(size). // take documents 0-9
		Pretty(true).          // pretty print request and response JSON
		Do(ctx)                // execute
	if err != nil {
		// Handle error
		return nil, err
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// var user User
	// for _, item := range searchResult.Each(reflect.TypeOf(user)) {
	// 	if t, ok := item.(User); ok {
	// 		fmt.Printf("User by %s: %s\n", t.Name, t.Message)
	// 	}
	// }

	return searchResult, nil
}

// ESSearchBoolQuery ...
func ESSearchBoolQuery(esIndex, esType, esSort string, from, size int, boolQuery *elastic.BoolQuery) (*elastic.SearchResult, error) {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(esURL))
	if err != nil {
		// Handle error
		return nil, err
	}

	// Search with a term query
	//matchQuery := elastic.NewMatchQuery("message", "申す")
	searchResult, err := client.Search().
		Index(esIndex).        // search in index "twitter"
		Query(boolQuery).      // specify the query
		Type(esType).          // search in type
		Sort(esSort, true).    // sort by "user" field, ascending
		From(from).Size(size). // take documents 0-9
		Pretty(true).          // pretty print request and response JSON
		Do(ctx)                // execute
	if err != nil {
		// Handle error
		return nil, err
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	return searchResult, nil
}
