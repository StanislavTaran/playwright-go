package playwright_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/h2non/filetype"
	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestVideoShouldWork(t *testing.T) {
	recordVideoDir := t.TempDir()
	newContextWithOptions(t, playwright.BrowserNewContextOptions{
		RecordVideo: &playwright.BrowserNewContextRecordVideo{
			Dir: playwright.String(recordVideoDir),
		},
	})
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Reload()
	require.NoError(t, err)
	_, err = page.Reload()
	require.NoError(t, err)
	require.NoError(t, context.Close())

	files, err := ioutil.ReadDir(recordVideoDir)
	require.NoError(t, err)
	require.Equal(t, len(files), 1)
	videoFileLocation := filepath.Join(recordVideoDir, files[0].Name())
	require.Equal(t, videoFileLocation, page.Video().Path())
	content, err := ioutil.ReadFile(videoFileLocation)
	require.NoError(t, err)
	require.True(t, filetype.IsVideo(content))
}
