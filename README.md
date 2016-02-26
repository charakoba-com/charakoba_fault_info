# charakoba_fault_info

## HOW TO USE
1. configure with config file :: set your tokens of twitter api to configsample.json, and set API KEY of this script.
then, rename configsample.json to config.json.
1. configure in api.py :: set your mysql username and password in api.py. you can find it where next to import statements, "# define constants"
1. make database :: make database with mysql, which named "faultinfo_db"
1. setup :: set api.py, config.json to your cgi-runnable dir. chmod 755 api.py and use it.

## API

| METHOD |              REQUIRED PARAMS              |         DESCRIPTION        |
|:------:|:-----------------------------------------:|:--------------------------:|
|  POST  | infotype, service, schedule, body, apikey | save to database and tweet |
|  GET   | view                                      | show record                |

### save to database and tweet
need REQUIRED PARAMS with json named "data"

### show record
need REQUIRED PARAM with query string
if you set view=null, show all record
