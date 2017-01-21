FROM mesosphere/mesos-dns:v0.6.0

LABEL description="official mesos-dns docker image extended to allow configuration using environment variables"
LABEL url="https://github.com/xebia/mesos-dns"
 

ADD mesos-dns /usr/bin/mesos-dns-starter


ADD empty /tmp/

ENTRYPOINT ["/usr/bin/mesos-dns-starter", "/usr/bin/mesos-dns" ]
CMD ["-config=/etc/mesos-dns/config.json"]
