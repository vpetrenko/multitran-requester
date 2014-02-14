#!/bin/sh

if test -z "$1"
then
echo "Usage: $0 <list-of-words-file>"
exit 1
fi

for WORD in `cat $1`
do
   go run multitran.go $WORD
done
