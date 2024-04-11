package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a app cli pronta para executar
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App cli"
	app.Usage = "Busca IPs e nome de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nome de servidores na internet",
			Flags:  flags,
			Action: buscarServers,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host") //pega flag passada

	//slice de ips
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	//para cada ip no range de ips
	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServers(c *cli.Context) {
	host := c.String("host") //pega flag passada

	servidores, erro := net.LookupNS(host) //server name
	if erro != nil {
		log.Fatal(erro)
	}

	// para cada servidor
	for _, servidores := range servidores {
		fmt.Println(servidores.Host)
	}
}
