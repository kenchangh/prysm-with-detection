package spectest

import (
	"path"
	"testing"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/blocks"
	"github.com/prysmaticlabs/prysm/shared/params/spectest"
	"github.com/prysmaticlabs/prysm/shared/testutil"
)

func runVoluntaryExitTest(t *testing.T, config string) {
	if err := spectest.SetConfig(config); err != nil {
		t.Fatal(err)
	}

	testFolders, testsFolderPath := testutil.TestFolders(t, config, "operations/voluntary_exit/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			exitFile, err := testutil.BazelFileBytes(folderPath, "voluntary_exit.ssz")
			if err != nil {
				t.Fatal(err)
			}
			voluntaryExit := &ethpb.SignedVoluntaryExit{}
			if err := ssz.Unmarshal(exitFile, voluntaryExit); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			body := &ethpb.BeaconBlockBody{VoluntaryExits: []*ethpb.SignedVoluntaryExit{voluntaryExit}}
			testutil.RunBlockOperationTest(t, folderPath, body, blocks.ProcessVoluntaryExits)
		})
	}
}