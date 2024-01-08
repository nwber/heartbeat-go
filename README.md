Heartbeat service written in go. Continuously polls the endpoint for a given duration. Saves output to a CSV (Status, endpoint, UTC timestamp).

I have basically no idea what I'm doing but it works I guess 

Setup:
- `go install`
- add go/bin to your path (like `export PATH=$PATH:$(go env GOPATH)/bin`) and `source .bashrc` or whatever you do for your shell
- alternatively just use it in this directory 
- start using the command :)

Example usage: `heartbeat-go --endpoint https://google.com --duration 10`
- Endpoint is the FQDN you'd like to check the heartbeat of
- Duration in seconds to query for (optional). Default is 600s
- run `heartbeat-go` for help!

![Alt text](image-1.png)

![Alt text](image.png)