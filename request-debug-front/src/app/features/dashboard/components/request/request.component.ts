import {Component, computed, inject, OnInit, signal} from '@angular/core';
import {Request} from '../../../../core/models';
import {RequestGroupStore} from '../../store/request-group.store';
import {ActivatedRoute} from '@angular/router';
import {JsonPipe} from '@angular/common';

@Component({
  selector: 'app-request',
  imports: [JsonPipe],
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

  ngOnInit() {
    this.activatedRoute.params.subscribe(p => {
      if (p['requestId']) {
        this.requestId.set(p['requestId']);
      }
    });
  }
}
