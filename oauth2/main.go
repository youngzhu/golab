package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

const (
	clientID     = "xx"
	clientSecret = "xx"
)

var (
	AccessTypeWebServer = oauth2.SetAuthURLParam("type", "web_server")
	//AccessTypeWebServer = oauth2.SetAuthURLParam("access_type", "web_server")
	redirectURL = "http://localhost:8080/callback" // 确保此重定向URL已在Basecamp应用程序设置中注册
	//redirectURL         = "http://myapp.com/oauth" // 经测试，需要有效的地址才行

	Endpoint = oauth2.Endpoint{
		AuthURL:  "https://launchpad.37signals.com/authorization/new",
		TokenURL: "https://launchpad.37signals.com/authorization/token",
	}
)

func main() {
	/*
		fmt.Println("webhook")
		url := "https://3.basecamp.com/5161745/integrations/VA5bYy3rjSySL9LNRSRBFwAo/buckets/34862946/chats/6688955697/lines"

		url = "https://3.basecampapi.com/5161745/buckets/1/schedules/3/entries.json"

		result, err := doRequest(url, strings.NewReader(`{"summary":"Important Meeting","starts_at":"2015-06-04T00:00:00Z","ends_at":"2015-06-04T00:00:00Z"}`))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	*/

	getAccessTokenByChatGPT()
	//getAccessTokenByExample()
}

func getAccessTokenByChatGPT() {

	// 设置Basecamp的OAuth2配置
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		//Scopes:       []string{"read_write"}, // 根据Basecamp API的权限范围设置scopes
		Endpoint: Endpoint,
	}

	// 创建一个HTTP服务器用于处理回调
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		url := conf.AuthCodeURL("state", AccessTypeWebServer)
		fmt.Printf("Visit the URL for the auth dialog: %v\n", url)
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		// 使用授权码交换获取token
		// 就这个参数，搞了好久。。。
		// 请求的时候加了，回调的时候没加
		// 虽然提示信息是一样的，但没想到这一块
		//token, err := conf.Exchange(context.Background(), code)
		token, err := conf.Exchange(context.Background(), code, AccessTypeWebServer)
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		// 使用token进行操作，例如向Basecamp API发出请求

		fmt.Fprintf(w, "Token: %s", token.AccessToken)
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 参考 oauth2 里的测试方法
// 好像没效果
func getAccessTokenByExample() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		//Scopes:       []string{"SCOPE1", "SCOPE2"},
		RedirectURL: redirectURL,
		Endpoint:    Endpoint,
	}

	/*
		// use PKCE to protect against CSRF attacks
		// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
		verifier := oauth2.GenerateVerifier()

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

	*/

	//url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	url := conf.AuthCodeURL("", AccessTypeWebServer)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code, AccessTypeWebServer)
	if err != nil {
		log.Fatal(err)
	}
	//tok.TokenType

	client := conf.Client(ctx, tok)
	client.Get("https://launchpad.37signals.com/authorization.json")
}
