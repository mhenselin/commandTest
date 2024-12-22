package cli

import (
	"fmt"
	"time"
)

import "github.com/spf13/cobra"

func NewTimeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "time",
		Short: "time",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		TraverseChildren:      true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			tm, err := GetTimeInTimezone("Europe/Berlin")
			if err != nil {
				return err
			}
			fmt.Println(tm)
			return nil
		},
	}

	//cmd.AddCommand()
	return cmd
}

func GetTimeInTimezone(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	currentTime := time.Now().In(location)
	return currentTime.Format(time.RFC1123), nil
}
