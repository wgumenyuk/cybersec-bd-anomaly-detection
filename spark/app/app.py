import pickle
from os import environ
from sys import stderr
from typing import Union, cast
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, from_json, udf
from pyspark.sql.types import FloatType, StringType, StructField, StructType

KAFKA_URI = environ.get("KAFKA_URI")
KAFKA_TOPIC = environ.get("KAFKA_TOPIC")

if not KAFKA_URI:
	print("`KAFKA_URI` not found in environment variables", file=stderr)
	exit(1)

if not KAFKA_TOPIC:
	print("`KAFKA_TOPIC` not found in environment variables", file=stderr)
	exit(1)

model = pickle.load(open("model.pkl", "rb"))

spark = cast(SparkSession.Builder, SparkSession.builder) \
	.appName("KafkaStreamProcessing") \
	.getOrCreate()

schema = StructType([
	StructField("level", StringType(), True),
	StructField("id", StringType(), True),
	StructField("method", StringType(), True),
	StructField("endpoint", StringType(), True),
	StructField("status", FloatType(), True),
	StructField("ip", StringType(), True),
	StructField("ua", StringType(), True),
	StructField("ms", FloatType(), True),
	StructField("time", StringType(), True)
])

df = spark \
	.readStream \
	.format("kafka") \
	.option("kafka.bootstrap.servers", KAFKA_URI) \
	.option("subscribe", KAFKA_TOPIC) \
	.option("startingOffsets", "earliest") \
	.load()

def predict_udf(data: dict[str, Union[str, float]]):
	try:
		status = data["status"]
		ms = data["ms"]

		feature_list = [float(status), float(ms)]
		prediction = model.predict(feature_list)[0]

		match prediction:
			case 0:
				return "Normal"
			case 1:
				return "Bruteforce"
			case 2:
				return "DDoS"
			case _:
				return "Unknown"
	except KeyError as e:
		return f"Error: Missing key {str(e)}"
	except Exception as e:
		return "Prediction failed"

predict = udf(predict_udf, StringType())

df = df \
	.selectExpr("CAST(value AS STRING)") \
	.withColumn("data", from_json(col("value"), schema))

# df = df \
# 	.withColumn("is_valid", col("data").isNotNull()) \
# 	.filter(col("is_valid")) \
# 	.select("data.status")

df = df.withColumn("prediction", predict(col("data")))

query = df \
	.writeStream \
	.outputMode("append") \
	.format("console") \
	.start()

query.awaitTermination()
