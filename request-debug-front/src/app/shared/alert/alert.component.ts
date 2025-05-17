import {Component, inject} from '@angular/core';
import {AlertService} from './service/alert.service';
import {AlertType} from './model';

@Component({
  selector: 'app-alert',
  imports: [],
  templateUrl: './alert.component.html',
  standalone: true,
})
export class AlertComponent {
  alertService = inject(AlertService);

  getAlertTypeClass(alertType: AlertType) {
    switch (alertType) {
      case AlertType.Error:
        return 'alert-error';
      case AlertType.Info:
        return 'alert-info';
      case AlertType.Success:
        return 'alert-success';
      case AlertType.Warning:
        return 'alert-warning';
      default:
        return 'alert-primary';
    }
  }
}
