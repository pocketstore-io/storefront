<template>
  <div>
    <section class="grid grid-cols-6 gap-x-3 px-3 py-3">
      <div class="col-spam-6">
        <h2 class="font-bold text-lg px-3 py-3">Total</h2>
      </div>
      <div class="col-span-6 bg-base-200 px-3 py-3 flex justify-between">
        <label for="" class="label font-bold">Preis</label>
        <span>{{ parseFloat(total).toFixed(2) }} €</span>
      </div>
      <div class="col-span-6 bg-base-300 px-3 py-3 flex justify-between">
        <label for="" class="label font-bold">Steuern</label>
        <span>{{ (parseFloat(total) * 0.19).toFixed(2) }} €</span>
      </div>
      <div class="col-span-6 bg-base-200 px-3 py-3 flex justify-between items-center">
        <label for="" class="label font-bold">Gesamt *</label>
        <span>{{ parseFloat(total.toFixed(2) * 1.19).toFixed(2) }} €</span>
      </div>
      <div class="col-span-6 px-3 py-3">
        <p class="text-sm block text-center">
          * Total + plus shipping and handling
        </p>
      </div>
    </section>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const cart = useLocalStorage("cart", [], {});
const total = computed(() => {
    let total = 0;
    cart.value.map((item) => {
        total += item.qty * item.product.price;
    });
    return parseFloat(total);
});
</script>