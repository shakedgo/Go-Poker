<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const $router = useRouter()
const authStore = useAuthStore()

// console.log($router)
// defineProps<{
//   msg: string
// }>()
</script>

<template>
    <nav>
        <template v-for="route in $router.options.routes" :key="route.path">
            <router-link v-if="!route?.meta?.requiresAuth || authStore.isLoggedIn" :to="route.path">{{ route.name }}</router-link>
        </template>
        <!-- <router-link to="/">Home</router-link>
        <router-link to="/game">Game</router-link>
        <router-link to="/about">About</router-link> -->
    </nav>
</template>

<style lang="scss">
    nav {
        display: flex;gap: var(--gap);
    }
</style>