package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func Start() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplication search server"
	app.Usage = "Nomes do Servidor da Internet"
	flags := []cli.Flag { cli.StringFlag{ Name: "host" ,},

	}
	app.Commands = []cli.Command{
		{
			Name: "dns",
			Usage: "Pegar dns",
			Flags: flags,
			Action: findDns,
		},
		{
			Name: "ip",
			Usage: "Pegar IP",
			Flags: flags,
			Action: findIP,
		},
	}
return app
}

func findIP(app *cli.Context){
	host := app.String("host")
	ip, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("findIP error err: ", err)
		return
	}
	for _, value := range ip {
		fmt.Println("IP:",value)
	}
}
func findDns(app *cli.Context){
host := app.String("host")
ns, err := net.LookupNS(host)
if err != nil {
	fmt.Println("findDNs error err: ", err)
	return
}
for _, value := range ns {
	fmt.Println("DNS:",value.Host)
}

}