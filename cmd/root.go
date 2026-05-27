package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cc "github.com/ivanpirog/coloredcobra"
)

var port int32
var logger *slog.Logger

func init() {
	logger = internal.Logger
}

func getConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

var rootCmd = &cobra.Command{
	Use:   "autorun",
	Short: "Automatically run scripts on a schedule",
}

func Execute() {
	cc.Init(&cc.Config{
        RootCmd:  rootCmd,
        Headings: cc.Bold,
        Commands: cc.Bold,
        Example:  cc.Italic,
        ExecName: cc.Bold,
        Flags:    cc.Bold,
    })

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int32Var(&port, "port", 5678, "Port")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


