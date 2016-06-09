FROM mesosphere/mesos-dns:0.5.2

ADD mesos-dns /usr/bin/mesos-dns-starter

ENTRYPOINT ["/usr/bin/mesos-dns-starter", "/usr/bin/mesos-dns" ]
CMD ["-v=2", "-config=/etc/mesos-dns/config.json"]
