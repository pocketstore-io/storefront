<template>
  <td class="col-span-6 md:col-span-4">{{ item.name }}</td>
  <td v-if="item.price" class="col-span-6 md:col-span-3">
    {{ (item.price).toFixed(2) }} € * {{ qty }} = {{ (item.price * qty).toFixed(2) }} € </td>
</template>

<script lang="ts" setup>
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();

const props = defineProps({
    identifier: { type: String, requiered: true },
    qty: { type: Number, requiered: true },
});

const item = ref({});

const load = async () => {
    item.value = await pb.collection("products").getOne(props.identifier);
};

onMounted(() => {
    load();
});
</script>