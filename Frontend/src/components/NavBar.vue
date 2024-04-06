<script setup lang="ts">
// defineProps<{
//   msg: string
// }>()
</script>

<template>
    <nav>
        <router-link to="/">Home</router-link>
        <router-link to="/game">Game</router-link>
        <router-link to="/about">About</router-link>
    </nav>
</template>

<style lang="scss">
    nav {
        display: flex;gap: var(--gap);
    }
</style>