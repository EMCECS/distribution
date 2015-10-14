FROM registry:2



WORKDIR $DISTRIBUTION_DIR
COPY ../bin /bin


VOLUME ["/var/lib/registry"]
EXPOSE 5000
ENTRYPOINT ["registry"]
CMD ["/etc/docker/registry/config.yml"]

