package metrics

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type AccountServiceMetrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter

	CreateAccountGrpcRequests  prometheus.Counter
	UpdateAccountGrpcRequests  prometheus.Counter
	DeleteAccountGrpcRequests  prometheus.Counter
	GetAccountByIdGrpcRequests prometheus.Counter
	SearchAccountGrpcRequests  prometheus.Counter

	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter

	CreateAccountKafkaMessages prometheus.Counter
	UpdateAccountKafkaMessages prometheus.Counter
	DeleteAccountKafkaMessages prometheus.Counter
}

func NewAccountServiceMetrics(cfg *config.Config) *AccountServiceMetrics {
	return &AccountServiceMetrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of error grpc requests",
		}),
		CreateAccountGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_account_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of create account grpc requests",
		}),
		UpdateAccountGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_account_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of update account grpc requests",
		}),
		DeleteAccountGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_account_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of delete account grpc requests",
		}),
		GetAccountByIdGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_account_by_id_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of get account by id grpc requests",
		}),
		SearchAccountGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_account_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of search account grpc requests",
		}),
		CreateAccountKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_account_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of create account kafka messages",
		}),
		UpdateAccountKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_account_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of update account kafka messages",
		}),
		DeleteAccountKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_account_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of delete account kafka messages",
		}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of error kafka processed messages",
		}),
	}
}
