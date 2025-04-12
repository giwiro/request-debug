import {
  Component,
  inject,
  input,
  OnInit,
  OnDestroy,
  computed,
  signal,
} from '@angular/core';
import {RequestGroup} from '../../../../core/models';
import {RouterLink} from '@angular/router';
import {RequestGroupStore} from '../../store/request-group.store';
import {BadgeComponent} from '../../../../shared/badge/badge.component';
import {DatePipe} from '@angular/common';
import {debounceTime, Subject, Subscription} from 'rxjs';

@Component({
  selector: 'app-sidebar',
  imports: [RouterLink, BadgeComponent, DatePipe],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.css',
})
export class SidebarComponent implements OnInit, OnDestroy {
  store = inject(RequestGroupStore);
  requestId = input.required<string | undefined>();
  requestGroup = input.required<RequestGroup | null>();
  searchInputValue = signal<string>('');
  requests = computed(() => {
    const req = this.requestGroup()?.requests;
    const input = this.searchInputValue();

    if (input && req) {
      const found = req.find(r => r.id === input);
      if (!found) return [];

      return [found];
    }

    return req;
  });

  inputSubject: Subject<string>;
  input$: Subscription | undefined;

  constructor() {
    this.inputSubject = new Subject();
  }

  handleInputChange(event: Event) {
    if (this.inputSubject) {
      this.inputSubject.next((event.target as HTMLInputElement).value);
    }
  }

  ngOnInit() {
    this.input$ = this.inputSubject.pipe(debounceTime(500)).subscribe({
      next: input => {
        this.searchInputValue.set(input);
      },
    });
  }

  ngOnDestroy() {
    if (this.input$) {
      this.input$.unsubscribe();
    }
  }
}
