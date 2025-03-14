# Clipboard Monitor

Monitora a área de transferência do Windows e exibe notificações caso o conteúdo copiado esteja na lista pré-definida para ser noticado.
Ex.:
10.71.12.0/24 | 01 - Matriz (São Paulo)
10.71.13.0/24 | 02 - Filia de Osasco
10.71.14.0/24 | 03 - Brasília DF

Ao copiar algum IP, ele identificará e exibirá a notificação informando o resultado após o "|".


## Requisitos

- Go 1.16 ou superior
- Windows 10 e 11

## Desenvolvimento

1. Clone o repositório: git clone https://github.com/andersoncorso/clipboard-monitor

2. Entre no diretório do projeto: cd clipboard-monitor

3. Instale as dependências: go mod tidy

4. Renomeie o arquivo "build/data/data.example.txt" para ".../data.txt" (adicione as informações necessárias)


## Teste

Execute o programa com o seguinte comando:
```sh
go run cmd/clipboard-monitor/main.go
```

O programa começará a monitorar a área de transferência e exibirá notificações quando novos conteúdos forem copiados.


## Compilação
Para compilar o executável, use o seguinte comando:

```sh
go build -ldflags "-H=windowsgui" -o ClipboardMonitor.exe cmd/clipboard-monitor/main.go
```


## Como Criar o Instalador Windows

Para orientações sobre como criar o instalador usando o Inno Setup, consulte o arquivo installer/INSTALLER.md.



## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.