package servers

import (
	"log"
	"net/http"

	"github.com/dindasigma/my-playground/go-grpc/orders"
	"github.com/gin-gonic/gin"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
)

// RestServer implements a REST server for the Order Service
type RestServer struct {
	server       *http.Server
	orderService orders.OrderServiceServer // the same order service we injected into the gRPC server
	errCh        chan error
}

// NewRestServer is a convenience func to create a RestServer
func NewRestServer(orderService orders.OrderServiceServer, port string) RestServer {
	router := gin.Default()

	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
		errCh:        make(chan error),
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
func (r RestServer) Start() {
	go func() {
		r.errCh <- r.server.ListenAndServe()
	}()
}

// Stop stops the server
func (r RestServer) Stop() error {
	return r.server.Close()
}

// Error returns the server's error channel
func (r RestServer) Error() chan error {
	return r.errCh
}

func (r RestServer) create(c *gin.Context) {

	// use the order service to create the order from the req
	resp, err := r.orderService.Create(c.Request.Context(), &wrapper.StringValue{Value: "macbook pro"})
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	c.String(http.StatusOK, "ok", resp)
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
	// use the order service to create the order from the req
	resp, err := r.orderService.List(c.Request.Context(), &wrapper.StringValue{Value: "macbook pro"})
	log.Print(resp)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}

	c.String(http.StatusOK, "ok", resp)
}
