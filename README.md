# Go Learning

Projeto simples para aprender a sintaxe de Go.

## A linguagem

**Escrita em**: Go foi boostrappado em C, então o primeiro compilador de Go foi escrito em C, depois foi reescrito em Go e foi utilizado esse compilador feito em C para compilá-lo e apartir dai as próximas versões foram escritas em Go e foi utilizado as versões anteriores para compilá-las.

## Conceitos úteis

**package**: Os códigos são organizados em pacotes, esses pacotes são determinados onde acabam e começam por pastas. Funções/Tipos/Variáveis/Constantes/Campos(dentro de tipos)  que começam com letra maiúscula são exportadas e outros pacotes podem usar, funções que começam em minúsculo são privadas somente ao pacote.

## Arquivos importantes

**go.mod**: Arquivo para configurar o projeto e as dependências nele. É utilizado para indicar a versão da linguagem, caminho do projeto e indicar as bibliotecas externas que ele depende. Equivalente do package.json.

**main.go**: Apesar do nome do arquivo não importar, para um arquivo poder ser compilado e por tanto executado em Go é preciso que ele tenha package main e uma função main() sem parâmetros e sem retorno

**go.sum**: Arquivo que tem checksums criptografados das dependências externas baixadas pelo projeto. Ele serve pra garantir que uma dependência externa não foi alterada ou corrompida. Deve ser versionado no git. Não é um arquivo de lock, as versões das dependências continuam sendo determinadas pelo go.mod e pelo algoritmod e seleção de versões do Go.

## Comandos úteis
go run main.go ou go run ./calculadora(diretorio caso dentro do diretorio tennha um package main com função main    )

**go mod tidy**: Comando que arruma/sincroniza as dependências do módulo.
    - Adiciona ao go.mod os módulos necessários pelos imports do projeto
    - Remove dependências que não tão sendo utilizadas
    - Adiciona dependências indiretas quando precisa
    - Cria ou atualiza o go.sum com os checksums pra verificar as dependências

**go install .**: Comando que compila o módulo e instala ele no path, fica disponível pra ser executado apartir do nome do modulo, ex.: nesse caso seria go install e ai go-learning pq eh o nome do módulo. Instala no GOBIN. Também da pra usar para instalar programas publicados.

**which (nome-do-módulo)**: Comando que lista o path do binario de 1 programa instalado fora da pasta do próprio projeto.

**go list -m all**: Lista todos módulos instalados no m

**go get (pacote@referencia)**: Comando para adicionar/atualizar dependências ao projeto, adiciona ao cache global do Go, se outro projeto pedir o mesmo pacote na mesma versão reutiliza e não baixa de novo. Os pacotes nunca são

**go clean**: Comando que limpa o executável criado no diretorio local.
    Flags: 
        * -modcache: Apague todo o GOMODCACHE que eh o teu cache global de pacotes Go instalados.
        * -cache: Apaga o cache global de compilação
        * -testcache: Invalida os resultados armazenados de testes, executando os testes novamente sem reutilizar resultado nenhum
        * -cache -testcache -modcache: combina tudo, da pra ver oq seria deletado com isso usando go clean -n

**go mod why (nome-do-modulo)
