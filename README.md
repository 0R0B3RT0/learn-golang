Configuge VSCode to use Go: https://code.visualstudio.com/docs/languages/go

Go commands:
- go env: Exibe as variáveis de ambiente;
  - go env -w para alterar o valor da variável de ambiente 
  - Configuração da variável de ambiente `GOBIN` ex. `go env -w GOBIN=$HOME/Projects/goworkspace/bin`
- go help: `go help [command]` ensina a utilizar o comando;
- go fmt: 
  - `go fmt main.go` para formatar um arquivo específico;
  - `go fmt .` para formtar todos os arquivos do diretório;
- go run: 
  - `go run main.go` (arquivo);
  - `go run *.go`;
- go build:
  - efetua o build;
  - exibe os erros;
  - executa os testes;
  - cria o executável em caso de sucesso;
- go install:
  - compila o arquivo;
  - gera o executável no diretório configurado na variável de ambiente `GOBIN`;
