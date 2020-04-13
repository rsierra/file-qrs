<template>
  <div class="game-card">
    <span class="game-card-close">x</span>
    <p class="game-name">{{game_name}}</p>
    <img src="/" alt class="game-cover-url" />
    <img src="/" alt class="game-qr-code-url" />
    <p class="game-desc"></p>
  </div>
</template>

<script>
module.exports = {
  data: function() {
    return {};
  },
  props: ["game_name", "game_cover_url", "game_qr_code_url", "game_desc"],
  mounted() {
    console.log("Game card rendered");
  }
};
</script>