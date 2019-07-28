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
          <div class="col s6 offset-s3 center">
            <div class="row">
              <h4>Preencha os dados abaixo</h4>
            </div>
            <form method="post" action="listsla.php">
              <div class="row">
                <div class="input-field">
                <i class="material-icons prefix">language</i>
                  <input id="dtfrom" type="text" class="datepicker" name="dtfrom">
                  <label for="dtfrom">Date from:</label>
                </div>
              </div>
              <div class="row">
                <div class="input-field">
                  <i class="material-icons prefix">language</i>
                  <input id="dttill" type="text" class="datepicker" name="dttill">
                  <label for="dttill">Date till:</label>
                </div>
              </div>

              <button class="btn waves-effect waves-light" type="submit" name="action">Submit
                <i class="material-icons right">send</i>
              </button>
            </form>
          </div>
        </div>
      </div>

      <!--JavaScript at end of body for optimized loading-->
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
      <script type="text/javascript">

        M.AutoInit();

        $(document).ready(function () {
          $(".datepicker").pickadate({
            closeOnSelect: true,
            format: "dd-mm-yyyy"
          });
        });
      </script>
    </body>
  </html>