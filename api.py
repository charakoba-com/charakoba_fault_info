#!/usr/bin/env python
# -*- coding:utf-8 -*-

import cgi
from datetime import datetime
import json
import os
from peewee import *

API_KEY = "DBC38B518A83DC98147A132895A4837DF89AB0CDF9F2A57D100D8E1312719EC63309C6055F8880481128B54388472D848AAF945149E1936FB5CC17EFDA0A5193"
DATETIME_FORMAT_STRING = "%Y-%m-%d %H:%M:%S"
DB_PASSWD = ""
DB_USER = "root"
requirement_params = ["infotype", "service", "schedule", "body", "apikey"]

db = MySQLDatabase("failinfo_db", **{"passwd":DB_PASSWD, "host":"localhost", "user": DB_USER})
class Failinfo(Model):
    infotype = CharField()
    service = CharField()
    schedule_begin = DateTimeField()
    schedule_end = DateTimeField()
    body = CharField()
    class Meta:
        database = db

if not os.environ['REQUEST_METHOD']=="POST":
    print("Status: 400 Bad Request\r\n")
    print("This API accepts only POST Method.")
    quit()

form = json.loads(cgi.FieldStorage()["data"].value)
for requirement in requirement_params:
    if not requirement in form:
        print("Status: 400 Bad Request\r\n")
        print(requirement + " parameter is required")
        quit()
    if requirement=="schedule":
        if not "begin" in form["schedule"] or not "end" in form["schedule"]:
            print("Status: 400 Bad Request\r\n")
            print('"begin" and "end" parameters are required in schedule')
            quit()

if not form["apikey"]==API_KEY:
    print("Status: 400 Bad Request\r\n")
    print("API key is not valid.")
    quit()

try:
    db.create_table(Failinfo, True)
    with db.transaction():
        info = Failinfo.create(infotype=form["infotype"], service=form["service"], schedule_begin=datetime.strptime(form["schedule"]["begin"], DATETIME_FORMAT_STRING), schedule_end=datetime.strptime(form["schedule"]["end"], DATETIME_FORMAT_STRING), body=form["body"])
except:
    print("Status: 500 Internal ServerError\r\n")
    print("An Error occured")
    db.rollback()
    quit()

print("Status: 201 Created\r\n")
print("failinfo_id: " + info.id)
