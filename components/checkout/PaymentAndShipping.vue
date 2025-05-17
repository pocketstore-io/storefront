<template>
	<section v-if="checkoutStep == 'payment'" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3">
		<CheckoutPaymentView />
		<CheckoutShippingView />
		<div class="col-span-6 flex justify-between py-3">
			<button class="btn btn-primary" @click="checkoutStep = 'addresses'">{{ $t('checkout.back.addresses')
				}}</button>
			<button class="btn btn-primary" @click="checkoutStep = 'confirm'"><span>{{ $t('checkout.continue.confirm')
					}}</span>
				<Fa :icon="faChevronCircleRight" />
			</button>
		</div>
	</section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";
import { faChevronCircleRight } from "@fortawesome/free-solid-svg-icons";
const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
</script>