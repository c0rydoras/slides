package main

import (
	_ "embed"
	"os"
	"time"

	"github.com/c0rydoras/slides/internal/model"
	"github.com/c0rydoras/slides/internal/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "slides <file.md>",
		Short: "Terminal based presentation tool",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var fileName string

			if len(args) > 0 {
				fileName = args[0]
			}

			presentation := model.Model{
				Page:     0,
				Date:     time.Now().Format("2006-01-02"),
				FileName: fileName,
				Search:   navigation.NewSearch(),
			}
			err = presentation.Load()
			if err != nil {
				return err
			}

			p := tea.NewProgram(presentation, tea.WithAltScreen())
			_, err = p.Run()
			return err
		},
	}
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
