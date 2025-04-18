<template>
  <tr class="grid grid-cols-12 px-3 py-3 gap-3">
    <td class="col-span-6 md:col-span-2">{{ item.qty }}</td>
    <CustomCartItemRow :identifier="item.id" :qty="item.qty" />
    <td class="col-span-6 md:col-span-3 join flex justify-end">
      <button class="btn btn-primary join-item btn-sm" @click="increaseCart(index)">
        <Fa :icon="faPlus" />
      </button>
      <button class="btn btn-error join-item btn-sm" @click="cart.splice(index, 1)">
        <Fa :icon="faTrash" />
      </button>
      <button class="btn btn-neutral join-item btn-sm" @click="decreaseCart(index)">
        <Fa :icon="faMinus" />
      </button>
    </td>
  </tr>
</template>

<script lang="ts" setup>
import { faMinus, faPlus, faTrash } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const props = defineProps({
  item: { type: Object, required: true },
  index: { type: Number, required: true },
});

const cart = useLocalStorage("cart", [], {});
const valid = useLocalStorage(
  "checkout-valid",
  { cart: false, addresses: false, payment: false, shipping: false, confirm: false,customer: false },
  {}
);

const increaseCart = (index) => {
  cart.value[index].qty++;
  if (cart.value.length > 0) {
    valid.value.cart = true;
  }
};

const decreaseCart = (index) => {
  if (cart.value[index].qty > 1) {
    cart.value[index].qty--;
  }
  if (cart.value.length > 0) {
    valid.value.cart = true;
  }
};
</script>