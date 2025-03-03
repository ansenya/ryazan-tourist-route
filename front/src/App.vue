<script setup>
import {onMounted, ref} from "vue";
import ScrollComponent from "@/components/ScrollComponent.vue";
import ReferencesComponent from "@/components/ReferencesComponent.vue";

const cards = ref([]);
const isLoading = ref(true);

// const baseUrl = "http://localhost:8000";
const baseUrl = "https://rtr.hipahopa.ru/api";

const fetchCards = async () => {
  try {
    const response = await fetch(baseUrl);
    if (!response.ok) throw new Error("Failed to fetch videos");

    cards.value = await response.json();

    setTimeout(() => {
      isLoading.value = false;
    }, 1500);
  } catch (error) {
    console.error("Error fetching cards:", error);
    isLoading.value = false;
  }
};

onMounted(fetchCards);
</script>

<template>
  <div v-if="isLoading" class="loading-screen">
    <img src="./assets/logo.png" alt="Logo" class="logo"/>
    <div class="spinner"></div>
  </div>

  <div>
    <ScrollComponent :cards="cards" :baseUrl="baseUrl"/>
    <ReferencesComponent/>
  </div>
</template>

<style scoped>
.loading-screen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: #ffffff;
  z-index: 9999;
}

img {
  padding: 30px;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 5px solid rgba(0, 0, 0, 0.3);
  border-top-color: #c6e577;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
