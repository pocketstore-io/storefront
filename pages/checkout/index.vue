<template>
    <section class="page" :class="{
        'bg-white rounded': bgWhite
    }">
        <section v-if="loaded && checkoutStep != 'confirm'" class="grid grid-cols-10 px-3 py-3 gap-3">
            <div class="col-span-10">
                <div id="headline" class="divider divider-primary font-bold text-lg">{{ $t('checkout.checkout') }}</div>
            </div>
            <div class="col-span-10 md:col-span-2">
                <button class="btn btn-block"
                    :class="{ 'btn-primary': checkoutStep == 'cart', 'btn-neutral': checkoutStep != 'cart' }">{{
                        $t('checkout.cart.headline') }}</button>
            </div>
            <div class="col-span-10 md:col-span-2">
                <button class="btn btn-block"
                    :class="{ 'btn-primary': checkoutStep == 'customer', 'btn-neutral': checkoutStep != 'customer' }">{{
                        $t('checkout.customer') }}</button>
            </div>
            <div class="col-span-10 md:col-span-2">
                <button class="btn btn-block"
                    :class="{ 'btn-primary': checkoutStep == 'addresses', 'btn-neutral': checkoutStep != 'addresses' }">{{
                        $t('checkout.addresses') }}</button>
            </div>
            <div class="col-span-10 md:col-span-2">
                <button class="btn btn-block"
                    :class="{ 'btn-primary': checkoutStep == 'payment', 'btn-neutral': checkoutStep != 'payment' }">{{
                        $t('checkout.payment-and-shipping') }}</button>
            </div>
            <div class="col-span-10 md:col-span-2">
                <button class="btn btn-block"
                    :class="{ 'btn-primary': checkoutStep == 'confirm', 'btn-neutral': checkoutStep != 'confirm' }">{{
                        $t('checkout.confirm') }}</button>
            </div>
        </section>

        <CheckoutCart />
        <CheckoutAddresses />
        <CheckoutCustomer />
        <CheckoutPaymentAndShipping />
        <CheckoutConfirm />
    </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';

const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});

const loaded = ref(false);

const bgWhite = computed(()=>{
    return checkoutStep.value != 'confirm';
})

watch(checkoutStep, () => {
    document.getElementById('headline')?.scrollIntoView({
        behavior: "smooth"
    });
});

onMounted(() => {
    loaded.value = true;
});
</script>