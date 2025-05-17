import {
  ChangeDetectionStrategy,
  Component,
  computed,
  inject,
  OnInit,
  signal,
} from '@angular/core';
import {Request} from '../../../../core/models';
import {RequestGroupStore} from '../../store/request-group.store';
import {ActivatedRoute} from '@angular/router';
import {BadgeComponent} from '../../../../shared/badge/badge.component';
import {DatePipe} from '@angular/common';
import {DateAgoPipe} from '../../../../shared/date-ago/date-ago.pipe';
import {BytesToHumanPipe} from '../../../../shared/bytes-to-human/bytes-to-human.pipe';

@Component({
  selector: 'app-request',
  changeDetection: ChangeDetectionStrategy.OnPush,
  imports: [BadgeComponent, DatePipe, DateAgoPipe, BytesToHumanPipe],
  templateUrl: './request.component.html',
  styleUrl: './request.component.css',
})
export class RequestComponent implements OnInit {
  store = inject(RequestGroupStore);
  activatedRoute = inject(ActivatedRoute);

  requestId = signal<string | undefined>(undefined);
  request = computed<Request | undefined>(() => {
    const rg = this.store.requestGroup();
    if (rg && this.requestId()) {
      return rg.requests.find(r => r.id === this.requestId());
    }

    return undefined;
  });
  form = computed(() => {
    const request = this.request();
    if (!request || !request.form) return undefined;

    return new Map(Object.entries(request.form));
  });
  files = computed(() => {
    const request = this.request();
    if (!request || !request.files) return undefined;

    return new Map(Object.entries(request.files));
  });
  headers = computed(() => {
    const request = this.request();
    if (!request || !request.headers) return undefined;

    return new Map(Object.entries(request.headers));
  });
  queryParams = computed(() => {
    const request = this.request();
    if (!request || !request.queryParams) return undefined;

    return new Map(Object.entries(request.queryParams));
  });

  ngOnInit() {
    this.activatedRoute.params.subscribe(p => {
      if (p['requestId']) {
        this.requestId.set(p['requestId']);
      }
    });
  }
}
