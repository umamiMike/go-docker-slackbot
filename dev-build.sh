#!/bin/bash

build_cmd () {
$(go build) > /dev/null && echo -e "\e[1m\e[36mapp rebuilt successfully\e[0m" || echo -e "\n\n\e[1m[91mrebuild failed\n\n"
./testbot & 
disown
}

killall testbot  &> /dev/null

rebuild () {
killall testbot &> /dev/null
build_cmd
fswatch --event Updated --one-event $(pwd) 
rebuild 
 }

rebuild

