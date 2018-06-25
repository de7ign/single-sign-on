package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type keys struct {
	Gg web `json:"gg"`
	Fb web `json:"fb"`
}


type web struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	AuthProviderx509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectURIs            []string `json:"redirect_uris"`
	JavascriptOrigins       []string `json:"javascript_origins"`
}

func main() {
	readFile("keys.json")
}

func readFile(fileName string) {
	file, e := ioutil.ReadFile("./" + fileName)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
	var keys keys
	e = json.Unmarshal(file, &keys)
	log.Println(keys)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
