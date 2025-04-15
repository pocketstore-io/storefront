<template>
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
</template>

<script setup lang="ts">
import { faEdit } from '@fortawesome/free-solid-svg-icons';
</script>