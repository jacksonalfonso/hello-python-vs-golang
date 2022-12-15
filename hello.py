import os 
import pandas as pd
from time import sleep
import requests
import logging
import validators
from datetime import datetime

qtde_monitoramentos = 3
delayMonitoramento = 5

def exibeIntroducao():
    nomePessoa = "Jackson"
    versao = 1.1

    print("Olá Sr.", nomePessoa, ".")
    print("Este programa esta na versão: ", versao)

def exibeMenu():
    print("1- Iniciar Monitoramento")
    print("2- Exibir os Logs")
    print("0- Sair do Programa")

def lerComando():
    escolha = int(input("Escolha uma opção:")) 
    print("O Comando escolhido foi:", escolha)

    return escolha

def leSitesDoArquivo():
    lines = []
    file = 'listasites.txt'
    with open(file) as f:
        [lines.append(line.rstrip('\n')) for line in f.readlines()] 
    
    return lines

def registraLog(site, status):
    
    logging.basicConfig(filename='logHello_Py.log', format='%(asctime)s - %(name)s - %(levelname)s - %(message)s', level = logging.INFO)

    if status:
        logging.info(f"Testando {site} - Online {status}")
    else:
        logging.error(f"Testando {site} - Online {status}")

def testaSiteOnline(site_para_teste):
    
    if validators.url(site_para_teste):
        res=requests.get(site_para_teste)
        if res.status_code == 200:
            print("site:", site_para_teste, "foi carregado com sucesso !!")
            registraLog(site_para_teste, True)
        else:
            print("O Site:", site_para_teste, "esta com problemas. Status code:", res.status_code)
            registraLog(site_para_teste, False)
    else:
        logging.error(f"Testando: {site_para_teste} - Status: URL é Inválida")

def iniciarMonitoramento():

    print("Monitorando.....")
    sites = list(leSitesDoArquivo())

    for i in range(qtde_monitoramentos):
        for site in sites:
            print("Testando Site", i, ":", site)
            testaSiteOnline(site)

        sleep(delayMonitoramento)
        print("")

    print("")

def imprimirLogs():
    try:
        with open('logHello_Py.log', 'r') as f:
            print(f.read())
    except IOError:
        print(u'Arquivo de log não encontrado!')

if __name__ == '__main__':
    
    exibeIntroducao()

    while True:
        exibeMenu()

        match lerComando():
            case 1:
                iniciarMonitoramento()
            case 2:
                imprimirLogs()
            case 0:
                print("Saindo...tchau...Pythoniacos")
                os._exit(0)
            case _:
                print("Opção escolhida não é valida.")
                os._exit(-1)
	