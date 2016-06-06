#!/usr/bin/env python
# -*- coding:utf-8 -*-

from bottle import Bottle, HTTPResponse, request, static_file
from bottle import TEMPLATE_PATH, jinja2_view as view
import json
import os
import api

configfile = os.path.join(os.path.dirname(__file__), 'config.json')
with open(configfile, 'r') as f:
    cfg = json.load(f)

application = Bottle()
get = application.get

application.mount('/api', api.application)
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
TEMPLATE_PATH.append(BASE_DIR + '/statics/template')


@get('/')
@view('index.tpl')
def index():
    return {"infos": api.get_all_info()}


@get('/detail/<id_:int>/')
@view('detail.tpl')
def detail(id_):
    return {"info": api.get_info(id_).reverse()}


@get('/statics/<filename:path>')
def static(filename):
    return static_file(filename, root=BASE_DIR+'/statics')

if __name__ == '__main__':
    application.run(host='localhost', port=8080, debug=True, reloader=True)
