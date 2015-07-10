# Golang @ Bity - NSQ Workshop

### 1. Install NSQ locally

http://nsq.io/deployment/installing.html

` $ brew install nsq`

### 2. Launch nsqlookupd

In one shell, start nsqlookupd:

`$ nsqlookupd`

### 3. Launch nsqd

In another shell, start nsqd:

`$ nsqd --lookupd-tcp-address=127.0.0.1:4160`

### 4. Manually pump a message into the queue

`$ curl -d "hello world 1" "http://127.0.0.1:4151/put?topic=test"`

### 5. Now what? What is happening?

We have two daemons running. What are they doing?

### 6. Launch nsqadmin

`$ nsqadmin --lookupd-http-address=127.0.0.1:4161`

Now you can load the NSQ Admin UI in a borwser: http://127.0.0.1:4171/

### 7. Writing a consumer

First get the client library for import:

`$ go get -u -v github.com/bitly/go-nsq`


