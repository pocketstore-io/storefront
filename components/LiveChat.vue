<template>
  <dialog id="my_modal_3" class="modal" :open="open">
    <div class="modal-box max-h-[80vh]">
      <h3 class="text-xl font-bold block text-center mb-6">
        <span>Live Chat by JMSE</span>
      </h3>
      <section class="chat w-full grid grid-cols-6">
        <ul class="w-full col-span-6">
          <li v-for="item in items">
            <div class="chat" :class="{ 'chat-start': item.from == 'from', 'chat-end': item.from != 'from' }">
              <div class="chat-bubble"
                :class="{ 'chat-bubble-primary': item.from == 'from', 'chat-bubble-neutral': item.from != 'from' }"
                @click="beep()">
                {{ item.message }}
              </div>
            </div>
          </li>
        </ul>
      </section>
    </div>
    <form method="dialog" class="modal-backdrop bg-black opacity-75" @click="open = !open;">
      <button>close</button>
    </form>
  </dialog>
  <section
    class="live-chat-icon fixed bottom-0 right-0 mr-3 mb-3 bg-red-600 text-white rounded-full h-16 w-16 flex items-center justify-center z-10"
    @click="open = !open">
    <Fa :icon="faComments" size="2x" />
  </section>
</template>

<script lang="ts" setup>
import { faComments } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const open = ref(false);
const cookie = useLocalStorage('cookie', null, {});
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const items: Ref = ref([]);

const load = async function () {
  items.value = (await pb.collection('support_chat').getList(1, 50)).items;
}

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


onMounted(() => {
  load();
  if (!items.value.length) {
    items.value = items.value.push({ message: 'Hallo Welt' });
  }
  setInterval(() => {
    load();
  }, 10000);
});

const beep = function () {
  const snd = new Audio('/sounds/bling.mp3');
  snd.play();
}
</script>


<style>
.live-chat-icon:hover {
  cursor: pointer;
}
</style>