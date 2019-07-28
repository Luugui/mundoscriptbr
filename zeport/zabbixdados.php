<?php
require_once 'lib/ZabbixApi.class.php';
include("config.inc.php");

// load ZabbixApi
use ZabbixApi\ZabbixApi;

              
 // connect to Zabbix API
 $api = new ZabbixApi($z_server, $z_user, $z_pass);

 $version = $api->apiinfoVersion();
?>

<!DOCTYPE html>
  <html>
    <head>
      <!--Import Google Icon Font-->
      <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
      <!--Import materialize.css-->
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">

      <!--Let browser know website is optimized for mobile-->
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    </head>

    <body>
      
      <nav>
        <div class="nav-wrapper blue">
          <a href="zabbixdados.php" class="brand-logo center">Report Zabbix | API v.<?=$version?></a>
        </div>
      </nav>
      <div class="container">
        <div class="row">
          <div class="col s12 center">


            <div class="col s12 m4">
                <div class="icon-block">
                    <h2 class="center light-blue-text">
                    <i class="material-icons">computer</i>
                    </h2>
                    <h5 class="center">Listagem de Hosts</h5>
                    <p class="light">Gera uma listagem com todos os hosts cadastrados no Zabbix.</p>
                    <a class="waves-effect waves-light btn" href="listhosts.php">Selecionar</a>
                </div>
            </div>
            <div class="col s12 m4">
                <div class="icon-block">
                    <h2 class="center light-blue-text">
                    <i class="material-icons">event</i>
                    </h2>
                    <h5 class="center">Listagem de Eventos</h5>
                    <p class="light">Gera uma listagem com todos os eventos gerados no period</p>
                    <a class="waves-effect waves-light btn" href="form-events.php">Selecionar</a>
                </div>
            </div>
            <div class="col s12 m4">
                <div class="icon-block">
                    <h2 class="center light-blue-text">
                    <i class="material-icons">insert_chart</i>
                    </h2>
                    <h5 class="center">Listagem de Eventos SLA</h5>
                    <p class="light">Gera uma listagem com o SLA dos eventos do periodo.</p>
                    <a class="waves-effect waves-light btn" href="form-sla.php">Selecionar</a>
                </div>
            </div>
                        
                    
            </div>
          </div>
        </div>
      </div>

      <!--JavaScript at end of body for optimized loading-->
      <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
    </body>
  </html>