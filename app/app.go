package app

import "github.com/urfave/cli"

// Gerar retorna a app cli pronta para executar
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App cli"
	app.Usage = "Busca IPs e nome de servidores na internet"

	return app

}
