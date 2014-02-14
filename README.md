# Multitran Requester

Multitran Requester let you get translation of word from English to Russian using [www.multitran.ru](http://www.multitran.ru). It may
be useful to make Flashcards to memorize word's translation.

## Installation

You need Go 1.1 installation. Also you will need to install go-charset packages to deal with windows-1251:

    go get code.google.com/p/go-charset/charset

## Usage

It's simple:

    go run multitran.go <word>

Example:

    go run multitran.go hello

Output:

    hello    приветственный возглас; оклик; восклицание удивления;

Also you can use helper script to prepare batch of cards to [upload](http://orangeorapple.com/Flashcards/Upload.aspx) to [Flashcards Deluxe website](http://orangeorapple.com/Flashcards/):

    ./words.sh <list-of-words-file>

## License

Distributed under the MIT license.
