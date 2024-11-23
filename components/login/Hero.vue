<template>
  <div class="overflow-hidden bg-white py-6 rounded-lg">
    <div class="mx-auto max-w-7xl px-6 lg:px-8">
      <div class="mx-auto grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-2">
        <div class="lg:pr-8 lg:pt-4">
          <div class="lg:max-w-lg">
            <h2 class="text-base/7 font-semibold text-indigo-600">Schneller bestellen</h2>
            <p class="mt-2 text-pretty text-2xl font-semibold tracking-tight text-gray-900 sm:text-2xl">
              Einfach nach Produkten suchen und sie bestellen
            </p>
            <p class="mt-6 text-lg/6 text-gray-600">
              Lorem, ipsum dolor sit amet consectetur adipisicing elit. Recusandae omnis
              tempora possimus, vel voluptatibus explicabo iusto neque tenetur libero
              reprehenderit, totam molestiae cum? Modi odit tempora quis nostrum aliquam
              ad.
            </p>
            <div class="mt-10 max-w-xl space-y-8 text-base/7 text-gray-600 lg:max-w-none">
              <div>
                <div class="flex">
                  <section class="icon text-indigo-600 mr-5">
                    <Fa :icon="faEdit" />
                  </section>
                  <h3 class="font-bold">Schnellerer Checkout-Prozess</h3>
                </div>
                <p class="pl-9 text-sm">
                  Kunden können ihre Daten wie Adresse und Zahlungsinformationen
                  speichern, was den Bestellvorgang vereinfacht und beschleunigt.
                </p>
              </div>
              <div>
                <div class="flex">
                  <section class="icon text-indigo-600 mr-5">
                    <Fa :icon="faEdit" />
                  </section>
                  <h3 class="font-bold">Übersicht über Bestellungen</h3>
                </div>
                <p class="pl-9 text-sm">
                  Kunden können vergangene Bestellungen einsehen, deren Status verfolgen
                  und bei Bedarf Rechnungen oder Details erneut abrufen.
                </p>
              </div>
              <div>
                <div class="flex">
                  <section class="icon text-indigo-600 mr-5">
                    <Fa :icon="faEdit" />
                  </section>
                  <h3 class="font-bold">Personalisierte Angebote</h3>
                </div>
                <p class="pl-9 text-sm">
                  Kunden erhalten personalisierte Empfehlungen, Rabatte oder Angebote, die
                  auf ihren Vorlieben und bisherigen Käufen basieren.
                </p>
              </div>
            </div>
          </div>
        </div>
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
                <label for="" class="label text-sm font-bold">Hallo Welt</label>
                <input type="text" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6 md:col-span-3">
                <label for="" class="label text-sm font-bold">Hallo Welt</label>
                <input type="text" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6">
                <label for="" class="label text-sm font-bold">Hallo Welt</label>
                <input type="text" class="input" placeholder="Name">
              </div>
              <div class="form-control col-span-6">
                <label for="" class="label text-sm font-bold">Hallo Welt</label>
                <input type="text" class="input" placeholder="Name">
              </div>
              <section class="actions flex justify-between">
                <button class="btn btn-primary">Hallo</button>
                <button class="btn btn-primary">Welt</button>
              </section>
            </section>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { faGithub, faGoogle } from "@fortawesome/free-brands-svg-icons";
import { faEdit } from "@fortawesome/free-solid-svg-icons";

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const loginViaGoogle = async function () {
  const authData = await pb.collection('customer').authWithOAuth2({ provider: 'google' });
  console.log(authData);
}
const loginViaGithub = async function () {
  const authData = await pb.collection('customer').authWithOAuth2({ provider: 'github', redirectUrl: 'https://demo.pocketstore.io/de/login' });
  console.log(authData);
}
</script>