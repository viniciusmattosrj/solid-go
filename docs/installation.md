# Instalação

Faça um git clone do projeto template-go:
```
git@github.com:viniciusmattosrj/template-go.git
```

## Encoding

Todos os arquivos estão em **UTF-8**.


#### Importante para que o git não considere alterações de permissão como modificações a serem rastreadas

```
git config core.fileMode false
```

#### Utilizando o docker ao invés do Host

Caso você utilize docker ao invés do NVM será necessário realizar a cópia do arquivo example:
```
cp -v docker-compose.yml.example docker-compose.yml
cp -v .env.example .env
```

Dentro do projeto **template-go** execute:
```
docker-compose up -d
```