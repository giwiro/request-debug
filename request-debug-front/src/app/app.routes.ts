import {Routes} from '@angular/router';
import {DashboardComponent} from './features/dashboard/dashboard.component';
import {StoredGroupService} from './features/dashboard/service/stored-group/stored-group.service';
import {inject} from '@angular/core';
import {LandingComponent} from './features/landing/landing.component';

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
    children: [
      {
        path: '',
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
        path: ':groupId',
        component: DashboardComponent,
      },
    ],
  },
];
