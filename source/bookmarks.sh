#!/bin/bash
export PATH="/Users/adria/homebrew/bin:/usr/bin"
ba() {
     bookmarks_path=~/Library/Application\ Support/Google/Chrome/Default/Bookmarks

     jq_script='
        def ancestors: while(. | length >= 2; del(.[-1,-2]));
        . as $in | paths(.url?) as $key | $in | getpath($key) | {name,url, path: [$key[0:-2] | ancestors as $a | $in | getpath($a) | .name?] | reverse | join("/") } | .path + "/" + .name + "\t" + .url'

    jq -r "$jq_script" < "$bookmarks_path" \
        | sed -E $'s/(.*)\t(.*)/\\1\t\\2/g' | sed 's/\// /g' | sed 's/_/ /g' | rg -S -w $1 || true # ignore error when empty
}
ba $1
