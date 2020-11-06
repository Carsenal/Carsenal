#/bin/bash

watch -n 1 "cmark-gfm \
    --unsafe \
    -e footnotes \
    -e table \
    -e strikethrough \
    -e autolink \
    -e tagfilter \
    -e tasklist \
    README.md > tmp.html"

