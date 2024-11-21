package prometheus

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	prometheusclient "github.com/prometheus/client_golang/prometheus"
)

const (
	existStatusFailure = "failure"
	exitStatusSuccess  = "success"
)

var (
	requestStartedCounter    prometheusclient.Counter
	requestCompletedCounter  prometheusclient.Counter
	resolverStartedCounter   *prometheusclient.CounterVec
	resolverCompletedCounter *prometheusclient.CounterVec
	timeToResolveField       *prometheusclient.HistogramVec
	timeToHandleRequest      *prometheusclient.HistogramVec
)

type Tracer struct{}

var _ graphql.HandlerExtension = Tracer{}
var _ graphql.OperationInterceptor = Tracer{}
var _ graphql.ResponseInterceptor = Tracer{}
var _ graphql.FieldInterceptor = Tracer{}

// Register all the Prometheus metrics
func Register() {
	RegisterOn(prometheusclient.DefaultRegisterer)
}

// RegisterOn allows you to register the metrics with a custom Prometheus registerer
func RegisterOn(registerer prometheusclient.Registerer) {
	requestStartedCounter = prometheusclient.NewCounter(
		prometheusclient.CounterOpts{
			Name: "graphql_request_started_total",
			Help: "Total number of requests started on the graphql server.",
		},
	)

	requestCompletedCounter = prometheusclient.NewCounter(
		prometheusclient.CounterOpts{
			Name: "graphql_request_completed_total",
			Help: "Total number of requests completed on the graphql server.",
		},
	)

	resolverStartedCounter = prometheusclient.NewCounterVec(
		prometheusclient.CounterOpts{
			Name: "graphql_resolver_started_total",
			Help: "Total number of resolver started on the graphql server.",
		},
		[]string{"object", "field"},
	)

	resolverCompletedCounter = prometheusclient.NewCounterVec(
		prometheusclient.CounterOpts{
			Name: "graphql_resolver_completed_total",
			Help: "Total number of resolver completed on the graphql server.",
		},
		[]string{"object", "field"},
	)

	timeToResolveField = prometheusclient.NewHistogramVec(prometheusclient.HistogramOpts{
		Name:    "graphql_resolver_duration_ms",
		Help:    "The time taken to resolve a field by graphql server.",
		Buckets: prometheusclient.ExponentialBuckets(1, 2, 11),
	}, []string{"exitStatus", "object", "field"})

	timeToHandleRequest = prometheusclient.NewHistogramVec(prometheusclient.HistogramOpts{
		Name:    "graphql_request_duration_ms",
		Help:    "The time taken to handle a request by graphql server.",
		Buckets: prometheusclient.ExponentialBuckets(1, 2, 11),
	}, []string{"exitStatus"})

	// Register the metrics
	registerer.MustRegister(
		requestStartedCounter,
		requestCompletedCounter,
		resolverStartedCounter,
		resolverCompletedCounter,
		timeToResolveField,
		timeToHandleRequest,
	)
}

// UnRegister removes all the metrics from the default registerer
func UnRegister() {
	UnRegisterFrom(prometheusclient.DefaultRegisterer)
}

// UnRegisterFrom removes all the metrics from a custom Prometheus registerer
func UnRegisterFrom(registerer prometheusclient.Registerer) {
	registerer.Unregister(requestStartedCounter)
	registerer.Unregister(requestCompletedCounter)
	registerer.Unregister(resolverStartedCounter)
	registerer.Unregister(resolverCompletedCounter)
	registerer.Unregister(timeToResolveField)
	registerer.Unregister(timeToHandleRequest)
}

// NewTracer creates a new instance of the Prometheus Tracer for GraphQL
func NewTracer() graphql.HandlerExtension {
	return Tracer{}
}

// ExtensionName returns the name of the extension
func (a Tracer) ExtensionName() string {
	return "Prometheus"
}

// Validate is called once when the schema is validated
func (a Tracer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// InterceptOperation intercepts the operation request and starts the request counter
func (a Tracer) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	requestStartedCounter.Inc() // This increments the "graphql_request_started_total" counter
	return next(ctx)
}

// InterceptResponse intercepts the response and measures the duration
func (a Tracer) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	errList := graphql.GetErrors(ctx)

	var exitStatus string
	if len(errList) > 0 {
		exitStatus = existStatusFailure
	} else {
		exitStatus = exitStatusSuccess
	}

	// Measure request handling time
	oc := graphql.GetOperationContext(ctx)
	observerStart := oc.Stats.OperationStart
	timeToHandleRequest.With(prometheusclient.Labels{"exitStatus": exitStatus}).Observe(float64(time.Since(observerStart).Nanoseconds() / int64(time.Millisecond)))

	requestCompletedCounter.Inc() // This increments the "graphql_request_completed_total" counter
	return next(ctx)
}

// InterceptField intercepts the field execution, measures the time, and increments the counter
func (a Tracer) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)

	// Start resolver
	resolverStartedCounter.WithLabelValues(fc.Object, fc.Field.Name).Inc()

	observerStart := time.Now()

	// Execute the resolver
	res, err := next(ctx)

	// Measure field resolution time
	var exitStatus string
	if err != nil {
		exitStatus = existStatusFailure
	} else {
		exitStatus = exitStatusSuccess
	}

	timeToResolveField.WithLabelValues(exitStatus, fc.Object, fc.Field.Name).
		Observe(float64(time.Since(observerStart).Nanoseconds() / int64(time.Millisecond)))

	// Complete resolver
	resolverCompletedCounter.WithLabelValues(fc.Object, fc.Field.Name).Inc()

	return res, err
}
