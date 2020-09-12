package cmd

import (
	"github.com/chriswalz/bit/util"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "bit switch [branch-name]",
	Long: `For existing branches simply run bit switch [branch-name]. 

For creating a new branch it's the same command! You'll simply be prompted to confirm that you want to create a new branch
`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Runwithcolor([]string{"fetch"})
		if !localOrRemoteBranchExists(args[0]) {
			resp := promptUser("Branch does not exist. Do you want to create it? Y/n")
			if isYes(resp) {
				util.Runwithcolor([]string{"checkout", "-b", args[0]})
			}
		}
	},
	Args: cobra.ExactArgs(1),
}

func isYes(resp string) bool {
	return resp == "YES" || resp == "Y" || resp == "yes" || resp == "y"
}

func init() {
	rootCmd.AddCommand(switchCmd)
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func localOrRemoteBranchExists(branch string) bool {
	msg, err := exec.Command("git", "checkout", branch).CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	return !strings.Contains(string(msg), "did not match any file")
}