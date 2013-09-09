// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
This is a keyboard backend for the Linux evdev kernel API.

evdev should be initialized by the host application.
This package expects the host application to take care of
appropriate calls to `PollEvents()` for each input event.
Refer to `keyboard_test.go` for an example.
*/
package evdev
