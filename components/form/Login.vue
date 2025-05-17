<template>
      <!--| TODO style Template nice |-->
  <form class="py-6 px-3" @submit.prevent="send()">
    <h2 class="font-bold text-lg ml-1 mb-3 text-center">Login Formular</h2>
    <div class="grid grid-cols-6">
      <div class="col-span-6 md:col-span-3">
        <label for="email" class="font-bold text-sm mb-1 ml-1">Email</label>
        <input v-model="form.login.email" type="email" name="email" laceholder="Deine Emailadresse"
          class="input input-bordered w-full">
      </div>
      <div class="col-span-6 md:col-span-3">
        <label for="password" class="font-bold text-sm mb-1 ml-1">Passwort</label>
        <input v-model="form.login.password" type="password" name="password" placeholder="Dein Passwort"
          class="input input-bordered w-full text-sm">
      </div>
    </div>
    <section class="form-action text-right mt-3">
      <button type="submit" class="btn btn-primary">abschicken</button>
    </section>
  </form>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { usePocketBase } from "~/util/pocketbase";

const router = useRouter();
const pb = usePocketBase();

const form = ref({
    login: {
        email: "",
        password: "",
    },
});

const send = async function () {
    await pb
        .collection("customers")
        .authWithPassword(form.value.login.email, form.value.login.password);
    if (pb.authStore.isValid) {
        router.push("/customer/profile");
    }
};
</script>
