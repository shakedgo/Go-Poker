<template>
    <div class="game_menu">
        <div class="create_player" @click="newPlayer">New Player</div>
        {{player}}
        <div>Join Table</div>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import { useFetch } from '@/composables/api'

    let player = ref(null)

    async function newPlayer(){
        const res = await useFetch('new-player',{name: 'sd'}, 'POST');/*  */
        if (res) player.value = res.player
    }
</script>

<style lang="scss" scoped>

</style>
