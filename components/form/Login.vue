<template>
  <form class="py-6 px-3" @submit.prevent="send()">
    <h2 class="font-bold text-lg ml-1 mb-3 text-center">Login Formular</h2>
    <section class="form-control">
      <label for="email" class="font-bold text-sm mb-1 ml-1">Email</label>
      <input
v-model="form.login.email" type="email" name="email" laceholder="Deine Emailadresse"
        class="input input-bordered w-full max-w-xs" >
    </section>
    <section class="form-control mt-3">
      <label for="password" class="font-bold text-sm mb-1 ml-1">Passwort</label>
      <input
v-model="form.login.password" type="password" name="password" placeholder="Dein Passwort"
        class="input input-bordered w-full max-w-xs" >
    </section>
    <section class="form-action text-right mt-3">
      <button type="submit" class="btn btn-primary">abschicken</button>
    </section>
  </form>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useRouter } from 'vue-router';

const router = useRouter();
const store = usePocketbaseStore();

const { url } = storeToRefs(store);

const pb = new PocketBase(url.value);

const form = ref({
  login: {
    email: '',
    password: ''
  },
});

const send = async function () {
  await pb.collection('customer').authWithPassword(form.value.login.email, form.value.login.password)
  if (pb.authStore.isValid) {
    router.push('/customer/profile');
  }
}
</script>
