build:
	@sh -c "'$(CURDIR)/build.sh'"

push:
	@sh -c "'$(CURDIR)/push.sh'"

build_develop:
	docker build -t "libonomy/libonomy:develop" -f Dockerfile.develop .

build_testing:
	docker build --tag libonomy/testing -f ./Dockerfile.testing .

push_develop:
	docker push "libonomy/libonomy:develop"

.PHONY: build build_develop push push_develop
