<script setup>

</script>

<template>
  <div class="references-container">
    <div class="header"><span style="color: #F4062A;">R</span><span style="color: #F8D00E">T</span><span style="color: #276C0A">R</span> Team Members</div>
    <ul>
      <li><strong>Director:</strong> Sobolev Alexander</li>
      <li><strong>Development Engineer:</strong> Arseniy Anoshin</li>
      <li><strong>Designers:</strong> Arseniy Arsentiev, Sobolev Alexander</li>
      <li><strong>Text Editors:</strong> Tatyana Safronova, Mikhail Kondrashov, Sobolev Alexander</li>
      <li><strong>Video Editors:</strong> Maria Sudarikova, Ivan Mirov</li>
      <li><strong>Video operator:</strong> Mikhail Kondrashov</li>
      <li><strong>Narrators:</strong> Ivan Mirov, Sobolev Alexander</li>
    </ul>
  </div>
</template>

<style scoped>

.references-container {
  display: flex;
  height: 100%;
  flex-direction: column;
  padding-inline: 15px;
  padding-bottom: 25px;
}

@media (min-width: 768px) {
  .references-container {
    max-width: 800px;
    width: 100%;
    margin: 0 auto;
    height: 100%;
    min-height: 100vh;
  }
}

.header {
  text-align: center;
  font-family: 'HeaderFont', sans-serif;
  font-size: 40px;
  color: #c6e577;
  -webkit-text-stroke: 1px rgba(0, 0, 0, 0.5);
}

ul {
  list-style-type: georgian;
}
li {
  margin: 4px;
}



</style>