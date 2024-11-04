import { defineStore } from 'pinia'

export const usePocketbaseStore = defineStore({
  id: 'pocketbaseStore',
  state: () => ({
    url: 'https://admin.pocketstore.io/'
  }),
  actions: {}
})
