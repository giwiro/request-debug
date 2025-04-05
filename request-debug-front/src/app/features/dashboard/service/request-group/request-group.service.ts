import {inject, Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {RequestGroup} from '../../../../core/models';
import {catchError, of, retry, tap, throwError} from 'rxjs';
import {StoredGroupService} from '../stored-group/stored-group.service';
import {Router} from '@angular/router';
import {NotifiableError} from '../../../../core/error/exc';
import {AlertType} from '../../../../shared/alert/model';
import {AlertService} from '../../../../shared/alert/service/alert.service';

@Injectable({
  providedIn: 'root',
})
export class RequestGroupService {
  private http = inject(HttpClient);
  private storedGroupService = inject(StoredGroupService);
  private router = inject(Router);
  private alertService = inject(AlertService);

  getRequestGroup(id: string) {
    return this.http.get<RequestGroup>(`/api/request/group/${id}`).pipe(
      /*retry({
        count: 1,
        delay: (error: HttpErrorResponse) => {
          if (error.status === 500) {
            return of(null);
          }
          return throwError(() => error);
        },
      }),*/
      catchError((error: HttpErrorResponse) => {
        if (
          error.status === 404 &&
          id === this.storedGroupService.storedRequestGroupId
        ) {
          this.storedGroupService.storedRequestGroupId = null;
        }

        /*this.router
          .navigate(['/'])
          .then(() => console.log(`Redirecting to '/'`))
          .catch(() => console.log("Could not redirect to '/'"));*/

        this.alertService.triggerAlert({
          type: AlertType.Error,
          message: `Could not get request group ${id}`,
          // closeAfter: 2500,
        });

        return throwError(() => error);
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
      }),
      catchError((error: HttpErrorResponse) => {
        this.alertService.triggerAlert({
          type: AlertType.Error,
          message: 'Could not get create group',
          // closeAfter: 2500,
        });

        return throwError(() => error);
      })
    );
  }
}
