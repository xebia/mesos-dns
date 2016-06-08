# mesos-dns 

This docker image generates a config.json before starting the real mesos-dns. The arguments to the utility
will be passed on to the real mesos-dns. 

You can configure mesos-dns using environment variables as show in the following template. 

```
{
  "zk": 	"$MESOS_DNS_ZK",
  "masters":	[for i in $MESOS_DNS_MASTERS; do echo \"$i\" ; done | paste -s -d, -],
  "refreshSeconds": $MESOS_DNS_REFRESH_SECONDS,
  "ttl":	$MESOS_DNS_TTL,
  "domain":	"$MESOS_DNS_DOMAIN",
  "port":	$MESOS_DNS_PORT,
  "resolvers":	[for i in $MESOS_DNS_RESOLVERS; do echo \"$i\" ; done | paste -s -d, -],
  "timeout":	$MESOS_DNS_TIMEOUT, 
  "httpon":	$MESOS_DNS_HTTPON,
  "dnson":	$MESOS_DNS_DNSON,
  "httpport":	$MESOS_DNS_HTTP_PORT,
  "externalon":	$MESOS_DNS_EXTERNALON,
  "listener":	"$MESOS_DNS_LISTENER",
  "SOAMname":	"$MESOS_DNS_SOA_MNAME",
  "SOARname":	"$MESOS_DNS_SOA_RNAME",
  "SOARefresh":	$MESOS_DNS_SOA_REFRESH,
  "SOARetry":	$MESOS_DNS_SOA_RETRY,
  "SOAExpire":	$MESOS_DNS_SOA_EXPIRE,
  "SOAMinttl":	$MESOS_DNS_SOA_MIN_TTL,
  "IPSources":	[for i in $MESOS_DNS_IP_SOURCES; do echo \"$i\" ; done | paste -s -d, -],
}
```

- At least MESOS_DNS_ZK or MESOS_DNS_MASTERS has to be spcified for mesos-dns to start
- If an environment variable is not specified, no value will be generated for it. 
- for MESOS_DNS_IP_SOURCES, MESOS_DNS_RESOLVERS and MESOS_DNS_MASTERS a space separated list should be set.

```
MESOS_DNS_MASTERS="zoo01:2181 zoo02:2181 zoo03:2181"
```

An example command line:

```
docker run -e MESOS_DNS_ZK=zk://zookeeper:2181/mesos -net host xebia/mesos-dns 
```
