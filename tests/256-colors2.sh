for i in {0..256} ; do echo -en "\e[38;5;${i}m#\e[0m" ; done ; echo
for i in {0..256} ; do echo -en "\e[48;5;${i}m \e[0m" ; done ; echo
