package rest

import (
	"net/http"

	"github.com/dindasigma/my-playground/go-grpc/orders"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// RestServer implements a REST server for the Order Service
type RestServer struct {
	server       *http.Server
	orderService orders.OrderServiceServer // the same order service we injected into the gRPC server
}

var (
	router = gin.Default()
)

// NewRestServer is a convenience func to create a RestServer
func NewRestServer(orderService orders.OrderServiceServer, port string) RestServer {
	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
	}

	// register routes
	router.POST("/order", rs.create)
	router.GET("/order/:id", rs.retrieve)
	router.PUT("/order", rs.update)
	router.DELETE("/order", rs.delete)
	router.GET("/order", rs.list)

	return rs
}

// Start starts the server
func (r RestServer) Start() error {
	return r.server.ListenAndServe()
}

func (r RestServer) create(c *gin.Context) {
	var req orders.CreateOrderRequest

	// unmarshal the order request
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}

	// use the order service to create the order from the req
	resp, err := r.orderService.Create(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}

func (r RestServer) retrieve(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) list(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
