package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/garyburd/go-oauth/oauth"
	yaml "gopkg.in/yaml.v2"
)

type ts struct {
	ConsumerKey    string `yaml:"TWITTER_KEY,required"`
	ConsumerSecret string `yaml:"TWITTER_SECRET,required"`
	AccessToken    string `yaml:"TWITTER_ACCESSTOKEN,required"`
	AccessSecret   string `yaml:"TWITTER_ACCESSSECRET,required"`
}

var conn net.Conn
var reader io.ReadCloser
var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func (t *ts) readKeys() *ts {
	yamlFile, err := ioutil.ReadFile("/Users/Leland/go/go_code/blueprints/socialpoll/api_keys/secrets.yml")

	if err != nil {
		log.Printf("Error reading YAML: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, t)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return t
}

//dial ensures that con is closed and then opens a new connection, keeping
//the conn variable updated with the current connection. If the connection dies,
// we can redial without worrying about zombie connections
func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn = netc
	return netc, nil
}

//we will use closeconn to break the ongoing connection with Twitter and tidy things up
func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}
}

func setupTwitterAuth() {
	var ts = &ts{}
	ts = ts.readKeys()
	creds = &oauth.Credentials{
		Token:  ts.AccessToken,
		Secret: ts.AccessSecret,
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
}
