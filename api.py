#!/usr/bin/env python
# -*- coding:utf-8 -*-

import cgi
from datetime import datetime
import json
import os
from peewee import *
from requests_oauthlib import OAuth1Session

def main():
    if check_method("POST"):
        with open("config.json") as f:
            conf = json.loads(f)
        failinfo_id = store(conf)
        tweet(conf, failinfo_id)
    elif check_method("GET"):
        try:
            form = cgi.FieldStorage()
            if "view" in form:
                if form["view"]=="null":
                    info = Failinfo.select()
                else:
                    key = int(form["view"])
                    info = Failinfo.select().where(Failinfo.id==key)
                print("Content-Type: text/plain\r\n")
                print(info)
            else:
                raise BadRequestError('key "view" is required')
        except BadRequestError as e:
            print(str(e))
            quit()
        except ValueError:
            print("Status: 400 Bad Request\r\n")
            print('key "view" should be integer')
            quit()

    def check_method(method):
        if os.environ['REQUEST_METHOD'] == method:
            return True
        else:
            return False

db = MySQLDatabase("failinfo_db", **{"passwd":DB_PASSWD, "host":"localhost", "user": DB_USER})
class Failinfo(Model):
    id = IntegerField()
    infotype = CharField()
    service = CharField()
    schedule_begin = DateTimeField()
    schedule_end = DateTimeField()
    body = CharField()
    end = BoolField()
    created_date = DateTimeField()
    class Meta:
        database = db

class BadRequestError(Exception):
    def __init__(self, msg):
        self._msg = str(msg)

    def __str__(self):
        return ("Status: 400 Bad Request\r\r\n\n" + self._msg)

def store(conf):
    # define constants
    API_KEY = "DBC38B518A83DC98147A132895A4837DF89AB0CDF9F2A57D100D8E1312719EC63309C6055F8880481128B54388472D848AAF945149E1936FB5CC17EFDA0A5193"
    DATETIME_FORMAT_STRING = "%Y-%m-%d %H:%M:%S"
    DB_PASSWD = ""
    DB_USER = "root"
    REQUIREMENT_PARAMS = ["infotype", "service", "schedule", "body", "apikey"]

    try:
        form = get_formdata()
        check_params(form)
        check_api_key(form)

        for key in ["begin", "end"]:
            if not form["schedule"][key]=="null":
                form["schedule"][key] = datetime.strptime(form["schedule"][key], DATETIME_FORMAT_STRING)

        db.create_table(Failinfo, True)
        with db.transaction():
            info = Failinfo.create(infotype=form["infotype"],
                                   service=form["service"],
                                   schedule_begin=form["schedule"]["begin"],
                                   schedule_end=form["schedule"]["end"],
                                   body=form["body"],
                                   end=False,
                                   created_date=datetime.now())
        return ("failinfo_id: " + info.id)
    except BadRequestError as e:
        print(str(e))
        quit()
    except Exception as e:
        print("Status: 500 Internal ServerError\r\n")
        print("An Error occured")
        print("================")
        print(str(e))
        print("================")
        db.rollback()
        quit()

    def get_formdata():
        form = cgi.FieldStorage()
        if "data" in form:
            return json.loads(form["data"].value)
        else:
            raise BadRequestError('form data does not have body that named "data"')

    def check_params(form):
        for requirement in REQUIREMENT_PARAMS:
            if not requirement in form:
                raise BadRequestError(requirement + " parameter is required")
            if requirement=="schedule":
                if not type(form["schedule"])==dict:
                    raise BadRequestError('"schedule" parameter should be object(dict) that has parameters "begin" and "end"')
                if not "begin" in form["schedule"] or not "end" in form["schedule"]:
                    raise BadRequestError('"begin" and "end" parameters are required in schedule')
        return True

    def check_api_key(form):
        if not form["apikey"]==API_KEY:
            raise BadRequestError("API key is not valid.")

def tweet(conf, info_id):
    # define constants
    URL = "https://api.twitter.com/1.1/statuses/update.json"
    TEMPLATE = "【{infotype}情報】{begin}〜{end}の間に、メンテナンスを行います。影響サービス:{service}。詳細は→{url}"
    data = dict()

    info = Failinfo.select().where(Failinfo.id==info_id)

    if info.infotype=="maintainance": data["infotype"] = "メンテナンス"
    elif info.infotype=="trouble": data["infotype"] = "障害"
    else: data["infotype"] = info.infotype

    if info.schedule_begin=="null": data["begin"] = "未明"
    else: data["begin"] = str(info.schedule_begin)

    if info.schedule_end=="null": data["end"] = "未定"
    else: data["end"] = str(info.schedule_end)

    data["service"] = info.service
    data["url"] = "http://charakoba.com/info/?view=" + info.id
    twitter = OAuth1Session(conf["CK"], conf["CS"], conf["AT"], conf["AS"])
    req = twitter.post(URL, params=dict(status=TEMPLATE.format(data)))

    if req.status_code == 200:
        print("Status: 200 OK\r\n")
        print("ツイートが完了しました")
        return True
    else:
        print("Status: " + str(req.status_code) + "\r\n")
        return False

main()
