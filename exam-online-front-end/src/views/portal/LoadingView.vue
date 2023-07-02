<script setup>


import { useRouter } from "vue-router"
import {ref} from "vue";
let router = useRouter()

const loading = ref(true)


router.push('/home')


</script>

<template>
<div v-loading="loading" class="loading-container"></div>
</template>

<style scoped>
.loading-container {
  height: 100%;
}
</style>