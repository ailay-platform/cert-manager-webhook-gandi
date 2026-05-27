package main

import (
	"os"
	"testing"

	acmetest "github.com/cert-manager/cert-manager/test/acme"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

// TestMain exits cleanly when TEST_ZONE_NAME is unset so CI can run without
// live DNS credentials or kubebuilder test assets.
func TestMain(m *testing.M) {
	if zone == "" {
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestRunsSuite(t *testing.T) {
	solver := &gandiDNSProviderSolver{}
	fixture := acmetest.NewFixture(solver,
		acmetest.SetResolvedZone(zone),
		acmetest.SetAllowAmbientCredentials(false),
		acmetest.SetManifestPath("testdata/gandi"),
	)
	//need to uncomment and  RunConformance delete runBasic and runExtended once https://github.com/cert-manager/cert-manager/pull/4835 is merged
	//fixture.RunConformance(t)
	fixture.RunBasic(t)
	fixture.RunExtended(t)

}
