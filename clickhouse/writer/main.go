package main

import (
	"etl"
	"etl/clickhouse"
	"etl/clickhouse/types"
	"etl/contract"
	"flag"
	"fmt"
	"net"
	"strings"

	logger "github.com/AntonYurchenko/log-go"
	"google.golang.org/grpc"
)

const version = "v0.1.0"

// Configuration of reader.
var (
	endpoint = flag.String("endpoint", "0.0.0.0:24190", "gRPC enndpoint.")
	host     = flag.String("host", "0.0.0.0", "Connection host.")
	port     = flag.Uint("port", 8123, "Connection port.")
	user     = flag.String("user", "default", "User for connecting.")
	password = flag.String("password", "", "Password of an user.")
	workers  = flag.Uint("workers", 1, "Count of workers which will execute queries on clickhouse.")
	log      = flag.String("log", "INFO", "Level of logger.")
)

func init() {
	flag.Parse()
	logger.SetLevelStr(*log)

	logger.InfoF("Start of clickhouse gateway version %s", version)
	logger.InfoF("Endpoint of data consumer %s", *endpoint)
	logger.InfoF("Clickhouse endpoint %s:%d", *host, *port)
	logger.InfoF("Workers %d", *workers)

	// Check of arguments.
	var errorMessage string
	switch {
	case *user == "":
		errorMessage = "User should be not empty"
	}
	if errorMessage != "" {
		panic(errorMessage)
	}
}

func main() {

	// Initialisation.
	lis, err := net.Listen("tcp", *endpoint)
	if err != nil {
		panic(err)
	}

	consumer := etl.DataConsumer{
		Conn: &clickhouse.Conn{
			Address:  fmt.Sprintf("http://%s:%d", *host, *port),
			User:     *user,
			Password: *password,
		},
		Workers:   int(*workers),
		Converter: messageToQuery,
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	contract.RegisterDataConsumerServer(server, &consumer)

	// Up gRPC server.
	if err = server.Serve(lis); err != nil {
		panic(err)
	}

}

// messageToQuery creates an insert query from message.
func messageToQuery(message *contract.Message) (query etl.InsertBatch) {
	count := len(message.Batch.Values) / len(message.Batch.Names)
	logger.InfoF("%d row(s) have been read from stream", count)
	sql := createHeader(message.Target, message.Batch.Names)

	// Adding all values of batch to sql query.
	for idx, value := range message.Batch.Values {
		typeIdx := idx % len(message.Batch.Names)
		if typeIdx == 0 {
			sql += "\n"
		} else {
			sql += "\t"
		}
		strValue, err := types.FromUniversal(message.Batch.Types[typeIdx], string(value))
		if err != nil {
			panic(err)
		}
		sql += string(strValue)
	}

	return etl.InsertBatch{Query: sql, CountRows: count}
}

// createHeader creates a header of SQL insert for clickhouse.
func createHeader(target string, names []string) (header string) {
	header = fmt.Sprintf("INSERT INTO %s FORMAT TSV", target)
	if len(names) != 0 {
		header = fmt.Sprintf("INSERT INTO %s (%s) FORMAT TSV", target, strings.Join(names, ","))
	}
	return header
}