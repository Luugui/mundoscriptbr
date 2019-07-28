<?php
include("config.inc.php");
require_once 'lib/ZabbixApi.class.php';
require_once $z_path.'/include/config.inc.php';
require_once $z_path.'/include/hosts.inc.php';

// load ZabbixApi
use ZabbixApi\ZabbixApi;

              
 // connect to Zabbix API
 $api = new ZabbixApi($z_server, $z_user, $z_pass);

 $version = $api->apiinfoVersion();

 $dt_from = new DateTime($_POST["dtfrom"]);
 $time_from = $dt_from->format('U');
 $dt_till = new DateTime($_POST["dttill"]);
 $time_till = $dt_till->format('U');

 $events = $api->eventGet(array(
  'output'    => 'extend',
  'time_from' => $time_from,
  'time_till' => $time_till,
  'value'     => 1,
  'severities'=> [4,5],
  'selectHosts' => ['name']
));

$event_day = array();

foreach($events as $day){
    if(in_array($day->objectid, $event_day)) {
       continue;
   } else {
        $triggerid = $day->objectid;
        $get_percent = calculateAvailability($triggerid, $time_from, $time_till);
        $event_day[$triggerid] = array(
            "host" => $day->hosts[0]->name,
            "name" => $day->name,
            "ok" => sprintf('%.2f%%', $get_percent['false']),
            "problem" => sprintf('%.2f%%', $get_percent['true']));    
       }
       
       
}


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
          <div class="col s12 center">

          <table id="events" class="highlight"> 
            <thead>
              <tr>
                  <th>Host</th>
                  <th>Trigger</th>
                  <th>SLA Ok</th>
                  <th>SLA Problem</th>
              </tr>
            </thead>

            <tbody>
                <?php
                  
                  foreach ($event_day as $key => $valor) {

                    echo "<tr>";
                    foreach($valor as $key => $valor) {
                        
                      echo "<td>".$valor."</td>";

                    }                    
                    echo "</tr>";

                  }

                ?>
                  
                        
            </tbody>
          </table>
              
                 
                        
                    
                    
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
              $('#events').DataTable( {
                  dom: 'Bfrtip',
                  buttons: [
                      'copy', 'csv', 'excel', 'pdf', 'print'
                  ],
                  responsive: true
              } );
          } );
      </script>
    </body>
  </html>