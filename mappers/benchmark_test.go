package mappers_test

import (
	"testing"

	"github.com/arran4/strings2"
	"github.com/arran4/strings2/mappers"
)

func BenchmarkAcronymify(b *testing.B) {
	subs, _ := strings2.StringToSubParts("National Aeronautics And Space Administration")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mappers.Acronymify(subs)
	}
}

func BenchmarkAcronymify_CamelCase(b *testing.B) {
	subs, _ := strings2.StringToSubParts("nationalAeronauticsAndSpaceAdministration")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mappers.Acronymify(subs)
	}
}
