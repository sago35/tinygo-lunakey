# tinygo-keyboard firmware for Lunakey Pico

This repository is for creating unofficial firmware for the Lunakey Pico.
The firmware is created using sago35/tinygo-keyboard.

## build / flash

```shell
# build
$ make smoketest

# flash
# When performing a tinygo flash for the first time, or when changing the USB VID/PID settings,
# please transition to Bootloader mode before executing.
$ tinygo flash --target ./firmware/lunakey-pico.json --size short ./firmware/left/
$ tinygo flash --target ./firmware/lunakey-pico.json --size short ./firmware/right/
```

## Lunakey Pico

* Project
    * https://remap-keys.app/catalog/Kg5ww7zULE3TTt9PSQzw
* Schema / Board
    * https://github.com/yoichiro/lunakey
    * [open with kicanvs](https://kicanvas.org/?github=https%3A%2F%2Fgithub.com%2Fyoichiro%2Flunakey%2Ftree%2Fmain%2Fpico%2Frev2%2Fpcb)

## Link

* https://github.com/sago35/tinygo-keyboard
