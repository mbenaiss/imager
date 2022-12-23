package image

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/h2non/bimg"
)

// Processor is the image processor.
type Processor struct {
	httpClient *http.Client
}

// NewService returns a new instance of Service.
func NewProcessor() *Processor {
	return &Processor{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Operation is the operation to be performed on the image.
type Operation struct {
	OperationType string
	Width         int
	Height        int
	Quality       int
	Format        string
}

// ProcessFromURL processes an image from the given URL.
func (s *Processor) ProcessFromURL(ctx context.Context, url string, op Operation) ([]byte, string, error) {
	data, err := s.getOriginalImage(ctx, url)
	if err != nil {
		return nil, "", fmt.Errorf("error downloading image: %w", err)
	}

	return s.ProcessFromBuffer(data, op)
}

// ProcessFromBuffer processes an image from the given buffer.
func (s *Processor) ProcessFromBuffer(buffer []byte, op Operation) ([]byte, string, error) {
	options := bimg.Options{}

	if op.OperationType == "crop" {
		options.Crop = true
		options.Gravity = bimg.GravitySmart
	}

	if op.Width > 0 {
		options.Width = op.Width
	}

	if op.Height > 0 {
		options.Height = op.Height
	}

	if op.Format != "" {
		options.Type = getBimgType(op.Format)
	}

	if op.Quality > 0 {
		options.Quality = op.Quality
	}

	buffer, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		return nil, "", fmt.Errorf("error processing image: %w", err)
	}

	contentType := getContentType(buffer, op.Format)

	return buffer, contentType, nil
}

func (s *Processor) getOriginalImage(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error downloading image: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error downloading image: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image: %w", err)
	}

	return data, nil
}

func getContentType(data []byte, format string) string {
	contentType := http.DetectContentType(data)

	if contentType != "application/octet-stream" {
		return contentType
	}

	switch format {
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "webp":
		return "image/webp"
	case "gif":
		return "image/gif"
	case "avif":
		return "image/avif"
	default:
		return "image/jpeg"
	}
}

func getBimgType(format string) bimg.ImageType {
	switch format {
	case "jpeg":
		return bimg.JPEG
	case "png":
		return bimg.PNG
	case "webp":
		return bimg.WEBP
	case "gif":
		return bimg.GIF
	case "avif":
		return bimg.AVIF
	default:
		return bimg.JPEG
	}
}
