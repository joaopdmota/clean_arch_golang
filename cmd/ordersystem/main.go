package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/20-CleanArch/configs"
	"github.com/devfullcycle/20-CleanArch/internal/event/handler"
	"github.com/devfullcycle/20-CleanArch/internal/infra/graph"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/service"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web/webserver"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getOrdersUseCase := NewGetOrdersUseCase(db)

	runWebServer(configs.WebServerPort, db, eventDispatcher)
	rungRPC(configs.GRPCServerPort, createOrderUseCase, getOrdersUseCase)
	runGraphql(configs.GraphQLServerPort, createOrderUseCase, getOrdersUseCase)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func runWebServer(webServerPort string, db *sql.DB, eventDispatcher *events.EventDispatcher) {
	webserver := webserver.NewWebServer(webServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler(http.MethodPost, "/order", webOrderHandler.Create)
	webserver.AddHandler(http.MethodGet, "/orders", webOrderHandler.GetAll)

	fmt.Println("Starting web server on port", webServerPort)
	go webserver.Start()
}

func rungRPC(gRPCServerPort string, createOrderUseCase *usecase.CreateOrderUseCase, getOrdersUseCase *usecase.GetOrdersUseCase) {
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *getOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", gRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)
}

func runGraphql(graphQLServerPort string, createOrderUseCase *usecase.CreateOrderUseCase, getOrdersUseCase *usecase.GetOrdersUseCase) {
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		GetOrdersUseCase:   *getOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", graphQLServerPort)
	http.ListenAndServe(":"+graphQLServerPort, nil)
}
