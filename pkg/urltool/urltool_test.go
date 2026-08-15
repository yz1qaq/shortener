package urltool_test

import (
	"shortener/pkg/urltool"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		LongUrl string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "基本示例", LongUrl: "https://www.liwenzhou.com//posts/go/golang-meanu?yz1=qwe", want: "golang-meanu", wantErr: false},
		{name: "无效示例", LongUrl: "/xxx/12312312", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := urltool.GetBasePath(tt.LongUrl)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBasePath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetBasePath() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
