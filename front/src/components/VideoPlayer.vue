<script setup>
import {ref, onMounted} from "vue";
import Hls from "hls.js";

const props = defineProps({
  videoUrl: {
    type: String,
    required: true,
  },
});

const videoPlayer = ref(null);

const setupPlayer = () => {
  const video = videoPlayer.value;
  if (Hls.isSupported()) {
    const hls = new Hls({
          startPosition: 0,
          maxBufferLength: 3,
          maxMaxBufferLength: 3,
        }
    );
    hls.loadSource(props.videoUrl);
    hls.attachMedia(video);
  } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
    video.src = props.videoUrl;
  }
};

onMounted(setupPlayer);
</script>

<template>
  <video
      ref="videoPlayer"
      controls
      preload="none"
      playsinline
  ></video>
</template>

<style scoped>
video {
  height: 100%;
  width: calc(100vw - 20px);
  margin: 10px auto;
  max-width: 775px;
  border-radius: 6px;
  background: #b2b2b2;
}
</style>
