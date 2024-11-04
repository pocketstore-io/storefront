import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core';
const messages = useLocalStorage('messages', [], {})

export const useMessagesStore = defineStore({
  id: 'myMessagesStore',
  state: () => ({
    
  }),
  actions: {
    add: (message) => {
      messages.value.push({
        'message': message.message,
        timer: Math.floor(Date.now() / 1000) + 5
      })
    }
  }
})
