package image

import (
	"os"
	"testing"

	"github.com/h2non/bimg"
)

func TestProcessor_ProcessFromBuffer(t *testing.T) {
	t.Parallel()

	processor := NewProcessor()

	awsLogo, err := os.ReadFile("testdata/aws-logo.png")
	if err != nil {
		t.Fatal(err)
	}

	gotImageBytes, gotContentType, err := processor.ProcessFromBuffer(awsLogo, Operation{
		OperationType: "crop",
		Width:         100,
		Height:        100,
		Format:        "avif",
	})
	if err != nil {
		t.Fatal(err)
	}

	wantContentType := "image/avif"
	if gotContentType != wantContentType {
		t.Errorf("got %q, want %q", gotContentType, wantContentType)
	}

	image := bimg.NewImage(gotImageBytes)

	size, err := image.Size()
	if err != nil {
		t.Fatal(err)
	}

	wantWidth := 100
	if size.Width != wantWidth {
		t.Errorf("got %d, want %d", size.Width, wantWidth)
	}

	wantHeight := 100
	if size.Height != wantHeight {
		t.Errorf("got %d, want %d", size.Height, wantHeight)
	}
}
