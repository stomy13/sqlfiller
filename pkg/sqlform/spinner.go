package sqlform

import (
	"context"
	"time"

	"github.com/charmbracelet/huh/spinner"
)

func wait() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	return spinner.New().
		Type(spinner.Line).
		Title(" Generating your sql...").
		Context(ctx).
		Run()
}
