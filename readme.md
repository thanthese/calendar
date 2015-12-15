Toggle STDIN between two different custom plain text calendar agenda formats.

# Philosophy

It's easiest to read a plain text calendar if there's some helpful indenting, but it's easiest to write tools to process it if each line has the same format. This utility lets you have the best of both worlds by toggling between the two.

Before

    15.12.16w to boldly
    15.12.17r go
    15.12.19s where no one
    15.12.19s has gone
    15.12.21m before

After

    15.12.16w
        to boldly
    15.12.17r
        go
    15.12.18f
    15.12.19s
        has gone
        where no one
    15.12.20u

    15.12.21m
        before

Blank lines are inserted between weeks (which start on Mondays). There's also some special logic involving prints dates before or near today which mostly exist to satisfy my own peculiarities.

# Installation

    $ go get https://github.com/thanthese/toggle-calendar

# License

MIT
