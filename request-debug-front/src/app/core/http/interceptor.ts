import {HttpEvent, HttpHandlerFn, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../../environments/environment';
import {join} from '../../shared/utils/path';

export function httpInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn
): Observable<HttpEvent<unknown>> {
  const url = join(environment.baseApiUrl, req.url);

  const newReq = req.clone({
    url,
  });

  return next(newReq);
}
