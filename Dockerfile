FROM mesosphere/mesos-dns:0.5.2

ADD mesos-dns /usr/bin/mesos-dns-starter


ADD empty /tmp/

ENTRYPOINT ["/usr/bin/mesos-dns-starter", "/usr/bin/mesos-dns" ]
CMD ["-config=/etc/mesos-dns/config.json"]
