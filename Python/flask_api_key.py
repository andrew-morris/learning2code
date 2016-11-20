#!/usr/bin/env python

"""
$ curl "http://localhost:5000/protected?apikey=CCCC"
SOMETHING SECRET!!!
$ curl "http://localhost:5000/"
HAYYY
$ curl "http://localhost:5000/protected"
Invalid API Key. Please contact someone@shittystartup.io
$ curl "http://localhost:5000/protected?apikey=ZZZZ"
Invalid API Key. Please contact someone@shittystartup.io
"""

from flask import Flask, request, abort
from functools import wraps
from werkzeug.wrappers import Response

app = Flask(__name__)

def some_validation_function(apikey):
    apikeys = ["AAAA", "BBBB", "CCCC"]
    return apikey in apikeys

def require_appkey(view_function):
    @wraps(view_function)
    def decorated_function(*args, **kwargs):
        if request.args.get('apikey') and some_validation_function(request.args.get('apikey')):
            return view_function(*args, **kwargs)
        else:
            abort(401)
    return decorated_function

@app.route('/')
def unprotected():
    return 'HAYYY'

@app.route('/protected', methods=["POST","GET"])
@require_appkey
def protected():
    return 'SOMETHING SECRET!!!'

@app.errorhandler(401)
def custom_401(error):
    return Response('Invalid API Key. Please contact someone@shittystartup.io', 401, {'WWWAuthenticate':'Basic realm="Login Required"'}) 

