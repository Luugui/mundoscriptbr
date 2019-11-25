# Mundo Script BR
Git destinado aos scripts compartilhados com a comunidade
# Repositório de códigos:

### Get Event

Script para extração de eventos do Zabbix. É só executar o script e preencher com os dados necessários.
É necessário ter as bibliotecas pyzabbix e openpyxl instalados
```
pip install pyzabbix
pip install openpyxl
```


### ZabbixInstall

Um programa escrito em Go para facilitar a instalação do Zabbix Agent em Ambientes Windows. O proppósito além de estudo da linguagem foi gerar um instalador que fosse necessário apenas uma conexão com a internet para baixar, configurar e instalar o serviço do Zabbix Agent em servidores Windows.

```
ZabbixInstall\main.go ----> codigo fonte da aplicação junto com suas dependencias
ZabbixInstall\ZabbixInstall.exe ----> executavel para instalação do Agent
```
