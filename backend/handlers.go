package faultinfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/charakoba-com/fault_info/backend/model"
	"github.com/charakoba-com/fault_info/backend/service"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/fsnotify/fsnotify"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type jsonErr struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type config struct {
	Token          string
	BaseURI        string
	TwAPIKey       string
	TwAPISecret    string
	TwAccessToken  string
	TwAccessSecret string
}

var c config

func init() {
	log.Printf("initialize token")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/faultinfo/")
	viper.AddConfigPath("$HOME/.faultinfo")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		log.Printf("token is now configured. (token: %s)", c.Token)
	})
	viper.SetDefault("token", "")
	viper.SetDefault("DBUser", "root")
	viper.SetDefault("DBPasswd", "")
	viper.SetDefault("DBAddr", "127.0.0.1:3306")
	viper.SetDefault("DBName", "faultinfo_db")
	c = config{
		Token:          viper.GetString(`token`),
		BaseURI:        viper.GetString(`baseuri`),
		TwAPIKey:       viper.GetString(`twitter.apikey`),
		TwAPISecret:    viper.GetString(`twitter.apisecret`),
		TwAccessToken:  viper.GetString(`twitter.accesstoken`),
		TwAccessSecret: viper.GetString(`twitter.accesssecret`),
	}
	dbcfg := &mysql.Config{
		User:      viper.GetString(`mysql.user`),
		Passwd:    viper.GetString(`mysql.password`),
		Net:       "tcp",
		Addr:      viper.GetString(`mysql.address`),
		DBName:    viper.GetString(`mysql.database`),
		ParseTime: true,
	}

	log.Printf("token is now configured. (token: %s)", c.Token)
	if err := db.Init(dbcfg); err != nil {
		panic(err)
	}
	log.Printf("connected database")

}

func httpJSONWithStatus(w http.ResponseWriter, st int, response interface{}) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(response); err != nil {
		httpError(w, http.StatusInternalServerError, `encode response to json`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(st)
	log.Printf("%d %s", st, buf.String())
	buf.WriteTo(w)
}

func httpJSON(w http.ResponseWriter, response interface{}) {
	httpJSONWithStatus(w, http.StatusOK, response)
}

func httpError(w http.ResponseWriter, st int, message string, err error) {
	v := jsonErr{
		Message: message,
	}
	if err != nil {
		v.Error = err.Error()
	}
	httpJSONWithStatus(w, st, v)
}

func auth(token string) bool {
	log.Printf("given token: %s", token)
	return token == c.Token
}

func tweet(message string) error {
	log.Printf("tweet")
	cfg := oauth1.NewConfig(c.TwAPIKey, c.TwAPISecret)
	tk := oauth1.NewToken(c.TwAccessToken, c.TwAccessSecret)
	httpclient := cfg.Client(oauth1.NoContext, tk)
	client := twitter.NewClient(httpclient)
	_, _, err := client.Statuses.Update(message, nil)

	return err
}

// PostInfoHandler is a HTTP handler
// the path is `/`
func PostInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("PostInfoHandler")
	// process request
	var request model.PostInfoHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httpError(w, http.StatusBadRequest, `invalid json request`, nil)
		return
	}
	if !auth(request.Token) {
		httpError(w, http.StatusForbidden, `invalid token`, nil)
		return
	}
	var svc service.InfoService
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, http.StatusInternalServerError, `database connection error`, err)
		return
	}
	id, err := svc.Create(tx, request)
	if err != nil {
		httpError(w, http.StatusInternalServerError, `record creation`, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, http.StatusInternalServerError, `commit transaction`, err)
		return
	}
	// tweet
	msg := fmt.Sprintf("%s :: %s | Service: %s / Date: %s - %s ", request.InfoType, request.Detail, request.Service, request.Begin, request.End)
	if len(msg+c.BaseURI) > 140 {
		msg = msg[:135-len(c.BaseURI)] + `...`
	}
	msg = msg + c.BaseURI
	if err := tweet(msg); err != nil {
		httpError(w, http.StatusInternalServerError, `successfully commited, but cannot tweet`, err)
		return
	}
	// generate response
	response := model.PostInfoHandlerResponse{
		Message: "success",
		ID:      id,
	}

	httpJSON(w, response)
}

// GetInfoListHandler is a HTTP handler
// the path is `/`
func GetInfoListHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetInfoHandler")
	// process request
	var svc service.InfoService
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, http.StatusInternalServerError, `database connection error`, err)
		return
	}
	infoList, err := svc.Listup(tx)
	if err != nil {
		httpError(w, http.StatusInternalServerError, `listup info`, err)
		return
	}

	// generate response
	response := model.GetInfoListHandlerResponse{
		Info: infoList,
	}
	httpJSON(w, response)
}

// UpdateInfoHandler is a HTTP handler
// the path is `/{id:[0-9]+}`
func UpdateInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateInfoHandler")

	if r.Header.Get("Content-Type") != "application/json" {
		httpError(w, http.StatusBadRequest, `Content-Type: application/json is required`, nil)
		return
	}

	// process request
	var request model.UpdateInfoHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httpError(w, http.StatusBadRequest, `invalid json request`, nil)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		httpError(w, http.StatusInternalServerError, `id must be int`, err)
		return
	}
	if !auth(request.Token) {
		httpError(w, http.StatusForbidden, `invalid token`, nil)
		return
	}

	var svc service.InfoService
	var info model.Info
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, http.StatusInternalServerError, `database error`, err)
		return
	}
	if err := svc.Update(tx, id, request); err != nil {
		httpError(w, http.StatusInternalServerError, `record update`, err)
		return
	}
	if err := info.Load(tx, id); err != nil {
		httpError(w, http.StatusInternalServerError, `loading info`, err)
		return
	}
	if err := tx.Commit(); err != nil {
		httpError(w, http.StatusInternalServerError, `commit transaction`, err)
		return
	}
	// tweet
	msg := fmt.Sprintf("=updated= %s :: %s | Service: %s / Date: %s - %s ", info.Type, info.Detail, info.Service, info.Begin, info.End)
	if len(msg+c.BaseURI) > 140 {
		msg = msg[:135-len(c.BaseURI)] + `...`
	}
	msg = msg + c.BaseURI
	if err := tweet(msg); err != nil {
		httpError(w, http.StatusInternalServerError, `successfully commited, but cannot tweet`, err)
		return
	}

	// generate response
	response := model.UpdateInfoHandlerResponse{
		Message: "success",
	}

	httpJSON(w, response)
}

// GetTypesHandler returns information type list
func GetTypesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetTypesHandler")
	var svc service.TypeService
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, http.StatusInternalServerError, `database error`, err)
		return
	}
	types, err := svc.Listup(tx)
	if err != nil {
		httpError(w, http.StatusInternalServerError, `listup types`, err)
		return
	}
	response := model.GetTypesHandlerResponse{
		Types: types,
	}
	httpJSON(w, response)
}

// GetServicesHandler returns service list
func GetServicesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetServicesHandler")

	var svc service.ServiceService
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, http.StatusInternalServerError, `database error`, err)
		return
	}
	services, err := svc.Listup(tx)
	if err != nil {
		httpError(w, http.StatusInternalServerError, `listup services`, err)
		return
	}
	response := model.GetServicesHandlerResponse{
		Services: services,
	}
	httpJSON(w, response)
}
