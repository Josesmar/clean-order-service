package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/josesmar/20-clean-arch/configs"
	"github.com/josesmar/20-clean-arch/internal/event/handler"
	"github.com/josesmar/20-clean-arch/internal/infra/graph"
	"github.com/josesmar/20-clean-arch/internal/infra/grpc/pb"
	"github.com/josesmar/20-clean-arch/internal/infra/grpc/service"
	"github.com/josesmar/20-clean-arch/internal/infra/web/webserver"
	"github.com/josesmar/20-clean-arch/pkg/events"
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
	listOrderUseCase := NewListOrderUseCase(db)
	getOrderUseCase := NewGetOrderUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.Router.HandleFunc("/order/{id}", webOrderHandler.Get).Methods("GET")
	webserver.Router.HandleFunc("/order", webOrderHandler.Create).Methods("POST")
	webserver.Router.HandleFunc("/orders", webOrderHandler.List).Methods("GET")

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			CreateOrderUseCase: *createOrderUseCase,
			GetOrderUseCase:    *getOrderUseCase,
			ListOrderUseCase:   *listOrderUseCase,
		}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	var ch *amqp.Channel

	for i := 0; i < 10; i++ {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				log.Println("Conectado ao RabbitMQ e canal aberto")
				return ch
			}
			log.Printf("Falha ao abrir canal: %v\n", err)
		} else {
			log.Printf("Tentativa %d: RabbitMQ não disponível, tentando novamente...\n", i+1)
		}
		time.Sleep(5 * time.Second)
	}

	panic(fmt.Sprintf("Falha ao conectar ao RabbitMQ após 10 tentativas"))
}
