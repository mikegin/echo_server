# GO TCP echo server

## A simple tcp echo server written in Go



### run the server
```
go build
./echo_server &
```

### access the server via telnet
```
telnet localhost 8080
```

### to quit telnet
```
Ctrl + ]
then type quit and press enter
```


### see the process running
```
ps ax | grep echo_server
```

### kill the process
```
kill echo_server
```

### GCP deployment
- create vm in compute engine (can be smallest instance like N1 class f1-micro )
- Create a firewall rule allowing all traffic to port 8080
```
gcloud compute firewall-rules create rule-allow-tcp-8080 --source-ranges 0.0.0.0/0 --target-tags allow-tcp-8080 --allow tcp:8080
```
- Add the firewall rule to the VM
```
gcloud compute instances add-tags CHANGE_ME_TO_VM_NAME  --tags allow-tcp-8080
```
- ssh into the VM and build/run the server
- access the server via telnet
```
telnet CHANGE_ME_TO_VM_IP 8080
```