package loghelper

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	// make printed logs testable
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// prevent printing date that makes difficult to test
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// revert the above changes
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	cases := map[string]struct {
		before string
		after  string

		wantOutput string
	}{
		"success": {
			before:     "🐟",
			after:      "🍣",
			wantOutput: "======== 🐟\n======== 🍣\n",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// reset the output log for the next test case
			defer func() {
				buf.Reset()
			}()

			// execure both the target func and the returned func object
			resfunc := Log(tc.before, tc.after)
			resfunc()

			// evaluation
			if tc.wantOutput != buf.String() {
				t.Errorf("unexpected output, want:%s, got: %s", tc.wantOutput, buf.String())
			}
		})
	}
}
