import {Component, inject, OnInit} from '@angular/core';
import {NavigationEnd, Router, RouterOutlet} from '@angular/router';
import {AlertComponent} from './shared/alert/alert.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, AlertComponent],
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
  router = inject(Router);

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
