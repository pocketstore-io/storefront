<template>
  <ModalLiveChat />
  <section
    class="hidden cursor-pointer fixed bottom-0 right-0 mr-3 mb-3 bg-red-600 text-white rounded-full h-16 w-16 flex items-center justify-center z-10"
    @click="open = !open">
    <Fa :icon="faComments" size="2x" />
  </section>
</template>

<script lang="ts" setup>
import { faComments } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const open = useLocalStorage('open-livechat', false, {});
</script>
