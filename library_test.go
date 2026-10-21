package pas2

import (
	"context"
	"errors"
	"testing"
)

func TestResults(t *testing.T) {
	bad := errors.New("unhealthy")
	r, e := Run(context.Background(), map[string]Probe{"b": func(context.Context) error { return bad }, "a": func(context.Context) error { return nil }}, 2)
	if e != nil || r[0].Name != "a" || r[0].Err != nil || !errors.Is(r[1].Err, bad) {
		t.Fatal(r, e)
	}
}
func TestCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r, _ := Run(ctx, map[string]Probe{"x": func(context.Context) error { t.Fatal("must not run"); return nil }}, 1)
	if !errors.Is(r[0].Err, context.Canceled) {
		t.Fatal(r)
	}
}
func TestInvalid(t *testing.T) {
	if _, e := Run(context.Background(), nil, 0); e == nil {
		t.Fatal("invalid concurrency")
	}
}
