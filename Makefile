all: smoketest

smoketest: FORCE
	mkdir -p out
	tinygo build -o ./out/tinygo-lunakey-left.uf2  --target ./firmware/lunakey-pico.json --size short ./firmware/left/
	tinygo build -o ./out/tinygo-lunakey-right.uf2 --target ./firmware/lunakey-pico.json --size short ./firmware/right/

flash-left: FORCE
	tinygo flash --target ./firmware/lunakey-pico.json --size short ./firmware/left/

flash-right: FORCE
	tinygo flash --target ./firmware/lunakey-pico.json --size short ./firmware/right/

FORCE:

gen-def-with-find:
	find . -name vial.json | xargs -n 1 go run github.com/sago35/tinygo-keyboard/cmd/gen-def

gen-def:
	go run github.com/sago35/tinygo-keyboard/cmd/gen-def ./firmware/left/vial.json

test-gen-def: gen-def-with-find
	git status
	test -z "$$(git status -s)"
