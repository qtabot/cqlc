package main

import (
	"fmt"
	"github.com/qtabot/gocql"
	"github.com/qtabot/cqlc/cqlc"
	"github.com/qtabot/cqlc/integration"
	"log"
	"os"
)

// Alias this to keep things short
var COUNTER = BASIC_COUNTER

func main() {
	session := integration.TestSession("127.0.0.1", "cqlc")
	integration.Truncate(session, COUNTER)

	result := "FAILED"

	ctx := cqlc.NewContext()

	err := ctx.UpdateCounter(COUNTER).
		Increment(COUNTER.COUNTER_COLUMN, 13).
		Having(COUNTER.ID.Eq("x")).
		Exec(session)

	if err != nil {
		log.Fatalf("Could not execute counter increment: %v", err)
	}

	found, counter := readCounter(session, "x")
	if found && counter == 13 {
		result = "PASSED"
	}

	c := BasicCounter{
		Id:            "x",
		CounterColumn: 11,
	}

	err = ctx.Add(COUNTER.Bind(c)).Exec(session)

	if err != nil {
		log.Fatalf("Could not execute counter increment: %v", err)
	}

	found, counter = readCounter(session, "x")
	if !found && counter != 24 {
		result = fmt.Sprintf("Expected 24, but counter was %d", counter)
	}

	os.Stdout.WriteString(result)
}

func readCounter(session *gocql.Session, key string) (bool, int64) {

	var counter int64

	ctx := cqlc.NewContext()
	found, err := ctx.Select(COUNTER.COUNTER_COLUMN).
		From(COUNTER).
		Where(COUNTER.ID.Eq(key)).
		Bind(COUNTER.COUNTER_COLUMN.To(&counter)).
		FetchOne(session)

	if err != nil {
		log.Fatalf("Could not bind data: %v", err)
		os.Exit(1)
	}

	return found, counter
}
