<p align="center"><a href="https://www.verygoodsecurity.com/"><img src="https://avatars0.githubusercontent.com/u/17788525" width="128" alt="VGS Logo"></a></p>
<p align="center"><b>Golang snippets to use VGS proxy to redact and reveal</b></p>

## How to run

### Import routes
1. Go to VGS dashboard -> Routes -> Import Routes and import `routes/inbound.yaml` and `routes/outbound.yaml`

### Inbound
```bash
go build example-inbound.go
```
Then run binary
```bash
./example-inbound
```

Take a look at response payload:
```
"json": {
  "card": "tok_sandbox_czyCsWKgmFJGFeL76PsBig", 
  "cvc": "tok_sandbox_4Se6PmhRrQZBEyLrGD1vbV"
}
```

### Outbound
Copy created card/cvc aliases from the above step and put to the `example-outbound.go`

```bash
go build example-outbound.go
```
Then run binary
```bash
./example-outbound
```

![go-demo](go-demo.png "go-demo")