<template>
  <div class="folder-item">
    <button v-on:click="fetchFolder">{{ folder.name }}</button>
  </div>
</template>

<script>
module.exports = {
  props: ["folder"],
  methods: {
    fetchFolder: function(event) {
      console.log("Fetch data item");
    }
  }
};
</script>
