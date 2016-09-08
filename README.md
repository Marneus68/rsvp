# rsvp

`rsvp` is a really simple virtual network printer. It will attempt to print everything that is sent to its open port (9100 by default) as a pdf and can optionally send you the resulting print via email.

## Usage

The `rsvp` executable expects parmeters to function, you can see all the possible arguments by inputing `rsvp -h`.

    Usage of ./rsvp:
      -dm string
            destination email address (default "user@domain.tld")
      -m	send prints via email
      -o string
            output directory (default "~/rsvp")
      -p string
            port rsvp listens on (default ":9100")
      -sa string
            smtp server and port (default "smtp.domain.tld:587")
      -sm string
            sender email address (default "rsvp-noreply@hostname")
      -sp string
            smtp user password (default "password")
      -su string
            smtp user name (default "user@domain.tld")
      -t int
            timeout in seconds (default 5)

Unless specified otherwise during the invokation, the content to be printed as well as the pdf prints will be stored in two subdirectories of `~/rsvp` at the date at which they were received:

    rsvp/
    ├── pdf
    │   └── Thu Sep  8 09:46:43 CEST 2016.pdf
    └── pdl
        └── Thu Sep  8 09:46:43 CEST 2016.pdl

### Examples

The following set of parameters will start `rsvp` on the 9800 port and instruct it to store its output in `~/output`:

    rsvp -p 9800 -o ~/output

The following set of parameters will start `rsvp` on the 9100 port and instruct it to forward every print via email:

    rsvp -p 9100 -sa mail.cock.li:587 -sm noreply@rsvp.net -su bob@cock.li -sp password

## Building

To build `rsvp`, your Go development environment must be set up correctly (your `$GOPATH` must be set correctly and `$GOPATH/bin` must be into your path).

You can fetch the sources with:

    go get github.com/Marneus68/rsvp

This will take care of the dependencies and install `rsvp` in your `$GPATH/bin` directory.

## Dependencies

`rsvp` uses [draw2d](https://github.com/llgcode/draw2d) and [ps](https://github.com/llgcode/ps).

## TODOs

A proper PDL language should be implemented at some point. I'll have to get my hands on "standard" network printers for that.

## License

This project is under the [WTFPL](http://www.wtfpl.net/), see [LICENSE](https://raw.githubusercontent.com/Marneus68/rsvp/master/LICENSE) for more details.

