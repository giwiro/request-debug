import {Component, ComponentRef, inject, OnInit, signal} from '@angular/core';
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

@Component({
  selector: 'app-dashboard',
  imports: [RouterLink, SidebarComponent, RouterOutlet],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent implements OnInit {
  router = inject(Router);
  store = inject(RequestGroupStore);
  activatedRoute = inject(ActivatedRoute);
  groupId = signal<string | undefined>(undefined);
  requestId = signal<string | undefined>(undefined);

  childrenActivatedRoute$: Subscription | undefined;

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
      }
    });
  }
}
