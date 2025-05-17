import {
  Component,
  inject,
  input,
  OnInit,
  OnDestroy,
  computed,
  signal,
  ChangeDetectionStrategy,
} from '@angular/core';
import {Request} from '../../../../core/models';
import {Router, RouterLink} from '@angular/router';
import {RequestGroupStore} from '../../store/request-group.store';
import {BadgeComponent} from '../../../../shared/badge/badge.component';
import {DatePipe} from '@angular/common';
import {debounceTime, Subject, Subscription} from 'rxjs';

@Component({
  selector: 'app-sidebar',
  changeDetection: ChangeDetectionStrategy.OnPush,
  imports: [RouterLink, BadgeComponent, DatePipe],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.css',
  standalone: true,
})
export class SidebarComponent implements OnInit, OnDestroy {
  store = inject(RequestGroupStore);
  router = inject(Router);
  requestId = input.required<string | undefined>();
  searchInputValue = signal<string>('');
  requests = computed(() => {
    const req = this.store.requestGroup()?.requests;
    const input = this.searchInputValue();

    if (input && req) {
      const found = req.find(r => r.id === input);
      if (!found) return [];

      return [found];
    }

    if (req) return [...req];

    return [];
  });

  inputSubject: Subject<string>;
  input$: Subscription | undefined;

  constructor() {
    this.inputSubject = new Subject();
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

  handleDelete(event: Event, request: Request) {
    const groupId = this.store.requestGroup()!.id;

    this.store.deleteRequest({
      requestGroupId: groupId,
      requestId: request.id,
    });

    this.router
      .navigate(['/dashboard/', {groupId}])
      .then(() => console.log(`Redirecting to '/dashboard/${groupId}/'`))
      .catch(() =>
        console.log(`Could not redirect to '/dashboard/${groupId}/'`)
      );
  }

  handleInputChange(event: Event) {
    if (this.inputSubject) {
      this.inputSubject.next((event.target as HTMLInputElement).value);
    }
  }
}
