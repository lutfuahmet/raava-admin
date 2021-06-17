package admin

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetDepVersionFromMod(t *testing.T) {
	goModDeps = []string{
		"github.com/qor/l10n v0.0.0-20181031091737-2ca95fb3b4dd // indirect",
		"github.com/qor/admin v1.1.1",
		"github.com/qor/publish2 v1.1.0 // indirect",
		"github.com/qor/media v0.0.0-20191022071353-19cf289e17d4 // indirect",
		"github.com/qor/i18n v2.0.7",
	}
	cases := []struct {
		view string
		want string
	}{
		{view: "github.com/qor/l10n/views", want: "pkg/mod/github.com/qor/l10n@v0.0.0-20181031091737-2ca95fb3b4dd/views"},
		{view: "github.com/qor/admin/views", want: "pkg/mod/github.com/TykTechnologies/raava-admin@v1.1.1/views"},
		{view: "github.com/qor/publish2/views", want: "pkg/mod/github.com/qor/publish2@v1.1.0/views"},
		{view: "github.com/qor/media/media_library/views", want: "pkg/mod/github.com/qor/media@v0.0.0-20191022071353-19cf289e17d4/media_library/views"},
		{view: "github.com/qor/i18n/exchange_actions/views", want: "pkg/mod/github.com/qor/i18n@v2.0.7/exchange_actions/views"},
	}
	for _, v := range cases {
		pth := strings.TrimSuffix(v.view, "/views")
		if got := getDepVersionFromMod(pth) + "/views"; v.want != got {
			t.Errorf("GetDepVersionFromMod-viewpath: %v, want: %v, got: %v", v.view, v.want, got)
		} else {
			fmt.Println(got)
		}
	}
}
