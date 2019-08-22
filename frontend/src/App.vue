<template>
  <v-app>

    <v-app-bar app fixed clipped-left text-center color="green accent-3" dark>
    <v-toolbar-side-icon v-if="sidebar" @click.stop="drawer = !drawer"></v-toolbar-side-icon>
  <v-spacer></v-spacer>
  <v-toolbar-title class="font-weight-medium">
    <a href="/" class="white--text">OPENCHAT グループ検索</a>
  </v-toolbar-title>
  <v-layout align-center justify-end>
  </v-layout>
</v-app-bar>


    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  components: {
    //HelloWorld,
  },
  data: () => ({
    //
  }),
};
</script>

<style>
a {
  text-decoration: none;
  color: white;
}

a:visited {
  color: white;
}

a:hover {
  text-decoration: none;
}
</style>
