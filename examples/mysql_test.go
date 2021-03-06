package example

import (
	"github.com/go-sql-driver/mysql"

	"github.com/goph/logur"
)

func Example_mysqlDriver() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	_ = mysql.SetLogger(logur.NewErrorPrintLogger(logger))

	// Output:
}
