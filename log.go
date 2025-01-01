package datadogwriter

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

// TODO:
// add support for custom encoding options:
// datadogV2.NewSubmitLogOptionalParameters().WithContentEncoding(datadogV2.CONTENTENCODING_DEFLATE)
// datadogV2.NewSubmitLogOptionalParameters().WithContentEncoding(datadogV2.CONTENTENCODING_GZIP)

// DatadogWriter is a custom writer that sends logs to Datadog
type DatadogWriter struct {
	Service  string
	Hostname string
	Tags     string
	Source   string
	Site     string
	Encoding string
	ApiKey   string
}

func (w *DatadogWriter) Write(p []byte) (n int, err error) {
	message := string(p)

	// Construct log item
	body := []datadogV2.HTTPLogItem{
		{
			Message:  message,
			Service:  datadog.PtrString(w.Service),
			Ddsource: datadog.PtrString(w.Source),
			Ddtags:   datadog.PtrString(w.Tags),
			Hostname: datadog.PtrString(w.Hostname),
		},
	}

	// Send log to Datadog
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	// apiKey := os.Getenv("DD_API_KEY")
	if w.ApiKey == "" {
		return 0, fmt.Errorf("datadog api key is not set")
	}

	if w.Site == "" {
		w.Site = "datadoghq.com" // Default to datadoghq.com
	}
	configuration.Servers = datadog.ServerConfigurations{
		{
			URL: "https://http-intake.logs." + w.Site, // Set the site dynamically
		},
	}

	configuration.AddDefaultHeader("DD-API-KEY", apiKey)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)

	_, _, err = api.SubmitLog(ctx, body, *datadogV2.NewSubmitLogOptionalParameters())
	if err != nil {
		return 0, fmt.Errorf("failed to send log to Datadog: %w", err)
	}

	return len(p), nil
}
