# Clipboard Monitor

Este projeto monitora a área de transferência do Windows 11+ e exibe notificações discretas quando novos conteúdos são copiados.

## Requisitos

- Go 1.16 ou superior
- Windows 11+

## Desenvolvimento

1. Clone o repositório: git clone https://github.com/yourusername/clipboard-monitor.git

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

[MIT License](LICENSE)