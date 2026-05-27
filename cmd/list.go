package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/madhavan-raja/autorun/pb"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls", "l"},
	Short: "List all Processes",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getConn()
		if err != nil {
			logger.Error("Cannot Get gRPC Connection", "err", err)
			return
		}
		defer conn.Close()

		c := pb.NewArDaemonClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.List(ctx, &pb.ListRequest{})
		if err != nil {
			logger.Error("Cannot List Processes", "err", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)

		fmt.Fprintln(w, "ID\tName\tDescription\tCommand\tInterval")
		for _, p := range r.GetProcesses() {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\n", p.Id, p.Name, p.Description, p.Command, p.Interval)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
