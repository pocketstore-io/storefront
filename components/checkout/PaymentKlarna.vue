<template>
  <div v-if="!locked" id="klarna-container" class="w-full mt-3"></div>
  <button v-else class="btn btn-neutral w-full">{{$t('payment.klarna.locked')}}</button>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const { locked } = defineProps({
  locked: {
    type: Boolean,
    required: false,
    default: ()=>{
      return false;
    }
  }
});

const router = useRouter();
useHead({
  script: [
    {
      async: true,
      onload: () => {
        klarnaAsyncCallback();
      },
      onerror: (err) => {
        console.error('Klarna SDK failed to load:', err);
      },
      src: 'https://x.klarnacdn.net/kp/lib/v1/api.js'
    }
  ]
});

let klarnaAsyncCallback = function () {
  window.Klarna.Payments.Buttons.init({
    client_id: "klarna_test_client_cVMhVCU_b0RwVG4pcml3ay16SCUxcnJqeVVJKnkjbmQsY2YzZWM4NDEtOTNlOS00MDc5LWI2NDctNzE0ZmZjZTM5M2Q2LDEsQ2F5OUJIMndTOXd5b3k4dzFGdTlkbVk1YlNkcUdRVGFhb1hoTnpsT0tNYz0",
  }).load(
    {
      container: "#klarna-container",
      theme: "default",
      shape: "default",
      on_click: (authorize) => {
        // Here you should invoke authorize with the order payload.
        authorize(
          { collect_shipping_address: true },
          {
            "purchase_country": "DE",
            "purchase_currency": "EUR",
            "locale": "de-DE",
            "order_amount": 11900,          // 100.00 EUR in minor units
            "order_tax_amount": 1900,       // 19.00 EUR tax in minor units
            "order_lines": [
              {
                "type": "physical",
                "reference": "19-402",
                "name": "T-shirt",
                "quantity": 1,
                "unit_price": 11900,         // 100.00 EUR
                "tax_rate": 1900,            // 20% tax rate
                "total_amount": 11900,       // 100.00 EUR total for the item
                "total_tax_amount": 1900     // 19.00 EUR total tax
              }
            ]
          }, // order payload 
          (result) => {
            console.log('auth result: ', result);
            router.push('/checkout/?klarna=' + result.session_id);
          },
        );
      },
    },
    function load_callback(loadResult) {
      console.log('result', loadResult);
      // Here you can handle the result of loading the button
    },
  );
};
</script>
