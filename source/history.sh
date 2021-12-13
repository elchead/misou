#!/bin/bash
export PATH="/Users/adria/homebrew/bin:/usr/bin:/bin/"
ha() {
		local cols sep google_history open
		cols=$(( COLUMNS / 3 ))
		sep='{::}'

		if [ "$(uname)" = "Darwin" ]; then
		  google_history="$HOME/Library/Application Support/Google/Chrome/Default/History"
		  open=open
		else
		  google_history="$HOME/.config/google-chrome/Default/History"
		  open=xdg-open
		fi
		cp -f "$google_history" /tmp/h
		sqlite3 -separator $sep /tmp/h \
		  "select title, url
		   from urls order by last_visit_time desc" |
		awk -F $sep '{printf "%s;%s\n", $1, $2}' |
		rg -S -w $1 || true # ignore error when empty
	  }
ha $1
