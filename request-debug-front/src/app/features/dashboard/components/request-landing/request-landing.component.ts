import {Component, computed, inject, OnInit, OnDestroy} from '@angular/core';
import {RequestGroupStore} from '../../store/request-group.store';
import {environment} from '../../../../../environments/environment';
import {join} from '../../../../shared/utils/path';
import ClipboardJS from 'clipboard';

@Component({
  selector: 'app-request-landing',
  imports: [],
  templateUrl: './request-landing.component.html',
  styleUrl: './request-landing.component.css',
  standalone: true,
})
export class RequestLandingComponent implements OnInit, OnDestroy {
  private clipboard: ClipboardJS | undefined;

  store = inject(RequestGroupStore);
  uniqueUrl = computed(() => {
    const groupId = this.store.requestGroup()?.id;
    return join(environment.baseApiUrl, `/group/${groupId}`);
  });

  ngOnInit(): void {
    this.clipboard = new ClipboardJS('#clipboard-button');
  }

  ngOnDestroy(): void {
    if (this.clipboard) {
      this.clipboard.destroy();
    }
  }
}
