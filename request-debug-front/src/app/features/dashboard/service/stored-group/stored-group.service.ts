import {Injectable, signal, effect} from '@angular/core';
import {localStorage} from '../../../../shared/utils/local-storage';

const STORED_REQUEST_GROUP_ID_KEY = 'STORED_REQUEST_GROUP_ID_KEY';

@Injectable({
  providedIn: 'root',
})
export class StoredGroupService {
  private requestGroupId = signal<string | null>(
    localStorage.getItem(STORED_REQUEST_GROUP_ID_KEY)
  );

  private _ = effect(() => {
    if (this.storedRequestGroupId === null) {
      localStorage.removeItem(STORED_REQUEST_GROUP_ID_KEY);
    } else {
      localStorage.setItem(
        STORED_REQUEST_GROUP_ID_KEY,
        this.storedRequestGroupId
      );
    }
  });

  get storedRequestGroupId() {
    return this.requestGroupId();
  }

  set storedRequestGroupId(groupId) {
    this.requestGroupId.set(groupId);
  }
}
