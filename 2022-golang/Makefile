.PHONY: run
FILES = $(shell ls -d [0-9][0-9])

run:
ifeq (, $(shell which go))
	docker build -q -t gorun . && docker run --rm -it gorun make
else
	$(foreach file, $(FILES), \
		echo DAY\#${file}; \
		cd ${file}; go test .; go run .; cd ..;)
endif
