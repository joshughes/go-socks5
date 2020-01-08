/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joshughes/go-socks5/pkg/socks5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//ServerOpts Server Options
type ServerOpts struct {
	host string
	port string
}

func init() {
	so := ServerOpts{}
	// serveCmd represents the serve command
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			so.RunServer(cmd, args)
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringVarP(&so.host, "host", "", "127.0.0.1", "Host to bind to")
	serveCmd.Flags().StringVarP(&so.port, "port", "p", "8080", "Port to bind to")

	rootCmd.AddCommand(serveCmd)
}

// RunServer starts a Socks5 server
func (to ServerOpts) RunServer(cmd *cobra.Command, args []string) {
	// viper.GetStringMapStringSlice("users")
	users := viper.GetStringMapString("users")
	creds := socks5.StaticCredentials{}
	for k, v := range users {
		creds[k] = v
	}

	conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	if len(creds) > 0 {
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		conf.AuthMethods = []socks5.Authenticator{cator}
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting Proxy server on " + to.host + ":" + to.port)
	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", to.host+":"+to.port); err != nil {
		panic(err)
	}
}
