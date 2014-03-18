package main

import (
	"github.com/gocql/gocql"
	"github.com/qtabot/cqlc/cqlc"
	"github.com/qtabot/cqlc/integration"
	"log"
	"os"
)

func main() {

	session := integration.TestSession("127.0.0.1", "cqlc")
	integration.Truncate(session, BASIC_CLUSTERED)

	result := "FAILED"

	rows := 10

	ctx := cqlc.NewContext()
	batch := gocql.NewBatch(gocql.LoggedBatch)

	for i := 0; i < rows; i++ {
		err := ctx.Upsert(BASIC_CLUSTERED).
			SetInt64(BASIC_CLUSTERED.ID, int64(0)).
			SetInt32(BASIC_CLUSTERED.INT32_CLUSTER, int32(i)).
			SetInt32(BASIC_CLUSTERED.INT32_COLUMN, int32(i)).
			Batch(batch)

		if err != nil {
			log.Fatalf("Could not batch upsert: %v", err)
			os.Exit(1)
		}
	}

	err := session.ExecuteBatch(batch)

	if err != nil {
		log.Fatalf("Could not execute batch: %v", err)
		os.Exit(1)
	}

	iter, err := ctx.Select().From(BASIC_CLUSTERED).Where(BASIC_CLUSTERED.ID.Eq(0)).Fetch(session)

	basics, err := BindBasicClustered(iter)
	if err != nil {
		log.Fatalf("Could not bind data: %v", err)
		os.Exit(1)
	}

	err = iter.Close()
	if err != nil {
		log.Fatalf("Could not bind data: %v", err)
		os.Exit(1)
	}

	if len(basics) == rows {
		result = "PASSED"
	}

	os.Stdout.WriteString(result)
}
