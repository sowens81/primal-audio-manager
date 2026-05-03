package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sowens81/primal-audio-manager/pkg/discogs"
)

// allow overriding in tests
var (
	exitFunc           = os.Exit
	out      io.Writer = os.Stdout
)

// HandleError prints user-friendly output and exits
func HandleError(err error) {
	if err == nil {
		return
	}

	// Discogs structured error
	if apiErr, ok := err.(discogs.DiscogsError); ok {
		fmt.Fprintf(out, "❌ Discogs API error (status %d)\n", apiErr.Status())

		if msg := apiErr.MessageText(); msg != "" {
			fmt.Fprintf(out, "Message: %s\n", msg)
		}

		if details := apiErr.Details(); details != nil {
			fmt.Fprintln(out, "Details:")
			detailJSON, _ := json.MarshalIndent(details, "  ", "  ")
			fmt.Fprintln(out, string(detailJSON))
		}

		exitFunc(1)
	}

	// fallback
	fmt.Fprintf(out, "❌ Error: %v\n", err)
	exitFunc(1)
}
