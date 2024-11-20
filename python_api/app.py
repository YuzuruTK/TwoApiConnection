from flask import Flask
from flask_cors import CORS


DEBUG_MODE = False # Defina como True para habilitar o modo de depuração

app = Flask(__name__)
CORS(app)

@app.route('/items/<id>', methods=['GET'])
def get_item(id):
    return "jorge"

if __name__ == '__main__':
    app.run(debug=DEBUG_MODE)
