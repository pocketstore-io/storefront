<template>
  <dialog id="my_modal_3" class="modal" :open="open">
    <div class="modal-box max-h-[80vh]">
      <h3 class="text-xl font-bold block text-center mb-6">
        <span>Live Chat by JMSE</span>
      </h3>
      <section class="chat w-full grid grid-cols-6">
        <ul class="w-full col-span-6">
          <li v-for="item in items">
            <LivechatItem :item="item" />
          </li>
        </ul>
      </section>
    </div>
    <form method="dialog" class="modal-backdrop bg-black opacity-75" @click="open = !open;">
      <button>close</button>
    </form>
  </dialog>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import LivechatItem from '../livechat/Item.vue';

const open = useLocalStorage('open-livechat', false, {});
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const items: Ref = ref([]);

const load = async function () {
  items.value = (await pb.collection('support_chat').getList(1, 50)).items;
}
onMounted(() => {
  load();
  if (!items.value.length) {
    items.value = items.value.push({ message: 'Hallo Welt' });
  }
  setInterval(() => {
    load();
  }, 10000);
});

watch(items, () => {
  watch(
    () => items.value.length,
    (newLength, oldLength) => {
      if (newLength !== oldLength && oldLength == 0) {
        beep();
      }
    }
  );
});

const beep = function () {
  const snd = new Audio('/sounds/bling.mp3');
  snd.play();
}
</script>