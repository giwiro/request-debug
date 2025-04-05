import {RequestGroup} from '../../../core/models';
import {patchState, signalStore, withMethods, withState} from '@ngrx/signals';
import {tapResponse} from '@ngrx/operators';
import {inject} from '@angular/core';
import {RequestGroupService} from '../service/request-group/request-group.service';
import {rxMethod} from '@ngrx/signals/rxjs-interop';
import {exhaustMap, pipe, tap} from 'rxjs';

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
  withMethods((store, requestGroupService = inject(RequestGroupService)) => ({
    createRequestGroup: rxMethod<void>(
      pipe(
        tap(() => patchState(store, {isLoading: true, requestGroup: null})),
        exhaustMap(() => {
          return requestGroupService.createRequestGroup().pipe(
            tapResponse({
              next: requestGroup =>
                patchState(store, {requestGroup, isLoading: false}),
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
              error: () => {
                patchState(store, {isLoading: false});
              },
            })
          );
        })
      )
    ),
  }))
);
