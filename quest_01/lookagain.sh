find -name '*.sh' | cut -d"." -f2 | sort -r | xargs basename -a
#| sed "s|^\./||"