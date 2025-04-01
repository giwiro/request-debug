import {HttpEvent, HttpHandlerFn, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../../environments/environment';

export function httpInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn
): Observable<HttpEvent<unknown>> {
  const url = new URL(req.url, environment.baseApiUrl);

  const newReq = req.clone({
    url: url.href,
  });

  return next(newReq);
}
