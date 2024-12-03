<template>
  <section class="card bg-base-300 px-3 py-3">
      <div class="form">
          <h2 class="font-bold text-lg mb-3">Register via</h2>
          <section class="actions">
              <div class="grid grid-cols-6 gap-3">
                  <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-block" @click="loginViaGoogle()">
                          <Fa :icon="faGoogle" />
                          <span>Google</span>
                      </button>
                  </div>
                  <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-block" @click="loginViaGithub()">
                          <Fa :icon="faGithub" />
                          <span>Github</span>
                      </button>
                  </div>
              </div>
          </section>
          <div class="divider divider-primary font-bold my-9">or</div>
          <h2 class="font-bold text-lg mb-3">Form</h2>
          <section class="grid grid-cols-6 gap-3">
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Vorname</label>
                  <input type="text" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Nachname</label>
                  <input type="text" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Email</label>
                  <input type="email" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Email</label>
                  <input type="email" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Passwort</label>
                  <input type="password" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                  <label for="" class="label text-sm font-bold">Passwort bestätigen</label>
                  <input type="password" class="input" placeholder="Name">
              </div>
              <section class="actions flex justify-between col-span-6">
                  <button class="btn btn-neutral" @click="form='login'">Jetzt anmelden</button>
                  <button class="btn btn-primary">Konto erstellen</button>
              </section>
          </section>
      </div>
  </section>
</template>


<script setup lang="ts">
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { faGithub, faGoogle } from "@fortawesome/free-brands-svg-icons";

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const loginViaGoogle = async function () {
  const authData = await pb.collection('customer').authWithOAuth2({ provider: 'google' });
  console.log(authData);
}
const loginViaGithub = async function () {
  const authData = await pb.collection('customer').authWithOAuth2({ provider: 'github' });
  console.log(authData);
}
</script>