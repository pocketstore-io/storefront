<template>
    <div class="col-span-6 md:col-span-3 space-y-3">
        <h2 class="font-bold text-lg">{{ $t('checkout.shipping-method') }} {{ $t('general.select') }}</h2>
        <section v-for="option in optionsShipping" :key="option"
            class="rounded-lg py-3 px-3 w-full border-2 border-base-200">
            <div class="flex justify-between items-center">
                <section class="input-group flex border-red-400">
                    <input :id="'shipping.methods.' + option.code" v-model="shippingMethod" type="radio"
                        :value="option.code" name="shippingMethod"
                        class="ml-auto radio border-2 border-black radio-primary mt-2 mr-3">
                    <label :for="'shipping.methods.' + option.code" class="label">{{ $t('shipping.methods.' +
                        option.code)
                        }}</label>
                </section>
                <div class="price-tag">
                    <b>{{ option.price.toFixed(2) }} €</b>
                </div>
            </div>
        </section>
        <section v-if="shippingMethod == 'click-and-collect'" class="border-base-300 border-2 rounded-lg px-3 py-3">
            <label for="" class="label font-bold text-lg">{{ $t('general.stores') }}</label>
            <select v-model="selectedStore" class="select w-full border-2 border-base-300 rounded-lg bg-white">
                <option v-for="store in stores" :key="store.id" :value="store.id"
                    :class="{ 'bg-base-200': store.id != selectedStore, 'bg-primary': store.id == selectedStore }">
                    {{ store.street }} {{ store.number }}, {{ store.zip }} {{ store.city }}
                </option>
            </select>
        </section>
    </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { usePocketBase } from '~/util/pocketbase';

const pb = usePocketBase();

const stores = ref([]);
const optionsCountry = ref([]);
const selectedStore = ref(null);

const optionsShipping = ref([]);
const shippingMethod = useLocalStorage('shippingMethod', 'click-and-collect', {});
const shippingMethodInfo = useLocalStorage('shippingMethodInfo', {}, {});

watch(selectedStore, (value) => {
    shippingMethodInfo.value = {
        store: value
    }
});

watch(shippingMethod, (value) => {
    if (value != 'click-and-collect') {
        shippingMethodInfo.value = {}
    }
});

onMounted(async () => {
    stores.value = await pb.collection('stores').getFullList();

    optionsShipping.value = await pb.collection('shipping_methods').getFullList({
        filter: 'active=true'
    });
    optionsCountry.value = await pb.collection('countries').getFullList();
    if (shippingMethod.value === 'click-and-collect') {
        shippingMethodInfo.value = {
            store: { "id": stores.value[0].id, "name": stores.value[0].name }
        }
    }
    else {
        shippingMethodInfo.value = {}
    }

    selectedStore.value = stores.value[0].id;
});

</script>