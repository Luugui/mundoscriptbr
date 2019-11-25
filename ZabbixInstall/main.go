package main

import (
    "net/http"
    "fmt"
    "github.com/mattn/go-colorable"
    "github.com/dimiro1/banner"
    "os"
    "io"
    "io/ioutil"
    "strings"
    "github.com/dustin/go-humanize"
  	"os/exec"
    "bytes"
    "runtime"
    "github.com/AlecAivazis/survey"
    "archive/zip"
    "path/filepath"
)

func header() {
	isEnabled := true
  isColorEnabled := true
  banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(`
{{ .AnsiColor.Red }},-------.          ,--.    ,--.    ,--.              ,--.                    ,--.            ,--. ,--.
{{ .AnsiColor.Red }}'--.   /   ,--,--. |  |-.  |  |-.  '--' ,--.  ,--.   |  | ,--,--,   ,---.  ,-'  '-.  ,--,--. |  | |  |
{{ .AnsiColor.Red }}  /   /   ' ,-.  | | .-. ' | .-. ' ,--.  \  ''  /    |  | |      \ (  .-'  '-.  .-' ' ,-.  | |  | |  |
{{ .AnsiColor.Red }} /   '--. \ '-'  | | '-' | | '-' | |  |  /  /.  \    |  | |  ||  | .-'  ')   |  |   \ '-'  | |  | |  |
{{ .AnsiColor.Red }}'-------'  '--'--'  '---'   '---'  '--' '--'  '--'   '--' '--''--' '----'    '--'    '--'--' '--' '--'
{{ .AnsiColor.Default }}

OS: {{ .GOOS }}
Now: {{ .Now "Monday, 2 Jan 2006" }}

`))
}

func CallClear(system string) {
	if system == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type WriteCounter struct {
    Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
    n := len(p)
    wc.Total += uint64(n)
    wc.PrintProgress()
    return n, nil
}

func (wc WriteCounter) PrintProgress() {
    // Clear the line by using a character return to go back to the start and remove
    // the remaining characters by filling it with spaces
    fmt.Printf("\r%s", strings.Repeat(" ", 50))

    // Return again and print current status of download
    // We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
    fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(url string, filepath string) error {

    // Create the file with .tmp extension, so that we won't overwrite a
    // file until it's downloaded fully
    out, err := os.Create(filepath + ".tmp")
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create our bytes counter and pass it to be used alongside our writer
    counter := &WriteCounter{}
    _, err = io.Copy(out, io.TeeReader(resp.Body, counter))
    if err != nil {
        return err
    }

    // The progress use the same line so print a new line once it's finished downloading
    fmt.Println()

    // Rename the tmp file back to the original file
    err = os.Rename(filepath+".tmp", filepath)
    if err != nil {
        return err
    }

    return nil
}

func Unzip(src string, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)

        // Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            // Make Folder
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        // Make File
        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        // Close the file without defer to close before next iteration of loop
        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}

func AgentInstall(version string) {

  fmt.Printf("Selecionada versão: %s\n", version)
  path := `C:\Zabbix`
  if _, err := os.Stat(path); os.IsNotExist(err) {
    os.Mkdir(path, os.ModePerm)
    fmt.Printf("Pasta %s criada", path)
  }

  switch version {
  case "3.4":
    // URL de Download do pacote do Zabbix Agent
    fileUrl := `https://www.zabbix.com/downloads/3.4.6/zabbix_agents_3.4.6.win.zip`

    // Download do Arquivo
    DownloadFile(fileUrl, "zabbix_agents_3.4.6.win.zip")
    oldName := "zabbix_agents_3.4.6.win.zip.tmp"
    newName := path+`\zabbix_agents_3.4.6.win.zip`

    // Movendo arquivo para a pasta C:\zabbix
    err := os.Rename(oldName, newName)
    if err != nil {
      fmt.Println(err)
    }

    // Descompactando Arquivos zipados
    files, err := Unzip(newName, path)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

    // Caminho do arquivo de configuração
    confFile := `C:\Zabbix\conf\zabbix_agentd.win.conf`

    // Limpando a Tela e começando a editar o arquivo de configuração
    CallClear(runtime.GOOS)
    header()
    server_ip := ""
    prompt := &survey.Input{
      Message: "Insira o ip do seu Zabbix Server: ",
    }
    survey.AskOne(prompt, &server_ip)

    input, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output := bytes.Replace(input, []byte("127.0.0.1"), []byte(server_ip), -1)

    if err = ioutil.WriteFile(confFile, output, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println("==> Server configurado")

    rc := false
    rc_prompt := &survey.Confirm{
      Message: "Deseja habilitar os comandos remotos?",
    }
    survey.AskOne(rc_prompt, &rc)

    if (rc == true) {
      input, err := ioutil.ReadFile(confFile)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      output_enable := bytes.Replace(input, []byte("# EnableRemoteCommands=0"), []byte("EnableRemoteCommands=1"), -1)
      if err = ioutil.WriteFile(confFile, output_enable, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("==> Comandos remotos habilitados")
    }

    hostName, err := os.Hostname()
    if err != nil {
      fmt.Println(err)
    }

    input_hostname, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_hostname := bytes.Replace(input_hostname, []byte("Windows host"), []byte(hostName), -1)
    if err = ioutil.WriteFile(confFile, output_hostname, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    input_log, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_log := bytes.Replace(input_log, []byte(`LogFile=c:\zabbix_agentd.log`), []byte(`LogFile=C:\Zabbix\zabbix_agentd.log`), -1)
    if err = ioutil.WriteFile(confFile, output_log, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("==> Hostname %s configurado\n", hostName)
    if(runtime.GOARCH == "amd64"){
      bin64 := path+`\bin\win64\zabbix_agentd.exe`

      cmd_install := exec.Command(bin64, "-c", confFile ,"-i")
  		cmd_install.Stdout = os.Stdout
  		cmd_install.Run()
      fmt.Println("==> Serviço instalado")

      cmd_run := exec.Command(bin64, "--start")
      cmd_run.Stdout = os.Stdout
      cmd_run.Run()
      fmt.Println("==> Serviço iniciado")
    }

    fmt.Println("Pressione qualquer tecla para sair....")
    fmt.Scanln()


  case "4.0":

    // URL de Download do pacote do Zabbix Agent
    fileUrl := `https://www.zabbix.com/downloads/4.0.14/zabbix_agents-4.0.14-win-amd64.zip`

    // Download do Arquivo
    DownloadFile(fileUrl, "zabbix_agents-4.0.14-win-amd64.zip")
    oldName := "zabbix_agents-4.0.14-win-amd64.zip.tmp"
    newName := path+`\zabbix_agents-4.0.14-win-amd64.zip`

    // Movendo arquivo para a pasta C:\zabbix
    err := os.Rename(oldName, newName)
    if err != nil {
      fmt.Println(err)
    }

    // Descompactando Arquivos zipados
    files, err := Unzip(newName, path)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

    // Caminho do arquivo de configuração
    confFile := `C:\Zabbix\conf\zabbix_agentd.conf`

    // Limpando a Tela e começando a editar o arquivo de configuração
    CallClear(runtime.GOOS)
    header()
    server_ip := ""
    prompt := &survey.Input{
      Message: "Insira o ip do seu Zabbix Server: ",
    }
    survey.AskOne(prompt, &server_ip)

    input, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output := bytes.Replace(input, []byte("127.0.0.1"), []byte(server_ip), -1)

    if err = ioutil.WriteFile(confFile, output, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println("==> Server configurado")

    rc := false
    rc_prompt := &survey.Confirm{
      Message: "Deseja habilitar os comandos remotos?",
    }
    survey.AskOne(rc_prompt, &rc)

    if (rc == true) {
      input, err := ioutil.ReadFile(confFile)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      output_enable := bytes.Replace(input, []byte("# EnableRemoteCommands=0"), []byte("EnableRemoteCommands=1"), -1)
      if err = ioutil.WriteFile(confFile, output_enable, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("==> Comandos remotos habilitados")
    }

    hostName, err := os.Hostname()
    if err != nil {
      fmt.Println(err)
    }

    input_hostname, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_hostname := bytes.Replace(input_hostname, []byte("Windows host"), []byte(hostName), -1)
    if err = ioutil.WriteFile(confFile, output_hostname, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    input_log, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_log := bytes.Replace(input_log, []byte(`LogFile=c:\zabbix_agentd.log`), []byte(`LogFile=C:\Zabbix\zabbix_agentd.log`), -1)
    if err = ioutil.WriteFile(confFile, output_log, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("==> Hostname %s configurado\n", hostName)

      bin64 := path+`\bin\zabbix_agentd.exe`

      cmd_install := exec.Command(bin64, "-c", confFile ,"-i")
  		cmd_install.Stdout = os.Stdout
  		cmd_install.Run()
      fmt.Println("==> Serviço instalado")

      cmd_run := exec.Command(bin64, "--start")
      cmd_run.Stdout = os.Stdout
      cmd_run.Run()
      fmt.Println("==> Serviço iniciado")


    fmt.Println("Pressione qualquer tecla para sair....")
    fmt.Scanln()
  case "4.2":
        // URL de Download do pacote do Zabbix Agent
        fileUrl := `https://www.zabbix.com/downloads/4.2.8/zabbix_agents-4.2.8-win-amd64.zip`

        // Download do Arquivo
        DownloadFile(fileUrl, "zabbix_agents-4.2.8-win-amd64.zip")
        oldName := "zabbix_agents-4.2.8-win-amd64.zip.tmp"
        newName := path+`\zabbix_agents-4.2.8-win-amd64.zip`

        // Movendo arquivo para a pasta C:\zabbix
        err := os.Rename(oldName, newName)
        if err != nil {
          fmt.Println(err)
        }

        // Descompactando Arquivos zipados
        files, err := Unzip(newName, path)
        if err != nil {
          fmt.Println(err)
        }
        fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

        // Caminho do arquivo de configuração
        confFile := `C:\Zabbix\conf\zabbix_agentd.conf`

        // Limpando a Tela e começando a editar o arquivo de configuração
        CallClear(runtime.GOOS)
        header()
        server_ip := ""
        prompt := &survey.Input{
          Message: "Insira o ip do seu Zabbix Server: ",
        }
        survey.AskOne(prompt, &server_ip)

        input, err := ioutil.ReadFile(confFile)
        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        output := bytes.Replace(input, []byte("127.0.0.1"), []byte(server_ip), -1)

        if err = ioutil.WriteFile(confFile, output, 0666); err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        fmt.Println("==> Server configurado")

        rc := false
        rc_prompt := &survey.Confirm{
          Message: "Deseja habilitar os comandos remotos?",
        }
        survey.AskOne(rc_prompt, &rc)

        if (rc == true) {
          input, err := ioutil.ReadFile(confFile)
          if err != nil {
            fmt.Println(err)
            os.Exit(1)
          }

          output_enable := bytes.Replace(input, []byte("# EnableRemoteCommands=0"), []byte("EnableRemoteCommands=1"), -1)
          if err = ioutil.WriteFile(confFile, output_enable, 0666); err != nil {
            fmt.Println(err)
            os.Exit(1)
          }

          fmt.Println("==> Comandos remotos habilitados")
        }

        hostName, err := os.Hostname()
        if err != nil {
          fmt.Println(err)
        }

        input_hostname, err := ioutil.ReadFile(confFile)
        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        output_hostname := bytes.Replace(input_hostname, []byte("Windows host"), []byte(hostName), -1)
        if err = ioutil.WriteFile(confFile, output_hostname, 0666); err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        input_log, err := ioutil.ReadFile(confFile)
        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        output_log := bytes.Replace(input_log, []byte(`LogFile=c:\zabbix_agentd.log`), []byte(`LogFile=C:\Zabbix\zabbix_agentd.log`), -1)
        if err = ioutil.WriteFile(confFile, output_log, 0666); err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        fmt.Printf("==> Hostname %s configurado\n", hostName)

          bin64 := path+`\bin\zabbix_agentd.exe`

          cmd_install := exec.Command(bin64, "-c", confFile ,"-i")
      		cmd_install.Stdout = os.Stdout
      		cmd_install.Run()
          fmt.Println("==> Serviço instalado")

          cmd_run := exec.Command(bin64, "--start")
          cmd_run.Stdout = os.Stdout
          cmd_run.Run()
          fmt.Println("==> Serviço iniciado")


        fmt.Println("Pressione qualquer tecla para sair....")
        fmt.Scanln()
  case "4.4":
    // URL de Download do pacote do Zabbix Agent
    fileUrl := `https://www.zabbix.com/downloads/4.4.1/zabbix_agents-4.4.1-win-amd64.zip`

    // Download do Arquivo
    DownloadFile(fileUrl, "zabbix_agents-4.4.1-win-amd64.zip")
    oldName := "zabbix_agents-4.4.1-win-amd64.zip.tmp"
    newName := path+`\zabbix_agents-4.4.1-win-amd64.zip`

    // Movendo arquivo para a pasta C:\zabbix
    err := os.Rename(oldName, newName)
    if err != nil {
      fmt.Println(err)
    }

    // Descompactando Arquivos zipados
    files, err := Unzip(newName, path)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

    // Caminho do arquivo de configuração
    confFile := `C:\Zabbix\conf\zabbix_agentd.conf`

    // Limpando a Tela e começando a editar o arquivo de configuração
    CallClear(runtime.GOOS)
    header()
    server_ip := ""
    prompt := &survey.Input{
      Message: "Insira o ip do seu Zabbix Server: ",
    }
    survey.AskOne(prompt, &server_ip)

    input, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output := bytes.Replace(input, []byte("127.0.0.1"), []byte(server_ip), -1)

    if err = ioutil.WriteFile(confFile, output, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println("==> Server configurado")

    rc := false
    rc_prompt := &survey.Confirm{
      Message: "Deseja habilitar os comandos remotos?",
    }
    survey.AskOne(rc_prompt, &rc)

    if (rc == true) {
      input, err := ioutil.ReadFile(confFile)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      output_enable := bytes.Replace(input, []byte("# EnableRemoteCommands=0"), []byte("EnableRemoteCommands=1"), -1)
      if err = ioutil.WriteFile(confFile, output_enable, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("==> Comandos remotos habilitados")
    }

    hostName, err := os.Hostname()
    if err != nil {
      fmt.Println(err)
    }

    input_hostname, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_hostname := bytes.Replace(input_hostname, []byte("Windows host"), []byte(hostName), -1)
    if err = ioutil.WriteFile(confFile, output_hostname, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    input_log, err := ioutil.ReadFile(confFile)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    output_log := bytes.Replace(input_log, []byte(`LogFile=c:\zabbix_agentd.log`), []byte(`LogFile=C:\Zabbix\zabbix_agentd.log`), -1)
    if err = ioutil.WriteFile(confFile, output_log, 0666); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Printf("==> Hostname %s configurado\n", hostName)

      bin64 := path+`\bin\zabbix_agentd.exe`

      cmd_install := exec.Command(bin64, "-c", confFile ,"-i")
      cmd_install.Stdout = os.Stdout
      cmd_install.Run()
      fmt.Println("==> Serviço instalado")

      cmd_run := exec.Command(bin64, "--start")
      cmd_run.Stdout = os.Stdout
      cmd_run.Run()
      fmt.Println("==> Serviço iniciado")


    fmt.Println("Pressione qualquer tecla para sair....")
    fmt.Scanln()
  }



}

func DesinstallAgent() {
  path_c := `C:\zabbix`
  path_file := os.Getenv("ProgramFiles")
  path_file = path_file+`\Zabbix Agent`
  path_file32 := os.Getenv("ProgramFiles(x86)")
  path_file32 = path_file32+`\Zabbix Agent`

  var option []string

  if _, err := os.Stat(path_c); !os.IsNotExist(err) {
    fmt.Printf("Pasta %s encontrada\n",path_c)
    option = append(option, path_c)

  }

  if _, err := os.Stat(path_file); !os.IsNotExist(err) {
    fmt.Printf("Pasta %s encontrada\n",path_file)
    option = append(option, path_file)
  }

  if _, err := os.Stat(path_file32); !os.IsNotExist(err) {
    fmt.Printf("Pasta %s encontrada\n",path_file32)
    option = append(option, path_file32)
  }

  option = append(option, "Outro")

  folder := ""
  prompt := &survey.Select{
    Message: "Escolha um folder",
    Options: option,
  }
  survey.AskOne(prompt, &folder)

  if(folder == "Outro"){
    other := ""
    prompt := &survey.Input{
      Message: "Insira o caminho:",
    }
    survey.AskOne(prompt, &other)
  } else {
    zabbix_agent := folder+`\zabbix_agentd.exe`
    if _, err := os.Stat(zabbix_agent); os.IsNotExist(err) {
      zabbix_agent = folder+`\bin\zabbix_agentd.exe`
    }

    fmt.Println(zabbix_agent)

    cmd_stop := exec.Command(zabbix_agent, "-x")
    cmd_stop.Stdout = os.Stdout
    cmd_stop.Run()
    fmt.Println("==> Serviço parado")

    cmd_run := exec.Command(zabbix_agent, "-d")
    cmd_run.Stdout = os.Stdout
    cmd_run.Run()
    fmt.Println("==> Serviço desinstaldo")

    os.RemoveAll(folder)
    fmt.Println("==> Pasta removida")
  }



}

func main() {
    CallClear(runtime.GOOS)
    header()

    var choose string

  	prompt := &survey.Select{
  		Message: "Escolha uma opção: ",
      Options: []string{"Instalar", "Desistalar", "Sair"},
    }
  	survey.AskOne(prompt, &choose)
      switch choose {
      case "Instalar":
        CallClear(runtime.GOOS)
        header()
        var version string

        prompt := &survey.Select{
          Message: "Escolha uma versão: ",
          Options: []string{"3.4", "4.0", "4.2", "4.4"},
        }
        survey.AskOne(prompt, &version)

        AgentInstall(version)

      case "Desistalar":
        DesinstallAgent()
      case "Sair":
        fmt.Println("Pressione uma tecla para sair...")
        fmt.Scanln()

    }


}
