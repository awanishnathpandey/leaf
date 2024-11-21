package routes

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterPrometheusRoute sets up the /metrics route with basic authentication
func RegisterPrometheusRoute(app *fiber.App) {

	// Fetch username and password from environment variables
	username := os.Getenv("PROMETHEUS_USERNAME")
	password := os.Getenv("PROMETHEUS_PASSWORD")

	if username == "" || password == "" {
		panic("PROMETHEUS_USERNAME or PROMETHEUS_PASSWORD is not set in .env")
	}

	// Encode the username and password into a base64 string
	authValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	// Create a Prometheus registry
	reg := prometheus.NewRegistry()

	// Optionally, register default metrics
	reg.MustRegister(collectors.NewGoCollector())
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// Create a Prometheus HTTP handler
	prometheusHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	// Add the /metrics route with basic authentication
	app.Get("/metrics", func(c *fiber.Ctx) error {
		// Validate the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader != authValue {
			c.Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		}

		// Serve Prometheus metrics
		return adaptor.HTTPHandler(prometheusHandler)(c)
	})
}
