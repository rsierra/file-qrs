<template>
  <div class="file-item">
    <button @click="fetchFile">{{ file.name }}</button>
  </div>
</template>

<script>
module.exports = {
  props: ["file"],
  methods: {
    fetchFile: function(event) {
      console.log("Fetch data item");
      this.$emit("update-game", "game");
      // axios
      //   .post(
      //     "https://api-v3.igdb.com/games",
      //     "fields cover,name,summary; where id = 375;",
      //     { headers: { "user-key": "830b1419090b76a1e37187b6c7d85891" } }
      //   )
      //   .then(response => {
      //     console.log(response);
      //   })
      //   .catch(error => {
      //     console.log(error);
      //   });
    }
  }
};
</script>
