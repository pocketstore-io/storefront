<template>
  <section
    v-if="checkoutStep == 'confirm'"
    class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3"
  >
    <div class="col-span-6 flex justify-between">
      <div
        class="mx-auto bg-gray-400 shadow-md rounded-lg overflow-hidden my-24 max-w-2xl"
      >
        <div class="bg-indigo-600 text-white text-center py-6">
          <h1 class="text-2xl font-bold">{{ $t("checkout.your-order") }}</h1>
          <p class="mt-2 text-lg">{{ $t("checkout.order.cta") }}</p>
        </div>
        <div class="p-6">
          <p class="text-gray-700 text-lg">
            {{ $t("Hallo") }} <strong>{{ pb.authStore.record.name }}</strong
            >,<br />
            {{ $t("checkout.blabla") }}:
          </p>

          <div class="divider divider-neutral">
            {{ $t("checkout.cart.headline") }}
          </div>
          <table class="w-full mt-4 border border-gray-200">
            <thead>
              <tr class="bg-gray-100 text-left">
                <th class="p-3 border-b border-gray-200">
                  {{ $t("checkout.cart.product") }}
                </th>
                <th class="p-3 border-b border-gray-200">
                  {{ $t("checkout.cart.qty") }}
                </th>
                <th class="p-3 border-b border-gray-200">
                  {{ $t("checkout.cart.price") }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in cart">
                <td class="p-3 border-b border-gray-200">
                  {{ item.product.name }}
                </td>
                <td class="p-3 border-b border-gray-200">{{ item.qty }}x</td>
                <td class="p-3 border-b border-gray-200">
                  {{ item.product.price.toFixed(2) }} €
                </td>
              </tr>
            </tbody>
          </table>
          <div class="mt-6">
            <p class="text-gray-700 text-lg flex justify-end">
              {{ $t("checkout.subtotal") }}:
              <strong class="ml-3">{{ total.toFixed(2) }} €</strong>
            </p>
             <p class="text-gray-700 text-lg flex justify-end">
              {{ $t("checkout.tax") }}:
              <strong class="ml-3">{{ (total * 0.19).toFixed(2) }} €</strong>
            </p>
            <p class="text-gray-700 text-lg flex justify-end">
              {{ $t("checkout.total") }}:
              <strong class="ml-3">{{ (total * 1.19).toFixed(2) }} €</strong>
            </p>
          </div>

          <div class="divider divider-neutral">
            {{ $t("checkout.addressen") }}
          </div>

          <div class="grid grid-cols-6">
            <div class="col-span-6 md:col-span-3">
              {{ shipping.name }}
              {{ shipping.surname }} <br />
              {{ shipping.street }}
              {{ shipping.number }}<br />
              {{ shipping.zip }}
              {{ shipping.city }}<br />
              {{ $t("country." + shipping.country) }}
            </div>
            <div v-if="!same" class="col-span-6 md:col-span-3">
              {{ payment.name }}
              {{ payment.surname }} <br />
              {{ payment.street }}
              {{ payment.number }}<br />
              {{ payment.zip }}
              {{ payment.city }}<br />
              {{ $t("country." + payment.country) }}
            </div>
            <div v-else class="col-span-6 md:col-span-3">
              {{ shipping.name }}
              {{ shipping.surname }} <br />
              {{ shipping.street }}
              {{ shipping.number }}<br />
              {{ shipping.zip }}
              {{ shipping.city }}<br />
              {{ $t("country." + shipping.country) }}
            </div>
          </div>

          <div class="divider divider-neutral">
            {{ $t("checkout.shipping-and-payment") }}
          </div>

          <section class="grid grid-cols-6">
            <div class="col-span-6 md:col-span-3">
              {{ $t("payment.methods." + paymentMethod) }}
              <div v-for="(value, key) in paymentMethodInfo">
                <b>{{ $t("payment.label." + key) }}</b
                >: {{ value }} <br />
              </div>
            </div>
            <div class="col-span-6 md:col-span-3">
              {{ $t("shipping.methods." + shippingMethod) }}
              <div v-for="(value, key) in shippingMethodInfo">
                <b>{{ $t("checkout.label." + key) }}</b
                >: {{ value }} <br />
              </div>
            </div>
          </section>
        </div>
        <div class="bg-gray-400 p-6 text-gray-600 block text-center">
          <p>
            {{ $t("checkout.need-help") }}
            <a
              :href="'mailto:' + config.settings.support"
              class="text-indigo-600 hover:underline"
            >
              {{ config.settings.support }}</a
            >.
          </p>
        </div>
      </div>
    </div>
    <div class="col-span-6 flex justify-between py-3">
      <button class="btn btn-secondary" @click="checkoutStep = 'payment'">
        {{ $t("checkout.back.payment") }}
      </button>
      <button class="btn btn-primary" @click="confirmOrder()">
        {{ $t("checkout.place-order") }}
      </button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";
import config from "@/pocketstore.json";

const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const cart = useLocalStorage("cart", [], {});
const { t } = useI18n();

const paymentMethod = useLocalStorage("paymentMethod", "vorkasse", {});
const paymentMethodInfo = useLocalStorage("paymentMethodInfo", {}, {});
const shippingMethod = useLocalStorage("shippingMethod", "dhl", {});
const shippingMethodInfo = useLocalStorage("shippingMethodInfo", {}, {});

const same = useLocalStorage("same", true, {});
const pb = usePocketBase();

const router = useRouter();
const shipping = useLocalStorage(
  "shipping",
  {
    name: "",
    surname: "",
    street: "",
    number: 1,
    zip: "",
    city: "",
    country: "de",
  },
  {}
);
const payment = useLocalStorage(
  "payment",
  {
    name: "",
    surname: "",
    street: "",
    number: 1,
    zip: "",
    city: "",
    country: "de",
  },
  {}
);

const total = computed(() => {
  let tmp = 0;
  cart.value.map((item) => {
    tmp += item.qty * item.product.price;
  });

  return tmp;
});

const confirmOrder = async () => {
  const order = await pb.collection("orders").create({
    customer: pb.authStore.record?.id,
    cart: cart.value,
    payment_method: paymentMethod.value,
    payment_method_info: paymentMethodInfo.value,
    shipping_method_info: shippingMethodInfo.value,
    shipping_method: shippingMethod.value,
    payment_address: payment.value,
    shipping_address: shipping.value,
  });
  router.push("/checkout/confirm?order=" + order.id);
  cart.value = [];
  checkoutStep.value = "cart";
  paymentMethod.value = null;
  paymentMethodInfo.value = {};
  shippingMethod.value = null;
  shippingMethodInfo.value = {};
  payment.value = null;
  shipping.value = null;
};

onMounted(() => {
  useHead({
    title: "Checkout Confirm - " + t("general.title"),
  });
});
</script>