package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delayMonitoramento = 5 * time.Second

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		//comando := lerComando()

		switch lerComando() {
		case 1:
			iniciarMonitoramento()
			break /* não é necessário, por default ele vai dar um break caso ache a opção desejada*/
		case 2:
			imprimirLogs()
		case 0:
			fmt.Println("Saindo...tchau...GoLangers")
			os.Exit(0)
		default:
			fmt.Println("Opção escolhida não é valida.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nomePessoa := "Jackson"
	versao := float32(1.1) /* se não informar o tipo, o Go irá criar como float64 - sempre*/

	fmt.Println("Olá Sr.", nomePessoa, ".")
	fmt.Println("Este programa esta na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var escolha int
	fmt.Scan(&escolha) /* & indica o ponteiro da variavel escolha, para que seja preenchido com o que o usuario informar */
	fmt.Println("O Comando escolhido foi:", escolha)

	return escolha
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando.....")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			fmt.Println("Testando Site", indice, ":", site)
			testaSiteOnline(site)
		}
		time.Sleep(delayMonitoramento)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSiteOnline(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um Erro : ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("site:", site, "foi carregado com sucesso !!")
		registraLog(site, true)
	} else {
		fmt.Println("O Site:", site, "esta com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("listasites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivoLog, err := os.OpenFile("logHello_Go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
	}

	horalog := time.Now().Format("02/01/2006 15:04:05")
	arquivoLog.WriteString(horalog + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")
	arquivoLog.Close()
}

func imprimirLogs() {

	arquivo, err := ioutil.ReadFile("logHello_Go.log") /* ler tudo de uma vez e imprimir */
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
