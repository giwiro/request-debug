import {Component, inject} from '@angular/core';
import {StoredGroupService} from '../dashboard/service/stored-group/stored-group.service';
import {RouterLink} from '@angular/router';
import {RequestGroupStore} from '../dashboard/store/request-group.store';

@Component({
  selector: 'app-landing',
  imports: [RouterLink],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css',
})
export class LandingComponent {
  storedGroupService = inject(StoredGroupService);
  store = inject(RequestGroupStore);

  onCreate = () => {
    this.store.createRequestGroup();
  };
}
