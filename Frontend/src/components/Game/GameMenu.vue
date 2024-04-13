<template>
    <div class="game_menu">
      <div class="title">Login</div>
      <form @submit.prevent="login">
        <input type="text" v-model="signinForm.username" placeholder="Enter your name">
        <input type="password" v-model="signinForm.password" placeholder="Enter your password">
        <button type="submit">Login</button>
      </form>
      <div class="title">Sign Up</div>
      <form @submit.prevent="signup">
        <input type="text" v-model="signupForm.username" placeholder="Enter your name">
        <input type="password" v-model="signupForm.password" placeholder="Enter your password">
        <button type="submit">Login</button>
      </form>
        <!-- <div class="create_player" @click="newPlayer">New Player</div>
        {{player}}
        <div>Join Table</div> -->
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import { useFetch } from '@/composables/api'

    // let player = ref(null)
    // async function newPlayer(){
    //     const res = await useFetch('new-player',{name: 'sd'}, 'POST');
    //     if (res) player.value = res.player
    // }
    let signinForm = ref({
        username: '',
        password: ''
    })
    let signupForm = ref({
        username: '',
        password: ''
    })
    async function login() {
      try {
        const res = await useFetch('login', signinForm.value, 'POST');
        console.log(res);
      } catch (err) {
        console.error(err)
      }
    }

    async function signup() {
      try {
        const res = await useFetch('signup', signupForm.value, 'POST');
        console.log(res);
      } catch (err) {
        console.error(err)
      }
    }
</script>

<style lang="scss" scoped>

</style>
