package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/labstack/echo/v4"
)

// EchoLambda mais fácil enviar um  evento de Proxy API Gateway para o echo.Echo
// A lib transforma um evento proxy em uma requisição HTTp e então
// cria um um objeto proxy response a partir do http.ResponseWritter

type EchoLambda struct {
	// O objeto core.RequestAccessorALB é responsável por transformar um evento ALB em uma requisição HTTP
	core.RequestAccessor

	Echo *echo.Echo
}

var eLambda *EchoLambda

// New -- cria uma nova instancia do objeto echoLambdaAPI
// Recebe um objeto *echo.Echo inicializado -- normalmente criado com um echo.New()
// Então retorna a instancia inicializada do objeto EchoLambdaALB
func New(e *echo.Echo) *EchoLambda {
	return &EchoLambda{Echo: e}
}

// Proxy -- recebe um evento de próxy API Gateway e transforma em um objeto de requisição HTTP
// depois envia isso para o echo.Echo para roteamento
// Então retorna um objeto proxy response gerado a partir do http.ResponseWriter
func (e *EchoLambda) Proxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	echoRequest, err := e.ProxyEventToHTTPRequest(req)
	return e.proxyInternal(echoRequest, err)
}

// ProxyWithContext -- recebe um contexto e um evento de próxy API Gateway
// transforma eles em um objeto de requisição HTTP e envia para o echo.Echo para roteamento
// Então retorna um objeto proxy response gerado a partir do http.ResponseWriter
func (e *EchoLambda) ProxyWithContext(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	echoRequest, err := e.EventToRequestWithContext(ctx, req)
	return e.proxyInternal(echoRequest, err)
}

// proxyInternal -- função interna para lidar com a requisição e resposta
func (e *EchoLambda) proxyInternal(req *http.Request, err error) (events.APIGatewayProxyResponse, error) {
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Could not convert proxy event to request: %v", err)
	}

	respWriter := core.NewProxyResponseWriter()
	e.Echo.ServeHTTP(http.ResponseWriter(respWriter), req)

	proxyResponse, err := respWriter.GetProxyResponse()
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Error while generating proxy response: %v", err)
	}

	return proxyResponse, nil
}

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! First Echo lambda function =D")
}

func ServerlessHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if eLambda == nil {
		// stdout and stderr são enviados para AWS CloudWatch Logs
		log.Printf("Echo Lambda Cold Start")
		r := echo.New()
		r.GET("/", Handler)

		eLambda = New(r)
	}

	return eLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		// Roda no ambiente AWS Lambda
		log.Printf("[AWS LAMBDA] Echo Lambda Start")
		lambda.Start(ServerlessHandler)
	} else {
		// Roda local como uma função lambda
		log.Printf("[LOCAL] Echo Lambda Start")
		e := echo.New()
		e.GET("/", Handler)
		e.Start(":6969")
	}
}
