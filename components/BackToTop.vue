<template>
  <dialog id="my_modal_2" class="modal" :open="open">
    <div class="modal-box">
      <h3 class="text-lg font-bold">{{ $t('cookie.headline') }}</h3>
      <p class="py-4">
        {{ $t('cookie.note') }}
      </p>
      <section class="actions">
        <div class="join grid grid-cols-6">
          <button class="btn btn-error join-item col-span-6 md:col-span-3" @click="open = false; cookie = false">{{
            $t('cookie.disable') }}</button>
          <button class="btn btn-primary join-item col-span-6 md:col-span-3" @click="open = false; cookie = true">{{
            $t('cookie.accept') }}</button>
        </div>
      </section>
    </div>
    <form method="dialog" class="modal-backdrop bg-black opacity-75" @click="open = !open;">
      <button>close</button>
    </form>
  </dialog>
  <transition>
    <section v-if="show"
      class="back-to-top-icon fixed bottom-0 right-0 mr-24 mb-3 bg-black text-white rounded-full h-16 w-16 flex items-center justify-center z-10"
      @click="scrollToTop()">
      <Fa :icon="faArrowUp" size="2x" />
    </section>
  </transition>
</template>

<script lang="ts" setup>
import { faArrowUp } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const open = ref(false);
const show = ref(false);
const cookie = useLocalStorage('cookie', null, {});

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}

onMounted(() => {
  if (cookie.value === null) {
    open.value = true;
  }

  window.addEventListener("scroll", () => {
    if (document.documentElement.scrollTop > 250) {
      show.value = true;
    } else {
      show.value = false;
    }
  });
});
</script>

<style>
.back-to-top-icon:hover {
  cursor: pointer;
}

/* we will explain what these classes do next! */
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>