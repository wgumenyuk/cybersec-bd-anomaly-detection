FROM python:3.12-slim AS model-builder
WORKDIR /app
COPY model/ ./
RUN pip install -r requirements.txt
RUN python ./train.py

FROM python:3.12-slim
RUN apt update && \
	apt install -y openjdk-17-jre && \
	rm -rf /var/lib/apt/lists
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64
WORKDIR /app
COPY --from=model-builder /app/model.pkl ./
COPY app/ ./
RUN pip install -r requirements.txt
ENTRYPOINT [ "spark-submit", "--packages", "org.apache.spark:spark-sql-kafka-0-10_2.12:3.5.4", "app.py" ]
