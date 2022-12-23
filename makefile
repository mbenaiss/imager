aws-registry := xxxx.dkr.ecr.us-east-1.amazonaws.com

build:
	docker build -t imager:latest -f cmd/app/dockerfile .

push-aws:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $(aws-registry)
	docker buildx build --platform linux/arm64 -t $(aws-registry)/imager:latest -f cmd/lambda/dockerfile --push .

deploy-cloudformation:
	aws cloudformation deploy \
		--template-file cloudformation/cloudformation.yml \
		--stack-name imager \
		--capabilities CAPABILITY_IAM \
		--region us-east-1 \
		--parameter-overrides \
			ImageUri=$(aws-registry)/imager:latest

