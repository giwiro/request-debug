import {Component, input} from '@angular/core';
import {RequestGroup} from '../../../../core/models';
import {RouterLink} from '@angular/router';

const formatter = new Intl.DateTimeFormat('en-US', {
  year: 'numeric',
  month: '2-digit',
  day: '2-digit',
  hour: '2-digit',
  minute: '2-digit',
  second: '2-digit',
});

@Component({
  selector: 'app-sidebar',
  imports: [RouterLink],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.css',
})
export class SidebarComponent {
  requestGroup = input.required<RequestGroup | null>();

  parseDate(date: string): string {
    return formatter.format(new Date(date));
  }
}
