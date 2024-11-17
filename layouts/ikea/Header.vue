<script setup lang="ts">
import { faCamera, faEdit, faGift, faGlobe, faHeart, faHouse, faMagnifyingGlass, faShoppingCart, faTruck, faUser } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome';
</script>

<template>
    <div class="w-full bg-black px-3 py-3 text-white flex justify-between">
        <section class="lang-select flex items-center">
            <section class="icon mx-3">
                <Fa :icon="faGlobe" />
            </section>
            <select name="" id="" class="select select-xs bg-black text-white">
                <option value="de">DE - Deutsch</option>
            </select>
        </section>
        <section class="features text-sm flex items-center">
            <section class="icon mt-1 mx-3">
                <Fa :icon="faGift" />
            </section>
            <span>50 Jahre IKEA Deutschland</span>
        </section>
        <section class="store-select flex items-center">
            <section class="icon mt-1 mx-3">
                <Fa :icon="faTruck" />
            </section>
            <button class="btn border-none hover:bg-transparent hover:border-none text-white bg-black">PLZ
                eingeben</button>
            <span class="px-3 py-3">|</span>
            <section class="icon mt-1 mx-3">
                <Fa :icon="faHouse" />
            </section>
            <button class="btn border-none hover:bg-transparent hover:border-none text-white bg-black">Standort
                auswählen</button>
        </section>
    </div>
    <div class="w-full bg-white mb-3 px-3 py-3 flex justify-between max-w-6xl mx-auto">
        <section class="logo-search flex">
            <img src="https://www.ikea.com/global/assets/logos/brand/ikea.svg" class="h-8" alt="">
            <div class="flex ml-3 bg-base-300 join">
                <div class="bg-red-400 join-item rounded-full mr-3 px-3 flex items-center">
                    <Fa :icon="faMagnifyingGlass" />
                </div>
                <input type="text" class="input input-sm bg-base-300 join-item" placeholder="Was suchen sie ?">
                <div class="bg-red-400 join-item rounded-full ml-3 px-3 flex items-center">
                    <Fa :icon="faCamera" />
                </div>
            </div>
        </section>
        <section class="icons flex justify-end space-x-6">
            <section class="icon mt-1 mr-2">
                <Fa :icon="faUser" />
                <span class="ml-3 text-sm">Hallo, Jonathan</span>
            </section>
            <section class="icon mt-1 mr-2">
                <Fa :icon="faHeart" />
            </section>
            <section class="icon mt-1 mr-2">
                <Fa :icon="faShoppingCart" />
            </section>
        </section>
    </div>
</template>