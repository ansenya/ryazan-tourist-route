<script setup>
import VideoPlayer from "./VideoPlayer.vue";
import {computed, ref} from "vue";

const isExpanded = ref(false);

const toggleText = () => {
  isExpanded.value = !isExpanded.value;
};

defineProps({
  videoUrl: {
    type: String,
    required: true,
  },
  headerUrl: {
    type: String,
    required: true,
  },
  imageUrl: {
    type: String,
    required: false,
  },
  header: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  baseUrl: {
    type: String,
    required: true,
  }
});

</script>

<template>
  <div class="card-container">
    <div class="header-container">
      <img class="header-pattern" :src="`${baseUrl}${headerUrl}`" alt=""/>
      <div class="text" @click="toggleText">
        <h2 class="text-header" v-html="header"></h2>
      </div>
    </div>
    <div class="description-container">
      <div class="text-content"
           :class="{ expanded: isExpanded, gradient: isExpanded }"
           @click="!isExpanded && toggleText"
           v-html="text">
      </div>
    </div>
    <div class="read-more" @click="toggleText()">
      {{ isExpanded ? "read less" : "read more" }}
    </div>
    <div v-if="videoUrl" class="video-container">
      <VideoPlayer :videoUrl="`${baseUrl}${videoUrl}`"/>
    </div>
    <div v-if="imageUrl" class="image-container">
      <img class="image" :src="`${baseUrl}${imageUrl}`" alt=""/>
    </div>

    <div class="line"></div>
  </div>
</template>


<style scoped>
.card-container {
  display: flex;
  flex-direction: column;
}

.header-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0;
}

.text {
  font-size: 26px;
  padding: 0 0 0 10px;
  user-select: none;
  align-self: center;
  font-family: 'HeaderFont', sans-serif;
  margin: 0;
}

.text-header {
  margin: 10px;
  text-align: center;
  -webkit-text-stroke: 1px rgba(0, 0, 0, 0.5);
}

.text:hover {
  text-decoration: none;
  color: #6b6b6b;
}

.description-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-inline: 16px;
  height: 100%;
  overflow: hidden;
  position: relative;
  padding-bottom: 0;
}

.text-content {
  max-height: 30px;
  overflow: hidden;
  position: relative;
  transition: max-height 0.3s ease;
}

.text-content:after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 30px;
  pointer-events: none;
  background: linear-gradient(to bottom, rgba(0, 0, 0, 0), white);;
}

.text-content.gradient:after {
  background: none;
}

.text-content.expanded {
  max-height: none;
}

.read-more {
  color: rgba(128, 128, 128, 0.58);
  align-self: end;
  padding-inline: 16px;
  padding-bottom: 5px;
  text-decoration: underline;
  font-size: 14px;
}

.read-more:hover {
  cursor: pointer;
  color: #6b6b6b;
}

.description-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-inline: 16px;
  padding-bottom: 10px;
  height: 100%;
  overflow: hidden;
}


.video-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  height: 100%;
  padding: 4px;
}


.image-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  height: 30%;
  padding: 4px;
}

.header-pattern {
  width: 100%;
  object-fit: contain;
  max-width: 450px;
  margin: 0 auto;
}

.image {
  width: 100%;
  height: 100%;
  object-fit: fill;
  max-width: 500px;
  margin: 0 auto;
  border-radius: 6px;
}

.line {
  height: 1px;
  width: 100%;
  background-color: #494949;
  margin-bottom: 15px;
  margin-top: 10px;
}
</style>
