import {Component, computed, inject} from '@angular/core';
import {RequestGroupStore} from '../../store/request-group.store';
import {environment} from '../../../../../environments/environment';
import {join} from '../../../../shared/utils/path';

@Component({
  selector: 'app-request-landing',
  imports: [],
  templateUrl: './request-landing.component.html',
  styleUrl: './request-landing.component.css',
})
export class RequestLandingComponent {
  store = inject(RequestGroupStore);
  uniqueUrl = computed(() => {
    const groupId = this.store.requestGroup()?.id;
    return join(environment.baseApiUrl, `/group/${groupId}`);
  });
}
