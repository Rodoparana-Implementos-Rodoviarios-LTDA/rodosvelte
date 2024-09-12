import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

function createPersistedStore<T>(key: string, initialValue: T): Writable<T> {
    const store: Writable<T> = writable<T>(initialValue, () => {
      const storedValue = localStorage.getItem(key);
      if (storedValue) {
        store.set(JSON.parse(storedValue) as T);
      }
  
      return store.subscribe((value: T) => {
        localStorage.setItem(key, JSON.stringify(value));
      });
    });
  
    return store;
  }
  
  export const username = createPersistedStore<string>('username', '');
  export const token = createPersistedStore<string>('token', '');