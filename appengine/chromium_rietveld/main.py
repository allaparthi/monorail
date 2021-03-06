# Copyright 2008 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Main program for Rietveld.

This is also a template for running a Django app under Google App
Engine, especially when using a newer version of Django than provided
in the App Engine standard library.

The site-specific code is all in other files: urls.py, models.py,
views.py, settings.py.
"""

# Standard Python imports.
import os
import sys
import logging

# Log a message each time this module get loaded.
logging.info('Loading %s, app version = %s',
             __name__, os.getenv('CURRENT_VERSION_ID'))

import appengine_config

# AppEngine imports.
from google.appengine.ext.webapp import util

# Import webapp.template.  This makes most Django setup issues go away.
from google.appengine.ext.webapp import template  #pylint: disable=F0401

# Django 1.2+ requires DJANGO_SETTINGS_MODULE environment variable to be set
# http://code.google.com/appengine/docs/python/tools/libraries.html#Django
os.environ['DJANGO_SETTINGS_MODULE'] = 'settings'

# Import various parts of Django.
import django.core.handlers.wsgi
import django.core.signals
import django.db
import django.dispatch.dispatcher
import django.forms


def log_exception(*_args, **_kwds):
  """Django signal handler to log an exception."""
  cls, err = sys.exc_info()[:2]
  logging.exception('Exception in request: %s: %s', cls.__name__, err)


# Log all exceptions detected by Django.
django.core.signals.got_request_exception.connect(log_exception)

# Unregister Django's default rollback event handler.
django.core.signals.got_request_exception.disconnect(
    django.db._rollback_on_exception)

# Create a Django application for WSGI.
application = django.core.handlers.wsgi.WSGIHandler()


def real_main():
  """Main program."""
  # Run the WSGI CGI handler with that application.
  util.run_wsgi_app(application)


def profile_main():
  """Main program for profiling."""
  import cProfile
  import pstats
  import StringIO

  prof = cProfile.Profile()
  prof = prof.runctx('real_main()', globals(), locals())
  stream = StringIO.StringIO()
  stats = pstats.Stats(prof, stream=stream)
  # stats.strip_dirs()  # Don't; too many modules are named __init__.py.
  stats.sort_stats('time')  # 'time', 'cumulative' or 'calls'
  stats.print_stats()  # Optional arg: how many to print
  # The rest is optional.
  # stats.print_callees()
  # stats.print_callers()
  print '\n<hr>'
  print '<h1>Profile</h1>'
  print '<pre>'
  print stream.getvalue()[:1000000]
  print '</pre>'

# Set this to profile_main to enable profiling.
main = real_main


if __name__ == '__main__':
  main()
