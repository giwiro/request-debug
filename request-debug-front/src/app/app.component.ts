import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {AlertComponent} from './shared/alert/alert.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, AlertComponent],
  templateUrl: './app.component.html',
})
export class AppComponent {}
