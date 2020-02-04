<template>
  <v-container>
    <v-layout
    wrap
    justify-center
    >
    <v-flex xs12>
      <v-img
      src="https://snsdays.com/wp-content/uploads/2019/08/a10db1b69d8bef8a5d04288a206a7c76-768x672.png"
      class="my-3"
      contain
      height="200"
      ></v-img>
    </v-flex>

    <v-flex xs12>
      <div class="text-center register">
        <v-btn rounded color="green accent-4" dark href="/register">グループを登録</v-btn>
      </div>
    </v-flex>

    <v-card
    color="green accent-1"

    width="700"
    row
    >

    <v-card-text class="headline font-weight-bold grey lighten-5">

      <v-card-actions class="">


        <v-autocomplete
        color="grey darken-2"
        hide-no-data
        hide-selected
        v-model="model"
        :items="items"
        item-text="Description"
        item-value="API"
        cache-items
        :loading="isLoading"
        :search-input.sync="search"
        label="グループ名を検索"
        prepend-icon="mdi-magnify"
        return-object
        ></v-autocomplete>


      </v-card-actions>

      <v-row>
        <v-col
        v-for="(group, i) in groups"
        :key="i"
        >
        <v-card
        color="green accent-4"
        dark
        >
        <v-list-item three-line>
          <v-list-item-content class="align-self-start">
            <v-list-item-title
            class="headline mb-2"
            v-text="group.name"
            ></v-list-item-title>

            <v-list-item-subtitle v-text="group.description"></v-list-item-subtitle>
            <v-list-item-subtitle v-text="'参加人数 : ' + group.num"></v-list-item-subtitle>
          </v-list-item-content>


          <v-list-item-avatar
          size="125"
          tile
          >
          <v-img :src="group.img"></v-img>
        </v-list-item-avatar>
      </v-list-item>

      <v-card-actions>
        <v-btn color="light-green accent-4" :href=group.url target="_blank">参加する</v-btn>
      </v-card-actions>
    </v-card>

  </v-col>
</v-row>

</v-card-text>
</v-card>



</v-layout>
</v-container>
</template>

<script>
import axios from 'axios';

//axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
//axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
//axios.defaults.baseURL = 'http://localhost:3000/';
axios.interceptors.response.use(null, (error) => {
  return Promise.reject(error);
});

export default {
  data: () => ({
    groups: [],
    list_groups: [],
    descriptionLimit: 60,
    entries: [],
    isLoading: false,
    model: null,
    search: null,
    items: [],
    //search: "",
    //search_tmp: "",
  }),
  watch: {
    search(val) {
      // すでに読み込み済みの場合は、何もしない
      if (this.items.length > 0) return
      // 読み込み中の場合も、何もしない
      if (this.isLoading) return

      this.isLoading = true
      this.groups = []
      // APIから、選択肢をfetchする
      axios.get('https://openchat-search.xyz/api/groups')
      //axios.get('http://localhost:3000/api/groups')
      //axios.get('http://jsonplaceholder.typicode.com/todos')
      .then(response => {
        this.list_groups = response.data
        for (var i of this.list_groups) {

          if ( i.name.indexOf(this.search) != -1 ) {
            this.groups.push(i)
          }
        }
      })
      .catch(err => {
        console.log(err)
        //読み込みが完了したので、loadingをfalseに
      }).finally(() => (this.isLoading = false))
    },
    computed: {
      fields () {
        if (!this.model) return []

        return Object.keys(this.model).map(key => {
          return {
            key,
            value: this.model[key] || 'n/a'
          }
        })
      },
      items () {
        return this.entries.map(entry => {
          const Description = entry.Description.length > this.descriptionLimit
          ? entry.Description.slice(0, this.descriptionLimit) + '...'
          : entry.Description
          return Object.assign({}, entry, { Description })
        })
      },
    },
  },
  methods: {
    getList() {
      console.log("welcome to openchat group search")
      axios.get('https://openchat-search.xyz/api/groups')
      //axios.get('http://localhost:3000/api/groups')
      //axios.get('http://jsonplaceholder.typicode.com/todos')
      .then(response => {
        this.groups = response.data
      })
      .catch(err => {
        console.log(err)
        //読み込みが完了したので、loadingをfalseに
      }).finally(() => (this.isLoading = false))
    }
  },
  created () {
    this.getList()
    alert("サービス終了しました。ポートフォリオのために一時的に公開しています。")
  }
};
</script>

<style>

div.register{
  margin-bottom: 10px;
}
</style>
