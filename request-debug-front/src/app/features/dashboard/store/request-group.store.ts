import {RequestGroup} from '../../../core/models';
import {patchState, signalStore, withMethods, withState} from '@ngrx/signals';
import {tapResponse} from '@ngrx/operators';
import {inject} from '@angular/core';
import {RequestGroupService} from '../service/request-group/request-group.service';
import {rxMethod} from '@ngrx/signals/rxjs-interop';
import {exhaustMap, pipe, tap} from 'rxjs';
import {HttpErrorResponse} from '@angular/common/http';
import {StoredGroupService} from '../service/stored-group/stored-group.service';
import {Router} from '@angular/router';

interface RequestGroupStoreState {
  isLoading: boolean;
  requestGroup: RequestGroup | null;
}

const initialState: RequestGroupStoreState = {
  isLoading: false,
  requestGroup: null,
};

export const RequestGroupStore = signalStore(
  {providedIn: 'root'},
  withState(initialState),
  withMethods(
    (
      store,
      requestGroupService = inject(RequestGroupService),
      storedGroupService = inject(StoredGroupService),
      router = inject(Router)
    ) => ({
      createRequestGroup: rxMethod<void>(
        pipe(
          tap(() => patchState(store, {isLoading: true, requestGroup: null})),
          exhaustMap(() => {
            return requestGroupService.createRequestGroup().pipe(
              tapResponse({
                next: requestGroup => {
                  patchState(store, {requestGroup, isLoading: false});
                  storedGroupService.storedRequestGroupId = requestGroup.id;
                  router
                    .navigate(['/dashboard', requestGroup.id])
                    .then(() =>
                      console.log(`Redirect to '/dashboard/${requestGroup.id}'`)
                    )
                    .catch(() =>
                      console.log(
                        `Could not redirect to '/dashboard/${requestGroup.id}'`
                      )
                    );
                },
                error: () => {
                  patchState(store, {isLoading: false});
                },
              })
            );
          })
        )
      ),
      getRequestGroup: rxMethod<string>(
        pipe(
          tap(() => patchState(store, {isLoading: true, requestGroup: null})),
          exhaustMap((id: string) => {
            return requestGroupService.getRequestGroup(id).pipe(
              tapResponse({
                next: requestGroup =>
                  patchState(store, {requestGroup, isLoading: false}),
                error: (err: HttpErrorResponse) => {
                  patchState(store, {isLoading: false});

                  if (
                    err.status === 404 &&
                    id === storedGroupService.storedRequestGroupId
                  ) {
                    storedGroupService.storedRequestGroupId = null;
                  }

                  router
                    .navigate(['/'])
                    .then(() => console.log(`${id} not found`))
                    .catch(() => console.log('Could not redirect'));
                },
              })
            );
          })
        )
      ),
    })
  )
);
