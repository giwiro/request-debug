import {Component, effect, inject, OnInit} from '@angular/core';
import {NavigationEnd, Router, RouterOutlet} from '@angular/router';
import {AlertComponent} from './shared/alert/alert.component';
import {ThemeStore} from './features/dashboard/store/theme.store';
import {DOCUMENT} from '@angular/common';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, AlertComponent],
  templateUrl: './app.component.html',
  standalone: true,
})
export class AppComponent implements OnInit {
  router = inject(Router);
  document = inject(DOCUMENT);
  themeStore = inject(ThemeStore);

  private _ = effect(() => {
    const t = this.themeStore.theme();

    const h = document.getElementsByTagName('html')[0];

    if (h) {
      h.setAttribute('data-theme', t);
    }
  });

  ngOnInit() {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        setTimeout(() => {
          if (typeof window !== 'undefined' && window.HSStaticMethods) {
            window.HSStaticMethods.autoInit();
          }
        }, 100);
      }
    });
  }
}
