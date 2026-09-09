# Toolchain e comandos úteis

Este documento reúne a referência de toolchain.
Não precisa decorar todos os comandos: `go help <comando>` é comum utilizar.
Toolchain é o conjunto de ferramentas usadas via CLI para buildar, compilar, analisar, executar, transformar, testar código go(geralmente) também existem ferramentas mais baixo nível nesse conjunto

## Convenção dos exemplos de terminal

Os blocos marcados como **PowerShell e Bash** funcionam da mesma forma nos dois
shells. Eles usam `shell` como identificador Markdown por ser um nome genérico
para comandos de terminal. Muitos desses comandos também funcionam em Zsh e
outros shells compatíveis, mas Bash é a referência de Linux deste documento.

Quando nomes de executáveis, caminhos ou comandos do sistema operacional
diferirem, serão mostrados blocos separados `powershell` e `bash`.

## Modelo mental: módulo e pacote

- Um **módulo** é a unidade versionada descrita por um `go.mod`.
- Um módulo contém um ou mais **pacotes**.
- Normalmente, os arquivos `.go` de um mesmo diretório formam um pacote.
- Um pacote chamado `main` com `func main()` representa um **comando** capaz de
  gerar um executável.
- Um pacote comum é uma biblioteca e não pode ser executado diretamente.

Identificadores iniciados por letra maiúscula são exportados pelo pacote:

```go
func Somar(a, b int) int // exportada
func subtrair(a, b int) int // não exportada
```

“Não exportado” é mais preciso que “privado”: todos os arquivos do mesmo pacote
podem acessar o identificador.

## Convenções de nomes e arquivos

- Pacotes usam nomes curtos, minúsculos e, normalmente, uma única palavra.
- Arquivos usam nomes curtos e minúsculos; `_` pode separar partes quando isso
  melhora a leitura.
- Funções e variáveis não exportadas normalmente usam `camelCase`.
- Identificadores exportados usam `PascalCase`.
- Arquivos e diretórios iniciados por `.` ou `_` são ignorados pelo toolchain.

Sufixos com comportamento especial:

- `nome_test.go`: incluído somente nos testes.
- `nome_windows.go`: incluído quando `GOOS=windows`.
- `nome_linux.go`: incluído quando `GOOS=linux`.
- `nome_amd64.go`: incluído quando `GOARCH=amd64`.
- `nome_windows_amd64.go`: exige simultaneamente Windows e AMD64.

Exemplos de valores possíveis incluem `windows`, `linux` e `darwin` para
`GOOS`, e `amd64`, `arm64` e `riscv64` para `GOARCH`.

![Extensões reconhecidas pelo toolchain](../assets/images/extensions.png)

Go também pode trabalhar com Assembly, C via cgo e alguns outros formatos.

## Arquivos e diretórios importantes

### `go.mod`

Define, entre outras configurações:

- O caminho do módulo.
- A versão mínima de Go declarada pelo módulo.
- Dependências diretas e indiretas.
- Substituições ou exclusões de módulos, quando existirem.

O módulo deste repositório é:

```go
module github.com/Garsaa/go-learning
```

### `go.sum`

Armazena hashes criptográficos usados para autenticar versões de dependências.
Deve ser versionado. Não é um arquivo de lock: as versões selecionadas são
determinadas pelo `go.mod` e pelo grafo de módulos.

### `GOBIN`

Diretório em que `go install` grava comandos compilados. Quando `GOBIN` está
vazio, o destino padrão normalmente é o diretório `bin` do primeiro `GOPATH`.

**PowerShell e Bash:**

```shell
go env GOBIN GOPATH
```

### Caches

**PowerShell e Bash:**

```shell
go env GOCACHE GOMODCACHE
```

- `GOCACHE`: resultados intermediários de compilação e testes.
- `GOMODCACHE`: módulos externos baixados e seus códigos-fonte.

## Padrões de pacotes

Muitos comandos aceitam os mesmos padrões:

```text
.       pacote do diretório atual
./...   pacote atual e pacotes abaixo dele, dentro do módulo
all     pacotes relevantes do módulo e de suas dependências
```

Exemplos:

**PowerShell e Bash:**

```shell
go test ./...
go build ./cmd/calculadora
go vet ./projects/...
```

Um diretório abaixo que possua outro `go.mod` pertence a outro módulo e não é
incluído pelo `./...` do módulo atual.

## Executar, compilar e instalar

### `go run`

Compila temporariamente e executa um comando:

**PowerShell e Bash:**

```shell
go run ./cmd/calculadora
```

Prefira informar o pacote (`.` ou um diretório) em vez de um único arquivo. Por
exemplo, `go run .` considera todos os arquivos que compõem o pacote atual,
enquanto `go run main.go` considera somente os arquivos listados.

### `go build`

Verifica e compila pacotes:

**PowerShell e Bash:**

```shell
go build ./...
```

Para escolher o nome do executável, a extensão convencional difere:

**PowerShell:**

```powershell
go build -o calculadora.exe ./cmd/calculadora
```

**Bash no Linux:**

```bash
go build -o calculadora ./cmd/calculadora
```

A flag `-o` escolhe o arquivo ou diretório de saída.

### `go install`

Compila e instala um comando em `GOBIN`:

**PowerShell e Bash:**

```shell
go install ./cmd/calculadora
go install golang.org/x/tools/cmd/goimports@latest
```

O sufixo `@versão` permite instalar uma ferramenta publicada sem adicionar a
ferramenta como dependência normal do projeto.

Para localizar o comando instalado, use o mecanismo do seu shell:

**PowerShell:**

```powershell
Get-Command calculadora
where.exe calculadora
```

**Bash:**

```bash
command -v calculadora
```

Em Bash, `command -v` é preferível a `which` porque é um recurso do próprio
shell e também reconhece aliases e funções.

## Formatação e análise estática

### `go fmt`

Formata os pacotes no padrão da linguagem:

**PowerShell e Bash:**

```shell
go fmt ./...
```

Para inspecionar ou controlar o `gofmt` diretamente:

**PowerShell e Bash:**

```shell
gofmt -d .
gofmt -w .
```

- `-d`: exibe as diferenças.
- `-w`: grava o resultado nos arquivos.

### `goimports`

Formata o código e também adiciona, remove e organiza imports:

**PowerShell e Bash:**

```shell
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### `go vet`

Procura construções suspeitas que compilam, mas provavelmente representam um
erro:

**PowerShell e Bash:**

```shell
go vet ./...
```

## Testes

Arquivos de teste terminam em `_test.go`. Um teste unitário possui a forma:

```go
func TestSomar(t *testing.T) {
	got := Somar(2, 3)
	want := 5

	if got != want {
		t.Errorf("Somar(2, 3) = %d; esperado %d", got, want)
	}
}
```

Comandos frequentes:

**PowerShell e Bash:**

```shell
go test                         # pacote atual
go test ./projects/calculadora # pacote específico
go test ./...                   # todos os pacotes abaixo
go test -v ./...                # saída detalhada
go test -run '^TestSomar$' ./... # seleciona por expressão regular
go test -count=1 ./...          # não reutiliza resultado em cache
go test -race ./...             # detector de data races
go test -cover ./...            # resumo de cobertura
```

Relatório de cobertura:

**PowerShell e Bash:**

```shell
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

O detector de corridas só encontra problemas nos caminhos executados durante o
teste. A ausência de relatório não prova que todo o programa está livre de
data races.

## Dependências e módulos

### `go get`

Ajusta dependências no `go.mod`; não é o comando usado para instalar programas:

**PowerShell e Bash:**

```shell
go get github.com/google/uuid@v1.6.0
go get github.com/google/uuid@latest
go get github.com/google/uuid@none
go get -u ./...
```

### `go mod tidy`

Sincroniza os imports do código com `go.mod` e `go.sum`:

**PowerShell e Bash:**

```shell
go mod tidy
go mod tidy -diff
```

`-diff` mostra o que seria alterado e não grava as mudanças.

### Outros comandos de módulos

**PowerShell e Bash:**

```shell
go mod download
go mod verify
go mod graph
go mod why -m github.com/google/uuid
```

- `download`: baixa os módulos necessários.
- `verify`: verifica se módulos existentes no cache foram alterados após o
  download.
- `graph`: imprime uma relação `módulo -> requisito` por linha.
- `why -m`: mostra por que um módulo é necessário.

## Inspeção com `go list`

Sem `-m`, `go list` trabalha com pacotes:

**PowerShell e Bash:**

```shell
go list .
go list ./...
go list -deps ./cmd/calculadora
go list -json ./projects/calculadora
```

Com `-m`, trabalha com módulos ativos no projeto atual:

**PowerShell e Bash:**

```shell
go list -m all
go list -m -u all
go list -m -versions github.com/google/uuid
go list -m -json all
```

Flags importantes:

- `-json`: saída estruturada e detalhada.
- `-deps`: inclui pacotes importados transitivamente.
- `-compiled`: informa arquivos efetivamente enviados ao compilador.
- `-m`: lista módulos, em vez de pacotes.
- `-u`: consulta atualizações de módulos.
- `-versions`: mostra versões conhecidas de um módulo.

`go list -m all` não lista tudo que existe no computador. Ele mostra o módulo
principal e os módulos selecionados no grafo atual.

## Documentação com `go doc`

Mostra a documentação de pacotes e símbolos:

**PowerShell e Bash:**

```shell
go doc fmt
go doc fmt.Println
go doc github.com/Garsaa/go-learning/projects/calculadora
go doc github.com/Garsaa/go-learning/projects/calculadora.Somar
```

Para documentar um pacote e uma função, escreva comentários imediatamente antes
das declarações. O comentário de um identificador exportado começa com o nome
dele:

```go
// Package calculadora fornece operações matemáticas simples.
package calculadora

// Somar retorna a soma de a e b.
func Somar(a, b int) int {
	return a + b
}
```

Não pode haver uma linha vazia entre o comentário e a declaração. Comentários
de documentação devem ser frases úteis e completas.

Flags:

**PowerShell e Bash:**

```shell
go doc -all github.com/Garsaa/go-learning/projects/calculadora
go doc -u github.com/Garsaa/go-learning/projects/calculadora
```

- `-all`: mostra toda a documentação do pacote.
- `-u`: inclui identificadores não exportados.

## Ambiente com `go env`

Exibe a configuração efetiva do toolchain:

**PowerShell e Bash:**

```shell
go env
go env GOOS GOARCH GOROOT GOPATH GOBIN
go env GOMOD GOMODCACHE GOCACHE GOPROXY GOPRIVATE
go env -json
go env -changed
```

Configurações persistentes podem ser gravadas e removidas com cuidado:

**PowerShell e Bash:**

```shell
go env -w NOME=valor
go env -u NOME
```

## Limpeza de caches

**PowerShell e Bash:**

```shell
go clean
go clean -cache
go clean -testcache
go clean -modcache
go clean -fuzzcache
go clean -n -cache -testcache -modcache
```

- `-cache`: apaga o cache de compilação.
- `-testcache`: invalida resultados de testes.
- `-modcache`: apaga módulos externos baixados e extraídos.
- `-fuzzcache`: apaga entradas preservadas pelo fuzzing.
- `-n`: mostra o que seria executado sem remover.

Limpar caches manualmente não deve fazer parte da rotina normal; o Go os
gerencia e recompila ou baixa novamente quando necessário.

## Ajuda e informações de versão

**PowerShell e Bash:**

```shell
go help
go help test
go help testflag
go version
```

Para inspecionar um executável já compilado:

**PowerShell:**

```powershell
go version -m ./calculadora.exe
```

**Bash no Linux:**

```bash
go version -m ./calculadora
```

`go version -m` inspeciona informações de build e módulos incorporadas em um
executável Go.

## Redirecionamento do shell

Redirecionamento não é uma funcionalidade específica do Go; é fornecido pelo
shell:

**PowerShell e Bash:**

```shell
go env > ambiente.txt
go test ./... > testes.txt
go test ./... >> historico-testes.txt
```

- `>` cria ou sobrescreve o arquivo.
- `>>` acrescenta ao final do arquivo.
