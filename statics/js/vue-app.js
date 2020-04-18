// App
let vm = new window.Vue({
  el: '#vue-app',

  components: {
    // Components
    'folder-item': window.httpVueLoader('/_app/statics/js/components/FolderItem.vue'),
    'file-item': window.httpVueLoader('/_app/statics/js/components/FileItem.vue'),
    'game-card': window.httpVueLoader('/_app/statics/js/components/GameCard.vue'),
  },

  data: {
    folders: [
      { id: 1, name: "PSX" },
      { id: 2, name: "NES" }
    ],
    files: [
      { id: 1, name: "Metal Gear Solid [igdb-375].cia", url: "/_app/files/Metal Gear Solid [igdb-375].cia" },
      { id: 4, name: "Metal Gear Solid 2 [igdb-375].cia", url: "/_app/files/Metal Gear Solid 2 [igdb-375].cia" }
    ],
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
    onUpdateGame: function (event) {
      console.log(event);
    },
    onGameCard: function (obj) {
      console.log(obj);
    }
  },
  created: function () {
    console.log('App ready!');
  }
});

// Add
document.addEventListener("click", function (e) {
  if (e.target.id === "sample-game-link-url") {
    e.preventDefault();
    vm.$root.onGameCard(e.target.href);
  }
});
