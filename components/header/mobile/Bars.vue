<template>
  <button class="btn btn-primary" @click="open = !open">
    <span v-if="!open">
      <Fa :icon="faBars" class="text-white w-5" />
    </span>
    <span v-else>
      <Fa :icon="faTimes" class="text-white w-5" />
    </span>
  </button>
</template>

<script lang="ts" setup>
import { faBars, faTimes } from "@fortawesome/free-solid-svg-icons";
const open = ref(false);

watch(open, (value) => {
    if (value) {
        modal_mobile_menu.showModal();
        document.activeElement.blur();
        modal_mobile_menu.addEventListener("close", () => {
            open.value = false;
        });
    }
});
</script>
