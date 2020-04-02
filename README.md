<p align="center"><a href="https://www.verygoodsecurity.com/"><img src="https://avatars0.githubusercontent.com/u/17788525" width="128" alt="VGS Logo"></a></p>
<p align="center"><b>Golang snippets to use VGS proxy to redact and reveal</b></p>

## Inbound redaction snippet
To test with the example below, please replace the Tenant ID with yours.
```go
  package main

  import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
  )

  type Payload struct {
    Account string 'json:"account_number"'
  }

  func main() {
    data := Payload{
      // fill struct
    }
    payloadBytes, err := json.Marshal(data)
    if err != nil {
      log.Fatal(err)
    }

    body := bytes.NewReader(payloadBytes)

    req, err := http.NewRequest("POST", "https://<Tenant ID>.SANDBOX.verygoodproxy.com/post", body)
    if err != nil {
      log.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
      log.Fatal(err)
    }

    defer resp.Body.Close()

    respB, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    log.Println(string(respB))
  }
  
```

## Outbound revealing snippet
This is a example to send POST request to Stripe API  https://stripe.com/docs/api/tokens/create_card  
To test with the example, copy the SANDBOX cert.pem in VGS dashboard to your local environment, reset the path to your cert and replace the Tenant ID, Username, Password with yours. You should also use the token generated from your vault.
```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "flag"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "net/url"
)


func main() {

    proxyUrl, err := url.Parse("http://<Username>:<Password>@<Tenant ID>.SANDBOX.verygoodproxy.com:8080")


	var caFile   = flag.String("CA", "/.../sandboxcert.pem", "A PEM eoncoded CA's certificate file.")

	flag.Parse()

    // Load CA cert
    caCert, err := ioutil.ReadFile(*caFile)
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // Setup HTTPS client
    tlsConfig := &tls.Config{
        RootCAs: caCertPool,
    }
    tlsConfig.BuildNameToCertificate()
    // To turn off TLS verification: &tls.Config{InsecureSkipVerify : true}
    transport := &http.Transport{TLSClientConfig: tlsConfig, Proxy: http.ProxyURL(proxyUrl), }
    req, err := http.NewRequest("POST", "https://api.stripe.com/v1/tokens", strings.NewReader("card[number]=tok_sandbox_2ssJThVetq4Tjn6QhViKaw&card[cvc]=tok_sandbox_5Uddqh8pY1sCuYbpqtU7xw&card[exp_month]=04&card[exp_year]=2021"))
    if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Authorization", "Bearer sk_test_4eC39HqLyjWDarjtT1zdp7dc")

    client := &http.Client{Transport: transport}

    // Do POST to stripe API
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Dump response
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
}
```

## Demo

Run `go build example-outbound.go`  
You will see it builds an executable named example-outbound in the current directory alongside your source code. Execute it to see the result: 

![go-demo](go-demo.png "go-demo")