# <p align="center">lrcsnc But Removed Romanization</p>
Gets the currently playing song's synced lyrics and displays them in sync with song's actual position!

This version of lrcsnc got rid of the music playing animation and replaced with shortened word like SNF (Sound not found), you will like it if you don't want tonns of animation going around in your waybar!

lrcsnc is primarily designed for bars like [Waybar](https://github.com/Alexays/Waybar).

https://github.com/user-attachments/assets/1bc93e59-385f-41cb-a23e-49298e5887b0

## Features

- Syncing to any player that supports MPRIS
- A decent level of customization and configuration using TOML
- Full integration to Waybar

## Build
```
git clone https://github.com/xNeonCatGirlx/lrcsnc-without-romanization.git
cd lrcsnc-without-romanization
make # or `sudo make all` for automatic install
```
Make sure to have go v1.23 or above.

## Usage
```
lrcsnc [OPTION]
```
Get more info on on available options with `lrcsnc -h`. waybar folder

## Setting up for waybar
This is a kinda ok-ish solution, maybe not the best

You will need to add a custom module in waybar, for example:
```sh
"custom/lyrics"
```

(Optional) Create a scripts folder in your waybar config folder

Create a file, can be names anything, for example
```sh
/home/neoncatgirl/.config/waybar/scripts/lrcsnc_waybar.sh
```
Use lrcsnc -o /replace/with/path/to/lrcsnc_waybar.sh

So lrcsnc outputs lyrics to the .sh file

Make it executable with 
```sh
chmod +x /replace/with/path/to/lrcsnc_waybar.sh
```
Finish the rest of the waybar module with 
```sh
    "custom/lyrics": {
    "exec": "/home/neoncatgirl/.config/waybar/scripts/lrcsnc_to_file.sh",
    "interval": 0,
    "return-type": "stdout"
```
