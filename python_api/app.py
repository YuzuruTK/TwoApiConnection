import requests # Pcs conversando na mesma lingua
from flask import Flask # importando flask e cors           
from flask_cors import CORS
from os import getenv
from dotenv import load_dotenv

load_dotenv()


# CONSTANTES E VARIAVEIS --------------------------------------------------------------------------------------

URL = "https://api.uploadthing.com/v6/listFiles" # constante URL
DEBUG_MODE = False # Defina como True para habilitar o modo de depuração
APITOKEN = getenv("ACESS_TOKEN_UPLOADTHING_OVOSECRETO")
PADRAOURL = "https://utfs.io/f/"

# FLASK -------------------------------------------------------------------------------------------

app = Flask(__name__) # flask tranforma o teu pc num servidor http
CORS(app) # cross origin fodace n sei pesquisar oq é depois (outros pcs podem usar esse link)

# batata ------------------------------------------------------------------------------------------

@app.route("/items", methods=["GET"]) # definir as rotas de acesso ao servidor ; rota de acesso /items
def list_items(): # funcao que define a lista de itens
    
    payload = {} # conteudo da caixa
    headers = { # rotulo da caixa
        "Content-Type": "application/json", # tipo de conteudo (json)
        "X-Uploadthing-Api-Key": APITOKEN # token de acesso da api
    }

    # requisicao post pra url, com aqule payload e aquele header configurados acima 
    response = requests.post(URL, json=payload, headers=headers) 

    return response.json()

@app.route('/items/<id>', methods=['GET']) # rota de acesso /tems/id
def get_item(id): # funcao que define como pega os itens 
    resposta = requests.get(PADRAOURL+id)
    return resposta.json()
    

if __name__ == '__main__':
    app.run(debug=DEBUG_MODE)

# REQUESTS ----------------------------------------------------------------------------------------
