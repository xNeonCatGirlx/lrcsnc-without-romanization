package mpris

import (
	"lrcsnc/internal/pkg/log"

	"github.com/Endg4meZer0/go-mpris"
	"github.com/godbus/dbus/v5"
)

var conn *dbus.Conn
var player *mpris.Player
var playerSignalReceiver = make(chan *dbus.Signal)
var nameOwnerChangedSignalReceiver = make(chan *dbus.Signal)

// Connect connects to D-Bus and returns any error it encounters.
// The connection is then stored privately in this module.
func Connect() error {
	var err error

	// Get a private connection from D-Bus
	conn, err = dbus.SessionBusPrivate()
	if err != nil {
		log.Error("mpris/Connect", err.Error())
		return err
	}
	log.Debug("mpris/Connect", "Got a private connection from D-Bus")

	// Do the needed procedures like auth...
	err = conn.Auth(nil)
	if err != nil {
		Disconnect()
		log.Error("mpris/Connect", err.Error())
		return err
	}
	log.Debug("mpris/Connect", "Authentificated the connection")

	// ...and hello
	err = conn.Hello()
	if err != nil {
		Disconnect()
		return err
	}
	log.Debug("mpris/Connect", "Greeted D-Bus")

	// Also register the NameOwnerChanged signal receiver channel to see
	// if there are new players or old are removed
	err = mpris.RegisterNameOwnerChanged(conn, nameOwnerChangedSignalReceiver)
	if err != nil {
		Disconnect()
		log.Fatal("mpris/Connect", "Cannot watch for player signals. More: "+err.Error())
		return err
	}
	log.Debug("mpris/Connect", "Registered the name owner changed signal receiver channel")

	// And now we can get the active player (if there is any)
	err = ChangePlayer()
	if err != nil {
		log.Error("mpris/Connect", err.Error())
	}
	log.Debug("mpris/Connect", "Got the initial player info from MPRIS")

	MPRISMessageChannel <- Message{SignalReady, nil}
	log.Debug("mpris/Connect", "Sent a SignalReady to MPRISMessageChannel")

	// And deploy the watchers
	go playerSignalWatcher()
	go nameOwnerChangeWatcher()

	log.Debug("mpris/Connect", "Deployed the MPRIS/D-Bus signal watchers")

	log.Info("mpris/Connect", "Successfully connected to D-Bus")

	return nil
}

// Disconnect disconnects from D-Bus. Any error is counted as fatal.
func Disconnect() {
	// We can close the connection since the method
	// also will close all the channels used for signals
	err := conn.Close()
	if err != nil {
		log.Fatal("mpris/Disconnect", err.Error())
	}
	log.Info("mpris/Disconnect", "Successfully disconnected from DBus")
}
