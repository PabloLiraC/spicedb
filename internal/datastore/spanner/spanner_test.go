//go:build ci
// +build ci

package spanner

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/authzed/spicedb/internal/datastore"
	"github.com/authzed/spicedb/internal/datastore/test"
	testdatastore "github.com/authzed/spicedb/internal/testserver/datastore"
)

func TestSpannerDatastore(t *testing.T) {
	b := testdatastore.RunSpannerForTesting(t, "")
	test.All(t, test.DatastoreTesterFunc(func(revisionQuantization, gcWindow time.Duration, watchBufferLength uint16) (datastore.Datastore, error) {
		ds := b.NewDatastore(t, func(engine, uri string) datastore.Datastore {
			ds, err := NewSpannerDatastore(uri, RevisionQuantization(revisionQuantization), GCWindow(gcWindow), WatchBufferLength(watchBufferLength))
			require.NoError(t, err)
			return ds
		})
		return ds, nil
	}))
}
