# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

.PHONY: test clean

pylint_parser: python pylint
	go build -o pylint_parser

test: pylint_parser
	./pylint_parser -input=test -output=out

python:
	echo "infra/python/cpython/linux-amd64 version:2.7.17.chromium.21.3pp" | \
			cipd ensure -ensure-file - -root python

pylint:
	pip install --ignore-installed --root pylint --no-compile pylint

clean:
	rm -rf pylint_parser pylint python out
