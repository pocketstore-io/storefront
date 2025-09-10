<template>
  <dialog id="modal_mobile_menu" class="modal px-6" :open="navigation">
    <div class="modal-box">
      <HeaderMobileLogo />
      <HeaderMobileSearchBar />
      <HeaderMobileMenu />
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn">{{ $t('general.close') }}</button>
        </form>
      </div>
    </div>
  </dialog>
</template>

<script setup lang="ts">
import {useLocalStorage} from "@vueuse/core";
const navigation = useLocalStorage('navigation', false)
onMounted(()=>{
  open.value = false;
});
</script>
