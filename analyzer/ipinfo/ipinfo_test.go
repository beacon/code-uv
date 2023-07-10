package ipinfo

import (
	"reflect"
	"testing"
)

func Test_IPRegexp(t *testing.T) {
	tests := []struct {
		s      string
		expect []string
	}{
		{`"192.168.0.1"`, []string{`"192.168.0.1"`, "192"}},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := ipv4Reg.FindStringSubmatch(tt.s)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Fatalf("Expect: %v Got:%v", tt.expect, got)
			}
		})
	}
}
