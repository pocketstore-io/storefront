<template>
  <div class="overflow-hidden bg-white py-6 rounded-lg grid grid-cols-6">
    <div class="col-span-6 md:col-span-3 px-8">
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
    <div class="col-span-6 md:col-span-3 px-8">
      <LoginCardRegisterForm v-if="form == 'register'" />
      <LoginCardLoginForm v-else />
    </div>
  </div>
</template>


<script setup lang="ts">
import { faEdit } from '@fortawesome/free-solid-svg-icons';

const form = ref('register');
</script>