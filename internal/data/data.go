package data

import (
	"review-job/internal/conf"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewReviewRepo, NewESClient, NewKafkaReader)

// Data .
type Data struct {
	kafkaReader *kafka.Reader // kafka reader
	esClient    *ESClient     // ES client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, es *ESClient, kafka *kafka.Reader) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources...")
		// todo 关闭kakfaReader
		log.NewHelper(logger).Info("closing the kafka.Reader")
		kafka.Close()
	}
	return &Data{
		kafkaReader: kafka,
		esClient:    es,
	}, cleanup, nil
}

func NewKafkaReader(cfg *conf.Kafka) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		GroupID: cfg.GroupId, // 指定消费者组id
		Topic:   cfg.Topic,
	})
}

type ESClient struct {
	*elasticsearch.TypedClient
	Idx string // 索引库
}

func NewESClient(cfg *conf.Elasticsearch) (*ESClient, error) {
	// ES 配置
	c := elasticsearch.Config{
		Addresses: cfg.Addresses,
	}

	// 创建客户端连接
	client, err := elasticsearch.NewTypedClient(c)
	if err != nil {
		return nil, err
	}
	return &ESClient{
		TypedClient: client,
		Idx:         cfg.Index,
	}, nil
}
