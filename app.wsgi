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
services = ["RS", "VPS"]


@get('/')
@view('index.tpl')
def index():
    infos = list(api.get_all_info())
    infos.reverse()
    return {"infos": infos}


@get('/detail/<id_:int>/')
@view('detail.tpl')
def detail(id_):
    return {"info": api.get_info(id_)}


@get('/detail/<id_:int>/edit/')
@view('edit.tpl')
def edit(id_):
    info = api.get_info(id_)
    info['begin_date'] = info['begin'].date()
    info['begin_hour'] = info['begin'].time().hour
    info['begin_minute'] = info['begin'].time().minute
    info['begin_second'] = info['begin'].time().second
    info['end_date'] = info['end'].date()
    info['end_hour'] = info['end'].time().hour
    info['end_minute'] = info['end'].time().minute
    info['end_second'] = info['end'].time().second
    return {"services": services, "info": info}


@get('/post/')
@view('post.tpl')
def postpage():
    return {"services": services}


@get('/statics/<filename:path>')
def static(filename):
    return static_file(filename, root=BASE_DIR+'/statics')

if __name__ == '__main__':
    application.run(host='localhost', port=8080, debug=True, reloader=True)
