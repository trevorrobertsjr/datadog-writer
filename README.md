# datadog-writer
A module that enables go programs to write observability data directly to DataDog endpoints rather than depending on the DataDog Agent

## Using this module

```go
package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/trevorrobertsjr/datadogwriter"
)

func main() {

	// Configure logger
	datadogWriter := &datadogwriter.DatadogWriter{
		Service:  "your-service",
		Hostname: "your-server",
		Tags:     "env:staging,version:1.0",
		Source:   "your-source",
	}
	log.Logger = zerolog.New(io.MultiWriter(os.Stdout, datadogWriter)).With().Timestamp().Logger()

	log.Info().Msg("Reading config...")

}
```

Your program needs to provide your API key and assigned datadog site as environment variables. If your datadog site is the default (datadoghq.com), you do not need to set the `DD_SITE` environment variable.

```bash
DD_API_KEY=YourApiKey DD_SITE=us5.datadoghq.com go run main.go
```