package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	TokenInfoURI            string   `json:"token_info_uri"`
	AuthProviderx509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectURIs            []string `json:"redirect_uris"`
	JavascriptOrigins       []string `json:"javascript_origins"`
}

type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}
type tokenInfo struct {
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Exp           string `json:"exp"`
	Iss           string `json:"iss"`
	Iat           string `json:"iat"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
}

func googleHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, keys.Gg.AuthURI+"?scope=profile&access_type=offline"+
		"&include_granted_scopes=true&state=state_parameter_passthrough_value"+
		"&redirect_uri="+keys.Gg.RedirectURIs[1]+"&res&response_type=code"+
		"&client_id=303128741136-r2284o9d45b5c1lhpnhe3uvrd5df9def.apps.googleusercontent.com", http.StatusFound)
}

func oauth2callbackHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if q.Get("code") == "" {
		http.Redirect(w, r, "http://localhost:5000/v1/api/auth/google", http.StatusFound)
	} else {

		res, err := http.PostForm(keys.Gg.TokenURI, url.Values{
			"code":          {q.Get("code")},
			"client_id":     {keys.Gg.ClientID},
			"client_secret": {keys.Gg.ClientSecret},
			"redirect_uri":  {keys.Gg.RedirectURIs[1]},
			"grant_type":    {"authorization_code"}})

		if err != nil {
			log.Print("error in exchanging code, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Print("error in reading body, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var token token
		err = json.Unmarshal(data, &token)
		if err != nil {
			log.Println("error on unmarshaling, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Print(token)

		res, err = http.PostForm(keys.Gg.TokenInfoURI, url.Values{
			"id_token": {token.IDToken}})

		if err != nil {
			log.Printf("error in getting metadata")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err = ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Print("error in reading body, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var tokenInfo tokenInfo
		err = json.Unmarshal(data, &tokenInfo)
		if err != nil {
			log.Printf("error in unmarshalling tokenInfo, ", )
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(tokenInfo)

	}
}

func main() {
	readFile("keys.json")
	http.HandleFunc("/v1/api/auth/google", googleHandler)
	http.HandleFunc("/oauth2callback", oauth2callbackHandler)
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
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
