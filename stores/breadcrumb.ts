import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core';
const breadcrumbs = useLocalStorage('breadcrumbs', [], {})

export const useBreadcrumbStore = defineStore({
  id: 'myBreadcrumbStore',
  state: () => ({

  }),
  actions: {
    clear: () => {
      breadcrumbs.value = [];
    },
    add: (breadcrumb) => {
      breadcrumbs.value.push({
        link: breadcrumb.link,
        id: breadcrumb.id
      })
    }
  }
})
