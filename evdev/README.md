## evdev

This is a keyboard backend for the Linux evdev kernel API.

evdev should be initialized by the host application.
This package expects the host application to take care of
appropriate calls to `PollEvents()` for each input event.
Refer to `keyboard_test.go` for an example.

Nore that this backend does not cover anywhere near all of the keycodes
available in `evdev`. It only covers those defined in the keyboard package.


### Dependencies

    go get github.com/jteeuwen/evdev


### Usage

    go get github.com/jteeuwen/keyboard/evdev


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

