<template>
  <dialog id="my_modal_2" class="modal" :open="open">
    <div class="modal-box">
      <h3 class="text-lg font-bold">{{$t('cookie.headline')}}</h3>
      <p class="py-4">
        {{$t('cookie.note')}}
      </p>
      <section class="actions">
        <div class="join grid grid-cols-6">
          <button class="btn btn-error join-item col-span-6 md:col-span-3" @click="open = false; cookie = false">{{$t('cookie.disable')}}</button>
          <button class="btn btn-primary join-item col-span-6 md:col-span-3" @click="open = false; cookie = true">{{$t('cookie.accept')}}</button>
        </div>
      </section>
    </div>
    <form method="dialog" class="modal-backdrop bg-black opacity-75" @click="open = !open;">
      <button>close</button>
    </form>
  </dialog>
  <section
    class="cookie-icon fixed bottom-0 left-0 ml-3 mb-3 bg-[#1f6fed] text-white rounded-full h-16 w-16 flex items-center justify-center z-10"
    @click="open = !open">
    <Fa :icon="faCookieBite" size="2x" />
  </section>
</template>

<script lang="ts" setup>
import { faCookieBite } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const open = ref(false);
const cookie = useLocalStorage('cookie', null, {});

onMounted(() => {
  if (cookie.value === null) {
    open.value = true;
  }
});
</script>
