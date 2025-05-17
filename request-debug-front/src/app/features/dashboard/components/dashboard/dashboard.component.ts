import {
  Component,
  ComponentRef,
  inject,
  OnInit,
  OnDestroy,
  signal,
  ChangeDetectionStrategy,
} from '@angular/core';
import {RequestGroupStore} from '../../store/request-group.store';
import {
  ActivatedRoute,
  Router,
  RouterLink,
  RouterOutlet,
} from '@angular/router';
import {SidebarComponent} from '../sidebar/sidebar.component';
import {RequestComponent} from '../request/request.component';
import {Subscription} from 'rxjs';
import {EventSourceService} from '../../../../core/sse/event-source.service';
import {join} from '../../../../shared/utils/path';
import {environment} from '../../../../../environments/environment';
import {Request} from '../../../../core/models';

@Component({
  selector: 'app-dashboard',
  changeDetection: ChangeDetectionStrategy.OnPush,
  imports: [RouterLink, SidebarComponent, RouterOutlet],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
  standalone: true,
})
export class DashboardComponent implements OnInit, OnDestroy {
  router = inject(Router);
  store = inject(RequestGroupStore);
  activatedRoute = inject(ActivatedRoute);
  groupId = signal<string | undefined>(undefined);
  requestId = signal<string | undefined>(undefined);
  eventSourceService = inject(EventSourceService);

  childrenActivatedRoute$: Subscription | undefined;
  eventSource$: Subscription | undefined;
  sseEventSource: EventSource | undefined;

  routeActivated(componentRef: ComponentRef<never>) {
    if (componentRef instanceof RequestComponent) {
      this.childrenActivatedRoute$ =
        this.activatedRoute.firstChild?.params.subscribe(p => {
          if (p['requestId']) {
            this.requestId.set(p['requestId']);
          }
        });
    }
  }

  routeDeactivated() {
    if (this.childrenActivatedRoute$) {
      this.requestId.set(undefined);
      this.childrenActivatedRoute$.unsubscribe();
    }
  }

  ngOnInit() {
    this.activatedRoute.params.subscribe(p => {
      if (p['groupId']) {
        this.groupId.set(p['groupId']);
        this.store.getRequestGroup(p['groupId']);

        const sseUrl = join(
          environment.baseApiUrl,
          `/api/group/${p['groupId']}/sse`
        );
        const [obs, eventSource] = this.eventSourceService.connectSSE<string>(
          sseUrl,
          'sse-requests'
        );

        this.eventSource$ = obs.subscribe({
          next: event => {
            let req: Request | undefined = undefined;

            try {
              req = JSON.parse(event.data);
            } catch (e) {
              console.log(`Could not parse '${event.data}'`, e);
            }

            if (req) {
              this.store.addRequest(req);
            }
          },
          error: error => {
            console.log('error', error);
          },
        });
        this.sseEventSource = eventSource;
      }
    });
  }

  ngOnDestroy() {
    if (this.eventSource$) {
      this.eventSource$.unsubscribe();
    }

    if (this.sseEventSource) {
      this.sseEventSource.close();
    }
  }
}
