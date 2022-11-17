.PHONY: run
FILES = $(shell ls -d [0-9][0-9])

run:
	$(foreach file, $(FILES), \
		echo DAY\#${file}; \
		cd ${file}; go run .; cd ..;)

run-with-docker:
	docker build -q -t gorun . && docker run --rm -it gorun make
