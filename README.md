# Golang @ Bity - NSQ Workshop

### 1. Install NSQ locally

http://nsq.io/deployment/installing.html

` $ brew install nsq`

### 2. Launch nsqlookupd

In one shell, start nsqlookupd:

`$ nsqlookupd`

### 3. Launch nsq

In another shell, start nsqd:

`$ nsqd --lookupd-tcp-address=127.0.0.1:4160`

### 4. Manually pump a message into the queue

`$ curl -d "hello world 1" "http://127.0.0.1:4151/put?topic=test"`

### 5. Now what?