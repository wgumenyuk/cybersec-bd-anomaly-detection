FROM python:3.12-slim AS model-builder
WORKDIR /app
COPY model/ ./
RUN pip install -r requirements.txt
RUN python ./train.py

FROM python:3.12-slim
WORKDIR /app
COPY --from=model-builder /app/model.pkl ./
COPY app/ ./
RUN pip install -r requirements.txt
ENTRYPOINT [ "python", "/app/app.py" ]
