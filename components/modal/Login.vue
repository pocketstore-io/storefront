<template>
  <div class="modal-box">
    <h3 class="text-lg font-bold">Login</h3>
    <section class="form">
      <label class="floating-label my-3">
        <span>Your Email</span>
        <input v-model="form.email" type="email" placeholder="mail@site.com" class="input input-md"/>
      </label>
      <label class="floating-label">
        <span>Your Password</span>
        <input v-model="form.password" type="password" placeholder="Password123" class="input input-md"/>
      </label>
      <section class="text-left py-5 px-3">
        <a href="/user/password/lost" class="text-sm text-primary">Passwort vergessen ?</a>
      </section>
      <section class="flex justify-between mt-3">
        <button class="btn btn-error" @click="open =false">schließen</button>
        <button class="btn btn-primary" @click="login()">einloggen</button>
      </section>
    </section>
  </div>
</template>
<script setup lang="ts">
import {useLocalStorage} from "@vueuse/core";
import {usePocketBase} from "~/util/pocketbase";

const pb = usePocketBase();
const form = ref({
  email: '',
  password: ''
})

const login = async () => {
  await pb.collection('customers').authWithPassword(form.value.email, form.value.password);
  if(pb.authStore.isValid){
    open.value = false;
  }
}
const open = useLocalStorage('modal-login', false, {})
</script>