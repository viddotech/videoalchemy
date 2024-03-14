# Streaming Video to YouTube

Use FFmpeg to stream live video content directly to YouTube, enabling real-time broadcasting of events, presentations, or personal streams.

## Command

```bash
ffmpeg -re -i input.mp4 -c:v libx264 -preset veryfast -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -c:a aac -b:a 160k -ar 44100 -f flv rtmp://a.rtmp.youtube.com/live2/your-stream-key
```


## Parameters

- **`-re`**: Reads the input file at its native frame rate, simulating live streaming.
- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-c:v libx264`**: Uses the H.264 codec for video encoding, widely supported for streaming.
- **`-preset veryfast`**: Sets a balance between encoding speed and quality. Other presets can be used depending on CPU capability.
- **`-maxrate 3000k`**: Sets the maximum video bit rate, which helps control the stream's quality and bandwidth usage.
- **`-bufsize 6000k`**: Sets the buffer size, which can affect video quality and stability of the stream.
- **`-pix_fmt yuv420p`**: Ensures the pixel format is compatible with YouTube and most other services.
- **`-g 50`**: Sets the group of pictures (GOP) size, affecting video quality and keyframe frequency.
- **`-c:a aac`**: Uses the AAC codec for audio encoding, recommended for YouTube.
- **`-b:a 160k`**: Sets the audio bit rate, balancing audio quality and bandwidth usage.
- **`-ar 44100`**: Sets the audio sampling rate, standard for high-quality audio.
- **`-f flv`**: Sets the output format to FLV, compatible with YouTube's RTMP servers.
- **`rtmp://a.rtmp.youtube.com/live2/your-stream-key`**: Specifies the YouTube RTMP server and your unique stream key. Replace `your-stream-key` with your actual YouTube stream key.

## Possible Errors

- **Connection refused or failed**: Occurs if there's an issue connecting to YouTube's servers. Ensure your stream key is correct and your internet connection is stable.
- **Invalid input or codec parameters**: Happens if the input file or specified codecs are not compatible with YouTube's streaming requirements. Double-check your command parameters.
- **Permission denied**: Can arise if there are restrictions preventing FFmpeg from accessing network resources. Check your firewall settings or network permissions.

## Additional Information

- **Stream Key Security**: Keep your stream key private, as anyone with the key can stream to your channel.
- **Internet Bandwidth**: Ensure your internet connection has sufficient upload bandwidth to support the chosen bit rate comfortably. Test your connection speed if unsure.
- **Audio-Video Sync**: Streaming, especially at higher qualities or over unstable connections, can lead to AV sync issues. Monitor your stream and adjust parameters as necessary.
- **Latency**: YouTube live streaming has inherent latency. Choose the appropriate latency setting in your YouTube live dashboard to balance interaction and stream quality.
