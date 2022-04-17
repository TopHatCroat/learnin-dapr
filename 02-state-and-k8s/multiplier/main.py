import time
import requests
import os

dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)
dapr_url = "http://localhost:{}/increment".format(dapr_port)

print("Starting multiplier app...")
print("dapr_port: {}".format(dapr_port))
print("dapr_url: {}".format(dapr_url))

while True:
    try:
        resp = requests.post(
            dapr_url,
            timeout=5,
            headers = {"dapr-app-id": "counter-app"}
        )
        print("Called counter, now waiting...")
        time.sleep(5)
    except Exception as e:
        print(e)
        time.sleep(5)
