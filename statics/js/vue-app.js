// App
let vm = new window.Vue({
  el: '#vue-app',

  components: {
    // Components
    'game-card': window.httpVueLoader('/_app/statics/js/components/GameCard.vue')
  },

  data: {
    game: {
      id: 9874,
      name: 'Nombre del juego',
      cover_url: '//via.placeholder.com/200x320.jpg',
      qr_code_url: '//via.placeholder.com/300x300.jpg',
      desc: 'Morbi leo risus, porta ac consectetur ac, vestibulum at eros.',
    },
    another_var: 'value',
  },
  methods: {
    onGameCard: function (data) {

    },
  },
  created: function () {
    console.log('App ready!');

  }
});

document.addEventListener("click", function (e) {
  if (e.target.id === "sample-game-link-url") {
    e.preventDefault();
    console.log(e.target.href);
    // alert("Clicked");
    vm.$root.onGameCard();
  }
});