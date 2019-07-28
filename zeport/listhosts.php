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

      <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/v/dt/jq-3.3.1/jszip-2.5.0/dt-1.10.18/b-1.5.6/b-html5-1.5.6/b-print-1.5.6/datatables.min.css"/>

      <!--Let browser know website is optimized for mobile-->
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    </head>

    <body>
      
      <nav>
        <div class="nav-wrapper blue">
          <a href="zabbixdados.php" class="brand-logo center">Report Zabbix | API v.<?=$version?></a>
        </div>
        <ul id="nav-mobile" class="left hide-on-med-and-down">
            <li><a href="index.php">Sair</a></li>
        </ul>
      </nav>
      <div class="container">
        <div class="row">
          <div class="col s10 center">

          <table id="hosts"> 
            <thead>
              <tr>
                  <th>Host ID</th>
                  <th>Hostname</th>
                  <th>Grupo</th>
              </tr>
            </thead>

            <tbody>
              
              <?php

                  $hosts = $api->hostGet(array(
                    'output' => 'extend',
                    'selectGroups' => ['name'],
                  ));

                  foreach ($hosts as $host) {
                    echo "<tr>";
                    echo "<td>$host->hostid</td>";
                    echo "<td>$host->host</td>";
                    echo "<td>".$host->groups[0]->name."</td>";
                    echo "</td>";
                  }

              ?>
              
            </tbody>
          </table>
              
                 <?php

                    $hosts = $api->hostGet(array(
                      'output' => 'extend'
                    ));

                    foreach ($hosts as $host) {
                      
                    }
                 
                 ?>
                        
                    
                    
          </div>
        </div>
      </div>

      <!--JavaScript at end of body for optimized loading-->
      <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
      <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.36/pdfmake.min.js"></script>
      <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.36/vfs_fonts.js"></script>
      <script type="text/javascript" src="https://cdn.datatables.net/v/dt/jq-3.3.1/jszip-2.5.0/dt-1.10.18/b-1.5.6/b-html5-1.5.6/b-print-1.5.6/datatables.min.js"></script>
      <script type="text/javascript">
          $(document).ready(function() {
              $('#hosts').DataTable( {
                  dom: 'Bfrtip',
                  buttons: [
                      'copy', 'csv', 'excel', 'pdf', 'print'
                  ]
              } );
          } );
      </script>
    </body>
  </html>