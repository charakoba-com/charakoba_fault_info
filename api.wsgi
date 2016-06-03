#!/usr/bin/env python
# -*- coding:utf-8 -*-

from bottle import Bottle, request, HTTPResponse
from datetime import datetime
import json
import MySQLdb
from MySQLdb.cursors import DictCursor
import os
import requests
from requests_oauthlib import OAuth1Session as OAuth

configfile = os.path.join(os.path.dirname(__file__), 'config.json')
with open(configfile, 'r') as f:
    cfg = json.load(f)

verb = {
    'maintenance': 'メンテナンスを行います',
    'event': '障害が発生しました'
}

application = Bottle()
post = application.post
get = application.get


class RequireNotSatisfiedError(Exception):
    pass


def require(keys):
    params = {}
    for key in keys:
        value = request.forms.get(key)
        if value is not None:
            params[key] = value
        else:
            raise RequireNotSatisfiedError(key)
    return params


def optional(keys):
    params = {}
    for key in keys:
        params[key] = request.forms.get(key)
    return params


def badRequest(key):
    response = HTTPResponse()
    response.status = 400
    response.body = json.dumps(
        {
            'message': 'Failed',
            'BadRequest': key
        }
    ) + "\n"
    return response


def cannotSave():
    response = HTTPResponse()
    response.status = 500
    response.body = json.dumps(
        {
            'message': 'Failed',
            'Error': 'CannotSaveToDB'
        }
    ) + "\n"
    return response


def cannotTweet():
    response = HTTPResponse()
    response.status = 500
    response.body = json.dumps(
        {
            'message': 'Failed',
            'Error': 'CannotTweet'
        }
    ) + "\n"
    return response


def success():
    response = HTTPResponse()
    response.status = 200
    response.body = json.dumps(
        {
            'message': 'Success'
        }
    ) + "\n"
    return response


def get_info(id_):
    with MySQLdb.connect(
            cursorclass=DictCursor,
            **cfg['DB_INFO']) as cursor:
        cursor.execute(
            '''SELECT * FROM fault_info_log
            WHERE id=%s;
            ''',
            (id_,)
        )
        row = cursor.fetchone()
    return row

def get_all_info():
    with MySQLdb.connect(
            cursorclass=DictCursor,
            **cfg['DB_INFO']) as cursor:
        cursor.execute(
            '''SELECT * FROM fault_info_log'''
        )
        rows = cursor.fetchall()
        return rows


def get_uri(id_):
    uri = (
        cfg['BASE_URI'] +
        str(id_) +
        '/'
    )
    return uri


def get_status(info):
    status = (
        '【{0}】{1}〜{2}{3}、{4}. 影響サービス:{5} 詳細:{6}'
        .format(
            info['type'],
            info['begin'],
            info.get('end', ''),
            '' if info.get('end', '') == '' else 'の間に',
            verb[info['type']],
            info['service'],
            get_uri(info['id'])
        )
    )
    return status


def save(params):
    with MySQLdb.connect(
        cursorclass=DictCursor,
        **cfg['DB_INFO']) as cursor:
        cursor.execute(
            '''INSERT INTO fault_info_log
            (type, service, begin, end, detail)
            VALUES (%s, %s, %s, %s, %s);
            ''',
            (
                params['type'],
                params['service'],
                params['begin'],
                params.get('end', None),
                params.get('detail', '')
            )
        )
        cursor.execute(
            '''SELECT last_insert_id() AS id FROM fault_info_log;
            '''
        )
    return cursor.fetchone()


def tweet(status):
    ENDPOINT = 'https://api.twitter.com/1.1/statuses/update.json'
    CONFIG = cfg['TWITTER_INFO']
    twitter = OAuth(
        CONFIG['API_KEY'],
        CONFIG['API_SECRET'],
        CONFIG['ACCESS_TOKEN'],
        CONFIG['ACCESS_SECRET']
    )
    res = twitter.post(
        ENDPOINT,
        params={"status": status}
    )
    if res.status_code == 200:
        return True
    else:
        return False


def default_datetime_format(o):
    if isinstance(o, datetime):
        return o.strftime('%Y/%m/%d %H:%M:%S')
    raise TypeError(repr(o) + " is not JSON serializable")


@post('/api')
def api_post_info():
    required_key = ['type', 'service', 'begin']
    optional_key = ['end', 'detail']
    try:
        params = require(required_key)
    except RequireNotSatisfiedError as e:
            return badRequest(e.message)
    params.update(optional(optional_key))
    try:
        id_ = save(params)['id']
    except:
        return cannnotSave()
    if tweet(get_status(get_info(id_))):
        return success()
    else:
        return cannotTweet()


@get('/api')
def api_get_info():
    response = HTTPResponse()
    all_ = request.query.get('all')
    if all_ in ['1', 'True', 'true']:
        rows = get_all_info()
        response.body = json.dumps(rows, default=default_datetime_format) + "\n"
    else:
        row = get_info(request.query.get('issue'))
        response.body = json.dumps(row, default=default_datetime_format) + "\n"
    return response
