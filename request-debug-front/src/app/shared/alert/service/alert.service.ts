import {computed, Injectable, signal} from '@angular/core';
import {exhaustMap, of, Subject} from 'rxjs';
import {Alert, CreateAlertOptions} from '../model';

interface AlertServiceOperation {
  operation: 'CREATE' | 'DELETE';
  createAlert?: Alert;
  deleteAlertId?: string;
}

@Injectable({
  providedIn: 'root',
})
export class AlertService {
  private _alerts = signal<Alert[]>([]);
  private operationsSubject = new Subject<AlertServiceOperation>();

  readonly alerts = computed(() => this._alerts());

  constructor() {
    this.operationsSubject
      .pipe(
        exhaustMap((op: AlertServiceOperation) => {
          const {operation, deleteAlertId, createAlert} = op;
          if (operation === 'DELETE') {
            const filteredAlerts = this._alerts().filter(
              a => a.id !== deleteAlertId
            );
            this._alerts.update(() => filteredAlerts);
          } else if (operation === 'CREATE' && createAlert) {
            this._alerts.update(a => [...a, createAlert]);
          }

          return of(op);
        })
      )
      .subscribe(({operation, deleteAlertId, createAlert}) => {
        if (operation === 'DELETE') {
          console.log(`Alert '${deleteAlertId}' deleted`);
        } else if (operation === 'CREATE' && createAlert) {
          console.log(`Alert '${createAlert.id}' created`);
        }
      });
  }

  private generateRandomId() {
    return '_' + Math.random().toString(36).substring(2, 9);
  }

  triggerAlert({type, message, closeAfter}: CreateAlertOptions) {
    const id = this.generateRandomId();

    this.operationsSubject.next({
      operation: 'CREATE',
      createAlert: {
        id,
        type,
        message,
        closeAfter,
      },
    });

    if (typeof closeAfter === 'number') {
      setTimeout(() => {
        this.operationsSubject.next({
          operation: 'DELETE',
          deleteAlertId: id,
        });
      }, closeAfter);
    }
  }

  deleteAlert(deleteAlertId: string) {
    this.operationsSubject.next({
      operation: 'DELETE',
      deleteAlertId,
    });
  }
}
