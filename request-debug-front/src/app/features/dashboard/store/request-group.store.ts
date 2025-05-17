import {RequestGroup} from '../../../core/models';
import {patchState, signalStore, withMethods, withState} from '@ngrx/signals';
import {tapResponse} from '@ngrx/operators';
import {inject} from '@angular/core';
import {RequestGroupService} from '../service/request-group/request-group.service';
import {rxMethod} from '@ngrx/signals/rxjs-interop';
import {exhaustMap, pipe, tap} from 'rxjs';
import {Request} from '../../../core/models';

interface RequestGroupStoreState {
  isLoading: boolean;
  requestGroup: RequestGroup | null;
  isDeleting: boolean;
}

const initialState: RequestGroupStoreState = {
  isLoading: false,
  requestGroup: null,
  isDeleting: false,
};

export const RequestGroupStore = signalStore(
  {providedIn: 'root'},
  withState(initialState),
  withMethods((store, requestGroupService = inject(RequestGroupService)) => ({
    addRequest: (request: Request) =>
      patchState(store, state => {
        if (!state.requestGroup) return state;
        const copy: RequestGroup = JSON.parse(
          JSON.stringify(state.requestGroup)
        );

        copy.requests = [...copy.requests, request];

        return {
          isLoading: state.isLoading,
          requestGroup: copy,
          isDeleting: state.isDeleting,
        };
      }),
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
    deleteRequest: rxMethod<{requestGroupId: string; requestId: string}>(
      pipe(
        tap(({requestId}) => {
          const updatedRequestGroup = store.requestGroup();

          if (updatedRequestGroup) {
            updatedRequestGroup.requests = updatedRequestGroup.requests.filter(
              r => r.id !== requestId
            );
          }

          patchState(store, {
            isDeleting: true,
            requestGroup: {...updatedRequestGroup} as RequestGroup,
          });
        }),
        exhaustMap(({requestGroupId, requestId}) => {
          return requestGroupService
            .deleteRequest(requestGroupId, requestId)
            .pipe(
              tapResponse({
                next: () => patchState(store, {isDeleting: false}),
                error: () => {
                  patchState(store, {isDeleting: false});
                },
              })
            );
        })
      )
    ),
  }))
);
