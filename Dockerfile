FROM mesosphere/mesos-dns:0.5.2

ADD mesos-dns /usr/bin/mesos-dns-starter

ENTRYPOINT ["/usr/bin/mesos-dns-starter" ]
CMD ["/usr/bin/mesos-dns", "-config=/etc/mesos-dns/config.json"]
