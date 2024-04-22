
# Acey Ducey
[![License: ODbL](https://img.shields.io/badge/License-PDDL-brightgreen.svg)](https://opendatacommons.org/licenses/pddl/)

## Description
This is a Go port of "Acey Ducey", based on the 1978 edition of
BASIC Computer Games (BGC) from Creative Computing.

https://www.atariarchives.org/basicgames/showpage.php?page=2

NOTE: The game is actually "Acey Deucey", and sometimes called "In-Between".
"Acey Ducey" is simply the spelling used in BCG.

## Purpose
Just getting used to Go. When BASIC was a thing, this was the first code I ever typed in and debugged.
So why not in Go? 

## Design Process

Flowchart the logic using [mermaid-js](https://github.com/mermaid-js/mermaid).

1. Review actual card game rules (found discrepencies vs the ruleset hosted by Bicycle Cards)
1. Review code in the book
1. Review other implementations: rando Github, the `coding-horror` repo, and forums. 
1. Flowchart original logic

To study any of these BCG listings, it helps to annotate printed code, or flowchart it.
The BASIC language variant used in the book did not support GOSUB, or even "THEN" and "ELSE".
Their workaround was spaghetti GOTOs, sometimes statement cramming, which makes it harder to read.

## Game Logic (Approximate)

This version follows the original book listing's logic and rules. Note that the book deviates
from more traditional card rules, as found on the Bicycle Cards website or Wikipedia.

```mermaid
flowchart TD
    
    A[Introduction] --> B[Your money remaining]
    B --> E[Draw dealer cards]
    // Evaluate dealer cards
    E --> F[Display cards]
    F --> G[Get bet]
    G --> H[Draw player card]
    H --> I [Evaluate player card]
    H --> Y[Try Again?]
    Y --> |Yes| B
    Y --> |No| Z
    Z[END]
```

You get the picture.

## Run

`go run main.go`
NOTE: Requires Go v1.20+ (has math/rand with seed auto-set).

## Credits

Bill Palmby of Prairie View, Illinois wrote the original version in BCG.

## LICENSE
Public domain, as was default for the original Palmby version.
Additionally, David H . Ahl has explicitly released his work to the public.

## Contributing

This is probably done, but if you find a bug please file a PR or Issue.

If you want to enhance, it is public domain, just fork it.

## Possible Improvements

* Two-player (note 3+ players requires actual deck modeling/card-counting)
* Ruleset selection (Basic Computer Games, or Bicycle Cards)
* Track wins-losses when the user opts to play again
* Display probabilities/risks (and logic how this is calculated)
* Emojis for cards
* Support for [Bubble Tea](https://github.com/charmbracelet/bubbletea)
* i18n
* ..maybe use i18n to also toggle between 40 and 80-col, legacy UPPER case or modern Mixed like how some retro micros worked.
* Display David Beker's card-playing robot in the Terminal [go-sixel](https://github.com/mattn/go-sixel) (with attribution and link of course)
* A GUI [Ebitengine](https://ebitengine.org/)
* Combine this with other card games into a single app, using subcommands to select games [Cobra](https://github.com/spf13/cobra)
* Write as an HTTP API game server (let someone else write a client for it)
* Use [goRemi, a prototype for card-playing game](https://github.com/ibrdrahim/goRemi)

## References

* [Acey Ducey, from BASIC Computer Games](https://www.atariarchives.org/basicgames/showpage.php?page=2)
* [Basic Computer Games 1978 version (PDF)](https://annarchive.com/files/Basic_Computer_Games_Microcomputer_Edition.pdf)
* [Basic Computer Games 1975 version (PDF)](http://www.bitsavers.org/pdf/dec/_Books/101_BASIC_Computer_Games_Mar75.pdf)
* [In-Between: How to Play (Bicycle Cards)](https://bicyclecards.com/how-to-play/in-between/)
* [In-Between: Game Rules](https://playingcarddecks.com/blogs/how-to-play/in-between-game-rules)
* [Updating The Single Most Influential Book of the BASIC Era](https://blog.codinghorror.com/updating-the-single-most-influential-book-of-the-basic-era/)
* [coding-horror/basic-computer-games](https://github.com/coding-horror/basic-computer-games)
* [Red dog (card game)](https://en.wikipedia.org/wiki/Red_dog_(card_game)) (Slightly different game, but shares wiki page with Red Dog)
* [Troy Campbell wrote the Go version of Acey Ducey in coding-horror](https://github.com/coding-horror/basic-computer-games/commits/main/00_Alternate_Languages/01_Acey_Ducey)
* [David H. Ahl, PUBLIC DOMAIN](https://twitter.com/search?q=David%20ahl&src=typed_query&f=live)
* [10 PRINT"LET'S READ BASIC COMPUTER GAMES":GOTO 10](https://forums.somethingawful.com/showthread.php?threadid=3928712)
* [David H Ahl](https://en.wikipedia.org/wiki/David_H._Ahl)
* [David Beker's Bots](http://www.bekerbots.com/)

## Special Mentions

Thanks to the author of BGC, [David H Ahl](https://en.wikipedia.org/wiki/David_H._Ahl) for collecting and documenting
this amazing tome of simple (some not so simple) terminal games.

Thanks also to the book's illustrator, [David Beker](http://www.bekerbots.com/), whose robots attracted the eye
of many an eleven year old. BCG was not intended to be a kid's book, but thanks to the robots, it was.