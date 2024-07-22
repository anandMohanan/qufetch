package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/anandMohanan/qufetch/color"
	"github.com/anandMohanan/qufetch/style"
	"github.com/anandMohanan/qufetch/util"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"github.com/anandMohanan/qufetch/app"
	"github.com/anandMohanan/qufetch/filesystem"
	"github.com/anandMohanan/qufetch/icon"
	"github.com/anandMohanan/qufetch/where"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("version", "v", false, app.Name+" version")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   strings.ToLower(app.Name),
	Short: app.DescriptionShort,
	Long:  app.DescriptionLong,
	Run: func(cmd *cobra.Command, args []string) {
		if lo.Must(cmd.Flags().GetBool("version")) {
			versionCmd.Run(versionCmd, args)
		}
		systemInfo := struct {
			OS     string
			Kernel string
			Uptime string
			Pkgs   string
			App    string
			Memory string
			User   string
		}{
			OS:     util.OS(),
			Kernel: util.Kernel(),
			Uptime: util.Uptime(),
			Pkgs:   util.Pkgs(),
			Memory: util.Memory(),
			App:    app.Name,
			User:   util.User(),
		}

		t, err := template.New("system").Funcs(map[string]any{
			"faint":   lipgloss.NewStyle().Foreground(color.Black).Render,
			"bold":    lipgloss.NewStyle().Foreground(color.HiWhite).Render,
			"magenta": lipgloss.NewStyle().Foreground(color.Yellow).Render,
		}).Parse(`{{ magenta "▇▇▇" }} {{ magenta .User }} 

  {{ faint "OS" }}       {{ bold .OS }}
  {{ faint "Kernel" }}   {{ bold .Kernel }}
  {{ faint "Uptime" }}   {{ bold .Uptime }}
  {{ faint "Pkgs" }}     {{ bold .Pkgs }}
  {{ faint "Memory" }}   {{ bold .Memory }}
`)
		handleErr(err)
		handleErr(t.Execute(cmd.OutOrStdout(), systemInfo))

	},
}

func Execute() {
	// Setup colored cobra
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	// Clears temp files on each run.
	// It should not affect startup time since it's being run as goroutine.
	go func() {
		_ = filesystem.Api().RemoveAll(where.Temp())
	}()

	_ = rootCmd.Execute()
}

// handleErr will stop program execution and logger error to the stderr
// if err is not nil
func handleErr(err error) {
	if err == nil {
		return
	}

	log.Error(err)
	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s %s\n",
		style.Failure(icon.Cross),
		strings.Trim(err.Error(), " \n"),
	)
	os.Exit(1)
}
