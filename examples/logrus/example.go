package main

import (
	"time"

	"github.com/j0urneyk/oops"
	oopslogrus "github.com/j0urneyk/oops/loggers/logrus"
	"github.com/sirupsen/logrus"
)

// go run examples/logrus/example.go 2>&1 | jq
// go run examples/logrus/example.go 2>&1 | jq .stacktrace -r

func d() error {
	return oops.
		Code("iam_authz_missing_permission").
		In("authz").
		Time(time.Now()).
		With("user_id", 1234).
		With("permission", "post.create").
		Hint("Runbook: https://doc.acme.org/doc/abcd.md").
		User("user-123", "firstname", "john", "lastname", "doe").
		Errorf("permission denied")
}

func c() error {
	return d()
}

func b() error {
	return oops.
		In("iam").
		Trace("6710668a-2b2a-4de6-b8cf-3272a476a1c9").
		With("hello", "world").
		Wrapf(c(), "something failed")
}

func a() error {
	return b()
}

func main() {
	logrus.SetFormatter(oopslogrus.NewOopsFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	}))

	err := a()
	if err != nil {
		logrus.WithError(err).Error(err)
	}
}
