<template>
  <button class="btn btn-primary" @click="open = !open">
    <span v-if="!open">
      <Fa :icon="faBars" size="2x" class="text-white" />
    </span>
    <span v-else>
      <Fa :icon="faTimes" size="2x" class="text-white" />
    </span>
  </button>
</template>

<script lang="ts" setup>
import { faBars, faTimes } from '@fortawesome/free-solid-svg-icons';
const open = ref(false);

watch(open, (value) => {
  if (value) {
    modal_mobile_menu.showModal()
    document.activeElement.blur();
    modal_mobile_menu.addEventListener('close', ()=> {
      open.value = false;
    })
  }
});
</script>
