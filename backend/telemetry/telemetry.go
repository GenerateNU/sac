package telemetry

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	// uncomment the following line to use Jaeger
	// "go.opentelemetry.io/otel/exporters/jaeger"
	"github.com/gofiber/contrib/otelfiber"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitTracer() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	// uncomment the following line to use Jaeger
	// exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("sac-service"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func Use(app *fiber.App) *fiber.App {
	app.Use(otelfiber.Middleware(otelfiber.WithSpanNameFormatter(func(c *fiber.Ctx) string {
		return fmt.Sprintf("%s - %s", c.Method(), c.Route().Path)
	})))

	return app
}
