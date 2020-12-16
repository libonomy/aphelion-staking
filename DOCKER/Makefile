build:
	@sh -c "'$(CURDIR)/build.sh'"

push:
	@sh -c "'$(CURDIR)/push.sh'"

build_develop:
	docker build -t "aphelion/aphelion:develop" -f Dockerfile.develop .

build_testing:
	docker build --tag aphelion/testing -f ./Dockerfile.testing .

push_develop:
	docker push "aphelion/aphelion:develop"

.PHONY: build build_develop push push_develop
