import {
  ApplicationConfig,
  ErrorHandler,
  provideZoneChangeDetection,
} from '@angular/core';
import {provideRouter, withComponentInputBinding} from '@angular/router';
import {provideHttpClient, withInterceptors} from '@angular/common/http';
import {httpInterceptor} from './core/http/interceptor';
import {CustomErrorHandler} from './core/error/handler';

import {routes} from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({eventCoalescing: true}),
    provideHttpClient(withInterceptors([httpInterceptor])),
    provideRouter(routes, withComponentInputBinding()),
    {provide: ErrorHandler, useClass: CustomErrorHandler},
  ],
};
