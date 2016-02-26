#!/usr/bin/env python
# -*- coding:utf-8 -*-

import cgi
from datetime import datetime
import json
import os
from peewee import *

def main():
    with open("config.json") as f:
        conf = json.loads(f)
    failinfo_id = store(conf)

db = MySQLDatabase("failinfo_db", **{"passwd":DB_PASSWD, "host":"localhost", "user": DB_USER})
class Failinfo(Model):
    infotype = CharField()
    service = CharField()
    schedule_begin = DateTimeField()
    schedule_end = DateTimeField()
    body = CharField()
    class Meta:
        database = db

class BadRequestError(Exception):
    def __init__(self, msg):
        self._msg = str(msg)

    def __str__(self):
        return ("Status: 404 Bad Request\r\r\n\n" + self._msg)

def store(conf):
    # define constants
    API_KEY = "DBC38B518A83DC98147A132895A4837DF89AB0CDF9F2A57D100D8E1312719EC63309C6055F8880481128B54388472D848AAF945149E1936FB5CC17EFDA0A5193"
    DATETIME_FORMAT_STRING = "%Y-%m-%d %H:%M:%S"
    DB_PASSWD = ""
    DB_USER = "root"
    REQUIREMENT_PARAMS = ["infotype", "service", "schedule", "body", "apikey"]

    try:
        check_method("POST")
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
                                   body=form["body"])
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

    def check_method(method):
        if not os.environ['REQUEST_METHOD']==method:
            raise BadRequestError("This API accepts {} method only".format(method))
        return True

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
    ...
