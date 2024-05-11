<template>
    <h1>This is an protected view</h1>
    <h2>{{ ans }}</h2>
</template>


<script setup lang="ts">
    import { ref } from 'vue'
    import { useFetch } from '@/composables/api'

    let ans = ref()
    async function getTable() {
      try {
        const res = await useFetch('print-table/1', {}, 'GET');
        console.log(res);
        ans.value = res;
      } catch (err) {
        console.error(err)
      }
    }
    getTable();
</script>

<style lang="scss" scoped>

</style>
