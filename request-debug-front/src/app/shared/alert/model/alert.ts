export interface Alert {
  id: string;
  type: AlertType;
  message: string;
  closeAfter?: number;
}

export interface CreateAlertOptions {
  type: AlertType;
  message: string;
  closeAfter?: number;
}

export enum AlertType {
  Success,
  Error,
  Info,
  Warning,
}
