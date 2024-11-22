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

### Arquivo de ambiente (.env)

Acesse o painel do UploadThing e crie uma nova chave de API.

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```env
UPLOADTHING_TOKEN="YOUR_API_KEY"
```

### Docker Compose

```bash
docker-compose up
```


## Documentação APIs
- **GO:** A api gerencia a criação de NPCs/Monstros do RPG Print Weaver.
- **Python:** Lê quantas fichas estão salvas no bucket(/itens) e lê o conteúdo de uma ficha específica(/intens/id).

---
### Endpoints:
## API Go
### `GET|POST:/generate/monster/beast`
### `GET|POST:/generate/monster/human`

**Corpo Requisição:**
```json
{
    "save": true|false
}
```

## API Python
### `GET:/items`

> exemplo de resposta: 

```json
{"items":["human_monster_22-Nov-2024_231624.json","classes.json"]}

```
### `GET:/items/{id}>`

> exemplo de resposta:

```json
{"armor":"Filigree Tabard (as Medium Armor), a two-tone embroidered cloth coat.","disposition":"Curious","goals":"Recruit forces to hunt down a fearsome local legend.","ring":"A random ring from 41-50 (1d10) in Appendix B.","stats":{"dextery":4,"strength":8,"vitality":16,"willpower":4},"trait":"Knows one random Scroll. Any damage dealt by it instead heals an equal amount.","trinket":"Ankle Weight, Advantage on STR checks.","weapon":"Trident (as Medium Weapon). Deals Cold damage."}

```

