package valve

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	t.Run("download dota2 replay file should be success", func(t *testing.T) {
		// this replay url can be found using OpenDota API
		// for docs how to get the replayURL using OpenDota API
		// can be read here: https://docs.opendota.com/#tag/matches%2Fpaths%2F~1matches~1%7Bmatch_id%7D%2Fget
		replayURL := "http://replay153.valve.net/570/7132230434_1635105612.dem.bz2"

		r := NewReplay(http.DefaultClient)

		ctx := context.Background()

		destination := "/tmp"

		err := r.Download(ctx, replayURL, destination)

		assert.NoError(t, err)
	})

	t.Run("download dota2 replay file should be failed due to invalid replay url", func(t *testing.T) {
		replayURL := "http://replay153.valve.net/570"

		r := NewReplay(http.DefaultClient)

		ctx := context.Background()

		destination := "/tmp"

		err := r.Download(ctx, replayURL, destination)

		assert.Error(t, err)
		assert.Equal(t, err, ErrDownloadRelpay)
	})
}

func Test_getFileName(t *testing.T) {
	type args struct {
		replayURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get file name from replay url should success",
			args: args{
				replayURL: "http://replay153.valve.net/570/7132230434_1635105612.dem.bz2",
			},
			want: "7132230434_1635105612.dem.bz2",
		},
		{
			name: "get file name from replay url should success",
			args: args{
				replayURL: "http://replay153.valve.net/570/7132230435_1635105612.dem.bz2",
			},
			want: "7132230435_1635105612.dem.bz2",
		},
		{
			name: "get file name from replay url should success",
			args: args{
				replayURL: "http://replay153.valve.net/570/7132230455_1635105612.dem.bz2",
			},
			want: "7132230455_1635105612.dem.bz2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileName(tt.args.replayURL); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}