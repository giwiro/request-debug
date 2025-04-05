import {Routes} from '@angular/router';
import {DashboardComponent} from './features/dashboard/components/dashboard/dashboard.component';
import {StoredGroupService} from './features/dashboard/service/stored-group/stored-group.service';
import {inject} from '@angular/core';
import {LandingComponent} from './features/landing/landing.component';
import {RequestComponent} from './features/dashboard/components/request/request.component';
import {RequestLandingComponent} from './features/dashboard/components/request-landing/request-landing.component';

export const routes: Routes = [
  {
    title:
      'HTTP Request debug tool that displays all requests made to the endpoint.',
    path: '',
    pathMatch: 'full',
    component: LandingComponent,
  },
  {
    title:
      'HTTP Request debug tool that displays all requests made to the endpoint.',
    path: 'dashboard',
    pathMatch: 'full',
    redirectTo: () => {
      const storedGroupService = inject(StoredGroupService);

      if (storedGroupService.storedRequestGroupId) {
        return `/dashboard/${storedGroupService.storedRequestGroupId}`;
      }

      return `/`;
    },
  },
  {
    title:
      'HTTP Request debug tool that displays all requests made to the endpoint.',
    path: 'dashboard/:groupId',
    component: DashboardComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        component: RequestLandingComponent,
      },
      {
        path: ':requestId',
        component: RequestComponent,
      },
    ],
  },
];
