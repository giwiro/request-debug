import {Component, inject, Input} from '@angular/core';
import {RequestGroupStore} from './store/request-group.store';
import {JsonPipe} from '@angular/common';

@Component({
  selector: 'app-dashboard',
  imports: [JsonPipe],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent {
  store = inject(RequestGroupStore);

  @Input()
  set groupId(id: string) {
    this.store.getRequestGroup(id);
  }
}
