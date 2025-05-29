import {
  patchState,
  signalStore,
  withHooks,
  withMethods,
  withState,
} from '@ngrx/signals';
import {localStorage} from '../../../shared/utils/local-storage';

type AppTheme = 'light' | 'dark';

const STORED_MODE_KEY = 'STORED_MODE_KEY';

interface ThemeStoreState {
  theme: AppTheme;
}

const initialState: ThemeStoreState = {
  theme: 'light',
};

export const ThemeStore = signalStore(
  {providedIn: 'root'},
  withState(initialState),
  withMethods(store => ({
    loadThemeState: () => {
      const storedTheme = localStorage.getItem(STORED_MODE_KEY);

      const preferredTheme: AppTheme =
        window.matchMedia &&
        window.matchMedia('(prefers-color-scheme: dark)').matches
          ? 'dark'
          : 'light';

      if (storedTheme && (storedTheme === 'light' || storedTheme === 'dark')) {
        patchState(store, () => ({
          theme: storedTheme as AppTheme,
        }));
      } else {
        patchState(store, () => ({
          theme: preferredTheme,
        }));
      }
    },
    updateThemeState: (theme: AppTheme) => {
      localStorage.setItem(STORED_MODE_KEY, theme);
      patchState(store, () => ({theme}));
    },
    saveThemeState: () => {
      localStorage.setItem(STORED_MODE_KEY, store.theme());
    },
  })),
  withHooks({
    onInit: store => {
      store.loadThemeState();
    },
    onDestroy: store => {
      store.saveThemeState();
    },
  })
);
