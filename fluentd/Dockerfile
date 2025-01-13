FROM bitnami/fluentd:latest
USER root
RUN fluent-gem install fluent-plugin-kafka
COPY fluent.conf /fluentd/etc/
CMD exec fluentd -c /fluentd/etc/fluent.conf
