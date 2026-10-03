# Funcionalidades

## 1 - Iniciar Monitoramento

Faz uma requisição HTTP GET para um site e informa se ele foi carregado com sucesso (status 200) ou se está com problema.

- Site verificado atualmente: `https://www.alura.com.br` (fixo no código, em [main.go](../main.go))
- Implementado em `startMonitoring()`

## 2 - Exibir Logs

Opção reservada para exibição de logs de monitoramento. Ainda não implementada — apenas exibe uma mensagem indicando a ação.

## 0 - Sair do Programa

Finaliza a execução do programa com código de saída `0`.

## Comando inexistente

Qualquer opção digitada fora das listadas no menu (1, 2 ou 0) exibe a mensagem `Comando inexistente!` e finaliza o programa com código de saída `-1`.
