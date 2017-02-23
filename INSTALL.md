## Getting started ##

First of all, make sure that etcd runs on the server.

```
$ sudo systemctl start etcd.service
```

Etcd should listen on TCP port 2379 by default.
Build netcore and run.

```
$ go build
$ sudo ./netcore --etcd http://127.0.0.1:2379
```

If etcd runs on a different port, e.g. 4001, specify the port number to the `--etcd` option.

Test by sending DNS queries.

```
$ host google.com 127.0.0.1
google.com has address 172.217.17.142
google.com has IPv6 address 2a00:1450:400e:807::200e
google.com mail is handled by 30 alt2.aspmx.l.google.com.
google.com mail is handled by 50 alt4.aspmx.l.google.com.
google.com mail is handled by 20 alt1.aspmx.l.google.com.
google.com mail is handled by 40 alt3.aspmx.l.google.com.
google.com mail is handled by 10 aspmx.l.google.com.
```

On the console where netcore daemon is running, you can see like that:

```
NETCORE INITIALIZING
DNS ETCD CONFIG FETCH
DHCP ETCD CONFIG: [&{Instance:localhost.localdomain Enabled:true NetworkAddress:0.0.0.0:53 DefaultTTL:3h0m0s MinimumTTL:1m0s CacheRetention:0s Forwarders:[8.8.8.8:53 8.8.4.4:53]}]
NETCORE DNS STARTED
DNS Query [1/1] google.com. A from 127.0.0.1:50349
  [   0.0751ms] QUERY   google.com. A
  [   1.4706ms] FORWARD google.com. A
  [  42.0288ms] ANSWER  google.com. 299     IN      A       172.217.17.142
DNS Query [1/1] google.com. AAAA from 127.0.0.1:34469
  [   0.0574ms] QUERY   google.com. AAAA
  [   1.0601ms] FORWARD google.com. AAAA
  [  47.7450ms] ANSWER  google.com. 299     IN      AAAA    2a00:1450:400e:807::200e
DNS Query [1/1] google.com. MX from 127.0.0.1:48917
  [   0.0585ms] QUERY   google.com. MX
  [   1.6259ms] FORWARD google.com. MX
  [  42.4590ms] ANSWER  google.com. 599     IN      MX      30 alt2.aspmx.l.google.com.
  [  42.5917ms] ANSWER  google.com. 599     IN      MX      50 alt4.aspmx.l.google.com.
  [  42.6440ms] ANSWER  google.com. 599     IN      MX      20 alt1.aspmx.l.google.com.
  [  42.6959ms] ANSWER  google.com. 599     IN      MX      40 alt3.aspmx.l.google.com.
  [  42.7522ms] ANSWER  google.com. 599     IN      MX      10 aspmx.l.google.com.
```

If netcore doesn't run due to existing keys in etcd, try to clean up etcd registry completely.

```
$ sudo systemctl stop etcd.service
$ sudo rm -rf /var/lib/etcd/default.etcd/member/
$ sudo systemctl start etcd.service
```
