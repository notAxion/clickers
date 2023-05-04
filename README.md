# Clickers

## What is it ?
Clickers is a project to create a way to send internet packets through sound

#### Play `out.bin` with `ffplay`
We need `ffplay` to play the sound

Install `ffmpeg` with your favorite package manager 
or you can go to the [official Website](https://ffmpeg.org/download.html) to download the binary
```sh
ffplay -f f32le -ar 44100 -showmode 1 out.bin
```
This command would play the sound
## Roadmap
- create a way to encode and decode internet packets as a sound file
- create a server to test the functionality
