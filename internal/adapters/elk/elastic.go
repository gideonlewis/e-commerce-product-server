package elk

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
	"github.com/olivere/elastic/v7"
)

type ElasticsearchRepository struct {
	client *elastic.Client
}

func NewElasticsearchRepository(url string) (ports.ProductRepository, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &ElasticsearchRepository{client: client}, nil
}

func (repo *ElasticsearchRepository) IndexProduct(ctx context.Context, product domain.Product) error {
	_, err := repo.client.Index().
		Index("products").
		Id(string(product.ID)).
		BodyJson(product).
		Do(ctx)
	return err
}

func (repo *ElasticsearchRepository) SearchProducts(ctx context.Context, query string) ([]domain.Product, error) {
	result, err := repo.client.Search().
		Index("products").
		Query(elastic.NewMultiMatchQuery(query, "name", "description")).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, hit := range result.Hits.Hits {
		var product domain.Product
		if err := json.Unmarshal(hit.Source, &product); err != nil {
			log.Println("Error parsing product: ", err)
			continue
		}
		products = append(products, product)
	}
	return products, nil
}
