import {ErrorHandler, inject} from '@angular/core';
import {NotifiableError} from './exc';
import {AlertService} from '../../shared/alert/service/alert.service';
import {AlertType} from '../../shared/alert/model';

export class CustomErrorHandler implements ErrorHandler {
  alertService = inject(AlertService);

  handleError(error: Error) {
    if (error instanceof NotifiableError) {
      this.alertService.triggerAlert({
        type: AlertType.Error,
        message: error.message,
        closeAfter: 2500,
      });
    } else {
      console.error(error);
    }
  }
}
