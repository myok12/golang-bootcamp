package crawler

import "testing"

func TestCrawl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "A working site",
			args: args{"http://www.google.com"},
			wantErr: false,
			want: true,
		},
		{
			name: "A non working 404 page",
			args: args{"http://www.google.com/hmasgdjhasgdhjasgd"},
			wantErr: false,
			want: false,
		},
		{
			name: "An invalid scheme",
			args: args{"asdbdfasbdhjasb://www.google.com/"},
			wantErr: true,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := crawl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("crawl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("crawl() got = %v, want %v", got, tt.want)
			}
		})
	}
}