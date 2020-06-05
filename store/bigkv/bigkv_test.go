package bigkv

import (
	"io"
	"os"
	"testing"

	"github.com/dfuse-io/kvdb/store"
	"github.com/dfuse-io/kvdb/store/storetest"
	"github.com/stretchr/testify/require"
)

func init() {
	// Run tests with:
	//
	//     gcloud beta emulators bigtable start
	//

	//	if os.Getenv("DEBUG") != "" {
	//logging.Override(logging.MustCreateLoggerWithLevel("test", zap.NewAtomicLevelAt(zap.DebugLevel)))
	//	}
}

func TestAll(t *testing.T) {
	if os.Getenv("TEST_BIGKV") == "" {
		t.Skip("To run those tests, you need to have TEST_BIGKV environment variable set")
		return
	}

	storetest.TestAll(t, "bigkv", newTestFactory(t))
}

func newTestFactory(t *testing.T) storetest.DriverFactory {
	return func() (store.KVStore, storetest.DriverCleanupFunc) {
		kvStore, err := NewStore("bigkv://dev.dev/dev?createTable=true")
		if err != nil {
			t.Skip("bigtable unreachable, cannot run tests") // FIXME: this just times out
			return nil, nil
		}
		require.NoError(t, err)
		return kvStore, func() {
			kvStore.(io.Closer).Close()
		}
	}
}
