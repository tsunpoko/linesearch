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
        :search-input.sync="search"
        label="グループ名を検索"
        :placeholder=search
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
            <v-img :src="group.src"></v-img>
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
export default {
  data: () => ({
    groups: [],
    list_groups: [
      {
        color: '#1F7087',
        src: 'https://news.mynavi.jp/article/20181102-LINE/images/005.jpg',
        name: 'aaa',
        description: 'asdfasdfasdfasdf\n',
        num: 100,
        url: "https://google.com",
      },      {
        color: '#1F7087',
        src: 'https://news.mynavi.jp/article/20181102-LINE/images/005.jpg',
        name: 'aaaaaaaa',
        description: 'asdfasdfasdfasdf',
        num: 'Foster the People',
        url: "https://google.com",
      },      {
        color: '#1F7087',
        src: 'https://news.mynavi.jp/article/20181102-LINE/images/005.jpg',
        name: 'aaaaaaaaaa',
        description: 'asdfasdfasdfasdf',
        num: 'Foster the People',
        url: "https://google.com",
      },     {
        color: '#1F7087',
        src: 'https://news.mynavi.jp/article/20181102-LINE/images/005.jpg',
        name: 'aaaaaaaaaa',
        description: 'asdfasdfasdfasdf',
        num: 'Foster the People',
        url: "https://google.com",
      },{
        color: '#1F7087',
        src: 'https://news.mynavi.jp/article/20181102-LINE/images/005.jpg',
        name: 'ああ',
        description: 'asdfasdfasdfasdf\n',
        num: 'Foster the People',
        url: "https://google.com",
      },
    ],

    search: "",
    search_tmp: "",
    isLoading: false,
  }),
  watch: {
    search: function(val) {
      if ( typeof(val) == null ) {
        return
      }

      if ( val.length == 0 ) {
        this.isLoading = false
        return
      }

      if ( this.isLoading ) {
        return
      }
      this.groups = []
      for (var i of this.list_groups) {
        //console.log(i.name)
        if ( i.name.indexOf(this.search) != -1 ) {
          this.groups.push(i)
        }
      }
    }
  },

};
</script>

<style>

div.register{
  margin-bottom: 10px;
}
</style>
