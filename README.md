# ARCHIVED

This repository is no longer maintained. See the [Eos](https://github.com/Zibbp/eos) repository for an updated version.

<br />
<div align="center">
  <a>
    <img src=".github/eos-logo.webp" alt="Logo" width="80" height="80">
  </a>

  <h2 align="center">Eos</h2>

  <p align="center">
    Easily import import and view your archived YouTube collection.
  </p>
</div>

---

## Images

![eos-full](https://user-images.githubusercontent.com/21207065/213781832-848afbb9-baa6-40bb-b330-904ecbf61dd6.png)

![eos-channel](https://user-images.githubusercontent.com/21207065/213781835-2fee53e6-6fb2-4660-ae3b-78fdd878838a.png)

## Features

- Video Playback
  - Includes support for captions
- Description
- Comments
- Search
- Dark Theme
- 100% Offline. Eos exclusively uses your archive collection.

## About

Videos are scanned and imported into Eos. Eos does not touch or alter your archived files, it only reads them.

## Requirements

Your YouTube collection needs to follow the below file tree. [YT-DLP](https://github.com/yt-dlp/yt-dlp/) generate files like this. I'm using [TheFrenchGhostys-Ultimate-YouTube-DL-Scripts-Collection](https://github.com/TheFrenchGhosty/TheFrenchGhostys-Ultimate-YouTube-DL-Scripts-Collection)'s channel script to archive channels.

```
videos/
├── Zibbp/
│   ├── video_1/
│   │   ├── video_1.info.json
│   │   ├── video_1.mkv
│   │   ├── video_1.webp
│   │   └── video_1.vtt
│   ├── video_2/
│   │   ├── video_2.info.json
│   │   ├── video_2.mkv
│   │   ├── video_2.webp
│   │   └── video_2.vtt
│   └── video_2/
├── Asmongold_TV/
│   └── wow_video/
│       ├── wow_video.info.json
│       ├── wow_video.mkv
│       ├── wow_video.webp
│       └── wow_video.vtt
├── AnotherChannel/
│   └── video
└── OtherChannel/
    └── video
```

Three files **must** be present in a video's folder.

- Video (mkv or mp4).
- Thumbnail (webp, jpg, png).
- Info.json (JSON file with the video's information include comments).

## Deployment

Use the provided docker-compose.yml file.
