package main_6

import (
	"reflect"
	"testing"

	// "github.com/ToQoz/gopwt"
	// "github.com/ToQoz/gopwt/assert"
	"github.com/google/go-cmp/cmp"
)

/*
func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}
*/

func TestMakeGatewayInfoDeepEqual(t *testing.T) {
	got, want := MakeGatewayInfo()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MakeGatewayInfo() got = %v, want %v", got, want)
	}
}

func TestMakeGatewayInfoGoCmp(t *testing.T) {
	got, want := MakeGatewayInfo()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
	}
}

/*
func TestMakeGatewayInfoPowerAssert(t *testing.T) {
	got, want := MakeGatewayInfo()
	assert.OK(t, reflect.DeepEqual(got, want))
}
*/
