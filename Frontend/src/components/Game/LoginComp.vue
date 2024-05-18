<template>
  <div class="game_menu">
    <div class="title">{{ authStore.loginForm ? 'Login' : 'Sign Up' }}</div>
    <form @submit.prevent="authStore.loginForm ? authStore.login(form) : authStore.signup(form)">
      <input type="text" name="username" v-model="form.username" placeholder="Enter your name" />
      <input type="password" name="password" v-model="form.password" placeholder="Enter your password" />
      <div class="action">
        <button type="submit">{{ authStore.loginForm ? 'Login' : 'Sign Up' }}</button>
        <a class="signup" @click="authStore.toggleLoginForm">
          {{ authStore.loginForm ? "Don't have an account?" : "Have an account?" }}
        </a>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const form = reactive({
  username: '',
  password: '',
})
</script>