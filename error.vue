<template>
  <NuxtLayout>
    <section class="page bg-white px-3 py-3 mx-auto max-w-6xl">
      <main
        class="grid min-h-full place-items-center bg-white px-6 py-24 sm:py-48 lg:px-8"
      >
        <div class="text-center">
          <p class="text-3xl font-semibold text-indigo-600">{{ error.statusCode }}</p>
          <h1
            class="text-3xl mt-6"
          >
            {{ error.message }}
          </h1>
          <p
            class="mt-6 text-lg font-medium text-pretty text-gray-500 sm:text-xl/8"
          >
            Sorry, we couldn’t find the page you’re looking for.
          </p>
          <div class="mt-10 flex items-center justify-center gap-x-6">
            <a
              href="#"
              class="btn btn-primary"
              >Go back home</a
            >
            <a href="#" class="btn btn-secondary"
              >Contact support 
              <FontAwesomeIcon :icon="faEdit" />
              </a
            >
          </div>
        </div>
      </main>
    </section>
  </NuxtLayout>
</template>

<script setup>
import { faEdit } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
const error = useError();
</script>