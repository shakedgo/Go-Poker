<template>
    <div class="game_menu">
      <div class="title">{{loginForm ? 'Login' : 'Sign Up'}}</div>
      <form @submit.prevent="loginForm ? login() : signup()">
        <input type="text" name="username" v-model="form.username" placeholder="Enter your name">
        <input type="password" name="password" v-model="form.password" placeholder="Enter your password">
        <div class="action">
          <button type="submit">{{loginForm ? 'Login' : 'Sign Up'}}</button>
          <a class="signup" @click="loginForm = !loginForm">
            {{loginForm ? "Don't have an account?" :"Have an account?"}}
          </a>
        </div>
      </form>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import { useFetch } from '@/composables/api'

    let form = ref({
        username: '',
        password: ''
    })
    async function login() {
      try {
        const res = await useFetch('login', form.value, 'POST');
        console.log(res);
      } catch (err) {
        console.error(err)
      }
    }
    async function signup() {
      try {
        const res = await useFetch('signup', form.value, 'POST');
        console.log(res);
      } catch (err) {
        console.error(err)
      }
    }

    let loginForm = ref(true);
</script>

<style lang="scss" scoped>

</style>
