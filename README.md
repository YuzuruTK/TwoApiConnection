# Sistemas Distribuídos - Trabalho 4

## Avaliação do trabalho:  

1. Escolher duas linguagens de programação  
2. Analisar como JSON e requisições usando HTTP são implementados nas linguagens
3. Criar dois programas para para trocarem informações por meio de JSON e BAAS:
    - Uma linguagem implementa o método POST
    - Outra linguagem implementa o método GET

## Entrega do trabalho:

> GRUPO de 2 a 4 componentes (criticidade na avaliação proporcional ao número de componentes)

Entregar no `github` uma pasta com o nome dos componentes contendo: 
1. Fontes das implementaçôes (separar as linguagens em pastas)  
1. Passos para instalação e execução  
1. Dois prints da execução:
    - Execução do POST com a linguagem A  
    - Execução GET com a linguagem B

## Especificações do grupo

| Integrantes: | Linguagens: | BaaS: | RESTful: |  
| - | - | - | - | 
| Gabriel Buron | Golang | http (Go)  | net/http (Go) |
| Gustavo Campos | Python | Flask(Python) | Flask (Python) |
| Juan Fricke | | Uploadthing (storage) | Requests (Python) |
| Stefani Arnold |

# TwoApiConnection

## Executando Projeto

### Env File
### Docker Compose

## Documentação Go API

***Explicação API***

---
### Endpoints

***Exemplo de documentação de outra API***
### `POST:/admin/add-user`
Adiciona um usuário a base de dados.

**Corpo Requisição:**
```json
{
    "token": "<token de acesso do usuário>",
    "username": "<string> nome de usuario",
    "password": "<string> senha do usuario",
    "isadmin": "(opcional: false) <bool> define se o usuário será administrador"
}
```

**Resposta:**
```json
{
    "status": "<string> success | error",
    "msg": "<string> mensagem de resposta"
}
```

## Documentação Python API

### Endpoints

***Explicação API***

---
### Endpoints

***Exemplo de documentação de outra API***
### `POST:/admin/add-user`
Adiciona um usuário a base de dados.

**Corpo Requisição:**
```json
{
    "token": "<token de acesso do usuário>",
    "username": "<string> nome de usuario",
    "password": "<string> senha do usuario",
    "isadmin": "(opcional: false) <bool> define se o usuário será administrador"
}
```

**Resposta:**
```json
{
    "status": "<string> success | error",
    "msg": "<string> mensagem de resposta"
}
```