# <p align="center">lrcsnc</p>
Gets the currently playing song's synced lyrics and displays them in sync with song's actual position!

lrcsnc is primarily designed for bars like [waybar](https://github.com/Alexays/Waybar).

<!-- TODO: insert video here once it's pushed to GitHub -->

## Features

- Syncing to any player that supports MPRIS
- Romanization for Japanese, Chinese and Korean languages
- A decent level of customization and configuration using TOML

## Build
```
git clone https://github.com/Endg4meZer0/lrcsnc.git
cd lrcsnc
make # or `sudo make all` for automatic install
```
Make sure to have go v1.24.2 or above.
More tricks are available at the [wiki](https://github.com/Endg4meZer0/lrcsnc/wiki/Building).

## Usage
```
lrcsnc [OPTION]
```
Get more info on on available options with `lrcsnc -h` or on [wiki](https://github.com/Endg4meZer0/lrcsnc/wiki/Available-options).

## TODO
- [ ] Check [compatibility](https://github.com/Endg4meZer0/lrcsnc/wiki/Compatibility-with-players) with different players
- [ ] More lyrics providers
- [ ] Terminal User Interface
- [ ] More configuration options?
- [ ] There is definitely always more!

## Need help or want to contribute?
You can always make an issue for either a bug or a feature suggestment! If your question is more general, consider opening a discussion.

## Your song was not found?
Consider adding the lyrics for it! Currently lrcsnc uses *LrcLib*, which is a great open-source lyrics provider service that has its own easy-to-use [app](https://github.com/tranxuanthang/lrcget) to download or upload lyrics. Once the lyrics are uploaded, lrcsnc should be able to pick them up on the next play of the song (that is if the cached version is not available though).
