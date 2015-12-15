Toggle STDIN between two different custom plain text calendar agenda formats.

# Philosophy

It's easiest to read a plain text calendar if there's some helpful indenting, but it's easiest to write tools to process it if each line has the same format. This utility lets you have the best of both worlds by toggling between the two.

Before ("regular")

    15.12.16w to boldly
    15.12.17r go
    15.12.19s where no one
    15.12.19s has gone
    15.12.21m before

After ("irregular")

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

Regular mode only lists dates where events exist. Irregular mode always lists all dates.

Blank lines are inserted between weeks (which start on Mondays). Lines are sorted alphabetically after each transformation. There's also some special logic involving printing dates before or near today which mostly exist to satisfy my odd preferences.

The utility reads the format of current stream and toggles to the other one. To force `regular` or `irregular` mode, use the `-r` or `-i` flags.

# Installation

    $ go get https://github.com/thanthese/toggle-calendar

# License

MIT
