<script setup>
import CardContainer from "./CardContainer.vue";

defineProps({
  cards: {
    type: Array,
    required: true,
  },
  baseUrl: {
    type: String,
    required: true,
  }
});

</script>

<template>
  <div class="scroll-container">
    <CardContainer
        v-for="(card, index) in cards"
        :key="index"
        :videoUrl="card.videoUrl"
        :header="card.header"
        :text="card.text"
        :header-url="card.headerUrl"
        :image-url="card.imageUrl"
        :baseUrl="baseUrl"
    />
  </div>
</template>

<style scoped>
.scroll-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: #ffffff;
  margin: 0;
  padding: 0;
}

@media (min-width: 768px) {
  .scroll-container {
    max-width: 800px;
    width: 100%;
    margin: 0 auto;
    height: 100%;
    min-height: 100vh;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
}

</style>