<template>
    <section v-if="checkoutStep == 'addresses'" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3">
        <div class="col-span-6 md:col-span-3">
            <h2 class="font-bold text-lg">{{ $t('checkout.shipping-address') }}</h2>
            <select v-model="selectedShippingAddress" class="select w-full block">
                <option value="new-one">neue Adresse</option>
                <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
            </select>
            <div class="divider divider-primary">Adresse</div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.firstname') }}</label>
                <input v-model="shipping.name" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.surname') }}</label>
                <input v-model="shipping.surname" type="text" class="input input-primary border-2 input-bordered">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.street') }}</label>
                <input v-model="shipping.street" type="text" class="input input-primary border-2 input-bordered">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.number') }}</label>
                <input v-model="shipping.number" type="number" class="input input-primary border-2 input-bordered">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.zip') }}</label>
                <input v-model="shipping.zip" type="text" class="input input-primary border-2 input-bordered">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.city') }}</label>
                <input v-model="shipping.city" type="text" class="input input-primary border-2 input-bordered">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
                <select v-model="shipping.country" class="select select-bordered select-primary border-2">
                    <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
                </select>
            </div>
        </div>
        <div v-if="!same" class="col-span-6 md:col-span-3">
            <h2 class="font-bold text-lg mb-3">Payment</h2>
            <select v-model="selectedPaymentAddress" class="select w-full block">
                <option value="new-one">neue Adresse</option>
                <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
            </select>
            <div class="divider divider-primary">Adresse</div>
            <section class="alert alert-neutral input-primary border-2">
                <input v-model="same" type="checkbox" class="checkbox checkbox-primary border-2 bg-white">
                <span>{{ $t('checkout.same') }}</span>
            </section>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.firstname') }}</label>
                <input v-model="payment.name" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.surname') }}</label>
                <input v-model="payment.surname" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.street') }}</label>
                <input v-model="payment.street" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.number') }}</label>
                <input v-model="payment.number" type="number" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.zip') }}</label>
                <input v-model="payment.zip" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.city') }}</label>
                <input v-model="payment.city" type="text" class="input input-bordered input-primary border-2">
            </div>
            <div class="form-control">
                <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
                <select v-model="payment.country" class="select select-bordered select-primary border-2">
                    <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
                </select>
            </div>
        </div>
        <div v-else class="col-span-6 md:col-span-3">
            <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment-address') }}</h2>
            <section class="alert border-primary border-2">
                <input v-model="same" type="checkbox" class="checkbox">
                <span>{{ $t('checkout.same') }}</span>
            </section>
        </div>
        <div class="col-span-6 flex justify-between py-3">
            <button class="btn btn-neutral" @click="checkoutStep = 'customer'">{{ $t('checkout.back.customer')
            }}</button>
            <button class="btn btn-primary" :disabled="(!valid.payment && !same) || !valid.shipping"
                @click="checkoutStep = 'payment'"><span>{{ $t('checkout.continue.addresses') }}</span>
                <Fa :icon="faChevronCircleRight" />
            </button>
        </div>
    </section>
</template>

<script lang="ts" setup>
import { faChevronCircleRight } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core'
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const selectedShippingAddress = ref('new-one');
const selectedPaymentAddress = ref('new-one');
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const same = useLocalStorage('same', true, {});
const addressess = ref([]);
const shipping = useLocalStorage('shipping', {
    name: '',
    surname: '',
    street: '',
    number: 1,
    zip: '',
    city: '',
    country: 'de',
}, {});
const payment = useLocalStorage('payment', {
    name: '',
    surname: '',
    street: '',
    number: 1,
    zip: '',
    city: '',
    country: 'de'
}, {});
const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});
const countries = reactive({
    'de': 'Deutschland'
});

onMounted(async () => {
    addressess.value = (await pb.collection('customer_addresses').getFullList(25));
    console.log(addressess.value);
    console.log('user: ', pb.authStore.model);
});

watch(selectedShippingAddress, (value) => {
    if (value == 'new-one') {
        shipping.value.name = '';
        shipping.value.surname = '';
        shipping.value.street = '';
        shipping.value.number = 1;
        shipping.value.city = '';
        shipping.value.zip = '';
        shipping.value.country = 'de';
    }
    else {
        addressess.value.map((item) => {
            if (item.id = value) {
                shipping.value.name = item.name.split(' ')[0];
                shipping.value.surname = item.name.split(' ')[1];
                shipping.value.street = item.street;
                shipping.value.number = item.number;
                shipping.value.city = item.city;
                shipping.value.zip = item.postcode;
                shipping.value.country = item.country;
            }
        })
    }
});
watch(selectedPaymentAddress, (value) => {
    if (value == selectedShippingAddress.value) {
        same.value = true;
    }
    else if (value == 'new-one') {
        payment.value.name = '';
        payment.value.surname = '';
        payment.value.street = '';
        payment.value.number = 1;
        payment.value.city = '';
        payment.value.zip = '';
        payment.value.country = 'de';
    }
    else {
        addressess.value.map((item) => {
            if (item.id = value) {
                payment.value.name = item.name.split(' ')[0];
                payment.value.surname = item.name.split(' ')[1];
                payment.value.street = item.street;
                payment.value.number = item.number;
                payment.value.city = item.city;
                payment.value.zip = item.postcode;
                payment.value.country = item.country;
            }
        })
    }
});

watch(shipping, () => {
    if (shipping.value.city == '') {
        valid.value.shipping = false;
        return
    }
    if (shipping.value.name == '') {
        valid.value.shipping = false;
        return
    }
    if (shipping.value.number < 1) {
        valid.value.shipping = false;
        return
    }
    if (shipping.value.street == '') {
        valid.value.shipping = false;
        return
    }
    if (shipping.value.surname == '') {
        valid.value.shipping = false;
        return
    }
    if (shipping.value.zip == '') {
        valid.value.shipping = false;
        return
    }
    valid.value.shipping = true;
}, { deep: true });

watch(payment, () => {
    if (!same.value) {
        if (payment.value.city == '') {
            valid.value.payment = false;
            return
        }
        if (payment.value.name == '') {
            valid.value.payment = false;
            return
        }
        if (payment.value.number < 1) {
            valid.value.payment = false;
            return
        }
        if (payment.value.street == '') {
            valid.value.payment = false;
            return
        }
        if (payment.value.surname == '') {
            valid.value.payment = false;
            return
        }
        if (payment.value.zip == '') {
            valid.value.payment = false;
            return
        }
    }
    valid.value.payment = true;
}, { deep: true });
</script>