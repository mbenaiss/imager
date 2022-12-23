## Welcome to the Imager project!

Imager is a serverless application that allows you to process images on the fly using AWS Lambda and Amazon API Gateway. With Imager, you can easily crop, resize, and adjust the quality of images by simply providing a URL to the image.

## Usage

To use Imager, you can send a GET request to the API endpoint with the following parameters:

`https://d34bfvhzd89jzk.cloudfront.net/{url}?o=crop&w=100&h=100&q=80&f=webp`

- `url`: the url of the image to process (e.g. https://images.unsplash.com/photo-1537151625747-768eb6cf92b2)
- `o`: Operation to perform on the image (e.g. crop, resize)
- `w`: Width of the image (e.g. 100)
- `h`: Height of the image (e.g. 100)
- `q`: Quality of the image (e.g. 80)
- `f`: Format of the image (e.g. webp, png, jpeg, avif, gif)

## Examples

Here is an example of how you can use Imager to crop an image:

| Original image                                                            | Size    | Result                                                                                                                                     | Size   | Description               |
| ------------------------------------------------------------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------ | ------ | ------------------------- |
| ![unsplash](https://images.unsplash.com/photo-1537151625747-768eb6cf92b2) | 2,28 mo | ![unsplash Cropped](https://d34bfvhzd89jzk.cloudfront.net/https://images.unsplash.com/photo-1537151625747-768eb6cf92b2?o=crop&w=100&h=100) | 6 ko   | Cropped to 100x100 pixels |
| ![unsplash](https://images.unsplash.com/photo-1537151625747-768eb6cf92b2) | 2,28 Mo | ![unsplash Cropped](https://d34bfvhzd89jzk.cloudfront.net/https://images.unsplash.com/photo-1537151625747-768eb6cf92b2?f=webp)             | 553 ko | Converted to WebP format  |

## Deploying

To deploy Imager, you can use the cloudformation template in the `cloudformation` directory. This template will create the following resources:

- A docker image containing the Imager application
- A Lambda function to process images
- An API Gateway endpoint to invoke the Lambda function
