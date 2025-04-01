import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {RequestGroup} from '../../../../core/models';
import {catchError, of, retry, tap, throwError} from 'rxjs';
import {StoredGroupService} from '../stored-group/stored-group.service';
import {Router} from '@angular/router';
import {NotifiableError} from '../../../../core/error/exc';

@Injectable({
  providedIn: 'root',
})
export class RequestGroupService {
  private http = inject(HttpClient);
  private storedGroupService = inject(StoredGroupService);
  private router = inject(Router);

  getRequestGroup(id: string) {
    return this.http.get<RequestGroup>(`/api/request/group/${id}`).pipe(
      retry({
        count: 1,
        delay: (error: HttpErrorResponse) => {
          if (error.status === 500) {
            return of(null);
          }
          return throwError(() => error);
        },
      }),
      catchError((err: HttpErrorResponse) => {
        if (
          err.status === 404 &&
          id === this.storedGroupService.storedRequestGroupId
        ) {
          this.storedGroupService.storedRequestGroupId = null;
        }

        this.router
          .navigate(['/'])
          .then(() => console.log(`Redirecting to '/'`))
          .catch(() => console.log("Could not redirect to '/'"));

        return throwError(
          () => new NotifiableError(`Could not get request group ${id}`)
        );
      })
    );
  }

  createRequestGroup() {
    return this.http.post<RequestGroup>(`/api/request/group/`, null).pipe(
      tap(requestGroup => {
        this.storedGroupService.storedRequestGroupId = requestGroup.id;

        this.router
          .navigate(['/dashboard', requestGroup.id])
          .then(() =>
            console.log(`Redirect to '/dashboard/${requestGroup.id}'`)
          )
          .catch(() =>
            console.log(`Could not redirect to '/dashboard/${requestGroup.id}'`)
          );
      })
    );
  }
}
