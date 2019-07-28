# Report Zabbix

![Report Zabbix](https://img.shields.io/badge/Report%20Zabbix-v.01-information)![Report Zabbix](https://img.shields.io/badge/Zabbix-v%204.x-red)![Report Zabbix](https://img.shields.io/badge/PHP-v7-blue)

O **Report Zabbix** é uma ferramenta para auxiliar na extração de eventos e da disponibilidade dos eventos. A aplicação tem três tipos de relatórios para serem gerados. 

- **Hosts** - Uma listagem com todos os ICs cadastrados na Ferramenta.

![Tela inicial](https://i.imgur.com/1vcBVAD.png)

- **Eventos** - Listagem com todos os eventos gerados no periodo selecionado

![Tela de Eventos](https://i.imgur.com/SAYmu7L.png)

- **SLA** - Listagem geral da SLA dos eventos gerados no periodo selecionado

![Tela de SLA](https://i.imgur.com/TXXzOge.png)



# Arquivos

A pasta do **Report Zabbix** não precisa ficar obrigatoriamente junto ao frontend do Zabbix. Porém precisa estar no mesmo servidor para acesso de algumas funções do proprio frontend do Zabbix.

### Configuração

Baixe a pasta do **Report Zabbix** e cole junto ao frontend do Zabbix. Para conexão com o zabbix e coleta dos dados é necessário um usuario de acesso da API com permissão para vizualização de todos os hosts.

### Arquivo config.inc.php

Este é o único arquivo que precisa ser editado para se adequar ao seu ambiente. Nele você insere a URL do seu Zabbix, usuario e senha. Também insere o path onde seu frontend esta configurado:


    <?php
    
    //CONFIGURABLE
    # zabbix server info(user must have API access)
    
    $z_server = 'http://zabbixserver/zabbix/api_jsonrpc.php';
    $z_user = 'Admin';
    $z_pass = 'zabbix';
    $z_path = '/usr/share/zabbix';
    
    ?>
