FROM fluentd:latest
USER root
RUN fluent-gem install fluent-plugin-rabbitmq
USER fluent
COPY fluent.conf /fluentd/etc/
EXPOSE 24224
CMD exec fluentd -c /fluentd/etc/fluent.conf
