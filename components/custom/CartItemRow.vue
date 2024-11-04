<template>
  <td class="col-span-6 md:col-span-4">{{ item.name }}</td>
  <td v-if="item.price" class="col-span-6 md:col-span-3">
    {{ (item.price).toFixed(2) }} € * {{ qty }} = {{ (item.price * qty).toFixed(2) }} € </td>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const { identifier } = defineProps({
  identifier: { type: String, requiered: true },
  qty: { type: Number, requiered: true }
});

const item = ref({});

const load = async function () {
  item.value = await pb.collection('products').getOne(identifier);
}

onMounted(() => {
  load();
});
</script>

<style></style>