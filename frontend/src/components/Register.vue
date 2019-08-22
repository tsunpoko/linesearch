<template>
  <v-container>
    <v-layout
      text-center
      justify-center
      align-center
      wrap
      width="300"
    >


    <v-flex xs12>
      <v-img
      src="https://snsdays.com/wp-content/uploads/2019/08/a10db1b69d8bef8a5d04288a206a7c76-768x672.png"
      class="my-3"
      contain
      height="200"
      ></v-img>
    </v-flex>

      <v-flex
        md6
        xs12
      >


        <div class="divider">
          <v-divider></v-divider>
        </div>



        <v-form
          ref="form"
          v-model="valid"
          lazy-validation
        >


          <v-text-field
            v-model="charge_name"
            :counter="100"
            :rules="nameRules"
            label="グループのURL"
            required
          ></v-text-field>


          <div class="confirmbtn">
            <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
            >
              登録
            </v-btn>
          </div>
        </v-form>




      </v-flex>
    </v-layout>
    </v-container>
  </template>


<script>
  export default {
    data: () => ({
      valid: true,
      nameRules: [
        v => !!v || '必須項目です',
        v => (v && v.length <= 100) || '100文字以下で入力してください',
      ],
    }),
  }
</script>

<style>

div.divider {
  padding-bottom: 10px;
}
</style>
