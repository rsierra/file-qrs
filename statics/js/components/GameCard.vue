<template>
  <div class="game-card">
    <span @click="hideMe" class="game-card-close" title="Close">x</span>
    <p class="game-name">{{game_name}}</p>
    <img :src="game_cover_url" :alt="game_name" :title="game_name" class="game-cover-url" />
    <figure class="game-qr-code" :title="game_qr_code_url">
      <qrcode-vue :value="game_qr_code_url" :size="size" level="H"></qrcode-vue>
      <figcaption>{{game_qr_code_url}}</figcaption>
    </figure>
    <p class="game-desc">{{game_desc}}</p>
    <a
      :href="'//google.com/' + game_id"
      class="game-original-url"
      target="_blank"
    >Open game in {{game_id}} IGBD</a>
  </div>
</template>

<script>
module.exports = {
  data: function() {
    return {
      // value: "lorem",
      size: 300
    };
  },
  props: [
    "game_id",
    "game_name",
    "game_cover_url",
    "game_qr_code_url",
    "game_desc"
  ],
  mounted() {
    console.log("Game card rendered");
  },
  methods: {
    hideMe(event) {
      console.log("Hide game card");
      document.body.className = document.body.className.replace(
        "show-game-card",
        ""
      );
    }
  },

  components: {
    // Components
    QrcodeVue
  }
};
</script>
