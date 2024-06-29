# GO TCP echo server

## A simple tcp echo server written in Go



### run the server
```
go build -o server server/main.go
./server/main
```

### access the server via telnet
```
telnet localhost 8080
```

### to quit telnet
```
Ctrl + ] and press enter
then type quit and press enter
```

### run client to test simultaneous connections to the server
```
go build -o client client/main.go
./client/main
```

### GCP deployment
- create vm in compute engine (can be smallest instance like N1 class f1-micro )
- Create a firewall rule allowing all traffic to port 8080
```
gcloud compute firewall-rules create rule-allow-tcp-8080 --source-ranges 0.0.0.0/0 --target-tags allow-tcp-8080 --allow tcp:8080
```
- Add the firewall rule to the VM
```
gcloud compute instances add-tags <vm name>  --tags allow-tcp-8080
```
- ssh into the VM and build/run the server
- access the server via telnet
```
telnet <vm ip> 8080
```