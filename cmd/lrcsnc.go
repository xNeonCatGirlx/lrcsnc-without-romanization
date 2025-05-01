package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"lrcsnc/internal/config"
	"lrcsnc/internal/mpris"
	"lrcsnc/internal/output/piped"
	"lrcsnc/internal/output/tui"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/log"
	"lrcsnc/internal/setup"
	"lrcsnc/internal/sync"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	// Handle all the general setup...
	setup.Setup()
	// ...and check for dependencies
	setup.CheckDependencies()

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

	// Deploy the main watchers
	sync.Start()

	// Initialize the player listener session
	err := mpris.Connect()
	if err != nil {
		log.Fatal("cmd", "Error when configuring MPRIS. Check logs for more info.")
	}
	defer mpris.Disconnect()

	// Initialize the output
	switch global.Config.C.Output.Type {
	case "piped":
		piped.Init()
		defer piped.Close()

		exitSigs := make(chan os.Signal, 1)
		signal.Notify(exitSigs, syscall.SIGINT, syscall.SIGTERM)

		log.Info("cmd", "Piped output initialized.")

		<-exitSigs
		log.Info("cmd", "Exit signal received, bye!")
		os.Exit(0)
	case "tui":
		p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatal("cmd", "Error running TUI: "+err.Error())
			os.Exit(1)
		}
	}
}
