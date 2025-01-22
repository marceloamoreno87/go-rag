package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

type WeaviateService struct {
	client *weaviate.Client
}

func NewWeaviateService(client *weaviate.Client) *WeaviateService {
	return &WeaviateService{client: client}
}

func (s *WeaviateService) Search(ctx context.Context, prompt string) (*models.GraphQLResponse, error) {

	generatePrompt := "Questionamento do usuário: " + prompt + " Responder em português com no máximo 3 frases e conte uma piada sobre o assunto. {question}"

	gs := graphql.NewGenerativeSearch().SingleResult(generatePrompt)

	response, err := s.client.GraphQL().Get().
		WithClassName("QuestionTest").
		WithFields(
			graphql.Field{Name: "question"},
			graphql.Field{Name: "answer"},
			graphql.Field{Name: "category"}).
		WithGenerativeSearch(gs).
		WithNearText((&graphql.NearTextArgumentBuilder{}).
			WithConcepts([]string{prompt})).
		WithLimit(1).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *WeaviateService) CreateSchema(ctx context.Context) error {
	schema := &models.Class{
		Class: "QuestionTest",
		Properties: []*models.Property{
			{
				Name:     "category",
				DataType: []string{"string"},
			},
			{
				Name:     "question",
				DataType: []string{"string"},
			},
			{
				Name:     "answer",
				DataType: []string{"string"},
			},
		},
	}

	err := s.client.Schema().ClassCreator().WithClass(schema).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *WeaviateService) AddDefaultData(ctx context.Context) error {

	s.CreateSchema(ctx)

	data, err := http.DefaultClient.Get("https://raw.githubusercontent.com/weaviate-tutorials/quickstart/main/data/jeopardy_tiny.json")
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()

	var items []map[string]string
	if err := json.NewDecoder(data.Body).Decode(&items); err != nil {
		panic(err)
	}

	objects := make([]*models.Object, len(items))
	for i := range items {
		objects[i] = &models.Object{
			Class: "QuestionTest",
			Properties: map[string]any{
				"category": items[i]["Category"],
				"question": items[i]["Question"],
				"answer":   items[i]["Answer"],
			},
		}
	}

	batchRes, err := s.client.Batch().ObjectsBatcher().WithObjects(objects...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, res := range batchRes {
		if res.Result.Errors != nil {
			panic(res.Result.Errors.Error)
		}
	}

	return nil
}
