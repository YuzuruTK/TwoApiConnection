import json
from os import getenv
from sys import getsizeof
import requests
from dotenv import load_dotenv

load_dotenv()

JSON_OBJ = {"mano": "batata"}
PRELOAD_URL = "https://api.uploadthing.com/v7/prepareUpload"
UPLOADTHING_TOKEN = getenv("UPLOADTHING_TOKEN")

preload_reponse = requests.post(
    url=PRELOAD_URL, 
    headers={
        "Content-Type": "application/json",
        "X-Uploadthing-Api-Key": UPLOADTHING_TOKEN
    }, 
    json={
        "fileName": "mano.json",
        "fileSize": getsizeof(json.dumps(JSON_OBJ)),
        "fileType": "json",
        "contentDisposition": "inline"
    }
)

preload = preload_reponse.json()
print(preload)

upload = requests.put(
    url=preload["url"], 
    headers={},
    files={'file': ('mano.json', json.dumps(JSON_OBJ), 'application/json')}
)

print(upload.json())
