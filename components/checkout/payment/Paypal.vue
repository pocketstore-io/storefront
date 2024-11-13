<template>
  <div>
    <!--| Move PayPal into component |-->
    <div id="paypal-button-container" class="mt-3" />
    <button v-if="locked" @click="resetLock()" class="btn btn-secondary w-full mt-3">
      <Fa :icon="faTrash" /> <span>Reset Stripe Payment</span>
      <Fa :icon="faTrash" />
    </button>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { loadScript } from "@paypal/paypal-js";
import { cartToPurchaseUnits } from '~/util/paypal';

const { locked } = defineProps({
  locked: {
    type: Boolean,
    required: false,
    default: () => {
      return false;
    }
  }
});

const paymentMethod = useLocalStorage('paymentMethod', 'vorkasse', {});
const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});

const resetLock = function(){

}

const initPaypal = () => {
  loadScript({ "client-id": 'AUtZarD95T0LXYT1Dn51mBIBswWY3QDb-S380C-_xgfhqqJ6IEJHZh-wuRnRzNy184Lj5fpbXokZBiOk', currency: 'EUR' })
    .then((paypal) => {
      paypal.Buttons({
        createOrder: function (data, actions) {
          // Directly create an order with the amount, no server needed
          return actions.order.create({
            purchase_units: cartToPurchaseUnits()
          });
        },
        onApprove: function (data, actions) {
          return actions.order.capture().then(function (details) {
            alert('Transaction completed by ' + details.payer.name.given_name);
            console.log(details); // You can log or handle the transaction details here
          });
          // TODO store details for order
        },
        onError: function (err) {
          console.error('An error occurred during the transaction', err);
        }
      }).render('#paypal-button-container');
    })
    .catch((err) => {
      console.error("failed to load the PayPal JS SDK script", err);
    });
}

onMounted(() => {
  if (paymentMethod.value == 'paypal' && checkoutStep.value == 'payment') {
    initPaypal();
  }
});

watch(paymentMethod, (value) => {
  if (value === 'paypal') {
    initPaypal();
  }
});

watch(checkoutStep, (value) => {
  if (paymentMethod.value == 'paypal' && value == 'payment') {
    initPaypal();
  }
});
</script>