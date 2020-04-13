let vm = new window.Vue({
  el: '#vue-app',

  components: {
    // Components
    'game-card': window.httpVueLoader('/statics/js/components/GameCard.vue')
  },

  data: {
    game: {
      name: '',
      cover_url: '',
      qr_code_url: '',
      desc: '',
    },
    another_var: 'value',
  },

  methods: {
    onUpdateConfigValue: function (key_value) {

    },
  },

  created: function () {
    console.log('App ready!');
  }
});