package repo

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/wot-oss/tmc/cmd"
	"github.com/wot-oss/tmc/cmd/completion"
	"github.com/wot-oss/tmc/internal/app/cli"
)

var attachmentImportCmd = &cobra.Command{
	Use:   "import <tmNameOrId> <attachmentFile>",
	Short: "Import an attachment",
	Long:  `Add or replace an attachment`,
	Args:  cobra.ExactArgs(2),
	Run:   attachmentImport,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		switch len(args) {
		case 0:
			return completion.CompleteTMNamesOrIds(cmd, args, toComplete)
		case 1:
			return comps, cobra.ShellCompDirectiveDefault
		default:
			comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments")
			return comps, cobra.ShellCompDirectiveNoFileComp
		}
	},
}

func attachmentImport(command *cobra.Command, args []string) {
	spec := cmd.RepoSpec(command)
	mediaType := command.Flag("media-type").Value.String()

	err := cli.AttachmentImport(context.Background(), spec, args[0], args[1], mediaType)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	attachmentCmd.AddCommand(attachmentImportCmd)
	attachmentImportCmd.Flags().StringP("media-type", "m", "", "Media type of the attachment")
}
