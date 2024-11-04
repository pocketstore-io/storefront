<template>
  <section v-if="checkoutStep == 'cart' && loaded" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl">
    <div class="col-span-6 md:col-span-4">
      <table v-if="cart.length > 0" class="w-full">
        <thead>
          <tr class="grid grid-cols-12 gap-3 font-bold text-sm text-left px-3 py-3 bg-gray-200 my-3 mx-3">
            <th class="col-span-6 md:col-span-2">Anzahl</th>
            <th class="col-span-6 md:col-span-10">Product - Name</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in cart" :key="item.id" class="grid grid-cols-12 px-3 py-3 gap-3">
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
        </tbody>
      </table>
      <section v-else class="px-3 py-3">
        <section v-if="loaded" class="alert alert-warning text-white">
          <p class="block text-center">
            Dein Warenkorb ist leer, füge etwas hinzu bevor du es kaufen kannst.
          </p>
        </section>
      </section>
    </div>

    <div v-if="cart.length > 0" class="col-span-6 md:col-span-2">
      <CustomCartTotal />
    </div>
    <div v-if="cart.length > 0" class="col-span-6 flex justify-end px-3 py-3">
      <button class="btn btn-primary" :disabled="!valid.cart" @click="checkoutStep = 'customer'">
        <span>Continue</span>
        <Fa :icon="faChevronCircleRight" />
      </button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { faPlus, faTrash, faMinus, faChevronCircleRight } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";
import { useLocalStorage } from "@vueuse/core";

const cart = useLocalStorage("cart", [], {});
const valid = useLocalStorage(
  "checkout-valid",
  { cart: false, addresses: false, payment: false, shipping: false, confirm: false,customer: false },
  {}
);
const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const loaded = ref(false);

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

onMounted(() => {
  loaded.value = true;

  if (cart.value.length > 0) {
    valid.value.cart = true;
  }
});
</script>
