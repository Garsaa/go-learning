# Go Learning

Projeto simples para aprender a sintaxe de Go.

## A linguagem

**Escrita em**: Go foi boostrappado em C, então o primeiro compilador de Go foi escrito em C, depois foi reescrito em Go e foi utilizado esse compilador feito em C para compilá-lo e apartir dai as próximas versões foram escritas em Go e foi utilizado as versões anteriores para compilá-las.

## Arquivos importantes

**go.mod**: Arquivo para configurar o projeto e as dependências nele. É utilizado para indicar a versão da linguagem, caminho do projeto e indicar as bibliotecas externas que ele depende. Equivalente do package.json.
**main.go**: Apesar do nome do arquivo não importar, para um arquivo poder ser compilado e por tanto executado em Go é preciso que ele tenha package main e uma função main() sem parâmetros e sem retorno

## Comandos úteis
go run main.go ou go run ./calculadora(diretorio caso dentro do diretorio tennha um package main com função main    )


## Conceitos úteis

**package**: Os códigos são organizados em pacotes, esses pacotes são determinados onde acabam e começam por pastas. Funções/Tipos/Variáveis/Constantes/Campos(dentro de tipos)  que começam com letra maiúscula são exportadas e outros pacotes podem usar, funções que começam em minúsculo são privadas somente ao pacote.