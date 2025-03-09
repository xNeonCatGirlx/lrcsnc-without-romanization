package mpris

import (
	"fmt"
	"strings"

	"lrcsnc/internal/log"

	"github.com/Endg4meZer0/go-mpris"
	"github.com/godbus/dbus/v5"
)

var conn *dbus.Conn
var player *mpris.Player
var PlayerSignalReceiver = make(chan *dbus.Signal)
var NameOwnerChangedSignalReceiver = make(chan *dbus.Signal)

// Connect connects to DBus and returns any error it encounters.
// The connection is then stored privately in this module.
func Connect() error {
	var err error

	// Get a private connection from DBus
	conn, err = dbus.SessionBusPrivate()
	if err != nil {
		log.Error("mpris/Connect", err.Error())
		return err
	}

	// Do the needed procedures like auth...
	err = conn.Auth(nil)
	if err != nil {
		Disconnect()
		log.Error("mpris/Connect", err.Error())
		return err
	}

	// ...and hello
	err = conn.Hello()
	if err != nil {
		Disconnect()
		return err
	}

	// Also register the NameOwnerChanged signal receiver channel to see
	// if there are new players or old are removed
	err = mpris.RegisterNameOwnerChanged(conn, NameOwnerChangedSignalReceiver)
	if err != nil {
		Disconnect()
		log.Fatal("mpris/Connect", "Cannot watch for player signals. More: "+err.Error())
		return err
	}

	// And now we can get the active player (if there is any)
	err = UpdatePlayer()
	if err != nil {
		log.Error("mpris/Connect", err.Error())
	}

	// And launch the watcher for these exact changes
	go nameOwnerChangesWatcher()

	log.Info("mpris/Connect", "Successfully connected to DBus")

	return nil
}

func nameOwnerChangesWatcher() {
	for {
		s, ok := <-NameOwnerChangedSignalReceiver
		// The channel should not be empty, but we must check to be sure.
		// If it is empty, then most likely it happened because of premature connection close.
		// That means we should reopen the connection.
		if !ok {
			log.Warn("mpris/nameOwnerChangesWatcher", "The channel is closed. Reconnecting to DBus...")
			_ = Connect()
			break
		}

		// Skipping any not NameOwnerChanged signals
		if signalType := mpris.GetSignalType(s); signalType != mpris.SignalNameOwnerChanged {
			log.Debug("mpris/nameOwnerChangesWatcher", fmt.Sprintf("Signal is not NameOwnerChanged (%s). Skipping...", signalType))
			continue
		}

		// Skipping any signals not related to MPRIS
		// See also: https://dbus.freedesktop.org/doc/dbus-specification.html#bus-messages-name-owner-changed
		if !strings.HasPrefix(s.Body[0].(string), mpris.BaseInterface) {
			log.Debug("mpris/nameOwnerChangesWatcher", fmt.Sprintf("Signal is not related to MPRIS (%s). Skipping...", s.Body[0].(string)))
			continue
		}

		err := UpdatePlayer()
		if err != nil {
			log.Error("mpris/nameOwnerChangesWatcher", err.Error())
		}
	}
}

// Disconnect disconnects from DBus. Any error is counted as fatal.
func Disconnect() {
	// We can close the connection since the method
	// also will close all the channels used for signals
	err := conn.Close()
	if err != nil {
		log.Fatal("mpris/Disconnect", err.Error())
	}
	log.Info("mpris/Disconnect", "Successfully disconnected from DBus")
}
