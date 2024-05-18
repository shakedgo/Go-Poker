<template>
    <Login v-if="!authStore.isLoggedIn"/>
</template>

<script setup lang="ts">
import Login from '@/components/Game/LoginComp.vue'
import { useAuthStore } from '@/stores/auth'
const authStore = useAuthStore()

</script>

<style lang="scss" scoped>
    
</style>