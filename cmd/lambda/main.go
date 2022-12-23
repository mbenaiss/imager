package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/mbenaiss/imager/image"
)

func main() {
	processor := image.NewProcessor()

	lambda.Start(handler(processor))
}

func handler(svc *image.Processor) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		query := request.QueryStringParameters
		url := strings.TrimLeft(request.Path, "/")

		op := image.Operation{
			OperationType: query["o"],
			Width:         strToInt(query["w"]),
			Height:        strToInt(query["h"]),
			Quality:       strToInt(query["q"]),
			Format:        query["f"],
		}

		newImage, contentType, err := svc.ProcessFromURL(ctx, url, op)
		if err != nil {
			return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to process image: %w", err)
		}

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       base64.StdEncoding.EncodeToString(newImage),
			Headers: map[string]string{
				"Content-Type":   contentType,
				"Cache-Control":  "public, max-age=86400",
				"Content-Length": strconv.Itoa(len(newImage)),
			},
			IsBase64Encoded: true,
		}, nil
	}
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}
