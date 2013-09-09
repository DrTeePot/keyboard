// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package evdev

import (
	"fmt"
	"github.com/jteeuwen/evdev"
	"os"
	"os/signal"
	"testing"
)

func findKeyboard(t *testing.T) *evdev.Device {
	keyboards, err := evdev.Find(evdev.Keyboard)
	if err != nil {
		t.Fatal(err)
	}

	if len(keyboards) == 0 {
		t.Fatal("No keyboards found")
	}

	kb := keyboards[0]

	// Close remaining keyboards. We do not need them.
	for _, dev := range keyboards[1:] {
		dev.Close()
	}

	return kb
}

func Test(t *testing.T) {
	signals := make(chan os.Signal, 1)
	dev := findKeyboard(t)
	defer dev.Close()

	kb := New()
	kb.Bind(func() {
		fmt.Println("pressed s")
	}, "s")

	kb.Bind(func() {
		fmt.Println("pressed ctrl+s or command+s")
	}, "ctrl+s", "command+s")

	kb.Bind(func() {
		fmt.Println("pressed 't e s t'")
	}, "t e s t")

	kb.Bind(func() {
		fmt.Println("pressed '3'")
	}, "3")

	kb.Bind(func() {
		fmt.Println("pressed '#'")
	}, "#")

	kb.Bind(func() {
		signals <- os.Kill
	}, "escape")

	// Print available keybindings.
	fmt.Printf("Known key bindings:\n")
	for _, b := range kb.Bindings() {
		fmt.Printf(" - %q\n", b)
	}

	// Go into event loop.
	signal.Notify(signals, os.Kill, os.Interrupt)

	for {
		select {
		case <-signals:
			return
		case evt := <-dev.Inbox:
			kb.Poll(evt)
		}
	}
}
