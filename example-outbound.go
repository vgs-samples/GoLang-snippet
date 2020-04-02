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
