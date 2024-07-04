package day9 
import (
 "github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
    day:= &cobra.Command{
        Use:   "9",
        Short: "Run the solutions for day 9",
    }
    
    day.AddCommand(aCommand())
    day.AddCommand(bCommand())
    root.AddCommand(day)
}
