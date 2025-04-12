import {Component, input} from '@angular/core';

@Component({
  selector: 'app-badge',
  imports: [],
  templateUrl: './badge.component.html',
})
export class BadgeComponent {
  method = input.required<string | undefined>();
}
