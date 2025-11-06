<template>
  <button class="btn btn-secondary" :disabled="props.qty < 1 || stock.length == 0" @click="addToCart(props.item.id)">
    {{ $t("checkout.add-to-cart") }}
  </button>
</template>

<script setup lang="ts">
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";

const props = defineProps({
    item: {
        required: true,
        type: Object,
    },
    stock: {
        required: true,
        type: Object,
    },
    qty: {
        required: true,
        type: Number,
    },
});

const qty = ref(1);
const cart = useLocalStorage("cart", [], {});
const pb = usePocketBase();

const addToCart = async function (id) {
    let found = false;
    if (typeof cart.value == "undefined") {
        cart.value = [];
    }

    // TODO Toast for Push item
    const product = await pb.collection("products").getOne(id);

    cart.value.map((item) => {
        if (item.id == id) {
            found = true;
        }
    });

    if (found) {
        cart.value.map((item) => {
            if (item.id == id) {
                item.qty += qty.value;
            }
        });
    } else {
        // TODO cart message
        cart.value.push({
            qty: qty.value,
            id,
            product,
        });
    }
};
</script>