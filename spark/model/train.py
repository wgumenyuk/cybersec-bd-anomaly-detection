import pickle
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import classification_report

normal_path = "dataset/normal_100.txt"
bruteforce_path = "dataset/bruteforce_100.txt"
ddos_path = "dataset/ddos_100.txt"

df_normal = pd.read_json(normal_path, lines=True)
df_bruteforce = pd.read_json(bruteforce_path, lines=True)
df_ddos = pd.read_json(ddos_path, lines=True)

df_normal["label"] = 0
df_normal["label"] = 1
df_normal["label"] = 2

bruteforce_sample = df_bruteforce.sample(frac=0.7, random_state=42)
ddos_sample = df_ddos.sample(frac=0.7, random_state=42)

data = pd.concat([bruteforce_sample, ddos_sample, df_normal], ignore_index=True)

x = data.drop(columns=["label"])
x = pd.get_dummies(x, drop_first=True)
x = x.fillna(0)

y = data["label"]
y = y.fillna(0)

x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.3, random_state=42)

model = RandomForestClassifier(random_state=42)

model.fit(x_train, y_train)
print("Training complete")

y_pred = model.predict(x_test)
print(classification_report(y_test, y_pred))

with open("model.pkl", "wb") as model_file:
	pickle.dump(model, model_file)

print("Model saved")
