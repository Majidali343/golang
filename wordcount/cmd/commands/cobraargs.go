package cobraargs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	
)

var (
	Routines int
	FileName   string
)


var RootCmd = &cobra.Command{
	Use:   "wordcount",
	Short: "My application will manuplate the file and data inside the file ",
	Long:  `Taking inpits for file name and goroutines to run the program.`,
	Run: func(cmd *cobra.Command, args []string) {
	
		fmt.Printf("Routines: %d\nFile Name: %s\n", Routines, FileName)
	},
}

func init() {

	RootCmd.Flags().IntVarP(&Routines, "number", "n", 0, "Routines")
	RootCmd.Flags().StringVarP(&FileName, "file", "f", "", "File name")


	RootCmd.MarkFlagRequired("number")
	RootCmd.MarkFlagRequired("file")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
