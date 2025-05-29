# Como Criar o Instalador do Clipboard Monitor

## Requisitos

- Inno Setup

## Passo a Passo

PRÉ-REQUISITOS: Compilar executável do projeto, seguindo orientações do REAMD.md

1. **Baixar e instalar o Inno Setup**:
   - Você pode baixar o Inno Setup  (https://jrsoftware.org/isdl.php).

2. **Abrir o script do Inno Setup**:
   - Abra o Inno Setup e carregue o script `ClipboardMonitor.iss` localizado na pasta `installer`.
   - Atualize a variável "BuildDir" e "ExeDir", conforme a localização dos arquivos de build do seu ambiente.

3. **Compilar o script**:
   - No Inno Setup, vá em `File > Open` e selecione `ClipboardMonitor.iss`.
   - Clique em `Compile` para gerar o instalador `Clipboard_Monitor_v1.0_Setup.exe`.

## Estrutura do Script do Inno Setup

O script `ClipboardMonitor.iss` está configurado para incluir os seguintes arquivos:

- `ClipboardMonitor.exe` (executável principal)
- `buid\data\data.txt` (arquivo de dados)
- `build\assets\icon.ico` (ícone da aplicação)
