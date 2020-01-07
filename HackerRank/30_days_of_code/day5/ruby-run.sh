#!/bin/bash
docker run -it --rm --volume `pwd`:/usr/local/src --workdir /usr/local/src --name ruby-run ruby:2.4-alpine3.9 ruby $1
