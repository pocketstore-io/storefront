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
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';

const open = useLocalStorage('open', false, {});
const cookie = useLocalStorage('cookie', false, {});
</script>