import requests # Pcs conversando na mesma lingua
from flask import Flask # importando flask e cors           
from flask_cors import CORS
from os import getenv
from dotenv import load_dotenv

load_dotenv()


# CONSTANTES E VARIAVEIS --------------------------------------------------------------------------------------

DEBUG_MODE = False # Defina como True para habilitar o modo de depuração
LIST_FILES_URL = "https://api.uploadthing.com/v6/listFiles" # constante URL
FILE_URL_PATTERN = "https://utfs.io/f/"
APITOKEN = getenv("UPLOADTHING_TOKEN")

# FLASK -------------------------------------------------------------------------------------------

app = Flask(__name__) # flask tranforma o teu pc num servidor http
CORS(app) # cross origin fodace n sei pesquisar oq é depois (outros pcs podem usar esse link)

# Functions ---------------------------------------------------------------------------------------
def request_items():
    payload = {} # conteudo da caixa
    headers = { # rotulo da caixa
        "Content-Type": "application/json", # tipo de conteudo (json)
        "X-Uploadthing-Api-Key": APITOKEN # token de acesso da api
    }
    # requisicao post pra url, com aqule payload e aquele header configurados acima 
    response = requests.post(LIST_FILES_URL, json=payload, headers=headers) 
    return response.json()

# Routes ------------------------------------------------------------------------------------------
@app.route("/items", methods=["GET"]) # definir as rotas de acesso ao servidor ; rota de acesso /items
def get_items(): # funcao que define a lista de itens
    items = request_items()
    return {"items": list(map(lambda n: n["name"], items["files"]))}

@app.route('/items/<id>', methods=['GET']) # rota de acesso /tems/id
def get_item(id): # funcao que define como pega os itens
    print(id)
    fmt_id = (f"{id}.json", id)[id.endswith(".json")]

    items = request_items()["files"]
    wanted_item = next((d for d in items if d["name"] == fmt_id), None) # Procura primeiro item com nome igual a id
    
    if wanted_item is None:
        return {"error": "Item not found"}, 404

    resposta = requests.get(f"{FILE_URL_PATTERN}{wanted_item['key']}")
    return resposta.json()
    

if __name__ == '__main__':
    app.run(debug=DEBUG_MODE)
