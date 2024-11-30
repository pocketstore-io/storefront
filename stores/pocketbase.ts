import { defineStore } from 'pinia'
import config from '../pocketstore.json';

export const usePocketbaseStore = defineStore({
  id: 'pocketbaseStore',
  state: () => ({
    url: config.domain
  }),
  actions: {}
})
