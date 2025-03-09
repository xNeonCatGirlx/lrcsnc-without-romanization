package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"lrcsnc/internal/config"
	"lrcsnc/internal/flags"
	"lrcsnc/internal/log"
	"lrcsnc/internal/mpris"
	"lrcsnc/internal/output/piped"
	"lrcsnc/internal/output/tui"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/sync"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	// Handle the -- flags
	flags.HandleFlags()

	// Start the USR1 signal listener for config updates
	// TODO: replace with live file watcher
	usr1Sig := make(chan os.Signal, 1)
	signal.Notify(usr1Sig, syscall.SIGUSR1)

	go func() {
		for {
			<-usr1Sig
			config.Update()
		}
	}()

	// Initialize the player listener session
	err := mpris.Connect()
	if err != nil {
		log.Fatal("main", "Error when configuring MPRIS. Check logs for more info.")
	}
	defer mpris.Disconnect()

	// Start the main loop
	sync.Loop()

	// Initialize the output
	switch global.Config.C.Global.Output {
	case "piped":
		piped.Init()
		defer piped.Close()

		exitSigs := make(chan os.Signal, 1)
		signal.Notify(exitSigs, syscall.SIGINT, syscall.SIGTERM)

		log.Info("main", "Piped output initialized.")

		<-exitSigs
		log.Info("main", "Exit signal received, bye!")
		os.Exit(0)
	case "tui":
		p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatal("main", "Error running TUI: "+err.Error())
			os.Exit(1)
		}
	}
}
