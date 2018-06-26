package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type key struct {
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

func googleHandler(w http.ResponseWriter, r *http.Request) {
	/*
	https://accounts.google.com/o/oauth2/v2/auth?scope=profile&access_type=offline&include_granted_scopes=true&
	state=state_parameter_passthrough_value&redirect_uri=${ protocol }%3A%2F%2F${ encodeURIComponent(req.get('host')) }%
	2Foauth2callback&response_type=code&client_id=303128741136-r2284o9d45b5c1lhpnhe3uvrd5df9def.apps.googleusercontent.com`)*/

	http.Redirect(w, r, keys.Gg.AuthURI + "?scope=profile&access_type=offline" + 
		"&include_granted_scopes=true&state=state_parameter_passthrough_value" +
		"&redirect_uri=" + keys.Gg.RedirectURIs[1] + "&res&response_type=code" + 
		"&client_id=303128741136-r2284o9d45b5c1lhpnhe3uvrd5df9def.apps.googleusercontent.com", http.StatusFound)
}

func main() {
	readFile("keys.json")
	http.HandleFunc("/v1/api/auth/google", googleHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

var keys key

func readFile(fileName string) {
	file, e := ioutil.ReadFile("./" + fileName)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	e = json.Unmarshal(file, &keys)
	log.Println(keys)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
