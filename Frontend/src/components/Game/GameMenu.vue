<template>
    <div class="game_menu">
        <div class="create_player" @click="newPlayer">New Player</div>
        {{player}}
        <div>Join Table</div>
        <button @click="takePicture">take pic</button>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue'
    import { useFetch } from '@/composables/api'
    import { Camera, CameraResultType } from '@capacitor/camera';

    const imageSrc = ref<string>('');

const takePicture = async () => {
  const image = await Camera.getPhoto({
    quality: 90,
    allowEditing: true,
    resultType: CameraResultType.Uri
  });
  const imageUrl = image.webPath;
  if (imageUrl) {
    imageSrc.value = imageUrl;
  } else {
    // Handle the case where imageUrl is undefined
    console.error('Image URL is undefined');
  }};

    let player = ref(null)

    async function newPlayer(){
        const res = await useFetch('new-player',{name: 'sd'}, 'POST');
        if (res) player.value = res.player
    }
</script>

<style lang="scss" scoped>

</style>
